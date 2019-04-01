package utils

import "context"

const testIdKey="testId"

func GetTestIdFromContext(ctx context.Context) string {
	testId := ctx.Value(testIdKey)
	if testId != nil {
		return testId.(string)
	}
	return "N/A"
}

func SetTestIdInContext(ctx context.Context, testId string) context.Context {
	return context.WithValue(ctx, testIdKey, testId)
}

