package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zuams/simple-reminder/models"
)

func GetNotes(c echo.Context) error {
	data, err := models.GetNotes()
	if err != nil {
		return err
	}

	res := make(map[string]interface{})
	res["results"] = data
	return c.JSON(http.StatusOK, res)
}

func PostNote(c echo.Context) error {
	data, err := models.PostNote(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func PutNote(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := models.PutNote(c, id)
	if err != nil {
		ret := make(map[string]interface{})
		ret["message"] = err.Error()
		return c.JSON(http.StatusNotFound, ret)
	}

	return c.JSON(http.StatusOK, data)
}

func DeleteNote(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteNote(c, id)
	if err != nil {
		ret := make(map[string]interface{})
		ret["message"] = err.Error()
		return c.JSON(http.StatusNotFound, ret)
	}

	res := make(map[string]interface{})
	res["message"] = "success delete note"
	return c.JSON(http.StatusOK, res)
}
