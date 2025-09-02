package config

import "os"

var (
	Mode         = os.Getenv("MODE")
	ClientID     = os.Getenv("DISCORD_CLIENT_ID")
	ClientSecret = os.Getenv("DISCORD_CLIENT_SECRET")
	RedirectURI  = os.Getenv("DISCORD_REDIRECT_URI")
	Reviewers    = os.Getenv("BOT_REVIEWERS")
	JwtSecret    = []byte(os.Getenv("JWT_SECRET"))
)

func init() {
	os.Setenv("GOEXPERIMENT", "jsonv2")
}
