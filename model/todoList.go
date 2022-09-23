package model

import ("time")

type LoginInfo struct {
	UserID     int64     `json:"userId"`
	ClientIP   string    `json:"clientIP"`
	LoginState string    `json:"loginState"`
	LoginTime  time.Time `json:"loginTime"`
}

type Sunshareboy struct {
    Name string
    Age  int
    City string
    }


func NewLoginInfo(id int64, clientIP string, loginState string) *LoginInfo{
	return &LoginInfo{
		UserID: id,
		ClientIP: clientIP,
		LoginState: loginState,
		LoginTime: time.Now(),
	}
}