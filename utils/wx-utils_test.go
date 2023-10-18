package utils

import (
	"encoding/json"
	"errors"
	"strconv"
	"testing"
)

func TestNormalOfOpenData(t *testing.T) {
	body := "{\"errcode\":0,\"errmsg\":\"ok\",\"data_list\":[{\"cloud_id\":\"73_D-GiBNM843CJ9DFbqaNKKhAGkw2Ma2ktzNsvAEeIy4IOBFnNElgWmN39yz8\",\"json\":\"{ \\\"cloudID\\\":\\\"73_D-GiBNM843CJ9DFbqaNKKhAGkw2Ma2ktzNsvAEeIy4IOBFnNElgWmN39yz8\\\", \\\"data\\\":{\\\"nickName\\\":\\\"P.P\\\",\\\"gender\\\":0,\\\"language\\\":\\\"zh_CN\\\",\\\"city\\\":\\\"\\\",\\\"province\\\":\\\"\\\",\\\"country\\\":\\\"\\\",\\\"avatarUrl\\\":\\\"https:\\/\\/thirdwx.qlogo.cn\\/mmopen\\/vi_32\\/Nr5Myql59Ztx8wbCAeGRfbDLKvIqCIwPzicgTILplI7YlDiaaSIRic59jMFbZnB3kA3icaPNNjJsryiaFh9X9IOO1YA\\/132\\\",\\\"watermark\\\":{\\\"timestamp\\\":1697601017,\\\"appid\\\":\\\"wxd173d395b4a6002f\\\"}} }\"}]}"
	parseOpenData(t, body)
}

func TestErrorOfOpenData(t *testing.T) {
	body := "{\"errcode\":0,\"errmsg\":\"ok\",\"data_list\":[{\"cloud_id\":\"73_gIKW83fLIbFfhPeWPsrhtv-uJddAeUpr2kgLH-pVUPt8ZsoKRfVQ5K1KR7s\",\"json\":\"{ \\\"cloudID\\\":\\\"73_gIKW83fLIbFfhPeWPsrhtv-uJddAeUpr2kgLH-pVUPt8ZsoKRfVQ5K1KR7s\\\", \\\"errCode\\\":-601007, \\\"errMsg\\\":\\\"cloudID not belong to user.\\\" }\"}]}"

	parseOpenData(t, body)
}

func parseOpenData(t *testing.T, body string) {
	var response WxCloudIDResp
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		t.Error(err)
	}

	t.Log(response)

	if response.ErrCode == 0 {
		var openData = &WxCloudOpenData{}
		if err := json.Unmarshal([]byte(response.DataList[0].Json), openData); err != nil {
			t.Error(err)
		}
		if openData.ErrCode != 0 {
			t.Error(errors.New("errCode: " + strconv.Itoa(response.ErrCode) + ", errMsg: " + response.ErrMsg))
		}
		t.Log(openData)
	}
}
