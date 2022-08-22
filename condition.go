package main

import (
	"context"
	"fmt"
	"time"
)

func (c *Connection) GetConditions() ([]Condition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	rows, queryErr := c.db.Query(ctx, GET_CONDITIONS_QUERY);
	if queryErr != nil {
		fmt.Println("line 14, condition.go", queryErr)
		return nil, queryErr
	}
	var conditions []Condition
	for rows.Next() {
		var condition Condition
		scanErr := rows.Scan(&condition.Id, &condition.Name, &condition.Color, &condition.Params)
		if scanErr != nil {
			fmt.Println("line 22, condition.go", scanErr)
			return nil, scanErr
		}
		conditions = append(conditions, condition)
	}
	rows.Close()
	return conditions, nil
}