package Login

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/bytedance/sonic"
	"github.com/playwright-community/playwright-go"
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

const cookiePath = `/v2/doudian/cookie`

func DoudianCookieInject() (err error) {
	sign, salt := doudian.DoudianSecret()
	ret := new(Net.Net).New().SetPostData(map[string]string{
		"appid": doudian.PlayWrightMain.Appid,
		"sign":  sign,
		"salt":  salt,
		//"data":  nil,
	}).SetUrl(doudian.ApiUrl + cookiePath + "/get").PostFormData()
	rtt := doudian.RetStruct[doudian.CookieStruct]{}
	ret.RetJson(&rtt)
	if rtt.Code != 0 {
		log.Fatalf("could not get cookie: %v", rtt.Echo)
		return
	} else {
		//log.Println(rtt.Echo, rtt.Data)
	}

	opck := []playwright.OptionalCookie{}
	for _, cookie := range rtt.Data {
		same := playwright.SameSiteAttribute(cookie.SameSite)
		opck = append(opck, playwright.OptionalCookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Domain:   &cookie.Domain,
			Path:     &cookie.Path,
			Expires:  &cookie.Expires,
			HttpOnly: &cookie.HttpOnly,
			Secure:   &cookie.Secure,
			SameSite: &same,
		})
	}
	doudian.PlayWrightMain.Context.AddCookies(opck)
	return
}

func DoudianLogin() (err error) {

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
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("jinritemai_ready")
		}
		_, errs = cfg.GetSection("online_jinritemai")
		if errs != nil {
			cfg.SetValue("online_jinritemai", "appkey", "youxiang")
			cfg.SetValue("online_jinritemai", "appsecert", "mima")
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("jinritemai_online_ready")
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
			fmt.Println(doudian.PlayWrightMain.Page.URL())

			log.Fatalf("could not click1: %v", err)
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
		time.Sleep(1 * time.Second)
		err = doudian.PlayWrightMain.Page.GetByRole(*playwright.AriaRoleButton).GetByText("登录").Click()
		if err != nil {
			log.Fatalf("could not click2: %v", err)
			return
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

		sign, salt := doudian.DoudianSecret()
		ret := new(Net.Net).New().SetPostData(map[string]string{
			"appid": doudian.PlayWrightMain.Appid,
			"sign":  sign,
			"salt":  salt,
			"data":  ck_save,
		}).SetUrl(doudian.ApiUrl + cookiePath + "/auto").PostFormData()

		rtt := doudian.RetStruct[any]{}
		ret.RetJson(&rtt)
		if rtt.Code != 0 {
			log.Fatalf("could not get cookie: %v", rtt.Echo)
			return
		} else {
			log.Println(rtt.Echo, ck_save)
		}

		//fmt.Println(doudian.PlayWrightMain.Page.InputValue("请输入邮箱"))
		//fmt.Println(doudian.PlayWrightMain.Page.Click())

	}

	return
}
