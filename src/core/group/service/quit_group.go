package service

import (
	"github.com/g3techlabs/revit-api/src/core/group/errors"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (gs *GroupService) QuitGroup(userId, groupId uint) error {

	if err := gs.groupRepo.QuitGroup(userId, groupId); err != nil {
		switch err.Error() {
		case "user is not a member":
			return errors.UserIsNotAMember()
		default:
			return generics.InternalError()
		}
	}

	return nil
}
