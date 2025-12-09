package repository

import (
	"time"

	"github.com/g3techlabs/revit-api/src/core/users/response"
	"gorm.io/gorm"
)

func (ur *userRepository) GetUserDetails(userId uint) (*response.GetUserResponse, error) {
	var userDetails response.GetUserResponse

	type UserRow struct {
		Name          string
		Nickname      string
		ProfilePicUrl *string
		CreatedAt     time.Time
	}
	var userRow UserRow
	if err := ur.db.Table("users").
		Select(`
			name,
			nickname,
			CASE 
				WHEN profile_pic IS NULL THEN NULL 
				ELSE CAST(? AS text) || profile_pic 
			END AS profile_pic_url,
			created_at
		`, cloudFrontUrl).
		Where("id = ?", userId).
		Scan(&userRow).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	userDetails.Name = userRow.Name
	userDetails.Nickname = userRow.Nickname
	userDetails.Since = userRow.CreatedAt
	userDetails.ProfilePicUrl = userRow.ProfilePicUrl

	// Buscar veículos
	type VehicleRow struct {
		ID       uint
		Nickname string
		Version  *string
		Year     uint
		Brand    string
		Model    string
	}
	var vehicleRows []VehicleRow
	if err := ur.db.Table("vehicle").
		Select("id, nickname, version, year, brand, model").
		Where("user_id = ? AND deleted_at IS NULL", userId).
		Order("created_at DESC").
		Scan(&vehicleRows).Error; err != nil {
		return nil, err
	}

	// Buscar fotos de veículos e montar resposta
	vehicles := make([]response.UserVehicle, 0, len(vehicleRows))
	for _, vRow := range vehicleRows {
		var photos []response.VehiclePhoto
		if err := ur.db.Table("photo").
			Select(`
				id,
				CAST(? AS text) || reference AS url
			`, cloudFrontUrl).
			Where("vehicle_id = ? AND deleted_at IS NULL", vRow.ID).
			Scan(&photos).Error; err != nil {
			return nil, err
		}

		vehicles = append(vehicles, response.UserVehicle{
			ID:       vRow.ID,
			Nickname: vRow.Nickname,
			Version:  vRow.Version,
			Year:     vRow.Year,
			Brand:    vRow.Brand,
			Model:    vRow.Model,
			Photos:   photos,
		})
	}
	userDetails.Vehicles = vehicles

	// Buscar até 3 grupos
	var groups []response.UserGroup
	groupsQuery := ur.db.Table("group_member gm").
		Select(`
			g.id AS group_id,
			g.name AS group_name,
			CASE 
				WHEN g.main_photo IS NULL THEN NULL 
				ELSE CAST(? AS text) || g.main_photo 
			END AS group_main_photo_url
		`, cloudFrontUrl).
		Joins("INNER JOIN groups g ON g.id = gm.group_id").
		Where("gm.user_id = ? AND gm.invite_status_id = ? AND gm.left_at IS NULL AND gm.removed_by IS NULL", userId, acceptedStatusId).
		Order("gm.member_since DESC").
		Limit(3)

	if err := groupsQuery.Scan(&groups).Error; err != nil {
		return nil, err
	}
	userDetails.Groups = groups

	// Buscar até 3 eventos futuros que o usuário está inscrito
	var events []response.UserEvent
	eventsQuery := ur.db.Table("event_subscriber es").
		Select(`
			e.id,
			e.name,
			CASE 
				WHEN e.photo IS NULL THEN NULL 
				ELSE CAST(? AS text) || e.photo 
			END AS event_photo_url
		`, cloudFrontUrl).
		Joins("INNER JOIN event e ON e.id = es.event_id").
		Where("es.user_id = ? AND es.invite_status_id = ? AND es.left_at IS NULL AND es.removed_by IS NULL", userId, acceptedStatusId).
		Where("e.date >= NOW() AND e.canceled = FALSE").
		Order("e.date DESC").
		Limit(3)

	if err := eventsQuery.Scan(&events).Error; err != nil {
		return nil, err
	}
	userDetails.Events = events

	return &userDetails, nil
}
