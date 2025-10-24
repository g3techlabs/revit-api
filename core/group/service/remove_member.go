package service

import (
	"github.com/g3techlabs/revit-api/core/group/errors"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (gs *GroupService) RemoveMember(groupAdminId, groupId, groupMemberId uint) error {
	isUserAdmin, err := gs.groupRepo.IsUserAdmin(groupAdminId, groupId)
	if err != nil {
		return err
	}

	if !isUserAdmin {
		return generics.Forbidden("Specified admin is not an admin/owner")
	}

	if err := gs.groupRepo.RemoveMember(groupAdminId, groupId, groupMemberId); err != nil {
		if err.Error() == "member not found" {
			return errors.UserIsNotAMember()
		}
		return err
	}

	return nil
}
