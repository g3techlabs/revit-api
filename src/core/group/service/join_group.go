package service

import (
	"github.com/g3techlabs/revit-api/src/core/group/errors"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (gs *GroupService) JoinGroup(userId, groupId uint) error {
	err := gs.groupRepo.InsertNewGroupMember(userId, groupId)
	if err != nil {
		if err.Error() == "group not found" || err.Error() == "group is private" {
			return errors.GroupNotFound()
		}
		if err.Error() == "user already member" {
			return errors.UserIsAlreadyMember()
		}
		return generics.InternalError()
	}

	return nil
}
