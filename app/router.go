package app

import (
	"direct-debit-api/controller"
	"direct-debit-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.PaymentController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/direct-debit/payment-host-to-host", categoryController.Payment)

	router.PanicHandler = exception.ErrorHandler

	return router
}
