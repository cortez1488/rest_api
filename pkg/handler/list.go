package handler

import (
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {

	var input todoServer.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := getUserId(c)
	if err != nil {
		return
		//newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.services.CreateList(idInt, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{"list id": id})

}

type getAllListsResponse struct {
	Data []todoServer.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	idInt, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.GetAllLists(idInt)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{Data: lists})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.GetList(listId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"your list": list})
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todoServer.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateList(listId, userId, input) //мб указатель?
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, responseStatus{Status: "ok"})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err = h.services.DeleteList(listId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, responseStatus{Status: "ok"})
}
