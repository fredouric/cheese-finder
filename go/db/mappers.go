package db

import (
	"encoding/json"

	"github.com/fredouric/cheese-finder-grpc/dataset"
	"github.com/rs/zerolog/log"
)

func serializeLait(lait []string) string {
	laitJSON, err := json.Marshal(lait)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to serialize lait")
	}
	return string(laitJSON)
}

func serializeGeoShape(geoShape dataset.GeoShape) string {
	geoShapeJSON, err := json.Marshal(geoShape)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to serialize geoshape")
	}
	return string(geoShapeJSON)
}

func serializeGeoPoint2D(geoPoint2D dataset.GeoPoint2D) string {
	geoPoint2DJSON, err := json.Marshal(geoPoint2D)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to serialize geopoint")
	}
	return string(geoPoint2DJSON)
}

func CheeseMapper(cheese dataset.Cheese) Cheese {
	return Cheese{
		Departement:   cheese.Departement,
		Fromage:       cheese.Fromage,
		Pagefrancaise: cheese.PageFrancaise,
		Englishpage:   cheese.EnglishPage,
		Lait:          serializeLait(cheese.Lait),
		Geoshape:      serializeGeoShape(cheese.GeoShape),
		Geopoint2d:    serializeGeoPoint2D(cheese.GeoPoint2D),
	}
}
