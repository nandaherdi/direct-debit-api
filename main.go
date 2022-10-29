package main

import (
	"direct-debit-api/app"
	"direct-debit-api/controller"
	"direct-debit-api/helper"
	"direct-debit-api/middleware"
	"direct-debit-api/repository"
	"direct-debit-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	paymentRepository := repository.NewPaymentRepository()
	paymentService := service.NewPaymentService(paymentRepository, db, validate)
	paymentController := controller.NewPaymentController(paymentService)
	router := app.NewRouter(paymentController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
