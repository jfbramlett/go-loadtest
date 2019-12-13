package naming

import (
	"fmt"
)

type simpleTestNamer struct {
}


func (s *simpleTestNamer) GetName(testNumber int) string {
	return fmt.Sprintf("%04d", testNumber)
}


func NewSimpleTestNamer() TestNamer {
	return &simpleTestNamer{}
}