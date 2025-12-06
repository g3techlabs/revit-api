package service

import (
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (gs *GroupService) GetGroup(userId, groupId uint) (*response.GroupResponse, error) {
	groupResponse, err := gs.groupRepo.GetGroup(userId, groupId)
	if err != nil {
		return nil, generics.InternalError()
	}

	return groupResponse, nil
}
