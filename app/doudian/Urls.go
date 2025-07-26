package doudian

import "github.com/tobycroft/Calc"

const ApiUrl = `https://upload.tuuz.cc:444`

func DoudianSecret() (sign string, salt string) {
	salt = Calc.GenerateOrderId()
	sign = Calc.Md5(PlayWrightMain.Appid + PlayWrightMain.AppSecert + salt)
	return
}

type RetStruct[ty any | string | CookieStruct] struct {
	Code int    `json:"code"`
	Data ty     `json:"data"`
	Echo string `json:"echo"`
}
type CookieStruct []struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Value    string  `json:"value"`
	Domain   string  `json:"domain"`
	Secure   bool    `json:"secure"`
	Expires  float64 `json:"expires"`
	HttpOnly bool    `json:"httpOnly"`
	SameSite string  `json:"sameSite"`
}
