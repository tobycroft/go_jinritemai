package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"main.go/app/doudian"
	"main.go/app/doudian/Login"
	"main.go/app/doudian/User"
	"main.go/config/app_conf"
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
	err := doudian.StartNormal()
	if err != nil {
		log.Fatal(err)
	}
	err = Login.DoudianCookieInject()
	if err != nil {
		log.Fatalf("could not inject cookie: %v", err)
	}
	//Login.DoudianCookieInject()
	//Login.DoudianLogin()
	fmt.Println(User.GetUserInfo("2519720661353812"))
	//time.Sleep(5 * time.Second)

}
