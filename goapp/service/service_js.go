package service

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/mlctrez/weather/goapp/compo"
)

func Entry() {
	compo.Routes()
	app.RunWhenOnBrowser()
}
