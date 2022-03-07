package entity

import "github.com/google/uuid"

func GenUuid() string {
	return uuid.New().String()
}

type UserID string

func (u UserID) String() string {
	return string(u)
}
