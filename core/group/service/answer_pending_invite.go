package service

import (
	"github.com/g3techlabs/revit-api/core/group/errors"
	"github.com/g3techlabs/revit-api/core/group/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (gs *GroupService) AnswerPendingInvite(userId, groupId uint, answer *input.AnswerPendingInvite) error {
	if err := gs.validator.Validate(answer); err != nil {
		return err
	}

	var err error
	switch answer.Answer {
	case "accept":
		err = gs.groupRepo.AcceptPendingInvite(groupId, userId)
	case "reject":
		err = gs.groupRepo.RejectPendingInvite(groupId, userId)
	}

	if err != nil {
		if err.Error() == "group invite not found" {
			return errors.GroupInviteNotFound()
		}
		return generics.InternalError()
	}

	return nil
}
