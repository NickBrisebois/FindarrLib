package TestData

type Config struct {
	TvInfo tvInfoTestConfig
	MovieInfo movieInfoTestConfig
}

type tvInfoTestConfig struct {
	Username string
	UserKey string
	ApiKey string
}

type movieInfoTestConfig struct {
	ApiKey string
}
