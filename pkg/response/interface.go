package response

import "context"

//go:generate mockgen --source=response_service.go --destination=./mock/mock_response_service.go
type ResponseService interface {
	CreateResponse(ctx context.Context, response CreateResponse) (string, error)
	GetResponse(ctx context.Context, uuid string) (*Response, error)
	GetAllResponses(ctx context.Context) ([]Response, error)
}

//go:generate mockgen --source=response_repository.go --destination=./mock/mock_response_repository.go
type ResponseRepository interface {
	CreateResponse(ctx context.Context, response CreateResponse) (string, error)
	GetAllResponses(ctx context.Context) ([]Response, error)
	GetResponse(ctx context.Context, uuid string) (*Response, error)
}
