package utils

import (
	"gorm.io/gorm"
)

type ColumnCount struct {
	Column string
	Count  int64
}

// CountByColumn cuenta los registros en una tabla específica basada en una columna específica.

type CountResult struct {
	Column string `json:"column"`
	Count  int    `json:"count"`
}

func GetGroupedColumnsCount(db *gorm.DB, table string, column string) ([]CountResult, error) {
	var results []CountResult
	query := `SELECT ` + column + ` AS ` + column + `, COUNT(*) AS count FROM ` + table + ` GROUP BY ` + column
	// query := `SELECT ` + column + ` AS column, COUNT(*) AS count FROM ` + table + ` GROUP BY ` + column
	if err := db.Raw(query).Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
