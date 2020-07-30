package idgenerator

import (
	"github.com/google/uuid"
)

type uuidTestIdGenerator struct {
}


func (s *uuidTestIdGenerator) GetId(testNumber int) string {
	return uuid.New().String()
}


func NewUUIDTestIdGenerator() TestIdGenerator {
	return &uuidTestIdGenerator{}
}
