package repository

import (
	"fmt"

	"github.com/g3techlabs/revit-api/src/config"
	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	georesponse "github.com/g3techlabs/revit-api/src/core/geolocation/response"
	"github.com/g3techlabs/revit-api/src/core/route/response"
	"github.com/g3techlabs/revit-api/src/db"
	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
)

const acceptedStatusId uint = 1

var cloudFrontUrl string = config.Get("AWS_CLOUDFRONT_URL")

type RouteRepository interface {
	CreateRoute(userId uint, startLocation, destination geoinput.Coordinates) (uint, error)
	GetRouteOwner(userId, routeId uint) (*OwnerInfo, error)
	GetFriendsToInvite(userId uint) (*[]response.OnlineFriendsResponse, error)
	GetNearbyUsersDetails(nearbyIds []uint) (*[]response.NearbyUserToRouteResponse, error)
	AcceptRouteInvite(userId, routeId uint, coordinates *geoinput.Coordinates) (*georesponse.UserDetails, error)
}

type routeRepository struct {
	db *gorm.DB
}

func NewRouteRepository() RouteRepository {
	return &routeRepository{
		db: db.Db,
	}
}

func (r *routeRepository) CreateRoute(userId uint, startLocation, destination geoinput.Coordinates) (uint, error) {
	route := models.Route{
		Destination: gorm.Expr("ST_SetSRID(ST_MakePoint(?, ?), 4326)", destination.Long, destination.Lat),
	}

	r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&route).Error; err != nil {
			return err
		}

		routeOwner := models.RouteParticipant{
			UserID:        userId,
			RouteID:       route.ID,
			StartLocation: gorm.Expr("ST_SetSRID(ST_MakePoint(?, ?), 4326)", startLocation.Long, startLocation.Lat),
			IsOwner:       true,
		}

		if err := tx.Create(&routeOwner).Error; err != nil {
			return err
		}

		return nil
	})

	return route.ID, nil
}

func (r *routeRepository) GetRouteOwner(userId, routeId uint) (*OwnerInfo, error) {
	var owner OwnerInfo

	if err := r.db.Model(&models.RouteParticipant{}).
		Select("u.nickname", "route_participant.is_owner", "ST_Y(event.location::geometry) as lat", "ST_X(event.location::geometry) as long").
		Joins("JOIN route r ON r.id = route_participant.route_id AND r.finished_at IS NULL and r.is_done = FALSE").
		Joins("JOIN users u ON u.id = route_participant.user_id").
		Where("user_id = ? AND route_id = ?", userId, routeId).
		Where("ended_at IS NULL AND is_owner = TRUE").Scan(&owner).Error; err != nil {
		return nil, err
	}

	return &owner, nil
}

func (r *routeRepository) GetFriendsToInvite(userId uint) (*[]response.OnlineFriendsResponse, error) {
	var friends []response.OnlineFriendsResponse

	subQuery := r.db.
		Model(&models.Friendship{}).
		Select(`
		CASE 
			WHEN requester_id = ? THEN receiver_id
			WHEN receiver_id = ? THEN requester_id
		END AS friend_id
	`, userId, userId).
		Where("requester_id = ? OR receiver_id = ?", userId, userId).
		Where("invite_status_id = ?", acceptedStatusId).
		Where("removed_at IS NULL AND removed_by_id IS NULL")

	if err := r.db.
		Table("users").
		Select("users.id AS friend_id, users.nickname, COALESCE(? || users.profile_pic) AS profile_pic", cloudFrontUrl).
		Joins("JOIN (?) AS sub ON sub.friend_id = users.id", subQuery).
		Scan(&friends).Error; err != nil {
		return nil, err
	}

	return &friends, nil
}

func (r *routeRepository) GetNearbyUsersDetails(nearbyIds []uint) (*[]response.NearbyUserToRouteResponse, error) {
	var nearbyUsers []response.NearbyUserToRouteResponse

	if err := r.db.Model(&models.User{}).
		Select("id, nickname, profile_pic").
		Where("id IN ?", nearbyIds).
		Scan(&nearbyUsers).Error; err != nil {
		return nil, err
	}

	return &nearbyUsers, nil
}

func (r *routeRepository) AcceptRouteInvite(userId, routeId uint, coordinates *geoinput.Coordinates) (*georesponse.UserDetails, error) {
	var userDetails georesponse.UserDetails

	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var routeCount int64
	if err := tx.Model(&models.Route{}).
		Where("id = ?", routeId).
		Where("is_done = FALSE AND finished_at IS NULL").
		Count(&routeCount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if routeCount == 0 {
		tx.Rollback()
		return nil, fmt.Errorf("route not found")
	}

	var participantCount int64
	if err := tx.Model(&models.RouteParticipant{}).
		Where("route_id = ? AND user_id = ?", routeId, userId).
		Count(&participantCount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if participantCount > 0 {
		tx.Rollback()
		return nil, fmt.Errorf("user already a participant")
	}

	routeParticipant := models.RouteParticipant{
		UserID:        userId,
		RouteID:       routeId,
		StartLocation: gorm.Expr("ST_SetSRID(ST_MakePoint(?, ?), 4326)", coordinates.Long, coordinates.Lat),
	}

	if err := tx.Create(&routeParticipant).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&models.User{}).
		Select("id AS user_id, nickname, COALESCE(? || profile_pic) AS profile_pic").
		Where("id = ?", userId).
		Scan(&userDetails).Error; err != nil {
		return nil, err
	}

	return &userDetails, tx.Commit().Error
}
