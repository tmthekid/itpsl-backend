package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/xuri/excelize/v2"
)

func (c *Connection) SaveResult(r *http.Request) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	var payload Payload
	jsonErr := json.NewDecoder(r.Body).Decode(&payload)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return jsonErr
	}
	userId := 0; 
	row := c.db.QueryRow(ctx, CREATE_USER_QUERY, &payload.Name, &payload.Proffesion, &payload.Gender, &payload.Age, &payload.FinalGrowth)
	scanErr := row.Scan(&userId)
	if scanErr != nil { 
		fmt.Println(scanErr)
		return scanErr
	}
	if userId > 0 {
		for i := 0; i < len(payload.Data); i++ {
			data := payload.Data[i]
			finalStatId := 0; 
			row := c.db.QueryRow(ctx, CREATE_FINAL_STATS_QUERY, &data.Year, &data.Growth, &data.ScoreId, &data.ScoreCategory, userId)
			finalStatScanErr := row.Scan(&finalStatId)
			if finalStatScanErr != nil { 
				fmt.Println(finalStatScanErr)
				return finalStatScanErr
			}
			_, initialInformationQueryErr := c.db.Exec(ctx, CREATE_INITIAL_INFORMATION, &data.CityValue, &data.TaxRate, &data.BondsAmount, &data.BondsInterestRate, finalStatId)
			if initialInformationQueryErr != nil {
				fmt.Println(initialInformationQueryErr)
				return initialInformationQueryErr
			}
			for j := 0; j < len(data.Conditions); j++ {
				condition := data.Conditions[j]
				_, probabilityQueryErr := c.db.Exec(ctx, CREATE_PROBABILITIES_QUERY, &condition.ConditionId, &condition.Condition, &condition.Amount, &condition.Percentage, finalStatId)
				if probabilityQueryErr != nil {
					fmt.Println(probabilityQueryErr)
					return probabilityQueryErr
				}
			}
			for k := 0; k < len(data.Expenditures); k++ {
				expenditure := data.Expenditures[k]
				_, expenditureQueryErr := c.db.Exec(ctx, CREATE_EXPENDITURE_ALLOCATIONS_QUERY, &expenditure.ExpenditureId, &expenditure.Expenditure, &expenditure.Amount, finalStatId)
				if expenditureQueryErr != nil {
					fmt.Println(expenditureQueryErr)
					return expenditureQueryErr
				}
			}
		}
	}
	return nil
}

