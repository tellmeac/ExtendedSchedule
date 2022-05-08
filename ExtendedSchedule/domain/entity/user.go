package entity

import "github.com/google/uuid"

type UserInfo struct {
	UserID         uuid.UUID `json:"userID"`
	UserIdentifier string    `json:"userIdentifier"`
}
