package main

import (
	"WeatherfForecast/configs"
	"WeatherfForecast/middleware"
	"WeatherfForecast/models"
	"WeatherfForecast/pkg/controller"
	"WeatherfForecast/pkg/repository"
	"WeatherfForecast/pkg/service"
	"WeatherfForecast/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	config, err := configs.Load()

	if err != nil {
		logrus.Fatalf("failed to load configs: %v", err)
	}

	db, err := initDb(config.DatabaseURL)

	if err != nil {
		logrus.Fatalf("Failed to connect to db: %v", err)
	}

	userRepository := repository.NewUserRepository(db)
	jwtUtil := util.NewJwtUtil(config.JwtSecret)
	authFilter := middleware.NewAuthMiddleware(jwtUtil)
	authService := service.NewAuthService(userRepository, jwtUtil)
	authController := controller.NewAuthController(authService)

	weatherClient := service.NewWeatherClient(config.WeatherApiKey)
	forecastService := service.NewForecastService(userRepository, weatherClient)
	forecastController := controller.NewForecastController(forecastService)

	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	authController.RegisterRoutes(router)
	forecastController.RegisterRoutes(router, authFilter)
	userController.RegisterRoutes(router, authFilter)

	if err := router.Run(config.ServerAddress); err != nil {
		logrus.Fatalf("failed to run server: %v", err)
	}
}

func initDb(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}, &models.Forecast{}); err != nil {
		return nil, err
	}

	return db, nil
}
