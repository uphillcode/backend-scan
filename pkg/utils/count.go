package utils

import (
	"fmt"

	"gorm.io/gorm"
)

// CountResult es una estructura genérica para mantener el resultado de la consulta
type CountResult struct {
	ColumnValue string `json:"column"`
	Count       int    `json:"count"`
}

// GetGroupedColumnsCount cuenta los registros en una tabla específica basada en una columna específica y devuelve los resultados en un formato genérico.
func GetGroupedColumnsCount(db *gorm.DB, table string, column string) ([]CountResult, error) {
	var results []CountResult
	// Construir la consulta usando fmt.Sprintf para insertar dinámicamente el nombre de la columna
	query := fmt.Sprintf("SELECT `%s`, COUNT(*) AS count FROM `%s` GROUP BY `%s` ORDER BY count DESC", column, table, column)
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result CountResult
		err = rows.Scan(&result.ColumnValue, &result.Count)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
