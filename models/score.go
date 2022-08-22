package models

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
		scanErr := rows.Scan(&score.Id, &score.Category, &score.Range, &score.Growth, &score.Rating)
		if scanErr != nil {
			fmt.Println("line 20, score.go", scanErr)
			return nil, scanErr
		}
		scores = append(scores, score)
	}
	return scores, nil
}