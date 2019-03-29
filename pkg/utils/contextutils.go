package utils

import "context"

func GetTestId(ctx context.Context) string {
	testId := ctx.Value("testId")
	if testId != nil {
		return testId.(string)
	}
	return "N/A"
}

