package User

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"log"
	"main.go/app/doudian"
	"main.go/tuuz"
)

type userInfoStruct struct {
	CheckSpamResp struct {
		Decision        string `json:"decision"`
		VerifyType      string `json:"verify_type"`
		CheckSpamDetail struct {
		} `json:"check_spam_detail"`
		SpamScene string `json:"spam_scene"`
	} `json:"check_spam_resp"`
	Code                   int           `json:"code"`
	Data                   []UsersStruct `json:"data"`
	Msg                    string        `json:"msg"`
	Page                   int           `json:"page"`
	RequestArrivedTime     int64         `json:"request_arrived_time"`
	ServerExecutionEndTime int64         `json:"server_execution_end_time"`
	Size                   int           `json:"size"`
	St                     int           `json:"st"`
	Total                  int           `json:"total"`
}
type UsersStruct struct {
	Id         string `json:"id"`
	ScreenName string `json:"screen_name"`
	AvatarUrl  string `json:"avatar_url"`
	DisPatch   bool   `json:"dis_patch"`
}

const userPath = `/v2/doudian/user`

const getUserInfo = `https://pigeon.jinritemai.com/backstage/getuserinfo?uids=`

func GetUserInfo(uids string) (err error, users []UsersStruct) {
	// 获取用户信息
	// 这里可以使用 Playwright 的 API 来获取用户信息
	// 例如，获取页面内容并解析用户信息
	resp, err := doudian.PlayWrightMain.Page.Goto(getUserInfo + uids)
	if err != nil {
		log.Fatalf("could not goto: %v", err)
		return err, nil
	} else {
		body, err := resp.Body()
		if err != nil {
			log.Fatalf("could not get body: %v", err, tuuz.FUNCTION_ALL())
			return err, nil

		}
		var us userInfoStruct
		err = sonic.Unmarshal(body, &us)
		if err != nil {
			log.Fatalf("could not unmarshal: %v", err, tuuz.FUNCTION_ALL())
			return err, nil

		}
		if us.Code != 0 {
			return fmt.Errorf("error code: %d, message: %s", us.Code, us.Msg), nil
		}
		users = us.Data
		for _, user := range users {
			sign, salt := doudian.DoudianSecret()
			ret := new(Net.Net).New().SetPostData(map[string]string{
				"appid":       doudian.PlayWrightMain.Appid,
				"sign":        sign,
				"salt":        salt,
				"uid":         user.Id,
				"screen_name": user.ScreenName,
				"avatar_url":  user.AvatarUrl,
			}).SetUrl(doudian.ApiUrl + userPath + "/auto").PostFormData()
			rtt := doudian.RetStruct[any]{}
			ret.RetJson(&rtt)
			if rtt.Code != 0 {
				log.Fatalf("could not get cookie: %v", rtt.Echo)
			} else {
				log.Println(rtt.Echo)
			}
		}

	}
	return err, nil
}
