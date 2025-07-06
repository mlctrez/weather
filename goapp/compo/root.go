package compo

import (
	"encoding/json"
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	fetch "github.com/mlctrez/wasm-fetch"
	"time"
)

var _ app.AppUpdater = (*Root)(nil)
var _ app.Mounter = (*Root)(nil)

type Root struct {
	app.Compo
	forecast *Forecast
	current  *Current
}

func (r *Root) Render() app.UI {
	app.Log("Root render")

	if r.forecast == nil {
		//return app.Div().Style("width", "100%").Body(&Chart{})
		return app.Div().Text("loading")
	}

	var rows []app.UI

	cu := r.current
	if cu != nil {
		ni := Item{
			Dt:      cu.Dt,
			Weather: cu.Weather,
		}
		ni.Main.Humidity = cu.Main.Humidity
		ni.Main.Temp = cu.Main.Temp
		ni.Main.FeelsLike = cu.Main.FeelsLike
		r.forecast.List = append([]Item{ni}, r.forecast.List...)
	}
	for _, item := range r.forecast.List {
		t := time.Unix(int64(item.Dt), 0)
		var str = ""
		var imgSrc = ""
		if len(item.Weather) > 0 {
			str = item.Weather[0].Description
			if item.Rain.H > 0 {
				str += fmt.Sprintf(" %0.1f mm/h", item.Rain.H)
			}
			imgSrc = "/web/images/" + item.Weather[0].Icon + "_t.png"
		}

		rows = append(rows, app.Tr().Body(
			app.Td().Body(app.Img().Src(imgSrc)),
			app.Td().Body(app.Text(t.Format("Mon 03:04PM"))),
			//app.Td().Body(app.Text(fmt.Sprintf("%3.0f", item.Main.Temp))),
			app.Td().Body(app.Text(fmt.Sprintf("%3.0f", item.Main.FeelsLike))),
			//app.Td().Body(app.Text(fmt.Sprintf("%5d", item.Main.Humidity))),
			app.Td().Body(app.Text(str)),
		))
	}
	return app.Div().Body(app.Table().Body(rows...))

}

func (r *Root) OnAppUpdate(ctx app.Context) {
	if app.Getenv("DEV") != "" && ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}

func (r *Root) OnMount(ctx app.Context) {
	if app.Getenv("DEV") != "" {
		ctx.Async(func() {
			timer := time.NewTicker(time.Second * 3)
			defer timer.Stop()
			for {
				select {
				case <-timer.C:
					app.TryUpdate()
				case <-ctx.Done():
					return
				}
			}
		})
	}

	var err error
	fc := &Forecast{}
	if err = bind("/api/forecast", fc); err == nil {
		r.forecast = fc
	}

	cu := &Current{}
	if err = bind("/api/current", cu); err == nil {
		r.current = cu
	}

}

func bind(url string, item any) (err error) {
	var response *fetch.Response
	if response, err = fetch.Fetch(url, &fetch.Opts{Method: fetch.MethodGet}); err != nil {
		app.Log("Failed to fetch", url, err)
		return
	}
	err = json.Unmarshal(response.Body, item)
	if err != nil {
		app.Log("Failed to unmarshal", url, err)
	}
	return err
}
