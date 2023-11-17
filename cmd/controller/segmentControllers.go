package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"test_avito/internal/data/request"
	"test_avito/internal/data/response"
	"test_avito/internal/helper"
	"test_avito/internal/service"
)

type SegmentControllers struct {
	SegmentService service.SegmentService
}

func NewSegmentController(segmentService service.SegmentService) *SegmentControllers {
	return &SegmentControllers{SegmentService: segmentService}
}

func (sc *SegmentControllers) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	segmentCreateRequest := request.SegmentCreateRequest{}
	helper.ReadRequestBody(r, &segmentCreateRequest)

	sc.SegmentService.Save(r.Context(), segmentCreateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(w, webResponse)
}

func (sc *SegmentControllers) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	segmentDeleteRequest := request.SegmentDeleteRequest{}
	helper.ReadRequestBody(r, &segmentDeleteRequest)

	sc.SegmentService.Delete(r.Context(), segmentDeleteRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(w, webResponse)
}
