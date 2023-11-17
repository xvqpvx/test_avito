package service

import (
	"context"
	"test_avito/internal/data/response"
	"test_avito/internal/helper"
	"test_avito/internal/repos"
)

type UserSegmentServiceImpl struct {
	UserSegmentRepository repos.UserSegmentRepo
}

func NewUserSegmentServiceImpl(userSegmentRepository repos.UserSegmentRepo) UserSegmentService {
	return &UserSegmentServiceImpl{UserSegmentRepository: userSegmentRepository}
}

func (uss *UserSegmentServiceImpl) AddRemoveSegments(ctx context.Context, idUser int, segmentsToAdd []string, segmentsToDelete []string) {
	for _, segmentName := range segmentsToAdd {
		idSegment, err := uss.UserSegmentRepository.GetIdSegmentByName(ctx, segmentName)
		helper.PanicIfError(err)

		err = uss.UserSegmentRepository.AddSegments(ctx, idUser, idSegment)
		helper.PanicIfError(err)
	}

	for _, segmentName := range segmentsToDelete {
		idSegment, err := uss.UserSegmentRepository.GetIdSegmentByName(ctx, segmentName)
		helper.PanicIfError(err)

		err = uss.UserSegmentRepository.DeleteSegments(ctx, idUser, idSegment)
		helper.PanicIfError(err)
	}
}

func (uss *UserSegmentServiceImpl) GetActiveUserSegments(ctx context.Context, idUser int) response.UserSegmentResponse {
	activeSegments, err := uss.UserSegmentRepository.GetActiveUserSegments(ctx, idUser)
	helper.PanicIfError(err)

	return response.UserSegmentResponse{
		IdUser:         idUser,
		ActiveSegments: activeSegments,
	}
}
