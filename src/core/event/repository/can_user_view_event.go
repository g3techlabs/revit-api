package repository

func (er *eventRepository) CanUserViewEvent(userId, eventId uint) (bool, error) {
	var count int64

	err := er.db.Table("event e").
		Joins("LEFT JOIN event_subscriber es ON es.event_id = e.id AND es.user_id = ? AND es.invite_status_id = ? AND es.left_at IS NULL AND es.removed_by IS NULL", userId, acceptedStatusId).
		Where("e.id = ? AND e.canceled = FALSE AND (e.visibility_id = ? OR es.user_id IS NOT NULL)", eventId, publicVisibility).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
