package routes

import (
	"github.com/Megidy/TaskManagmentSystem/pkj/controllers"
	"github.com/Megidy/TaskManagmentSystem/pkj/middleware"
	"github.com/Megidy/TaskManagmentSystem/pkj/producer"
	"github.com/gin-gonic/gin"
)

var InitRoutes = func(router gin.IRouter) {
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.LogIn)
	router.GET("/task/:taskId", middleware.RequireAuth, controllers.GetSingleTask)
	router.GET("/tasks", middleware.RequireAuth, controllers.GetAllTasks)
	router.POST("/task", middleware.RequireAuth, controllers.CreateTask)
	router.DELETE("/task/:taskId", middleware.RequireAuth, controllers.DeleteTask)
	router.GET("/tasks/sort", middleware.RequireAuth, controllers.SortTasks)
	router.PUT("/task/:taskId", middleware.RequireAuth, controllers.UpdateTask)
	router.GET("/dependencies", middleware.RequireAuth, controllers.GetAllDependencies)
	router.POST("/status/:taskId", middleware.RequireAuth, producer.ChangeStatus)
	router.DELETE("/admin/user/:userId", middleware.RequireAuth, middleware.RequireAdmin, controllers.DeleteUser)
	router.GET("/admin/logs", middleware.RequireAuth, middleware.RequireAdmin, controllers.GetAllLogs)
	router.GET("/admin/logs/:userId", middleware.RequireAuth, middleware.RequireAdmin, controllers.GetUserLogs)

}
