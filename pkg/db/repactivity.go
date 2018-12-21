package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type RepActivity struct {
	DBConnection 		*sql.DB
}


func (r RepActivity) AddActivity() {

	var ctx = context.Background()
	result, err := r.DBConnection.ExecContext(ctx,
		`	INSERT INTO  RepActivityLog (CreatedTime, CompanyId, RepId, ActivityType, AppName)
				VALUES
				(?, ?, ?, ?, ?);`,
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