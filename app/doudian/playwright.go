package doudian

import (
	"fmt"
	"log"
	"os"

	"github.com/Unknwon/goconfig"
	"github.com/playwright-community/playwright-go"
)

var PlayWrightMain struct {
	Appid           string
	AppSecert       string
	PlayWright      *playwright.Playwright
	Context         playwright.BrowserContext
	HeadLessContext playwright.BrowserContext
	Page            playwright.Page
	UserDir         string
}

func run() (err error) {
	if PlayWrightMain.PlayWright != nil {
		return nil
	}
	PlayWrightMain.PlayWright, err = playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
		return
	}
	return
}

func StartNormal() (err error) {
	err = doudian_ready()
	if err != nil {
		log.Fatalf("could not ready doudian: %v", err)
	}
	err = playwright_install()
	if err != nil {
		log.Fatalf("could not install playwright: %v", err)
		return err
	}
	err = run()
	if err != nil {
		log.Fatalf("could not run playwright: %v", err)
	}
	//browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
	//	Channel:  playwright.String("msedge"),
	//	Headless: playwright.Bool(false),
	//})
	if PlayWrightMain.UserDir == "" {
		if dir, err := os.UserCacheDir(); err != nil {
			PlayWrightMain.UserDir = dir + `\Microsoft\Edge\User Data\Default`
		}
	}
	PlayWrightMain.Context, err = PlayWrightMain.PlayWright.Chromium.LaunchPersistentContext(PlayWrightMain.UserDir, playwright.BrowserTypeLaunchPersistentContextOptions{
		Channel:  playwright.String("msedge"),
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
		return err
	}
	PlayWrightMain.Page, err = PlayWrightMain.Context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
		return err
	}
	return nil
}

func doudian_ready() (err error) {
	cfg, errs := goconfig.LoadConfigFile("conf.ini")
	if errs != nil {
		err = goconfig.SaveConfigFile(&goconfig.ConfigFile{}, "conf.ini")
	} else {
		value, errs := cfg.GetSection("online_jinritemai")
		if errs != nil {
			cfg.SetValue("online_jinritemai", "appid", "app_id")
			cfg.SetValue("online_jinritemai", "appsecert", "secret")
			err = goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("jinritemai_online_ready")
			return errs
		}
		PlayWrightMain.Appid = value["appid"]
		PlayWrightMain.AppSecert = value["appsecert"]
	}
	return
}

func playwright_install() (err error) {
	runOption := &playwright.RunOptions{
		SkipInstallBrowsers: true,
	}
	err = playwright.Install(runOption)
	return
}
