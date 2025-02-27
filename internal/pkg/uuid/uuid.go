package uuid

import "github.com/google/uuid"

//go:generate mockgen -source=uuid.go -destination=./mocks/uuid_mock.go -package=mocks

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
