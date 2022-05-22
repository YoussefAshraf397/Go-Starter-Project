package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" grom:"type:varchar(50)"`
	Email    string `json:"email" 	 grom:"type:varchar(50)`
	Password string `json:"password" grom:"type:varchar(50)`
	Token    string `json:"token" grom:"type:varchar(100)`
}
