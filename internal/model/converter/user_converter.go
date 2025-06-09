package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserWithContactsToResponse(user *entity.User, contacts []entity.Contact) *model.UserWithContactResponse {
	contactResponses := make([]model.ContactResponse, len(contacts))
	for i, contact := range contacts {
		contactResponses[i] = model.ContactResponse{
			ID:        contact.ID,
			FirstName: contact.FirstName,
			Phone:     contact.Phone,
		}
	}

	return &model.UserWithContactResponse{
		ID:       user.ID,
		Name:     user.Name,
		Contacts: contactResponses,
	}
}

func UserToTokenResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		Token: user.Token,
	}
}

func UserToEvent(user *entity.User) *model.UserEvent {
	return &model.UserEvent{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
