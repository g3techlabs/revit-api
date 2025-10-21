package service

import (
	"github.com/g3techlabs/revit-api/core/group/errors"
	"github.com/g3techlabs/revit-api/core/group/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (gs *GroupService) UpdateGroup(userId, groupId uint, data *input.UpdateGroup) error {
	if err := gs.validator.Validate(data); err != nil {
		return err
	}

	if err := gs.groupRepo.UpdateGroup(userId, groupId, data); err != nil {
		if err.Error() == "group not found" {
			return errors.GroupNotFound()
		}
		return generics.InternalError()
	}

	return nil
}
