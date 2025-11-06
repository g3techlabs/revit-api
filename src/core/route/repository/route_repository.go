package repository

import (
	"github.com/g3techlabs/revit-api/src/config"
	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
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
