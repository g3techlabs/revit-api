package service

import (
	"github.com/g3techlabs/revit-api/src/core/group/errors"
	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (gs *GroupService) UpdateGroup(userId, groupId uint, data *input.UpdateGroup) error {
	if err := gs.validator.Validate(data); err != nil {
		return err
	}

	if err := gs.groupRepo.UpdateGroup(userId, groupId, data); err != nil {
		switch err.Error() {
		case "group not found":
			return errors.GroupNotFound()
		case "city not found":
			return errors.CityNotFound()
		default:
			return generics.InternalError()
		}
	}

	return nil
}
