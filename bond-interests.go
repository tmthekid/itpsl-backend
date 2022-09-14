package main

import (
	"context"
	"fmt"
	"time"
)

func (c *Connection) GetBondInterests() ([]BondInterest, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := c.db.Query(ctx, GET_BOND_INTERESTS_QUERY)
	if queryErr != nil {
		fmt.Println("line 14, bond-interests.go", queryErr)
		return nil, queryErr
	}
	var bondInterests []BondInterest
	for rows.Next() {
		var bondInterest BondInterest
		scanErr := rows.Scan(&bondInterest.Id, &bondInterest.Min, &bondInterest.Max, &bondInterest.Interest)
		if scanErr != nil {
			fmt.Println("line 22, bond-interest.go", scanErr)
			return nil, scanErr
		}
		bondInterests = append(bondInterests, bondInterest)
	}
	return bondInterests, nil
}