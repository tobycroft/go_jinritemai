package main

import (
	"log"
	"main.go/app/doudian"
	"main.go/app/doudian/login"
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
	login.Doudian_Login()

	//page, err := doudian.PlayWrightMain.Context.NewPage()
	//if err != nil {
	//	log.Fatalf("could not create page: %v", err)
	//}
	//if rpp, err := page.Goto("https://fxg.jinritemai.com/ffa/morder/order/list"); err != nil {
	//	log.Fatalf("could not goto: %v", err)
	//} else {
	//	log.Println(rpp.Text())
	//}
	//fmt.Println(doudian.PlayWrightMain.Context.Cookies())
	//
	//fmt.Println("end")
	//time.Sleep(1 * time.Minute)
	//entries, err := page.Locator("table").All()
	//if err != nil {
	//	log.Fatalf("could not get entries: %v", err)
	//}
	//for i, entry := range entries {
	//	title, err := entry.Locator("tbody").TextContent()
	//	if err != nil {
	//		log.Fatalf("could not get text content: %v", err)
	//	}
	//	fmt.Printf("%d: %s\n", i+1, title)
	//}
	////if err = browser.Close(); err != nil {
	////	log.Fatalf("could not close browser: %v", err)
	////}
	//if err = doudian.PlayWrightMain.PlayWright.Stop(); err != nil {
	//	log.Fatalf("could not stop Playwright: %v", err)
	//}
}
