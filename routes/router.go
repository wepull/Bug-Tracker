package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wepull/Bug-Tracker/config"
	"github.com/wepull/Bug-Tracker/controllers"
	"github.com/wepull/Bug-Tracker/middlewares"
)

func SetupRouter(
	authController *controllers.AuthController,
	userController *controllers.UserController,
	teamController *controllers.TeamController,
	projectController *controllers.ProjectController,
	bugController *controllers.BugController,
	cfg *config.Config,
) *gin.Engine {
	r := gin.Default()

	// Auth
	r.POST("/auth/register", authController.Register)
	r.POST("/auth/login", authController.Login)
	r.POST("/auth/logout", authController.Logout)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware(cfg.JWTSecret))
	{
		// Users
		auth.GET("/users", userController.ListUsers)
		auth.GET("/users/lookup/:identifier", userController.GetUser)
		auth.PUT("/users/:id", userController.UpdateUser)
		auth.DELETE("/users/:id", userController.DeleteUser)

		// Teams
		auth.POST("/teams", teamController.CreateTeam)
		auth.GET("/users/:userId/teams", teamController.ListUserTeams)
		auth.PUT("/teams/:teamId", teamController.UpdateTeam)
		auth.DELETE("/teams/:teamId", teamController.DeleteTeam)
		auth.DELETE("/teams/:teamId/members/:userId", teamController.RemoveUserFromTeam)
		// For invites, you'd add: POST /teams/:teamId/invites, GET/POST invites, etc.

		// Projects
		auth.POST("/projects", projectController.CreateProject)
		auth.GET("/projects/:projectId", projectController.GetProject)
		auth.PUT("/projects/:projectId", projectController.UpdateProject)
		auth.DELETE("/projects/:projectId", projectController.DeleteProject)
		auth.GET("/users/:userId/projects", projectController.ListUserProjects)
		auth.GET("/teams/:teamId/projects", projectController.ListTeamProjects)

		// Bugs
		auth.POST("/projects/:projectId/bugs", bugController.CreateBug)
		auth.GET("/bugs/:bugId", bugController.GetBug)
		auth.PUT("/bugs/:bugId", bugController.UpdateBug)
		auth.DELETE("/bugs/:bugId", bugController.DeleteBug)
		auth.GET("/bugs/assigned/:userId", bugController.ListBugsAssigned)
		auth.GET("/bugs/created/:userId", bugController.ListBugsCreated)
		auth.GET("/projects/:projectId/bugs", bugController.ListBugsInProject)
	}

	return r
}
