package service

import (
	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (gs *GroupService) GetMembers(userId, groupId uint, query *input.GetMembersInput) (*response.GroupMembersResponse, error) {
	if err := gs.validator.Validate(query); err != nil {
		return nil, err
	}

	canCheck, err := gs.groupRepo.CanUserViewGroup(userId, groupId)
	if err != nil {
		return nil, generics.InternalError()
	}

	if !canCheck {
		return nil, generics.Forbidden("User cannot view group members")
	}

	members, err := gs.groupRepo.GetMembers(groupId, *query)
	if err != nil {
		return nil, generics.InternalError()
	}

	return members, nil
}
