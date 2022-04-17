package entity

import "github.com/google/uuid"

type JoinedGroups []GroupInfo

type GroupInfo struct {
	ID   uuid.UUID
	Name string
}
