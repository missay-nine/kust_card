package conf

//import "gopkg.in/yaml.v3"

const (
	SchoolList_url   = "https://gw.wozaixiaoyuan.com/basicinfo/mobile/login/getSchoolList"
	Login_url        = "https://gw.wozaixiaoyuan.com/basicinfo/mobile/login/username"
	GetMySignLog_url = "https://gw.wozaixiaoyuan.com/sign/mobile/receive/getMySignLogs"
	Tencent_url      = "https://apis.map.qq.com/ws/geocoder/v1"
	Punch_url        = "https://gw.wozaixiaoyuan.com/sign/mobile/receive/doSignByArea"
)

type Login_Params struct {
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Tencent_Key string `yaml:"tecent_key"`
}

// School represents the structure of school data from the API
type School struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// signlog 相关结构体
type Area struct {
	ID        string `json:"id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
	Radius    int    `json:"radius"`
	Shape     int    `json:"shape"`
	DataStr   string `json:"dataStr,omitempty"`
}

type SignData struct {
	UserArea   string `json:"userArea"`
	SignID     string `json:"signId"`
	ID         string `json:"id"`
	SignStatus int    `json:"signStatus"`
	AreaList   []Area `json:"areaList"`
}

type ResponseData struct {
	Code int        `json:"code"`
	Data []SignData `json:"data"`
}

// Location struct now uses strings for lat and lng
// 通过腾讯服务api 将文字转换为经纬度相关结构
type Location struct {
	Lat float64 `json:"lat"` // 改为 float64
	Lng float64 `json:"lng"` // 改为 float64
}

type GeocodeResult struct {
	Location Location `json:"location"`
}

type GeocodeResponse struct {
	Status int           `json:"status"`
	Result GeocodeResult `json:"result"`
}

// 逆向解析
type AdInfo struct {
	Province string `json:"province"`
	City     string `json:"city"`
	Adcode   string `json:"adcode"`
	District string `json:"district"`
}

type AddressComponent struct {
	Street string `json:"street"`
}

type AddressReference struct {
	Town struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	} `json:"town"`
}

type ReverseGeocodeResponse struct {
	Status int `json:"status"`
	Result struct {
		Location         Location         `json:"location"`
		AdInfo           AdInfo           `json:"ad_info"`
		AddressComponent AddressComponent `json:"address_component"`
		AddressReference AddressReference `json:"address_reference"`
	} `json:"result"`
}

// PunchData 结构体定义了返回的数据格式
type PunchData struct {
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Nationcode string `json:"nationcode"`
	Country    string `json:"country"`
	Province   string `json:"province"`
	Citycode   string `json:"citycode"`
	City       string `json:"city"`
	Adcode     string `json:"adcode"`
	District   string `json:"district"`
	Towncode   string `json:"towncode"`
	Township   string `json:"township"`
	Streetcode string `json:"streetcode"`
	Street     string `json:"street"`
	InArea     int    `json:"inArea"`
	AreaJSON   string `json:"areaJSON"`
}
