package server

import (
	"dailylife/conf"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetMySignLogs(headers map[string]string) (string, string, map[string]interface{}, error) {
	// url := "https://gw.wozaixiaoyuan.com/sign/mobile/receive/getMySignLogs"
	client := &http.Client{}
	req, err := http.NewRequest("GET", conf.GetMySignLog_url, nil)
	if err != nil {
		return "", "", nil, err
	}

	// 添加请求头
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// 设置查询参数
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("size", "10")
	req.URL.RawQuery = q.Encode()

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return "", "", nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", nil, err
	}

	// 解析 JSON 响应
	var response conf.ResponseData
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", "", nil, err
	}

	if len(response.Data) == 0 {
		return "", "", nil, fmt.Errorf("no sign data found")
	}

	data := response.Data[0]

	// 检查签到状态
	if data.SignStatus != 1 {
		fmt.Println("用户已打过卡！")
		return "", "", nil, nil
	}

	// 查找用户区域
	for _, area := range data.AreaList {
		if data.UserArea == area.Name {
			// 构建 dataJson
			dataStr := area.DataStr
			fmt.Println(area.DataStr)
			// fmt.Printf("Type of dataStr: %s\n", reflect.TypeOf(dataStr))
			fmt.Println("----------------------")
			if dataStr == "" {
				dataStr = fmt.Sprintf(`[{"longitude": %s, "latitude": %s}]`, area.Longitude, area.Latitude)
			}
			dataJson := map[string]interface{}{
				"type":    1,
				"polygon": dataStr,
				"id":      area.ID,
				"name":    area.Name,
			}
			// fmt.Println("datajson如下")
			// fmt.Println(dataJson)
			// fmt.Println("---------------------------------------")
			return data.SignID, data.ID, dataJson, nil
		}
	}

	return "", "", nil, nil
}
