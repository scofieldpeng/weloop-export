package app

type config struct {
	Debug bool
}

var (
	Config = config{Debug: false}
)
