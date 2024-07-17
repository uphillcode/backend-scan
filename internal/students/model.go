// package user

// type User struct {
//     ID    uint   `gorm:"primaryKey"`
//     Name  string `json:"name"`
//     Email string `json:"email"`
// }

package user

type Student struct {
	// ID    uint   `gorm:"primaryKey"`
	// Name  string `json:"name"`
	// Email string `json:"email"`
	Idstudent      uint   `gorm:"primaryKey"`
	Code           string `json:"code"`
	Carrer         string `json:"carrer"`
	Dni            string `json:"dni"`
	Fullname       string `json:"fullname"`
	Modality       string `json:"modality"`
	Fecha_registro string `json:"fecha_registro"`

	// idstudent int(11) AI PK
	// code varchar(45)
	// carrer varchar(45)
	// dni varchar(45)
	// fullname varchar(45)
	// modality varchar(45)
	// fecha_registro datetime
}

// func (Student) TableName() string {
// 	return "student"
// }
