package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type CompanyIssue struct {
	DBConnection 		*sql.DB
}


func (c CompanyIssue) AddActivity() {

	var ctx = context.Background()
	result, err := c.DBConnection.ExecContext(ctx,
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