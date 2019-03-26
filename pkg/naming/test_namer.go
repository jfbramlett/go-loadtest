package naming


type TestNamer interface {
	GetName(testNumber int) string
}