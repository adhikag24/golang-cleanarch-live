package app

import (
	"fmt"
	"log"
	"os"

	"bitbucket.org/adhikag24/golang-api/app/controllers"
	"bitbucket.org/adhikag24/golang-api/app/repository"
	"bitbucket.org/adhikag24/golang-api/app/routes"
	"bitbucket.org/adhikag24/golang-api/app/usecase"
	"github.com/joho/godotenv"
)

func StartApp() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	repo := repository.NewRepo()
	usecase := usecase.NewUC(repo)
	ctrl := controllers.NewControllers(usecase)
	router := routes.NewRouter(ctrl)
	app := router.Router()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	fmt.Println(port)

	if err := app.Run(port); err != nil {
		log.Fatalf("Error when start application: %s", err)
	}

}
