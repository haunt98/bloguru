package author

import (
	"context"

	authorv1 "github.com/haunt98/bloguru/gen/author/v1"
)

type server struct {
	authorv1.UnimplementedServiceServer
}

func NewServer() authorv1.ServiceServer {
	return &server{}
}

func (s *server) Get(ctx context.Context, req *authorv1.GetRequest) (*authorv1.GetResponse, error) {
	if req.Id == "1" {
		return &authorv1.GetResponse{
			Data: &authorv1.GetResponseData{
				Author: &authorv1.Author{
					Id:    "1",
					Name:  "1",
					Email: "1",
				},
			},
		}, nil
	}

	return nil, nil
}
