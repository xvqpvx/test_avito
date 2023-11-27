package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"test_avito/internal/data/request"
	"test_avito/internal/data/response"
	"test_avito/internal/helper"
	"test_avito/internal/repos"
	"time"
)

type UserSegmentServiceImpl struct {
	UserSegmentRepository repos.UserSegmentRepo
}

func NewUserSegmentServiceImpl(userSegmentRepository repos.UserSegmentRepo) UserSegmentService {
	return &UserSegmentServiceImpl{UserSegmentRepository: userSegmentRepository}
}

func (uss *UserSegmentServiceImpl) AddRemoveSegments(ctx context.Context, request request.AddSegmentsRequest) {

	log.Print("TTL is :", request.Ttl)

	for _, segmentName := range request.SegmentsToAdd {
		idSegment, err := uss.UserSegmentRepository.GetIdSegmentByName(ctx, segmentName)
		helper.PanicIfError(err)

		if request.Ttl != "" {
			if uss.UserSegmentRepository.SegmentExistsAndInactive(ctx, request.IdUser, idSegment) {
				err = uss.UserSegmentRepository.UpdateSegmentsTTL(ctx, request.IdUser, idSegment, request.Ttl)
				log.Info().Msgf("ttl isn't empty and segment already exist, need to update")
			} else if !uss.UserSegmentRepository.SegmentExistAndActive(ctx, request.IdUser, idSegment) {
				err = uss.UserSegmentRepository.AddSegmentsTTL(ctx, request.IdUser, idSegment, request.Ttl)
				log.Info().Msgf("ttl isn't empty and segment doesn't exist, need to add")
			}
		} else {
			if uss.UserSegmentRepository.SegmentExistsAndInactive(ctx, request.IdUser, idSegment) {
				err = uss.UserSegmentRepository.UpdateSegments(ctx, request.IdUser, idSegment)
			} else if !uss.UserSegmentRepository.SegmentExistAndActive(ctx, request.IdUser, idSegment) {
				err = uss.UserSegmentRepository.AddSegments(ctx, request.IdUser, idSegment)
			}

		}
	}

	for _, segmentName := range request.SegmentsToDelete {
		idSegment, err := uss.UserSegmentRepository.GetIdSegmentByName(ctx, segmentName)
		helper.PanicIfError(err)

		err = uss.UserSegmentRepository.DeleteSegments(ctx, request.IdUser, idSegment)
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

func (uss *UserSegmentServiceImpl) startTTLChecker() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			uss.UserSegmentRepository.CheckTTL()
		}
	}
}

func (uss *UserSegmentServiceImpl) StartTTLChecker() {
	go uss.startTTLChecker()
}
