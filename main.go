package main

import (
	"dailylife/server"
	"dailylife/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	school_id, err := utils.GetSchoolID("昆明理工大学")
	fmt.Println(school_id)
	if err != nil {
		panic("error")
	}
	username := os.Getenv("username")
	password := os.Getenv("password")
	tencent_key := os.Getenv("tencent_key")
	// dataBytes, err := os.ReadFile("config.yaml")
	// if err != nil {
	// 	fmt.Println("读取文件失败", err)
	// 	return

	// login_params := conf.Login_Params{}

	// err = yaml.Unmarshal(dataBytes, &login_params)
	// if err != nil {
	// 	fmt.Println("解析失败", err)
	// 	return
	// }
	// username := "1****39"
	// password := "****6"
	// tencent_key := "asdasd"
	// username := login_params.Username
	// password := login_params.Password
	// tencent_key := login_params.Tencent_Key

	location_area := "云南省昆明市呈贡区吴家营街道樱花大道"

	jws, err := server.Login(school_id, username, password)
	if err != nil {
		panic("jws")
	}
	if jws != "" {
		fmt.Println(jws)
	}
	headers := map[string]string{
		"Host":             "gw.wozaixiaoyuan.com",
		"Connection":       "keep-alive",
		"Accept":           "application/json, text/plain, */*",
		"jwsession":        jws,
		"Cookie":           fmt.Sprintf("JWSESSION=%s; WZXYSESSION=%s", jws, jws),
		"User-Agent":       "Mozilla/5.0 (Linux; Android 10; WLZ-AN00 Build/HUAWEIWLZ-AN00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/4343 MMWEBSDK/20220903 Mobile Safari/537.36 MMWEBID/4162 MicroMessenger/8.0.28.2240(0x28001C35) WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64 miniProgram/wxce6d08f781975d91",
		"Content-Type":     "application/json;charset=UTF-8",
		"X-Requested-With": "com.tencent.mm",
		"Sec-Fetch-Site":   "same-origin",
		"Sec-Fetch-Mode":   "cors",
		"Sec-Fetch-Dest":   "empty",
		"Referer":          "https://gw.wozaixiaoyuan.com/h5/mobile/health/0.3.7/health",
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7",
	}
	// signId, id, dataJson, err := GetMySignLogs(headers)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("SignId:", signId)
	// 	fmt.Println("Id:", id)
	// 	fmt.Println("DataJson:", dataJson)
	// }
	signID, id, dataJson, err := server.GetMySignLogs(headers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if signID == "" && id == "" && dataJson == nil {
		fmt.Println("用户已打过卡或无可用数据")
		return
	} else {
		fmt.Println("Sign ID:", signID)
		fmt.Println("ID:", id)
		fmt.Println("Data JSON:", dataJson)
	}

	lat, lng, err := server.GetPunchData_address(location_area, tencent_key)
	if err != nil {
		log.Fatalf("Failed to get punch data: %v", err)
	}

	//	fmt.Printf("Latitude: %f, Longitude: %f", lat, lng)
	PunchData_add, err := server.GetPunchData(lat, lng, tencent_key, dataJson)
	if err != nil {
		panic("punchData_add")
	}
	flag := server.Punch(headers, PunchData_add, username, id, signID, school_id)
	if flag {
		fmt.Printf("打卡成功\n")
	} else {
		fmt.Printf("打卡失败\n")
	}

}
