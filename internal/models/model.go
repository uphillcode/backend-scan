package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Setting struct {
	Id       uint      `gorm:"primaryKey" json:"id"`
	Table    string    `json:"table"`
	Semestre string    `json:"semestre"`
	CreateAt time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeleteAt string    `json:"delete_at"`
	State    string    `json:"state"`
}

type SettingAdd struct {
	Setting
	Id       *uint      `gorm:"primaryKey" json:"id"`
	CreateAt *time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeleteAt *string    `json:"delete_at"`
}

type Student struct {
	Idstudent     uint      `gorm:"primaryKey" json:"id_student"`
	Code          uint      `json:"code"`
	Carrer        string    `json:"carrer"`
	Dni           uint      `json:"dni"`
	Fullname      string    `json:"fullname"`
	Modality      string    `json:"modality"`
	FechaRegistro time.Time `json:"fecha_registro" gorm:"default:CURRENT_TIMESTAMP"`
}

type Identity struct {
	Ididentity    uint      `gorm:"primaryKey" json:"id_identity"`
	Code          uint      `json:"code"`
	Litho         uint      `json:"litho"`
	Value         string    `json:"value"`
	Increment     uint      `json:"increment"`
	FechaRegistro time.Time `json:"fecha_registro" gorm:"default:CURRENT_TIMESTAMP"`
}

type Response struct {
	StudentID      int    `json:"student_id"`
	QuestionNumber int    `json:"question_number"`
	Response       string `json:"response"`
}

type Responses []Response

func (r Responses) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Responses) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, r)
}

type StudentResponse struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	Litho         string    `json:"litho"`
	Tema          string    `json:"tema"`
	StudentID     int       `json:"student_id"`
	FechaRegistro time.Time `json:"fecha_registro" gorm:"default:CURRENT_TIMESTAMP"`
	Responses     Responses `json:"responses" gorm:"type:json"`
}

type ResponseCustom[T any] struct {
	State   string `json:"state"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type StudentAdd struct {
	Student
	Idstudent     *uint
	FechaRegistro *time.Time
}

type IdentityAdd struct {
	Identity
	Ididentity    *uint
	FechaRegistro *time.Time
}

func (StudentAdd) TableName() string {
	return "students"
}

func (IdentityAdd) TableName() string {
	return "identities"
}
func (SettingAdd) TableName() string {
	return "settings"
}

type StudentResponseAdd struct {
	Litho     string    `json:"litho"`
	Tema      string    `json:"tema"`
	StudentID int       `json:"student_id"`
	Responses Responses `json:"responses"`
}

func (StudentResponseAdd) TableName() string {
	return "student_responses"
}

type ColumnCount struct {
	Column string
	Count  int64
}

type Duplicate struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ColumnValue string `gorm:"column_value"`
	Count       int    `gorm:"count"`
	Table       string `gorm:"table_name"`
}
