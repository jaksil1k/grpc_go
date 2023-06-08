package usecase

import "grpc/services/contact/internal/domain"

type contact interface {
	create(contact domain.Contact) int64
	get(id int64) domain.Contact
	createAndAddToGroup(contactId int64, group domain.Group)
}
