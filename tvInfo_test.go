package findarrlib

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/NickBrisebois/FindarrLib/TestData"
	"testing"
)


func TestTVInfo_GetJWT(t *testing.T) {
	var conf TestData.Config
	if _, err := toml.DecodeFile("./TestData/configs.toml", &conf); err != nil {
		t.Error(err.Error())
	}

	username := conf.TvInfo.Username
	userKey := conf.TvInfo.UserKey
	apiKey := conf.TvInfo.ApiKey

	tvInfo := NewTVInfo()

	jwt, err := tvInfo.GetJWT(username, userKey, apiKey)

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(jwt)
}