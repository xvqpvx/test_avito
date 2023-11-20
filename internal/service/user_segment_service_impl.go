package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"test_avito/internal/data/request"
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

func (uss *UserSegmentServiceImpl) GetReport(ctx context.Context, request request.GetReport) (string, error) {
	reportRows, err := uss.UserSegmentRepository.GetReport(ctx, request.Year, request.Month)
	helper.PanicIfError(err)

	filename := fmt.Sprintf("Report_%s-%s.csv", request.Year, request.Month)
	file, err := os.Create(filename)
	helper.PanicIfError(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	header := []string{"id", "segment", "operation", "date"}

	err = writer.Write(header)
	if err != nil {
		return filename, err
	}

	for _, row := range reportRows {
		fields := strings.Split(row, ";")
		err := writer.Write(fields)

		if err != nil {
			return filename, err
		}
	}
	return filename, nil
}
