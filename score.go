package main

import (
	"context"
	"fmt"
	"time"
)

func (c *Connection) GetScores() ([]Score, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := c.db.Query(ctx, GET_SCORES_QUERY)
	if queryErr != nil {
		fmt.Println("line 14, score.go", queryErr)
		return nil, queryErr
	}
	var scores []Score
	for rows.Next() {
		var score Score
		scanErr := rows.Scan(&score.Id, &score.Category, &score.Score, &score.Growth, &score.Stars)
		if scanErr != nil {
			fmt.Println("line 22, score.go", scanErr)
			return nil, scanErr
		}
		scores = append(scores, score)
	}
	return scores, nil
}