package main

import (
	"fmt"
	"log"

	"github.com/wepull/Bug-Tracker/config"
	"github.com/wepull/Bug-Tracker/controllers"
	"github.com/wepull/Bug-Tracker/models"
	"github.com/wepull/Bug-Tracker/routes"
	"github.com/wepull/Bug-Tracker/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	// Build DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	// Auto-migrate models
	err = db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.TeamMember{},
		&models.TeamInvite{},
		&models.Project{},
		&models.Bug{},
	)
	if err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}

	// Initialize services
	authService := &services.AuthService{DB: db, Config: cfg}
	userService := &services.UserService{DB: db}
	teamService := &services.TeamService{DB: db}
	projectService := &services.ProjectService{DB: db}
	bugService := &services.BugService{DB: db}

	// Initialize controllers
	authController := &controllers.AuthController{
		AuthService: authService,
		Config:      cfg,
	}
	userController := &controllers.UserController{
		UserService: userService,
	}
	teamController := &controllers.TeamController{
		TeamService: teamService,
	}
	projectController := &controllers.ProjectController{
		ProjectService: projectService,
	}
	bugController := &controllers.BugController{
		BugService: bugService,
	}

	// Setup router
	r := routes.SetupRouter(
		authController,
		userController,
		teamController,
		projectController,
		bugController,
		cfg,
	)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
