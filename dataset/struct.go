package dataset

type Cheese struct {
	Departement   string     `json:"departement"`
	Fromage       string     `json:"fromage"`
	PageFrancaise string     `json:"page_francaise"`
	EnglishPage   string     `json:"english_page"`
	Lait          []string   `json:"lait"`
	GeoShape      GeoShape   `json:"geo_shape"`
	GeoPoint2D    GeoPoint2D `json:"geo_point_2d"`
}

type GeoShape struct {
	Geometry Geometry `json:"geometry"`
	Type     string   `json:"type"`
}

type Geometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type GeoPoint2D struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
