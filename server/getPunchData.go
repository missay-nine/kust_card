package server

import (
	"dailylife/conf"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPunchData(lat, lng float64, tencentKey string, dataJson map[string]interface{}) (conf.PunchData, error) {
	// 构建请求 URL
	url := fmt.Sprintf("%s?location=%f,%f&key=%s", conf.Tencent_url, lat, lng, tencentKey)

	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		return conf.PunchData{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return conf.PunchData{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// 解析 JSON 响应
	var reverseGeocodeData conf.ReverseGeocodeResponse
	err = json.Unmarshal(body, &reverseGeocodeData)
	if err != nil {
		return conf.PunchData{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if reverseGeocodeData.Status != 0 {
		return conf.PunchData{}, fmt.Errorf("API returned error status: %d", reverseGeocodeData.Status)
	}

	// 解析地址数据
	locationData := reverseGeocodeData.Result
	// dataJson["polygon"], _ = json.Marshal(dataJson["polygon"]) // 将 polygon 从字符串转换为 JSON 字符串

	dataJsonStr, err := json.Marshal(dataJson)
	if err != nil {
		return conf.PunchData{}, fmt.Errorf("failed to marshal dataJson: %w", err)
	}
	// 将 float64 类型的纬度和经度转换为 string 类型
	punchData := conf.PunchData{
		Latitude:   fmt.Sprintf("%f", locationData.Location.Lat),
		Longitude:  fmt.Sprintf("%f", locationData.Location.Lng),
		Nationcode: "",
		Country:    "中国",
		Province:   locationData.AdInfo.Province,
		Citycode:   "",
		City:       locationData.AdInfo.City,
		Adcode:     locationData.AdInfo.Adcode,
		District:   locationData.AdInfo.District,
		Towncode:   locationData.AddressReference.Town.ID,
		Township:   locationData.AddressReference.Town.Title,
		Streetcode: "",
		Street:     locationData.AddressComponent.Street,
		InArea:     1,
		AreaJSON:   string(dataJsonStr), // 将 JSON 字符串转换为字符串
	}

	return punchData, nil
}
