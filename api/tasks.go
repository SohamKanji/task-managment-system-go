package api

import (
	"net/http"
	"strconv"
	"time"

	db "github.com/SohamKanji/task-management-system-go/db/sqlc"
	"github.com/SohamKanji/task-management-system-go/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) listTasks(ctx *gin.Context) {
	tasks, err := server.store.ListTasks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, tasks)
}

type createTaskRequest struct {
	Due         string `json:"due" binding:"required"`
	Status      string `json:"status" binding:"required,validstatus"`
	Priority    string `json:"priority" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) createTask(ctx *gin.Context) {
	var req createTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	due_date, err := time.Parse("2006-01-02", req.Due)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid due date")
		return
	}

	priority, ok := utils.GetPriorityValue(req.Priority)
	if !ok {
		ctx.JSON(http.StatusBadRequest, "Invalid priority")
		return
	}

	arg := db.CreateTaskParams{
		Due:         due_date,
		Status:      req.Status,
		Priority:    priority,
		Title:       req.Title,
		Description: req.Description,
	}

	tasks, err := server.store.CreateTask(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

type updateTaskRequest struct {
	Due         string `json:"due" binding:"required"`
	Status      string `json:"status" binding:"required,validstatus"`
	Priority    string `json:"priority" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) updateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	var req updateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	due_date, err := time.Parse("2006-01-02", req.Due)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid due date")
		return
	}

	priority, ok := utils.GetPriorityValue(req.Priority)
	if !ok {
		ctx.JSON(http.StatusBadRequest, "Invalid priority")
		return
	}

	arg := db.UpdateTaskParams{
		ID:          taskID,
		Due:         due_date,
		Status:      req.Status,
		Priority:    priority,
		Title:       req.Title,
		Description: req.Description,
	}

	task, err := server.store.UpdateTask(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, task)

}

func (server *Server) getTask(ctx *gin.Context) {
	id := ctx.Param("id")
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	task, err := server.store.GetTask(ctx, taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, task)
}

func (server *Server) deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	err = server.store.DeleteTask(ctx, taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, "Task deleted successfully")
}
