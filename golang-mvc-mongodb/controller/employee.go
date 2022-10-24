package controller

import (
	"fmt"
	"net/http"

	"github.com/ellipizle/golang-mvc/model"
	"github.com/ellipizle/golang-mvc/pkg/structs"

	"github.com/ellipizle/golang-mvc/service"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	appSvc *service.Service
}

func New(srv *service.Service) *Controller {
	return &Controller{srv}
}

// GetAllEmployee func gets all exists employee.
func (contrl *Controller) GetAllEmployee(c echo.Context) error {
	emps, err := contrl.appSvc.GetAllEmployee()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, emps)
	// return c.Render(http.StatusOK, "default.html", emps)

}

// GetEmployee func gets all exists employee.
func (contrl *Controller) GetEmployee(c echo.Context) error {
	id := c.Param("id")
	emps, err := contrl.appSvc.GetEmployee(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, emps)
}

// GetEmployee func gets all exists employee.
func (contrl *Controller) AddEmployee(c echo.Context) error {
	r := new(model.Employee)
	if err := c.Bind(r); err != nil {
		return err
	}
	emps, err := contrl.appSvc.AddEmployee(r)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, emps)
	// return c.Render(http.StatusOK, "add,html", emps)
}

// GetEmployee func gets all exists employee.
func (contrl *Controller) UpdateEmployee(c echo.Context) error {
	r := new(model.Employee)
	fmt.Println(r)

	id := string(c.Param("id"))
	if id == "" {
		return nil
	}
	if err := c.Bind(r); err != nil {
		return err
	}
	findEm, err := contrl.appSvc.GetEmployee(id)
	if err != nil {
		return err
	}
	structs.Merge(findEm, r)
	emps, err := contrl.appSvc.EditEmployee(findEm)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, emps)
}

// GetEmployee func gets all exists employee.
func (contrl *Controller) DeleteEmployee(c echo.Context) error {
	id := c.Param("id")
	err := contrl.appSvc.DeleteEmployee(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
