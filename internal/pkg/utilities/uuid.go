package utilities

import "github.com/google/uuid"

type UUIDGenerator interface {
	GenerateUUID() string
}

func NewUUIDGenerator() UUIDGenerator {
	return &uuidGeneratorImpl{}
}

type uuidGeneratorImpl struct{}

func (u *uuidGeneratorImpl) GenerateUUID() string {
	return uuid.New().String()
}
