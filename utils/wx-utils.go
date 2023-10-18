package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// WxCloudIDResp 微信CloudID调用通用响应
type WxCloudIDResp struct {
	// ErrCode 响应码
	ErrCode int `json:"errcode"`
	// ErrCode 响应码说明
	ErrMsg string `json:"errmsg"`
	// DataList 数据列表响应
	DataList []struct {
		// CloudId
		CloudID int `json:"cloud_id"`
		// Json 数据
		Json string
	}
}

type WxCloudID struct {
	CloudID string `json:"cloudID"`
}

type UserInfoData struct {
	NickName  string `json:"nickName"`
	Gender    int8   `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	Watermark struct {
		Timestamp int    `json:"timestamp"`
		AppId     string `json:"appid"`
	}
}

// WxCloudOpenData 微信云用户开放数据结构
type WxCloudOpenData struct {
	WxCloudID
	Data UserInfoData `json:"data"`
}

// GetOpenData 根据CloudID开放数据
func GetOpenData(openid, cloudid string) (*WxCloudOpenData, error) {
	reqBody := fmt.Sprintf("{\"cloudid_list\": [\"%s\"]}", cloudid)
	resp, err := http.Post(fmt.Sprintf("http://api.weixin.qq.com/wxa/getopendata?openid=%s", openid), "application/json", bytes.NewBufferString(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	println("====== Body ==========")
	println(string(body))

	var response WxCloudIDResp
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if response.ErrCode == 0 {
		var openData = &WxCloudOpenData{}
		if err := json.Unmarshal([]byte(response.DataList[0].Json), openData); err != nil {
			return nil, err
		}
		return openData, nil
	}

	return nil, errors.New("errCode: " + strconv.Itoa(response.ErrCode) + ", errMsg: " + response.ErrMsg)
}
