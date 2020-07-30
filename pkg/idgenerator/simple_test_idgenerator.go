package idgenerator

import (
	"fmt"
)

type simpleTestIdGenerator struct {
}


func (s *simpleTestIdGenerator) GetId(testNumber int) string {
	return fmt.Sprintf("%04d", testNumber)
}


func NewSimpleTestIdGenerator() TestIdGenerator {
	return &simpleTestIdGenerator{}
}