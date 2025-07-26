package main

import (
	"fmt"
	"log"
	"main.go/app/doudian"
	"main.go/app/doudian/User"
	"main.go/config/app_conf"
	"os"
	"time"
)

func init() {
	time.Local = app_conf.TimeZone
	if app_conf.TestMode == false {
		s, err := os.Stat("./log/")

		if err != nil {
			os.Mkdir("./log", 0755)
		} else if s.IsDir() {
			os.Mkdir("./log", 0755)
		}
	}
}

func main() {

	//Calc.RefreshBaseNum()
	//go route.MainWsRouter()
	//mainroute := gin.Default()
	////gin.SetMode(gin.ReleaseMode)
	////gin.DefaultWriter = ioutil.Discard
	//mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	//mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	//route.OnRoute(mainroute)
	//mainroute.Run(":80")
	//err := playwright.Install()
	err := doudian.Start()
	if err != nil {
		log.Fatal(err)
	}
	//Login.DoudianCookieInject()
	//Login.DoudianLogin()
	err, users := User.GetUserInfo()
	fmt.Println(users)

}

type T struct {
	CheckSpamResp struct {
		Decision        string `json:"decision"`
		VerifyType      string `json:"verify_type"`
		CheckSpamDetail struct {
		} `json:"check_spam_detail"`
		SpamScene string `json:"spam_scene"`
	} `json:"check_spam_resp"`
	Code int `json:"code"`
	Data []struct {
		Id         string `json:"id"`
		ScreenName string `json:"screen_name"`
		AvatarUrl  string `json:"avatar_url"`
		DisPatch   bool   `json:"dis_patch"`
	} `json:"data"`
	Msg                    string `json:"msg"`
	Page                   int    `json:"page"`
	RequestArrivedTime     int64  `json:"request_arrived_time"`
	ServerExecutionEndTime int64  `json:"server_execution_end_time"`
	Size                   int    `json:"size"`
	St                     int    `json:"st"`
	Total                  int    `json:"total"`
}
