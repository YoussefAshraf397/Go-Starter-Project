package main

import (
	"fmt"
	"github.com/bykovme/gotrans"
	"go-starter/Application"
	"go-starter/Models"
	"go-starter/Routes"
)

func main() {
	// Gin framework for working with routes  (Routing , validate request , response , middleware)
	// gorm (connect database , quiers)
	// auth user
	// language
	// ------------------------------------------------------------------------------------------------- //
	_ = gotrans.InitLocales("public/languages") //  Path to the folder with localization files
	_ = gotrans.SetDefaultLocale("ar")          // Setting default locale
	fmt.Println(gotrans.T("hello_world"))

	app := Application.NewApp()
	// migrate project
	app.DB.AutoMigrate(&Models.User{})
	// close connection
	Application.CloseConnection(app)

	routerApp := Routes.RouterApp{app}
	routerApp.Routing()

	// start server app
	app.Gin.Run(":9999") //http://127.0.0.1:9999/ping

}
