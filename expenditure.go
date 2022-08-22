package main

import (
	"context"
	"fmt"
	"time"
)

func (c *Connection) GetExpenditures() ([]Expenditure, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	rows, queryErr := c.db.Query(ctx, GET_EXPENDITURE_QUERY)
	if queryErr != nil {
		fmt.Println("line 14, expenditure.go", queryErr)
		return nil, queryErr
	}
	var expenditures []Expenditure
	for rows.Next() {
		var expenditure Expenditure
		scanErr := rows.Scan(&expenditure.Id, &expenditure.Name, &expenditure.Placeholder, &expenditure.Color)
		if scanErr != nil {
			fmt.Println("line 22, expenditure.go", scanErr)
			return nil, scanErr
		}
		expenditures = append(expenditures, expenditure)
	}
	rows.Close()
	return expenditures, nil
}