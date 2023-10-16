package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// GetOpenData 根据CloudID开放数据获取服务
func GetOpenData(openid, cloudid string) (string, error) {
	reqBody := fmt.Sprintf("{\"cloudid_list\": [\"%s\"]}", cloudid)
	println("Request -> Body: " + reqBody)
	resp, err := http.Post(fmt.Sprintf("http://api.weixin.qq.com/wxa/getopendata?openid=%s", openid), "application/json", bytes.NewBufferString(reqBody))
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
