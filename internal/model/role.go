package model

type Role struct {
	Name string `json:"name" gorm_scope:"varchar;not_null;unique"`
	Common
}
