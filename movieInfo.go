package findarrlib

import "github.com/ryanbradynd05/go-tmdb"

type MovieInfo struct {
	tmdbConfig tmdb.Config
	tmdbAPI *tmdb.TMDb
}

func NewMovieInfo(apiKey string, proxies []tmdb.Proxy, useProxy bool) *MovieInfo {
	return &MovieInfo {
		tmdbConfig: tmdb.Config{
			APIKey: apiKey,
			Proxies: proxies,
			UseProxy: useProxy,
		},
	}
}

func (mi *MovieInfo) TmdbInit() {
	mi.tmdbAPI = tmdb.Init(mi.tmdbConfig)
}
