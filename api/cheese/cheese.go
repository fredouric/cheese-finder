package cheese

import (
	"github.com/fredouric/cheese-finder-grpc/pb/cheesev1"
	"golang.org/x/net/context"
)

type cheeseAPIServer struct {
	cheesev1.UnimplementedCheeseAPIServer
	// TODO
}

func New() *cheeseAPIServer {
	// TODO
	return &cheeseAPIServer{}
}

func GetOneCheese(ctx context.Context, req *cheesev1.GetOneCheeseRequest) (*cheesev1.GetOneCheeseResponse, error) {
	// TODO
	return &cheesev1.GetOneCheeseResponse{}, nil
}
