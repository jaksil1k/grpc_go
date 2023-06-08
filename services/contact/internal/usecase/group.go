package usecase

import (
	"context"
	"grpc/services/contact/internal/delivery"
	"grpc/services/contact/internal/domain"
)
import "grpc/services/contact/internal/repository"

type Group interface {
	CreateGroup(ctx context.Context, group delivery.CreateGroupDto) int64
	GetGroupById(ctx context.Context, id int64) domain.Group
	AddContactToGroup(ctx context.Context, dto delivery.AddContactToGroupDto) bool
}

type groupService struct {
	repository repository.GroupRepository
}

func newGroupService(groupRepository repository.GroupRepository) Group {
	return &groupService{repository: groupRepository}
}

func (g *groupService) GetGroupById(ctx context.Context, id int64) domain.Group {
	return g.repository.Get(id)
}

func (g *groupService) CreateGroup(ctx context.Context, group delivery.CreateGroupDto) int64 {
	return g.repository.Create(&domain.Group{Name: group.Name})
}

func (g *groupService) AddContactToGroup(ctx context.Context, dto delivery.AddContactToGroupDto) bool {
	return g.repository.AddContactToGroup(dto.GroupID, dto.Contact)
}
