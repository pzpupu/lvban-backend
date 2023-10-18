package service

import (
	"lvban/utils"
	"testing"
)

func init() {
	Setup()
}

func TestRegisterByOpenId(t *testing.T) {
	data := utils.UserInfoData{NickName: "NickName", Gender: 0, AvatarUrl: "data.AvatarUrl"}
	user, err := UserService.RegisterByOpenId("test", data)
	if err != nil {
		t.Error(err)
	}
	t.Log(user)

}
