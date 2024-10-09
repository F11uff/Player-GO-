package repository

import "player/internal/storage/postgresql/user"

type UserRep interface {
	GetUserId(userID string) (user.UserRegistration, error)
}
