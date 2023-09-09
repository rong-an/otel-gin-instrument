package models

import "time"

type UserInfo struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type User struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
