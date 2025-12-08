package service

import (
	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (gs *GroupService) GetAdminGroups(userId uint, query *input.GetAdminGroupsInput) (*response.GetAdminGroupsResponse, error) {
	if err := gs.validator.Validate(query); err != nil {
		return nil, err
	}

	groups, err := gs.groupRepo.GetAdminGroups(userId, *query)
	if err != nil {
		return nil, generics.InternalError()
	}

	return groups, nil
}
