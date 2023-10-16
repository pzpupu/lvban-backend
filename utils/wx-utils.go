package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// GetOpenData 根据CloudID开放数据获取服务
func GetOpenData(appid, accessToken, openid, cloudid string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/getopendata?from_appid=%s&openid=%s&cloudbase_access_token=%s", appid, openid, accessToken)
	println("Request -> URL: " + url)
	reqBody := fmt.Sprintf("{\"cloudid_list\": [\"%s\"]}", cloudid)
	println("Request -> Body: " + reqBody)
	resp, err := http.Post(url, "application/json", bytes.NewBufferString(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
