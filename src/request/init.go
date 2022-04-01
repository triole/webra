package request

// Req holds the main object
type Req struct {
	Settings Settings
}

// Settings holds what it says and is also imported into webra
type Settings struct {
	UserAgent   string
	TimeOut     int
	AuthEnabled bool
	AuthUser    string
	AuthPass    string
}

// Init does exactly what it says
func Init(settings Settings) (req Req) {
	return Req{
		Settings: settings,
	}
}
