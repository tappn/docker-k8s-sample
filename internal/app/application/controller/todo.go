package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tappn/docker-k8s-sample/internal/app/application/request"
	"github.com/tappn/docker-k8s-sample/internal/app/application/response"
	"github.com/tappn/docker-k8s-sample/internal/app/domain/model"
	"github.com/tappn/docker-k8s-sample/internal/app/domain/service"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Index(ctx echo.Context) error {
	todos, err := c.service.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "")
	}

	responses := []*response.Todo{}
	for _, v := range todos {
		resp := response.NewTodo(v)
		responses = append(responses, resp)
	}

	return ctx.JSON(http.StatusCreated, responses)
}

func (c *Controller) GetByID(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "id is int")
	}
	todo, err := c.service.GetByID(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "")
	}

	if todo == nil {
		return ctx.JSON(http.StatusNotFound, "todo is not found")
	}

	resp := response.NewTodo(todo)

	return ctx.JSON(http.StatusCreated, resp)
}

func (c *Controller) Create(ctx echo.Context) error {
	req := new(request.Todo)
	if err := ctx.Bind(req); err != nil {
		return err
	}
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	in := &model.Todo{
		Title:       req.Title,
		Description: req.Description,
		Deadline:    req.Deadline,
	}

	todo, err := c.service.Create(in)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "")
	}

	return ctx.JSON(http.StatusCreated, response.NewTodo(todo))
}

func (c *Controller) Edit(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "id is int")
	}
	req := new(request.Todo)
	if err := ctx.Bind(req); err != nil {
		return err
	}
	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	in := &model.Todo{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Deadline:    req.Deadline,
	}

	todo, err := c.service.Edit(in)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "")
	}

	return ctx.JSON(http.StatusCreated, response.NewTodo(todo))
}

func (c *Controller) Delete(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "id is int")
	}

	if err := c.service.Delete(id); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
