package routes

import (
	"todo-app/controller"
	"todo-app/middleware"
	"todo-app/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(todoService service.TodoService) *gin.Engine {
	router := gin.New()

	// Middleware
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Controllers
	todoController := controller.NewTodoController(todoService)

	// API routes
	api := router.Group("/api")
	{
		// Todo routes
		todos := api.Group("/todos")
		{
			todos.GET("", todoController.GetAllTodos)
			todos.POST("", todoController.CreateTodo)
			todos.GET("/:id", todoController.GetTodoByID)
			todos.PUT("/:id", todoController.UpdateTodo)
			todos.DELETE("/:id", todoController.DeleteTodo)
			todos.PATCH("/:id/toggle", todoController.ToggleTodoComplete)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Todo API is running",
		})
	})

	return router
}
