package response

import (
	"context"
)

type responseService struct {
	responseRepository ResponseRepository
}

func NewResponseService(p ResponseRepository) ResponseService {
	return &responseService{
		responseRepository: p,
	}
}

func (p *responseService) CreateResponse(ctx context.Context, response CreateResponse) (string, error) {
	return p.responseRepository.CreateResponse(ctx, response)
}

func (p *responseService) GetResponse(ctx context.Context, uuid string) ([]Response, error) {
	return p.responseRepository.GetResponse(ctx, uuid)
}

func (p *responseService) GetAllResponses(ctx context.Context) ([]Response, error) {
	return p.responseRepository.GetAllResponses(ctx)
}
func (p *responseService) DeleteResponses(ctx context.Context, uuid string) {
	p.responseRepository.DeleteResponses(ctx, uuid)
}
