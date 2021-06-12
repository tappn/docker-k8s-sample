package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tappn/docker-k8s-sample/internal/app/application/controller"
	"github.com/tappn/docker-k8s-sample/internal/app/domain/service"
	"github.com/tappn/docker-k8s-sample/internal/app/infrastructure/db"
	"github.com/tappn/docker-k8s-sample/internal/app/infrastructure/repository"
)

// Validator represents validator.
type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Validator = &Validator{validator: validator.New()}

	db := db.Connect()
	r := repository.NewRespository(db)
	s := service.NewService(r)
	c := controller.NewController(s)

	// Route
	e.GET("/", healthCheck)

	// Routes
	todo := e.Group("/todos")
	todo.GET("", c.Index)
	todo.GET("/:id", c.GetByID)
	todo.POST("", c.Create)
	todo.PUT("/:id", c.Edit)
	todo.DELETE("/:id", c.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
