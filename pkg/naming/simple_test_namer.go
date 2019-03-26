package naming

import "strconv"

type simpleTestNamer struct {
}


func (s *simpleTestNamer) GetName(testNumber int) string {
	return strconv.FormatInt(int64(testNumber), 10)
}


func NewSimpleTestNamer() TestNamer {
	return &simpleTestNamer{}
}