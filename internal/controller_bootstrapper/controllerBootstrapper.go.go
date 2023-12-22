package controllerBootstrapper

import (
	"database/sql"
	classificationController "datamining-be/internal/controller/classification"
	pingController "datamining-be/internal/controller/ping"
	classificationRepository "datamining-be/internal/repository/classification"
	classificationService "datamining-be/internal/service/classification"
	pingService "datamining-be/internal/service/ping"
	"datamining-be/pkg/config"

	"github.com/gorilla/mux"
)

func InitializeEndpoints(router *mux.Router, db *sql.DB, cfg config.Config) {
	pingService := pingService.NewPingService()

	pingController := pingController.NewPingController(router, pingService)
	pingController.InitEndpoints()

	classificationRepository := classificationRepository.NewClassificationRepository(db)
	classificationService := classificationService.NewClassificationService(classificationRepository)

	imageController := classificationController.NewClassificationController(router, cfg, classificationService)
	imageController.InitializeEndpoints()
}
