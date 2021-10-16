package app

import (
	"bitbucket.org/adhikag24/golang-api/app/migration"
	"bitbucket.org/adhikag24/golang-api/app/routes"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func init() {
	Router = gin.Default()
}

func PrepareApp() error {
	err := migration.RunMigration()
	return err
}

func StartApp() {
	routes.Routes.Router(Router)

	if err := Router.Run(":8080"); err != nil {
		panic(err)
	}

}
