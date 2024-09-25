package server

import (
	"bytes"
	"dailylife/conf"
	"encoding/json"
	"fmt"
	"net/http"
)

func Punch(headers map[string]string, punchData conf.PunchData, username, id, signId, schoolId string) bool {
	headers["Referer"] = "https://servicewechat.com/wxce6d08f781975d91/200/page-frame.html"
	// url := "https://gw.wozaixiaoyuan.com/sign/mobile/receive/doSignByArea"

	params := map[string]string{
		"id":       id,
		"schoolId": schoolId,
		"signId":   signId,
	}

	// 将 punchData 转换为 JSON
	punchDataBytes, err := json.Marshal(punchData)
	if err != nil {
		fmt.Println("Error marshaling punchData:", err)
		return false
	}

	// 创建请求
	req, err := http.NewRequest("POST", conf.Punch_url, bytes.NewBuffer(punchDataBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 设置 URL 参数
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	// 发送请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer res.Body.Close()

	// 解析响应
	var txt map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&txt); err != nil {
		fmt.Println("Error decoding response:", err)
		return false
	}

	// 检查响应代码
	if code, ok := txt["code"].(float64); ok && code == 0 {
		fmt.Printf("%s打卡成功！\n", username)
		//MsgSend("打卡成功！", fmt.Sprintf("%s归寝打卡成功！", username))
		return true
	} else {
		fmt.Printf("%s打卡失败！%v\n", username, txt)
		//MsgSend("打卡失败！", fmt.Sprintf("%s归寝打卡失败！%v", username, txt))
		return false
	}
}
