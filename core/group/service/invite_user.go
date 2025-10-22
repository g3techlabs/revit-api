package service

import (
	"github.com/g3techlabs/revit-api/core/group/errors"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (gs *GroupService) InviteUser(groupAdminId, groupId, invitedId uint) error {
	if err := gs.groupRepo.MakeGroupInvitation(groupAdminId, groupId, invitedId); err != nil {
		if err.Error() == "requester not a group admin" {
			return errors.RequesterIsNotAnAdmin()
		}
		if err.Error() == "invited already invited/a member" {
			return errors.UserAlreadyInvitedOrMember()
		}
		return generics.InternalError()
	}

	return nil
}
