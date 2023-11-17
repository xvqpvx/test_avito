package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"test_avito/internal/data/request"
	"test_avito/internal/data/response"
	"test_avito/internal/helper"
	"test_avito/internal/model"
	"test_avito/internal/repos"
)

type SegmentServiceImpl struct {
	SegmentRepository repos.SegmentRepo
}

func NewSegmentServiceImpl(segmentRepository repos.SegmentRepo) SegmentService {
	return &SegmentServiceImpl{SegmentRepository: segmentRepository}
}

func (s *SegmentServiceImpl) FindById(ctx context.Context, idSegment int) response.SegmentResponse {
	//TODO implement me
	panic("implement me")
}

func (s *SegmentServiceImpl) FindByName(ctx context.Context, name string) response.SegmentResponse {
	segment, err := s.SegmentRepository.FindByName(ctx, name)
	helper.PanicIfError(err)
	return response.SegmentResponse{
		Name: segment.Name,
	}
}

func (s *SegmentServiceImpl) FindAll(ctx context.Context) []response.SegmentResponse {
	//TODO implement
	panic("implement")
}

func (s *SegmentServiceImpl) Save(ctx context.Context, request request.SegmentCreateRequest) {
	segment := model.Segment{
		Name: request.Name,
	}
	s.SegmentRepository.Save(ctx, segment)
	log.Info().Msg("New segment created")
}

func (s *SegmentServiceImpl) Update(ctx context.Context, request request.SegmentUpdateRequest) {
	//TODO implement
}

func (s *SegmentServiceImpl) Delete(ctx context.Context, request request.SegmentDeleteRequest) {
	segment, err := s.SegmentRepository.FindByName(ctx, request.Name)
	helper.PanicIfError(err)

	s.SegmentRepository.Delete(ctx, segment)
	log.Info().Msg("Segment deleted")
}
