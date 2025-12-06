package repository

func (gr *groupRepository) CanUserViewGroup(userId, groupId uint) (bool, error) {
	var count int64

	err := gr.db.Table("groups g").
		Joins("LEFT JOIN group_member gm ON gm.group_id = g.id AND gm.user_id = ? AND gm.invite_status_id = ? AND gm.left_at IS NULL AND gm.removed_by IS NULL", userId, acceptedStatusId).
		Where("g.id = ? AND (g.visibility_id = ? OR gm.user_id IS NOT NULL)", groupId, PublicVisibility).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
