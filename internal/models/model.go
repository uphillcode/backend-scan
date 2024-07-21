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
	Code           uint      `json:"code"`
	Carrer         string    `json:"carrer"`
	Dni            uint      `json:"dni"`
	Fullname       string    `json:"fullname"`
	Modality       string    `json:"modality"`
	Fecha_registro time.Time `json:"fecha_registro" gorm:"default:CURRENT_TIMESTAMP"`
}

type Identity struct {
	Ididentity     uint      `gorm:"primaryKey"`
	Code           uint      `json:"code"`
	Litho          uint      `json:"litho"`
	Value          string    `json:"value"`
	Increment      uint      `json:"increment"`
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

type IdentityAdd struct {
	Identity
	Ididentity     *uint
	Fecha_registro *time.Time
}

func (StudentInsert) TableName() string {
	return "students"
}

func (IdentityAdd) TableName() string {
	return "identities"
}
