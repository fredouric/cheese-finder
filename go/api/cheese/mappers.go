package cheese

import (
	"github.com/fredouric/cheese-finder-grpc/db"
	"github.com/fredouric/cheese-finder-grpc/gen/cheesev1"
)

func DBCheeseToProtobuf(dbCheese *db.Cheese) *cheesev1.Cheese {
	return &cheesev1.Cheese{
		Id:            dbCheese.ID,
		Departement:   dbCheese.Departement,
		Fromage:       dbCheese.Fromage,
		PageFrancaise: dbCheese.Pagefrancaise,
		PageAnglaise:  dbCheese.Englishpage,
		Lait:          dbCheese.Lait,
		GeoShape:      dbCheese.Geoshape,
		GeoPoint2D:    dbCheese.Geopoint2d,
	}
}
