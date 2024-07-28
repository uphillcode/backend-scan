package utils

import (
	"fmt"

	"gorm.io/gorm"
)

type ColumnCount struct {
	Column string
	Count  int64
}

// CountByColumn cuenta los registros en una tabla específica basada en una columna específica.
func CountByColumn(db *gorm.DB, tableName, columnName string) (int64, error) {
	var count int64
	query := fmt.Sprintf("%s IS NOT NULL", columnName)
	if err := db.Table(tableName).Where(query).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
