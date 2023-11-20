package router

import (
	"github.com/julienschmidt/httprouter"
	"test_avito/cmd/controller"
)

func NewRouter(userController *controller.UserControllers, segmentController *controller.SegmentControllers,
	userSegmentController *controller.UserSegmentControllers) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/user/findById/:idUser", userController.FindById)
	router.GET("/api/user/findAll", userController.FindAll)
	router.POST("/api/user/create", userController.Save)
	router.PATCH("/api/user/update/:idUser", userController.Update)
	router.DELETE("/api/user/delete/:idUser", userController.Delete)

	router.POST("/api/segment/create", segmentController.Save)
	router.DELETE("/api/segment/delete", segmentController.Delete)

	router.GET("/api/usseg/active", userSegmentController.GetActiveSegments)
	router.POST("/api/usseg/add", userSegmentController.AddSegments)
	router.POST("/api/usseg/report", userSegmentController.GetReport)

	return router
}
