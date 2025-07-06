package compo

type Forecast struct {
	Cnt  int    `json:"cnt"`
	List []Item `json:"list"`
}

type Item struct {
	Dt   int `json:"dt"`
	Main struct {
		Temp        float64 `json:"temp"`
		TempMin     float64 `json:"temp_min"`
		TempMax     float64 `json:"temp_max"`
		FeelsLike   float64 `json:"feels_like"`
		Pressure    int     `json:"pressure"`
		SeaLevel    int     `json:"sea_level"`
		GroundLevel int     `json:"grnd_level"`
		Humidity    int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Rain struct {
		H float64 `json:"3h,omitempty"`
	} `json:"rain"`
}

type Current struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Message int    `json:"message"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Base    string `json:"base"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		FeelsLike float64 `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
	} `json:"rain"`
	Snow struct {
	} `json:"snow"`
	Dt       int    `json:"dt"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
	Timezone int    `json:"timezone"`
	Unit     string `json:"Unit"`
	Lang     string `json:"Lang"`
	Key      string `json:"Key"`
}
