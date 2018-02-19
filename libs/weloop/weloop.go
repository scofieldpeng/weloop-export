package weloop

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"time"

	"io/ioutil"

	"errors"

	"strconv"

	"github.com/pquerna/ffjson/ffjson"
	"go.zhuzi.me/config"
)

type path string

const (
	host = "https://v3.weloop.cn"

	pathLogin        path = "/weloop4/loginAction!userLogin.do"
	pathTokenIsValid path = "/weloop4/userAction!tokenIsValid.do"
	pathDailyDetail  path = "/weloop4/synchrolabel!*downloadSourceData.do"

	appVersion = "5.2.1"
	appkey     = "C6B00F79C39640D9"
	clientType = "2"
)

// 进行请求
func request(path path, userId int, requestData url.Values, response interface{}) (err error) {
	var (
		req            *http.Request
		resp           *http.Response
		respBytes      []byte
		commonResponse CommonResponse
		c              = http.Client{
			Timeout: time.Second * 30,
		}
		yfheader = map[string]interface{}{}
	)

	if req, err = http.NewRequest("POST", host+string(path), strings.NewReader(requestData.Encode())); err != nil {
		return
	}
	req.Header.Add("User-Agent", fmt.Sprintf("WeLoop/%s (iPhone; iOS 11.2.5; Scale/2.00)", appVersion))
	req.Header.Add("accept-language", "en-CN;q=1, zh-Hans-CN;q=0.9")
	req.Header.Add("yfflag", "")
	req.Header.Add("authority", "v3.weloop.cn")

	yfheader["releaseType"] = 1
	yfheader["systemType"] = 2
	yfheader["deviceId"] = config.String(config.Data("app").Get("weloop", "deviceId"))
	yfheader["appVersion"] = config.String(config.Data("app").Get("weloop", "appVersion"))
	yfheader["userId"] = userId
	if b, err := ffjson.Marshal(yfheader); err != nil {
		return err
	} else {
		req.Header.Add("yfheader", string(b))
	}

	if resp, err = c.Do(req); err != nil {
		return
	}

	if respBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	defer resp.Body.Close()

	// 先看下是否请求成功
	if err = ffjson.Unmarshal(respBytes, &commonResponse); err == nil && commonResponse.Message == "ok" && commonResponse.Result == "0000" {
		err = ffjson.Unmarshal(respBytes, &response)
	} else {
		err = errors.New(string(respBytes))
	}

	return
}

type LoginParams struct {
	Account     string
	DeviceToken string
	Password    string
}

// 登录信息
func Login(params LoginParams) (responseData LoginResponse, err error) {
	requestData := url.Values{}
	requestData.Add("account", params.Account)
	requestData.Add("appkey", appkey)
	requestData.Add("clientType", clientType)
	requestData.Add("deviceToken", params.DeviceToken)
	requestData.Add("pwd", params.Password)

	err = request(pathLogin, 0, requestData, &responseData)
	return
}

type TokenValidParams struct {
	Token string
}

// token是否有效
func IsTokenValid(params TokenValidParams) (valid bool, err error) {
	requestData := url.Values{}
	requestData.Add("accessToken", params.Token)
	responseData := TokenValidResponse{}

	err = request(pathTokenIsValid, 0, requestData, &responseData)
	if err != nil {
		return
	}
	if responseData.Valid == TokenValid {
		valid = true
	}

	return
}

// 每日详情参数
type DailyDetailParams struct {
	AccessToken string
	UserId      int
	DayCount    int
	EndTime     int64
}

// 每日详情
func DailyDetail(params DailyDetailParams) (responseData DailySourceResponse, err error) {
	requestData := url.Values{}
	requestData.Add("accessToken", params.AccessToken)
	requestData.Add("dayCount", strconv.Itoa(params.DayCount))
	requestData.Add("happenDate", strconv.FormatInt(params.EndTime, 10))

	err = request(pathDailyDetail, params.UserId, requestData, &responseData)
	return
}
