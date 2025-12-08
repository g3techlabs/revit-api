package service

import (
	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/core/group/repository"
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"github.com/g3techlabs/revit-api/src/infra/storage"
	"github.com/g3techlabs/revit-api/src/validation"
)

type IGroupService interface {
	CreateGroup(userId uint, data *input.CreateGroup) (*response.PresignedGroupPhotosInfo, error)
	ConfirmNewPhotos(userId, groupId uint, data *input.ConfirmNewPhotos) error
	GetGroups(userId uint, query *input.GetGroupsQuery) (*response.GetGroupsResponse, error)
	GetGroup(userId, groupId uint) (*response.GroupResponse, error)
	GetMembers(userId, groupId uint, query *input.GetMembersInput) (*response.GroupMembersResponse, error)
	UpdateGroup(userId, groupId uint, data *input.UpdateGroup) error
	RequestNewGroupPhotos(userId, groupId uint, data *input.RequestNewGroupPhotos) (*response.PresignedGroupPhotosInfo, error)
	JoinGroup(userId, groupId uint) error
	QuitGroup(userId, groupId uint) error
	InviteUser(groupAdminId, groupId, invitedId uint) error
	GetPendingInvites(userId uint, query *input.GetPendingInvites) (*response.GetPendingInvitesPaginatedResponse, error)
	AnswerPendingInvite(userId, groupId uint, answer *input.AnswerPendingInvite) error
	RemoveMember(adminGroupId, groupId, groupMemberId uint) error
	GetAdminGroups(userId uint, query *input.GetAdminGroupsInput) (*response.GetAdminGroupsResponse, error)
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
