package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"template/lib/log"
	"template/lib/response"
	"template/package/user/model"
	"template/package/user/usecase"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *userHandler {
	return &userHandler{uc}
}

func (h *userHandler) Mount(g *echo.Group) {
	g.GET("/", h.GetUsers)
	g.POST("/add", h.AddUser)
	g.PUT("/edit", h.EditUser)
}

func (h *userHandler) GetUsers(e echo.Context) error {
	log.Info("Get User ...")
	data, err := h.uc.GetUsers()
	if err != nil {
		return response.ToJson(e).InternalServerError(err.Error())
	}

	return response.ToJson(e).OK(data, "Get Users Success")
}

func (h *userHandler) AddUser(e echo.Context) error {
	log.Info("Add User ...")
	body, err := io.ReadAll(e.Request().Body)
	if err != nil {
		return response.ToJson(e).BadRequest("Failed get body")
	}

	user := model.User{}
	if err = json.Unmarshal(body, &user); err != nil {
		return response.ToJson(e).BadRequest("Failed unmarshal")
	}

	if err = h.uc.AddUser(user); err != nil {
		switch {
		case strings.Contains(err.Error(), "(SQLSTATE 23502)"):
			column := strings.Split(err.Error(), "\"")[1]
			return response.ToJson(e).BadRequest(fmt.Sprintf("Can't Null value on column \"%v\"", column))
		case strings.Contains(err.Error(), "(SQLSTATE 23505)"):
			column := strings.Split(err.Error(), "\"")[1]
			return response.ToJson(e).BadRequest(fmt.Sprintf("Duplicate value on column \"%v\"", column))
		default:
			return response.ToJson(e).InternalServerError(err.Error())
		}

	}

	return response.ToJson(e).OK(nil, "Add User Success")
}

func (h *userHandler) EditUser(e echo.Context) error {
	log.Info("EditUser ...")
	user := model.User{}

	if err := json.NewDecoder(e.Request().Body).Decode(&user); err != nil {
		return response.ToJson(e).BadRequest("Failed Decode: " + err.Error())
	}

	if err := h.uc.EditUser(user); err != nil {
		switch {
		case strings.Contains(err.Error(), "(SQLSTATE 23502)"):
			column := strings.Split(err.Error(), "\"")[1]
			return response.ToJson(e).BadRequest(fmt.Sprintf("Can't Null value on column \"%v\"", column))
		case strings.Contains(err.Error(), "(SQLSTATE 23505)"):
			column := strings.Split(err.Error(), "\"")[1]
			return response.ToJson(e).BadRequest(fmt.Sprintf("Duplicate value on column \"%v\"", column))
		default:
			return response.ToJson(e).InternalServerError(err.Error())
		}
	}

	return response.ToJson(e).OK(nil, "Update User Success")
}
