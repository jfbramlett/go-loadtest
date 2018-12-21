package db

/**
INSERT INTO `CompanyIssue` (`CreatedTime`, `IssueId`, `CompanyId`, `CompanyGroupId`, `CustomerId`, `PlatformType`, `RegionId`, `RepId`, `RepIssuePos`, `IssueSecret`, `IssueTopicId`, `FirstCompanyEventLogSeq`, `LastCompanyEventLogSeq`, `CustomerState`, `EndedTime`, `Resolved`, `IssueStatus`, `IssueStatusTime`, `SentimentId`, `UpdateTime`)
VALUES
	(1545252128, 28, 200001, 1, 14, 7, 1, NULL, NULL, 'hello world', NULL, 1, 10, 1, NULL, 0, 5, NULL, NULL, '2018-12-19 20:42:08');

 */

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type RepActivity struct {
	DBConnection 		*sql.DB
}


func (r RepActivity) AddActivity() {

	var ctx = context.Background()
	result, err := r.DBConnection.ExecContext(ctx,
		`	INSERT INTO  CompanyIssue (
								CreatedTime, 
  								CompanyId,
  								CompanyGroupId,
  								RegionId,
  								CustomerId,
  								CustomerState,
  								IssueSecret,
  								FirstCompanyEventLogSeq,
  								LastCompanyEventLogSeq,
  								PlatformType)
				VALUES
				(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		time.Now().UnixNano(),
		200001,
		990001,
		17,
		"desk2")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}