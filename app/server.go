package app

import (
	"bitbucket.org/adhikag24/golang-api/app/controllers"
	"bitbucket.org/adhikag24/golang-api/app/repository"
	"bitbucket.org/adhikag24/golang-api/app/routes"
	"bitbucket.org/adhikag24/golang-api/app/usecase"
)

func StartApp() {
	repo := repository.NewRepo()
	usecase := usecase.NewUC(repo)
	ctrl := controllers.NewControllers(usecase)
	router := routes.NewRouter(ctrl)
	app := router.Router()

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}

}
