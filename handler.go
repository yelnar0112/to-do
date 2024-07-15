package todo

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	todo "github.com/yelnar0112/todo-app/pkg"

	_ "github.com/yelnar0112/todo-app/docs"

	swaggerFiles "github.com/swaggo/files"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	list := router.Group("/api/todo-list")
	{
		list.POST("/tasks", h.createTask)
		list.PUT("/tasks/:id", h.updateTask)
		list.DELETE("/tasks/:id", h.deleteTask)
		list.PUT("/tasks/:id/done", h.doneTask)
		list.GET("/tasks", h.findTaskByStatus)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

var lists = []todo.TodoList{}
var id = 0

// @Summary Update an existing task
// @Description Update an existing task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body todo.TodoList true "Updated task"
// @Success 204
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Task not found"
// @Router /api/todo-list/tasks/{id} [put]
func (h *Handler) updateTask(c *gin.Context) {
	idstv := c.Param("id")

	id, err := strconv.Atoi(idstv)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask todo.TodoList
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, task := range lists {
		if task.ID == id {
			updatedTask.ID = id
			updatedTask.ActiveAt = time.Now().Format(time.RFC822)
			lists[i] = updatedTask
			c.IndentedJSON(http.StatusOK, updatedTask)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})

}

// @Summary Delete a task
// @Description Delete a task by ID
// @Tags tasks
// @Param id path int true "Task ID"
// @Success 204
// @Failure 404 {string} string "Task not found"
// @Router /api/todo-list/tasks/{id} [delete]
func (h *Handler) deleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for i, task := range lists {
		if task.ID == id {
			lists = append(lists[:i], lists[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// @Summary Mark a task as done
// @Description Mark a task as done by ID
// @Tags tasks
// @Param id path int true "Task ID"
// @Success 204
// @Failure 404 {string} string "Task not found"
// @Router /api/todo-list/tasks/{id}/done [put]
func (h *Handler) doneTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for i, task := range lists {
		if task.ID == id {
			lists[i].Completed = true
			c.JSON(http.StatusOK, lists[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})

}

// @Summary Create a new task
// @Description Create a new task with title, active date and isCompleated fields
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body todo.TodoList true "Task to create"
// @Success 201 {object} todo.TodoList
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Task already exists"
// @Router /api/todo-list/tasks [post]
func (h *Handler) createTask(c *gin.Context) {
	var newTask todo.TodoList
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = getId()
	newTask.ActiveAt = time.Now().Format(time.RFC822)
	lists = append(lists, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func (h *Handler) findTaskByStatus(c *gin.Context) {
	param := c.DefaultQuery("status", "active")

	var filteredTask []todo.TodoList
	for _, i := range lists {
		if (i.Completed == true && param == "done") || (i.Completed == false && param == "active") {
			filteredTask = append(filteredTask, i)
		}
	}

	c.IndentedJSON(http.StatusOK, filteredTask)
}

func getId() int {
	id++
	return id
}
