package doudian

import (
	"github.com/playwright-community/playwright-go"
	"log"
	"os"
)

var PlayWrightMain struct {
	PlayWright *playwright.Playwright
	Context    playwright.BrowserContext
	Page       playwright.Page
	UserDir    string
}

func Start() error {
	var err error
	PlayWrightMain.PlayWright, err = playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
		return err
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
	return nil
}
