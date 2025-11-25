package app

import (
	"net/http"
	"os"
	"synk/gateway/app/controller"
	"synk/gateway/app/util"
)

func Router(service *Service) {
	aboutController := controller.NewAbout(service.DB)

	http.HandleFunc("GET /about", aboutController.HandleAbout)

	port := os.Getenv("PORT")
	util.Log("app running on port " + port)

	err := http.ListenAndServeTLS(
		":"+port,
		"/cert/cert.pem",
		"/cert/key.pem",
		controller.Cors(http.DefaultServeMux),
	)
	if err != nil {
		util.Log("app failed on running on port " + port + ": " + err.Error())
	}
}
