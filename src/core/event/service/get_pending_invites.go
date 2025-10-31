package service

import (
	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/core/event/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (es *EventService) GetPendingInvites(userId uint, query *input.GetPendingInvitesFilters) (*[]response.GetPendingInvites, error) {
	if err := es.validator.Validate(query); err != nil {
		return nil, err
	}

	pendingInvites, err := es.eventRepo.GetPendingInvites(userId, query.Limit, query.Page)
	if err != nil {
		return nil, generics.InternalError()
	}

	return pendingInvites, nil
}
