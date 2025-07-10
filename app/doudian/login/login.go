package login

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/bytedance/sonic"
	"github.com/playwright-community/playwright-go"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"log"
	"main.go/app/doudian"
	"time"
)

//type _cookie struct {
//	Name     string                       `json:"name"`
//	Value    string                       `json:"value"`
//	Domain   string                       `json:"domain"`
//	Path     string                       `json:"path"`
//	Expires  float64                      `json:"expires"`
//	HttpOnly bool                         `json:"httpOnly"`
//	Secure   bool                         `json:"secure"`
//	SameSite playwright.SameSiteAttribute `json:"sameSite"`
//}

func DoudianLogin() (err error) {
	doudian.PlayWrightMain.Page, err = doudian.PlayWrightMain.Context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
		return
	}

	//doudian.PlayWrightMain.Page.OnDOMContentLoaded(func(page playwright.Page) {
	//	fmt.Println("aa")
	//	//locate := doudian.PlayWrightMain.Page.GetByText("加载中")
	//	//fmt.Println(locate.TextContent())
	//	//page.click('text="智能内容审核"')
	//	log.Println(rpp.Text())
	//
	//})
	cfg, errs := goconfig.LoadConfigFile("conf.ini")
	if errs != nil {
		goconfig.SaveConfigFile(&goconfig.ConfigFile{}, "conf.ini")
	} else {
		value, errs := cfg.GetSection("jinritemai")
		if errs != nil {
			cfg.SetValue("jinritemai", "mail", "youxiang")
			cfg.SetValue("jinritemai", "password", "mima")
			cfg.SetValue("jinritemai", "cookie", "cookie_DO_NOT_CHANGE")
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("jinritemai_ready")
		}
		value2, errs := cfg.GetSection("online_jinritemai")
		if errs != nil {
			cfg.SetValue("online_jinritemai", "appkey", "youxiang")
			cfg.SetValue("online_jinritemai", "appsecert", "mima")
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("jinritemai_online_ready")
		}
		salt := Calc.Any2String(Calc.Rand(1000, 9999))
		new(Net.Net).New().SetPostData(map[string]string{
			"appkey": value2["appkey"],
			"hash":   Calc.Md5(value2["appkey"] + value2["appsecert"] + salt + Calc.Any2String(time.Now().Unix())),
			"salt":   salt,
		})

		if value["cookie"] == "" || value["cookie_DO_NOT_CHANGE"] == "cookie_DO_NOT_CHANGE" {

		} else {
			//fmt.Println(value["cookie"])
			opck := []playwright.OptionalCookie{}
			//ck := []_cookie{}
			err = sonic.UnmarshalString(value["cookie"], &opck)
			if err != nil {
				log.Fatalf("could not unmarshal cookie: %v", err)
			} else {
				//opck := []playwright.OptionalCookie{}
				//for _, cookie := range ck {
				//	opck = append(opck, playwright.OptionalCookie{
				//		Name:     cookie.Name,
				//		Value:    cookie.Value,
				//		Domain:   &cookie.Domain,
				//		Path:     &cookie.Path,
				//		Expires:  &cookie.Expires,
				//		HttpOnly: &cookie.HttpOnly,
				//		Secure:   &cookie.Secure,
				//		SameSite: &cookie.SameSite,
				//	})
				//}
				doudian.PlayWrightMain.Context.AddCookies(opck)
			}
		}

		_, err = doudian.PlayWrightMain.Page.Goto("https://fxg.jinritemai.com/login/common")
		if err != nil {
			log.Fatalf("could not goto: %v", err)
			return
		} else {
			//fmt.Println(resp.StatusText())
			//log.Println(rpp.Text())
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
		time.Sleep(2 * time.Second)
		err = doudian.PlayWrightMain.Page.GetByRole(*playwright.AriaRoleButton).GetByText("登录").Click()
		if err != nil {
			log.Fatalf("could not click: %v", err)
			return
		}

		//fmt.Println(doudian.PlayWrightMain.Page.InputValue("请输入邮箱"))
		//fmt.Println(doudian.PlayWrightMain.Page.Click())

	}

	time.Sleep(5 * time.Second)
	cookie, errs := doudian.PlayWrightMain.Context.Cookies()
	if errs != nil {
		log.Fatalf("unable to get cookies: %v", errs)
	}
	ck_save, errs := sonic.MarshalString(cookie)
	if errs != nil {
		log.Fatalf("unable to save cookie: %v", errs)
	}
	cfg.SetValue("jinritemai", "cookie", ck_save)
	goconfig.SaveConfigFile(cfg, "conf.ini")
	return
}
