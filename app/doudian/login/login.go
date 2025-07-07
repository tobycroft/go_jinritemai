package login

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/playwright-community/playwright-go"
	"log"
	"main.go/app/doudian"
	"time"
)

func Doudian_Login() (err error) {
	doudian.PlayWrightMain.Page, err = doudian.PlayWrightMain.Context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
		return
	}
	_, err = doudian.PlayWrightMain.Page.Goto("https://fxg.jinritemai.com/login/common")
	if err != nil {
		log.Fatalf("could not goto: %v", err)
		return
	} else {
		//fmt.Println(resp.StatusText())
		//log.Println(rpp.Text())
	}
	//doudian.PlayWrightMain.Page.OnDOMContentLoaded(func(page playwright.Page) {
	//	fmt.Println("aa")
	//	//locate := doudian.PlayWrightMain.Page.GetByText("加载中")
	//	//fmt.Println(locate.TextContent())
	//	//page.click('text="智能内容审核"')
	//	log.Println(rpp.Text())
	//
	//})
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		goconfig.SaveConfigFile(&goconfig.ConfigFile{}, "conf.ini")
	} else {
		value, err := cfg.GetSection("jinritemai")
		if err != nil {
			cfg.SetValue("jinritemai", "mail", "youxiang")
			cfg.SetValue("jinritemai", "password", "mima")
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("jinritemai_ready")
		}
		err = doudian.PlayWrightMain.Page.GetByText("邮箱登录").Click()
		if err != nil {
			log.Fatalf("could not click: %v", err)
			return
		}
		err = doudian.PlayWrightMain.Page.GetByPlaceholder("请输入邮箱").Type(value["mail"])
		if err != nil {
			log.Fatalf("could not Type: %v", err)
			return
		}
		err = doudian.PlayWrightMain.Page.GetByPlaceholder("密码").Type(value["password"])
		if err != nil {
			log.Fatalf("could not Type: %v", err)
			return
		}
		err = doudian.PlayWrightMain.Page.GetByRole(*playwright.AriaRoleCheckbox).Check()
		if err != nil {
			log.Fatalf("could not Check: %v", err)
			return
		}
		err = doudian.PlayWrightMain.Page.GetByRole(*playwright.AriaRoleButton).GetByText("登录").Click()
		if err != nil {
			log.Fatalf("could not click: %v", err)
			return
		}
	}

	//fmt.Println(doudian.PlayWrightMain.Page.InputValue("请输入邮箱"))
	//fmt.Println(doudian.PlayWrightMain.Page.Click())

	time.Sleep(10 * time.Second)
	return
}
