package delivery

import "grpc/services/contact/internal/domain"

type CreateGroupDto struct {
	Name string
}

type AddContactToGroupDto struct {
	GroupID int64
	Contact domain.Contact
}
