package repository

import (
	"fmt"
	"time"

	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
)

func (r *routeRepository) LeaveRoute(userId, routeId uint, leftAt time.Time) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.updateUserEndedAt(tx, routeId, userId, leftAt); err != nil {
			return err
		}

		finishedParticipants, err := r.getArrivalOrder(tx, routeId)
		if err != nil {
			return err
		}

		mustRouteFinish, err := r.mustRouteFinish(tx, finishedParticipants, routeId)
		if err != nil {
			return err
		}
		if mustRouteFinish {
			if err := r.markRouteAsDone(tx, routeId); err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *routeRepository) markRouteAsDone(tx *gorm.DB, routeId uint) error {
	update := tx.Model(&models.Route{}).
		Where("id = ? AND is_done = FALSE AND finished_at IS NULL", routeId).
		Update("is_done", true)

	if update.RowsAffected == 0 {
		return fmt.Errorf("route not found")
	}

	return update.Error
}
