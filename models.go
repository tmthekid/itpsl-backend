package main

import "github.com/jackc/pgx/v4/pgxpool"

type Connection struct {
	db *pgxpool.Pool
}

func GetModels(db *pgxpool.Pool) Connection {
	return Connection {
		db: db,
	}
}

type User struct {
	Id         	int    `json:"id"`
	Name       	string `json:"name"`
	Proffesion 	string `json:"proffesion"`
	Gender 		string `json:"gender"`
	Age        	string `json:"age"`
	FinalGrowth string `json:"final_growth"`
}

type InitialInformation struct {
	Id         	int    `json:"id"`
	CityValue float64 `json:"city_value"`
	TaxRate int `json:"tax_rate"` 
	BondsAmount int `json:"bonds_amount"`
	BondsInterestRate int `json:"bonds_interest_rate"`
}

type Probability struct {
	Id         int    `json:"id"`
	ConditionId  int `json:"condition_id"`
	Condition  string `json:"condition"`
	Amount float64 `json:"amount"`
	Percentage float64    `json:"percentage"`
}

type Condition struct {
	Id         int    `json:"id"`
	Name string `json:"name"`
	Color string `json:"color"`
	Params any `json:"params"`
}

type Expenditure struct {
	Id         int    `json:"id"`
	Name string `json:"name"`
	Placeholder string `json:"placeholder"`
	Color string `json:"color"`
}

type Score struct {
	Id         int    `json:"id"`
	Category string `json:"category"`
	Range string `json:"range"`
	Growth float64 `json:"growth"`
	Rating int `json:"rating"`
}