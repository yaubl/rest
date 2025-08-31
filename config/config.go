package config

import "os"

var (
	ClientID     = os.Getenv("DISCORD_CLIENT_ID")
	ClientSecret = os.Getenv("DISCORD_CLIENT_SECRET")
	RedirectURI  = os.Getenv("DISCORD_REDIRECT_URI")
	JwtSecret    = []byte(os.Getenv("JWT_SECRET"))
)
