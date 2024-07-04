package cheese

import (
	"github.com/fredouric/cheese-finder-grpc/db"
	"github.com/fredouric/cheese-finder-grpc/gen/cheesev1"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

type cheeseAPIServer struct {
	cheesev1.UnimplementedCheeseAPIServer
	queries *db.Queries
}

func New(q *db.Queries) *cheeseAPIServer {
	return &cheeseAPIServer{
		queries: q,
	}
}

func (s *cheeseAPIServer) GetOneCheese(ctx context.Context, req *cheesev1.GetOneCheeseRequest) (*cheesev1.GetOneCheeseResponse, error) {
	cheese, err := s.queries.GetCheese(ctx, req.GetId())
	if err != nil {
		log.Err(err).Msg("no cheese found")
		return nil, err
	}
	return &cheesev1.GetOneCheeseResponse{
		Cheese: DBCheeseToProtobuf(&cheese),
	}, nil
}

func (s *cheeseAPIServer) GetAllCheeses(ctx context.Context, req *cheesev1.GetAllCheesesRequest) (*cheesev1.GetAllCheesesResponse, error) {
	cheeses, err := s.queries.GetAllCheeses(ctx)
	if err != nil {
		log.Err(err).Msg("failed to get cheeses")
		return nil, err
	}

	data := []*cheesev1.Cheese{}
	for _, cheese := range cheeses {
		protoCheese := DBCheeseToProtobuf(&cheese)
		data = append(data, protoCheese)
	}

	return &cheesev1.GetAllCheesesResponse{
		Cheeses: data,
	}, nil
}
