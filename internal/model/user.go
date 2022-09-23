package model

type User struct {
	Fullname string `json:"fullname" gorm_scope:"varchar;not_null"`
	Email    string `json:"email" gorm_scope:"varchar;not_null;unique"`
	Password string `json:"password" gorm_scope:"varchar;not_null"`
	RoleID   uint   `json:"role_id"`
	Role     Role
	Common
}
