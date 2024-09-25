package server

import (
	"dailylife/conf"
	"dailylife/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

// func Login

func Login(schoolId, username, password string) (string, error) {
	key := username + "00000"
	encrypt_text, err := utils.Encrypt(password, key)
	if err != nil {
		panic("login")
	}
	params := url.Values{}
	params.Add("schoolId", schoolId)
	params.Add("username", username)
	params.Add("password", encrypt_text)
	//	fmt.Println("Params:", params.Encode())
	// 构建完整的 URL
	loginURL := fmt.Sprintf("%s?%s", conf.Login_url, params.Encode())
	//fmt.Print(loginURL)
	//写请求
	req, err := http.NewRequest("POST", loginURL, nil)
	if err != nil {
		panic("post")
	}
	//fmt.Println("Request URL:", req.URL.String())
	// fmt.Println("Request Body:", params.Encode())
	// req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 10; WLZ-AN00 Build/HUAWEIWLZ-AN00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/4343 MMWEBSDK/20220903 Mobile Safari/537.36 MMWEBID/4162 MicroMessenger/8.0.28.2240(0x28001C35) WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64 miniProgram/wxce6d08f781975d91")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("post")
	}
	defer resp.Body.Close()

	// 开始读body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("read")
	}
	// 把得到的结果解析一下
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		//panic("dd")
		return "", err
	}
	//	fmt.Print(result)
	// 开始断言
	if result["code"].(float64) == 0 {
		fmt.Printf("%s账号登录成功! \n", username)
		//设置cookie
		setCookie := resp.Header.Get("Set-Cookie")
		re := regexp.MustCompile(`JWSESSION=(.*?);`)
		jws := re.FindStringSubmatch(setCookie)[1]
		return jws, nil
	} else {
		fmt.Printf("%s登陆失败，请检查账号密码！\n", username)
		return "", fmt.Errorf("login failed")
	}

}

/*
"accept": "application/json, text/plain, *",
    "user-agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6

		Mobile/15E148 Safari/604.1 Edg/119.0.0.0"}

*/

/*

 'User-Agent': 'Mozilla/5.0 (Linux; Android 10; WLZ-AN00 Build/HUAWEIWLZ-AN00; wv)
  AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/4343 MMWEBSDK/20220903 Mobile Safari/537.36 MMWEBID/4162 MicroMessenger/8.0.28.2240(0x28001C35) WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64 miniProgram/wxce6d08f781975d91',
*/
