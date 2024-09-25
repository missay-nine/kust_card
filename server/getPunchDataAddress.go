package server

import (
	"dailylife/conf"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func GetPunchData_address(location, tencentKey string) (float64, float64, error) {
	params := url.Values{}
	params.Add("address", location)
	params.Add("key", tencentKey)

	// 构建完整的 URL
	dizhi_url := fmt.Sprintf("%s?%s", conf.Tencent_url, params.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("GET", dizhi_url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return 0, 0, fmt.Errorf("failed to create request")
	}

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return 0, 0, fmt.Errorf("failed to create request")
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return 0, 0, fmt.Errorf("failed to read response body")
	}

	// 输出响应体内容，用于调试
	// log.Println("Response body:", string(body))

	// 解析 JSON
	var result conf.GeocodeResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return 0, 0, fmt.Errorf("failed to unmarshal JSON")
	}

	// 检查返回状态
	if result.Status != 0 {
		log.Println("Error response from API, status code:", result.Status)
		return 0, 0, fmt.Errorf("API returned an error, status: %d", result.Status)
	}

	// 返回经纬度
	return result.Result.Location.Lat, result.Result.Location.Lng, nil
}
