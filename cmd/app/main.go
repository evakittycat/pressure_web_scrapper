package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"math/rand"
	"rocketJump5G/internal/service"
)

func main() {
	//mainCtx, cancelMainCtx := context.WithCancel(context.Background())

	definingService := service.New(parserStub{})

	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// Routes
	router.GET("/service", definingService.Get)
	router.GET("/service/:category", definingService.GetThemeByCategory)
	router.GET("/service/:category/:theme", definingService.GetSiteByCategory)

	//router.POST("/service/", createUser)

	// Start server
	router.Logger.Fatal(router.Start(":1323"))
}

type parserStub struct {
}

func (ps parserStub) GetThemeByCategory(category string) []string {
	var themes []string

	themesN := rand.Intn(10) + 1
	for i := 0; i < themesN; i++ {
		themes = append(themes, gofakeit.BeerName())
	}

	return themes
}

func (ps parserStub) GetSiteByCategory(category, theme string) []string {
	var site []string

	// get themes from parser
	siteN := rand.Intn(10) + 1
	for i := 0; i < siteN; i++ {
		site = append(site, gofakeit.URL())
	}

	return site
}
