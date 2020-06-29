package findarrlib

import (
	"github.com/BurntSushi/toml"
	"github.com/NickBrisebois/FindarrLib/TestData"
	"testing"
)

func TestMovieInfo_TmdbInit(t *testing.T) {
	var conf TestData.Config
	if _, err := toml.DecodeFile("./TestData/configs.toml", &conf); err != nil {
		t.Error(err.Error())
	}

	apiKey := conf.MovieInfo.ApiKey
	useProxy := false

	movieInfo := NewMovieInfo(apiKey, nil, useProxy)

	movieInfo.TmdbInit()
}