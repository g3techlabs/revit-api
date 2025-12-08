package service

import (
	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (gs *GroupService) GetPendingInvites(userId uint, query *input.GetPendingInvites) (*response.GetPendingInvitesPaginatedResponse, error) {
	if err := gs.validator.Validate(query); err != nil {
		return nil, err
	}

	pendingInvitesResponse, err := gs.groupRepo.GetPendingInvites(userId, query.Page, query.Limit)
	if err != nil {
		return nil, generics.InternalError()
	}

	return pendingInvitesResponse, nil
}
