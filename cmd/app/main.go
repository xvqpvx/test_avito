package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"test_avito/cmd/controller"
	"test_avito/internal/config"
	"test_avito/internal/helper"
	"test_avito/internal/repos"
	"test_avito/internal/router"
	"test_avito/internal/service"
)

func main() {
	fmt.Println("Start controller")

	db := config.DatabaseConnection()

	userRepo := repos.NewUserRepository(db)
	segmentRepo := repos.NewSegmentRepository(db)
	userSegmentRepo := repos.NewUserSegmentRepository(db)

	userService := service.NewUserServiceImpl(userRepo)
	segmentService := service.NewSegmentServiceImpl(segmentRepo)
	userSegmentService := service.NewUserSegmentServiceImpl(userSegmentRepo)

	userController := controller.NewUserController(userService)
	segmentController := controller.NewSegmentController(segmentService)
	userSegmentController := controller.NewUserSegmentControllers(userSegmentService)

	routes := router.NewRouter(userController, segmentController, userSegmentController)

	serv := http.Server{Addr: "localhost:8080", Handler: routes}
	err := serv.ListenAndServe()
	helper.PanicIfError(err)
}
