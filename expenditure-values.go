package main

import (
	"context"
	"fmt"
	"time"
)

func (c *Connection) GetExpenditureValues() ([]ExpenditureValue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := c.db.Query(ctx, GET_EXPENDITURE_VALUES_QUERY)
	if queryErr != nil {
		fmt.Println("line 14, expenditure-values.go", queryErr)
		return nil, queryErr
	}
	var expenditureValues []ExpenditureValue
	for rows.Next() {
		var expenditureValue ExpenditureValue
		scanErr := rows.Scan(&expenditureValue.Id, &expenditureValue.Expenditure, &expenditureValue.ConditionName, &expenditureValue.Value)
		if scanErr != nil {
			fmt.Println("line 22, expenditure-values.go", scanErr)
			return nil, scanErr
		}
		expenditureValues = append(expenditureValues, expenditureValue)
	}
	return expenditureValues, nil
}