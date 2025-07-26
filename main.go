package main

import (
	"fmt"
	"log"
	"main.go/app/doudian"
	"main.go/app/doudian/Login"
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
	Login.DoudianCookieInject()
	//Login.DoudianLogin()
	err, users := User.GetUserInfo("2519720661353812")
	fmt.Println(users)

}
