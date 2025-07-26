package controller

import (
	"strconv"
	"todo-app/dto"
	"todo-app/models"
	"todo-app/service"
	"todo-app/utils"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body dto.CreateTodoRequest true "Todo object"
// @Success 201 {object} dto.APIResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /api/todos [post]
func (tc *TodoController) CreateTodo(c *gin.Context) {
	var req dto.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	todo, err := tc.todoService.CreateTodo(&req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to create todo: "+err.Error())
		return
	}

	utils.CreatedResponse(c, todo, "Todo created successfully")
}

// GetTodoByID godoc
// @Summary Get a todo by ID
// @Description Get a specific todo item by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} dto.APIResponse
// @Failure 404 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /api/todos/{id} [get]
func (tc *TodoController) GetTodoByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid todo ID")
		return
	}

	todo, err := tc.todoService.GetTodoByID(uint(id))
	if err != nil {
		if err.Error() == "todo not found" {
			utils.NotFoundResponse(c, "Todo not found")
			return
		}
		utils.InternalServerErrorResponse(c, "Failed to get todo: "+err.Error())
		return
	}

	utils.SuccessResponse(c, todo, "Todo retrieved successfully")
}

// GetAllTodos godoc
// @Summary Get all todos
// @Description Get all todo items with optional filtering and pagination
// @Tags todos
// @Accept json
// @Produce json
// @Param completed query bool false "Filter by completion status"
// @Param priority query string false "Filter by priority (LOW, MEDIUM, HIGH)"
// @Param limit query int false "Number of items per page (default: 10, max: 100)"
// @Param offset query int false "Number of items to skip (default: 0)"
// @Success 200 {object} dto.PaginatedResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /api/todos [get]
func (tc *TodoController) GetAllTodos(c *gin.Context) {
	// Parse query parameters
	completedStr := c.Query("completed")
	priorityStr := c.Query("priority")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	// Parse completed filter
	var completed *bool
	if completedStr != "" {
		completedVal, err := strconv.ParseBool(completedStr)
		if err != nil {
			utils.BadRequestResponse(c, "Invalid completed parameter")
			return
		}
		completed = &completedVal
	}

	// Parse priority filter
	var priority *models.Priority
	if priorityStr != "" {
		priorityVal := models.Priority(priorityStr)
		if priorityVal != models.LOW && priorityVal != models.MEDIUM && priorityVal != models.HIGH {
			utils.BadRequestResponse(c, "Invalid priority parameter")
			return
		}
		priority = &priorityVal
	}

	// Parse pagination parameters
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid limit parameter")
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid offset parameter")
		return
	}

	todos, total, err := tc.todoService.GetAllTodos(completed, priority, limit, offset)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get todos: "+err.Error())
		return
	}

	utils.PaginatedSuccessResponse(c, todos, total, limit, offset, "Todos retrieved successfully")
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update an existing todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body dto.UpdateTodoRequest true "Updated todo object"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 404 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /api/todos/{id} [put]
func (tc *TodoController) UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid todo ID")
		return
	}

	var req dto.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body: "+err.Error())
		return
	}

	todo, err := tc.todoService.UpdateTodo(uint(id), &req)
	if err != nil {
		if err.Error() == "todo not found" {
			utils.NotFoundResponse(c, "Todo not found")
			return
		}
		utils.InternalServerErrorResponse(c, "Failed to update todo: "+err.Error())
		return
	}

	utils.SuccessResponse(c, todo, "Todo updated successfully")
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo item by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} dto.APIResponse
// @Failure 404 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /api/todos/{id} [delete]
func (tc *TodoController) DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid todo ID")
		return
	}

	err = tc.todoService.DeleteTodo(uint(id))
	if err != nil {
		if err.Error() == "todo not found" {
			utils.NotFoundResponse(c, "Todo not found")
			return
		}
		utils.InternalServerErrorResponse(c, "Failed to delete todo: "+err.Error())
		return
	}

	utils.SuccessResponse(c, nil, "Todo deleted successfully")
}

// ToggleTodoComplete godoc
// @Summary Toggle todo completion status
// @Description Toggle the completion status of a todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} dto.APIResponse
// @Failure 404 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /api/todos/{id}/toggle [patch]
func (tc *TodoController) ToggleTodoComplete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid todo ID")
		return
	}

	todo, err := tc.todoService.ToggleTodoComplete(uint(id))
	if err != nil {
		if err.Error() == "todo not found" {
			utils.NotFoundResponse(c, "Todo not found")
			return
		}
		utils.InternalServerErrorResponse(c, "Failed to toggle todo: "+err.Error())
		return
	}

	utils.SuccessResponse(c, todo, "Todo completion status toggled successfully")
}
