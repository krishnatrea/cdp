package bootstrap

import "gorm.io/gorm"

type Application struct {
	DB     *gorm.DB
	Config *Config
}

func App() Application {
	app := &Application{}
	app.Config = NewConfig()
	app.DB, _ = ConnectDatabase(app.Config)
	return *app
}
