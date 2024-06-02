package cheese

import (
	"github.com/fredouric/cheese-finder-grpc/db"
	"github.com/fredouric/cheese-finder-grpc/pb/cheesev1"
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
