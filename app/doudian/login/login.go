package login

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
	"main.go/app/doudian"
	"time"
)

func Doudian_Login() (err error) {
	doudian.PlayWrightMain.Page, err = doudian.PlayWrightMain.Context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
		return err
	}
	_, err = doudian.PlayWrightMain.Page.Goto("https://fxg.jinritemai.com/login/common")
	if err != nil {
		log.Fatalf("could not goto: %v", err)
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
	fmt.Println(doudian.PlayWrightMain.Page.GetByText("邮箱登录").Click())
	fmt.Println(doudian.PlayWrightMain.Page.GetByPlaceholder("请输入邮箱").Type("asdasd"))
	fmt.Println(doudian.PlayWrightMain.Page.GetByPlaceholder("密码").Type("asdasd"))
	fmt.Println(doudian.PlayWrightMain.Page.GetByRole(playwright.AriaRole("button")).GetByText("登录").Click())
	//fmt.Println(doudian.PlayWrightMain.Page.InputValue("请输入邮箱"))
	//fmt.Println(doudian.PlayWrightMain.Page.Click())

	time.Sleep(10 * time.Second)

	//text, err = locate.TextContent()

	//err = locate.Click()
	if err != nil {
		log.Fatalf("could not click: %v", err)
	}
	//entries, err := doudian.PlayWrightMain.Page.GetByText("邮箱登录")
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

	return nil
}
