// package user

// type User struct {
//     ID    uint   `gorm:"primaryKey"`
//     Name  string `json:"name"`
//     Email string `json:"email"`
// }

package models

import (
	"time"
)

type Student struct {
	Idstudent      uint      `gorm:"primaryKey"`
	Code           string    `json:"code"`
	Carrer         string    `json:"carrer"`
	Dni            string    `json:"dni"`
	Fullname       string    `json:"fullname"`
	Modality       string    `json:"modality"`
	Fecha_registro time.Time `json:"fecha_registro" gorm:"default:CURRENT_TIMESTAMP"`
	// Fecha_registro time.Time `json:"fecha_registro"`

}

type Identity struct {
	Ididentity     uint      `gorm:"primaryKey"`
	Code           string    `json:"code"`
	Litho          string    `json:"litho"`
	Value          string    `json:"value"`
	Increment      string    `json:"increment"`
	Fecha_registro time.Time `json:"fecha_registro" gorm:"default:CURRENT_TIMESTAMP"`
}

type ResponseCustom[T any] struct {
	State   string `json:"state"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    T      `json:"data,omitempty"`
}

type StudentInsert struct {
	Student
	Idstudent      *uint
	Fecha_registro *time.Time
}

func (StudentInsert) TableName() string {
	return "students"
}