func (c *Connection) Download() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	rows, queryErr := c.db.Query(ctx, GET_RESULT_QUERY)
	if queryErr != nil {
		fmt.Println("line 68, result.go", queryErr)
		return nil, queryErr
	}
	var results []Result
	for rows.Next() {
		var result Result
		scanErr := rows.Scan(
			&result.Id,
			&result.Name,
			&result.Proffesion,
			&result.Gender,
			&result.Age,
			&result.Year,
			&result.Growth,
			&result.ScoreCategory,
			&result.CityValue,
			&result.TaxRate,
			&result.BondsAmount,
			&result.BondsInterestRate,
			&result.Drought,
			&result.Flood,
			&result.Landslide,
			&result.Tsunami,
			&result.SocialUnrest,
			&result.Normal,
			&result.PhysicalDevelopment,
			&result.EcoSystemPreservationProtection,
			&result.EconomicDevelopment,
			&result.PublicUtilitiesTransport,
			&result.WastePollutionManagement,
			&result.EducationHealthcare,
			&result.CulturalHeritageManagement,
			&result.SocialSupport,
			&result.DisasterMitigationManagement,
			&result.ResearchInnovationDevelopment,
			&result.MonitoringManagement,
			&result.FinalGrowth,
		)
		if scanErr != nil {
			fmt.Println("line 76, result.go", scanErr)
		}
		results = append(results, result)
	}
	f := excelize.NewFile()
	const sheetName = "data"
	f.SetSheetName("Sheet1", sheetName)
	firstCell, _ := excelize.JoinCellName("A", 1)
	f.SetSheetRow(sheetName, firstCell, &[]string{
		"#", 
		"Name",
		"Profession",
		"Gender",
		"Age",
		"Year",
		"Growth",
		"Score Category",
		"City Value",
		"Tax Rate",
		"Bonds Amount",
		"Bonds Interest Rate",
		"Drought",
		"Flood",
		"Landslide",
		"Tsunami",
		"Social Unrest",
		"Normal",
		"Physical Development",
		"Eco-System Preservation and Protection",
		"Economic Development",
		"Public Utilities and Transport",
		"Waste and Pollution Management",
		"Education and Healthcare",
		"Cultural and Heritage Management",
		"Social Support",
		"Disaster Mitigation and Management",
		"Research Innovation and Development",
		"Monitoring and Management",
		"Final Growth",
	})
	for i, row := range results {
		secondCell, _ := excelize.JoinCellName("A", i+2)
		f.SetSheetRow(sheetName, secondCell, &[]any{
			row.Id, 
			row.Name,
			row.Proffesion,
			row.Gender,
			row.Age,
			row.Year,
			row.Growth,
			row.ScoreCategory,
			row.CityValue,
			row.TaxRate,
			row.BondsAmount,
			row.BondsInterestRate,
			row.Drought,
			row.Flood,
			row.Landslide,
			row.Tsunami,
			row.SocialUnrest,
			row.Normal,
			row.PhysicalDevelopment,
			row.EcoSystemPreservationProtection,
			row.EconomicDevelopment,
			row.PublicUtilitiesTransport,
			row.WastePollutionManagement,
			row.EducationHealthcare,
			row.CulturalHeritageManagement,
			row.SocialSupport,
			row.DisasterMitigationManagement,
			row.ResearchInnovationDevelopment,
			row.MonitoringManagement,
			row.FinalGrowth,
		})
	}
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#12625E"}, Pattern: 1},
		Font: &excelize.Font{Bold: true, Color: "#ffffff"},
	})
	f.SetCellStyle(sheetName, "A1", "AD1", headerStyle)
	f.SetColWidth(sheetName, "C", "C", 11)
	f.SetColWidth(sheetName, "E", "E", 5)
	f.SetColWidth(sheetName, "F", "F", 5)
	f.SetColWidth(sheetName, "H", "H", 15)
	f.SetColWidth(sheetName, "I", "I", 10.15)
	f.SetColWidth(sheetName, "K", "K", 15)
	f.SetColWidth(sheetName, "L", "L", 19)
	f.SetColWidth(sheetName, "K", "K", 9.45)
	f.SetColWidth(sheetName, "O", "O", 9.45)
	f.SetColWidth(sheetName, "P", "P", 9.45)
	f.SetColWidth(sheetName, "Q", "Q", 12.80)
	f.SetColWidth(sheetName, "R", "R", 7.50)
	f.SetColWidth(sheetName, "S", "S", 21)
	f.SetColWidth(sheetName, "T", "T", 38)
	f.SetColWidth(sheetName, "U", "U", 28)
	f.SetColWidth(sheetName, "V", "V", 28)
	f.SetColWidth(sheetName, "W", "W", 32)
	f.SetColWidth(sheetName, "X", "X", 28)
	f.SetColWidth(sheetName, "Y", "Y", 33)
	f.SetColWidth(sheetName, "Z", "Z", 14)
	f.SetColWidth(sheetName, "AA", "AA", 35)
	f.SetColWidth(sheetName, "AB", "AB", 35)
	f.SetColWidth(sheetName, "AC", "AC", 30)
	f.SetColWidth(sheetName, "AD", "AD", 13)
	err := f.SaveAs("data.xlsx");
	if err != nil {
        fmt.Println(err)
		return nil, err
    }
	file, fileErr := ioutil.ReadFile("data.xlsx")
	if fileErr != nil {
        fmt.Println(fileErr)
		return nil, fileErr
    }
	return file, nil
}