package service

import (
	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (gs *GroupService) GetGroups(userId uint, query *input.GetGroupsQuery) (*response.GetGroupsResponse, error) {
	if err := gs.validator.Validate(query); err != nil {
		return nil, err
	}

	groupsResponse, err := gs.groupRepo.GetGroups(userId, query)
	if err != nil {
		return nil, generics.InternalError()
	}

	return groupsResponse, nil
}
