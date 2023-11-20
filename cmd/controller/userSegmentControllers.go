package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"test_avito/internal/data/request"
	"test_avito/internal/data/response"
	"test_avito/internal/helper"
	"test_avito/internal/service"
)

type UserSegmentControllers struct {
	UserSegmentService service.UserSegmentService
}

func NewUserSegmentControllers(userSegmentService service.UserSegmentService) *UserSegmentControllers {
	return &UserSegmentControllers{UserSegmentService: userSegmentService}
}

func (usc *UserSegmentControllers) AddSegments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	addSegments := request.AddSegmentsRequest{}
	helper.ReadRequestBody(r, &addSegments)

	usc.UserSegmentService.AddRemoveSegments(r.Context(), addSegments.IdUser, addSegments.SegmentsToAdd, addSegments.SegmentsToDelete)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (usc *UserSegmentControllers) GetActiveSegments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	activeSegments := request.GetActiveSegmentsRequest{}
	helper.ReadRequestBody(r, &activeSegments)

	result := usc.UserSegmentService.GetActiveUserSegments(r.Context(), activeSegments.IdUser)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(w, webResponse)
}

func (usc *UserSegmentControllers) GetReport(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	date := request.GetReport{}
	helper.ReadRequestBody(r, &date)

	filename, err := usc.UserSegmentService.GetReport(r.Context(), date)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   fmt.Sprintf("%s/%s", r.URL.String(), filename),
	}
	helper.WriteResponseBody(w, webResponse)
}
