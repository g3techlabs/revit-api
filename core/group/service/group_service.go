package service

import (
	"github.com/g3techlabs/revit-api/core/group/input"
	"github.com/g3techlabs/revit-api/core/group/repository"
	"github.com/g3techlabs/revit-api/core/group/response"
	"github.com/g3techlabs/revit-api/core/storage"
	"github.com/g3techlabs/revit-api/validation"
)

type IGroupService interface {
	CreateGroup(userId uint, data *input.CreateGroup) (*response.PresignedGroupPhotosInfo, error)
	ConfirmNewPhotos(userId, groupId uint, data *input.ConfirmNewPhotos) error
	GetGroups(userId uint, query *input.GetGroupsQuery) (*[]response.GetGroupsResponse, error)
}

type GroupService struct {
	groupRepo      repository.GroupRepository
	validator      validation.IValidator
	storageService storage.StorageService
}

func NewGroupService(groupRepository repository.GroupRepository, validator validation.IValidator, storageService storage.StorageService) IGroupService {
	return &GroupService{
		groupRepo:      groupRepository,
		validator:      validator,
		storageService: storageService,
	}
}
