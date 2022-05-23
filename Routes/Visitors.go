package Routes

import "go-starter/Controllers/Visitors"

func (app RouterApp) visitorsRoutes() {
	app.Gin.GET("/create-user", Visitors.CreateUser)
	app.Gin.GET("/view-user", Visitors.ViewUser)
}
