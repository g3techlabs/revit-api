package repository

import (
	"fmt"
	"time"

	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
)

func (r *routeRepository) FinishParticipant(userId, routeId uint, finishTime time.Time) (*FinishDetails, error) {
	var details FinishDetails

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.updateParticipant(tx, routeId, userId, finishTime); err != nil {
			return err
		}

		routeDuration, err := r.getRouteDuration(tx, routeId, finishTime)
		if err != nil {
			return err
		}

		arrivalOrder, err := r.getArrivalOrder(tx, routeId)
		if err != nil {
			return err
		}

		mustRouteFinish, err := r.mustRouteFinish(tx, arrivalOrder, routeId)
		if err != nil {
			return err
		}
		if mustRouteFinish {
			if err := r.finishRoute(tx, routeId); err != nil {
				return err
			}
		}

		details.ArrivalOrder = int(arrivalOrder)
		details.ArrivalTime = formatDuration(*routeDuration)

		return nil
	})

	return &details, err
}

func (r *routeRepository) updateParticipant(tx *gorm.DB, routeId, userId uint, finishTime time.Time) error {
	result := tx.Model(&models.RouteParticipant{}).
		Where("route_id = ? AND user_id = ? AND ended_at IS NULL", routeId, userId).
		Update("ended_at", finishTime)

	if result.RowsAffected == 0 {
		return fmt.Errorf("user not in this route")
	}

	return result.Error
}

func (r *routeRepository) getRouteDuration(tx *gorm.DB, routeId uint, finishTime time.Time) (*time.Duration, error) {
	var route models.Route
	if err := tx.First(&route, routeId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("route not found")
		}
		return nil, err
	}

	routeDuration := finishTime.Sub(*route.StartedAt)

	return &routeDuration, nil
}

func (r *routeRepository) getArrivalOrder(tx *gorm.DB, routeId uint) (int64, error) {
	var arrivalOrder int64
	err := tx.Model(&models.RouteParticipant{}).
		Where("route_id = ? AND ended_at IS NOT NULL", routeId).
		Count(&arrivalOrder).Error

	return arrivalOrder, err
}

func formatDuration(d time.Duration) string {
	duration := d.Round(time.Second)

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func (r *routeRepository) mustRouteFinish(tx *gorm.DB, finishedParticipants int64, routeId uint) (bool, error) {
	var totalParticipants int64

	if err := tx.Model(&models.RouteParticipant{}).
		Where("route_id = ?", routeId).
		Count(&totalParticipants).Error; err != nil {
		return false, err
	}

	return totalParticipants == finishedParticipants, nil
}

func (r *routeRepository) finishRoute(tx *gorm.DB, routeId uint) error {
	data := map[string]any{
		"finished_at": time.Now().UTC(),
		"is_done":     true,
	}

	result := tx.Model(&models.Route{}).
		Where("id = ? AND is_done = FALSE AND finished_at IS NULL", routeId).
		Updates(&data)

	if result.RowsAffected == 0 {
		return fmt.Errorf("route not found")
	}

	return result.Error
}
