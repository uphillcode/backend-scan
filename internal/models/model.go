package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Academic_calendar struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Activity  string    `json:"activity"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT time.Time `json:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at"`
	TermsId   uint      `json:"terms_id"`
}
type Term struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	State     string    `json:"state"`
	Year      uint      `json:"year"`
	Number    uint      `json:"number"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT time.Time `json:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TemdAdd struct {
	Term
	Id        *uint      `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT *time.Time `json:"deleted_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Academic_calendarAdd struct {
	Academic_calendar
	Id        *uint      `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT *time.Time `json:"deleted_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// Table: academic_calendars
// Columns:
// id int(11) AI PK
// activity varchar(45)
// created_at varchar(45)
// deleted_at varchar(45)
// updated_at varchar(45)
// state varchar(45)
// terms_id int(11)
type Observation struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Code        string    `json:"code"`
	Litho       string    `json:"litho"`
	Tema        string    `json:"tema"`
	State       string    `json:"state"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT   string    `json:"deleted_at"`
	UpdatedAt   string    `json:"updated_at"`
	CalendarsID uint      `json:"calendars_id"`
}
type ObservationAdd struct {
	Observation
	Id        *uint      `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT *string    `json:"deleted_at"`
	UpdatedAt *string    `json:"updated_at"`
}

type StudentAndIdentity struct {
	Code      string `json:"code"`
	Carrer    string `json:"carrer"`
	Dni       string `json:"dni"`
	Fullname  string `json:"fullname"`
	Tema      string `json:"tema"`
	Litho     string `json:"litho"`
	Increment string `json:"increment"`
	Value     string `json:"value"`
}
type FilterDto struct {
	Text string `json:"text" query:"text"`
	// InstitutionID uint   `json:"institution_id"`
}

type History struct {
	Id uint `gorm:"primaryKey" json:"id"`
	// Code        string    `json:"code_student"`
	Code        string    `json:"code_student" gorm:"column:code_student"`
	Litho       string    `json:"litho"`
	Tema        string    `json:"tema"`
	Unanswered  int       `json:"unanswered"`
	Correct     int       `json:"correct"`
	Incorrect   int       `json:"incorrect"`
	Score       float64   `json:"score"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   string    `json:"deleted_at"`
	UpdatedAt   string    `json:"updated_at"`
	CalendarsID uint      `json:"calendars_id"`
}

type Cypher_code struct {
	Idcypher_code   uint      `gorm:"primaryKey" json:"idcypher_code"`
	Number_question uint      `json:"number_question"`
	Litho           string    `json:"litho"`
	Tema            string    `json:"tema"`
	Response        Responses `json:"response"`
	CreatedAt       time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt       string    `json:"deleted_at"`
	UpdatedAt       string    `json:"updated_at"`
	CalendarsID     uint      `json:"calendars_id"`
}

type Cypher_codeAdd struct {
	Cypher_code
	Idcypher_code *uint      `gorm:"primaryKey" json:"idcypher_code"`
	CreatedAt     *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt     *string    `json:"deleted_at"`
	UpdatedAt     *string    `json:"updated_at"`
	// Number_question uint      `json:"number_question"`
	// Litho           string    `json:"litho"`
	// Tema            string    `json:"tema"`
	// Response        Responses `json:"response"`
	// CalendarsID     uint      `json:"calendars_id"`

}

func (Cypher_codeAdd) TableName() string {
	return "cypher_codes"
}

type Setting struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Table       string    `json:"table"`
	Semestre    string    `json:"semestre"`
	CreatedAt   time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   string    `json:"delete_at"`
	State       string    `json:"state"`
	CalendarsID uint      `json:"calendars_id"`
}

type SettingAdd struct {
	Setting
	Id          *uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   *time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   *string    `json:"deleted_at"`
	CalendarsID *uint      `json:"calendars_id"`
}

type Student struct {
	Idstudent   uint      `gorm:"primaryKey" json:"id_student"`
	Code        uint      `json:"code"`
	Carrer      string    `json:"carrer"`
	Dni         uint      `json:"dni"`
	Fullname    string    `json:"fullname"`
	Tema        string    `json:"tema"`
	Modality    string    `json:"modality"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT   string    `json:"deleted_at"`
	UpdatedAt   string    `json:"updated_at"`
	CalendarsID uint      `json:"calendars_id"`
}

type Identity struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Code        uint      `json:"code"`
	Litho       uint      `json:"litho"`
	Value       string    `json:"value"`
	Increment   uint      `json:"increment"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT   time.Time `json:"deleted_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CalendarsID int       `json:"calendars_id"`
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
	ID                  int       `gorm:"primaryKey" json:"id"`
	Litho               string    `json:"litho"`
	Tema                string    `json:"tema"`
	StudentID           int       `json:"student_id"`
	CreatedAt           time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT           string    `json:"deleted_at"`
	UpdatedAt           string    `json:"updated_at"`
	Responses           Responses `json:"responses" gorm:"type:json"`
	CalendarsID         uint      `json:"calendars_id"`
	TemaAccordingExam   string    `json:"tema_according_exam"`
	TemaAccordingCareer string    `json:"tema_according_career"`
	Code                string    `json:"code"`
}

type Terms struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Year      uint      `json:"year"`
	Number    uint      `json:"number"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT time.Time `json:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ResponseCustom[T any] struct {
	State   string `json:"state"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type StudentAdd struct {
	Student
	Idstudent *uint
	CreatedAt *time.Time
	DeletedAT *string
	UpdatedAt *string
}

type IdentityAdd struct {
	Identity
	Id        *uint
	CreatedAt *time.Time
	DeletedAt *time.Time
	UpdatedAt *time.Time
}

type HistoryAdd struct {
	History
	Id        *uint
	CreatedAt *time.Time
	DeletedAt *time.Time
	UpdatedAt *time.Time
}

func (HistoryAdd) TableName() string {
	return "histories"
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
	StudentResponse
	Id        *int       `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAT *string    `json:"deleted_at"`
	UpdatedAt *string    `json:"updated_at"`
	// TemaAccordingCareer *string    `json:"tema_according_career"`
}

func (StudentResponseAdd) TableName() string {
	return "student_responses"
}

type ColumnCount struct {
	Column string
	Count  int64
}

type Duplicate struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ColumnValue string    `gorm:"column_value"`
	Count       int       `gorm:"count"`
	Table       string    `gorm:"table_name"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   string    `gorm:"deleted_at"`
	UpdatedAt   string    `gorm:"updated_at"`
	CalendarsID uint      `json:"calendars_id"`
}

type HistoryAndCalifications struct {
	History
	StudentAndIdentity
}
