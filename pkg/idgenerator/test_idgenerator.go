package idgenerator


type TestIdGenerator interface {
	GetId(testNumber int) string
}