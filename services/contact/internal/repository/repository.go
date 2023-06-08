package repository

import "grpc/services/contact/internal/domain"

type GroupRepository interface {
	Get(id int64) domain.Group
	Create(group *domain.Group) int64
	AddContactToGroup(groupId int64, contact domain.Contact) bool
}
