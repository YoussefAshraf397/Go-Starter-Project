package Routes

import "go-starter/Controllers/Visitors"

func (app RouterApp) visitorsRoutes() {
	app.Gin.POST("/register", Visitors.Register)
}
