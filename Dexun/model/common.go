package model

type Api struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Role struct {
	RoleName string `json:"rolename"`
}
