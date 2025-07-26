package actions

import (
	"net/http"
	"strconv"

	"todo_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

// TodosResource is the resource for the Todo model
type TodosResource struct {
	buffalo.Resource
}

// List gets all Todos. This function is mapped to the path
// GET /todos
func (v TodosResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": "Database connection not found"}))
	}

	todos := &[]models.Todo{}

	// Pagination parameters
	page := c.Param("page")
	if page == "" {
		page = "1"
	}
	perPage := c.Param("per_page")
	if perPage == "" {
		perPage = "10"
	}

	// Convert to integers
	pageNum, _ := strconv.Atoi(page)
	perPageNum, _ := strconv.Atoi(perPage)

	// Build query
	q := tx.Paginate(pageNum, perPageNum)

	// Filter by completed status
	if completed := c.Param("completed"); completed != "" {
		if completed == "true" {
			q = q.Where("completed = ?", true)
		} else if completed == "false" {
			q = q.Where("completed = ?", false)
		}
	}

	// Filter by priority
	if priority := c.Param("priority"); priority != "" {
		q = q.Where("priority = ?", priority)
	}

	// Retrieve all Todos from the DB
	if err := q.All(todos); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": err.Error()}))
	}

	// Get total count
	count, err := q.Count(&models.Todo{})
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": err.Error()}))
	}

	// Return response with pagination info
	response := map[string]interface{}{
		"success": true,
		"data":    todos,
		"meta": map[string]interface{}{
			"page":        pageNum,
			"per_page":    perPageNum,
			"total":       count,
			"total_pages": (count + perPageNum - 1) / perPageNum,
		},
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

// Show gets the data for one Todo. This function is mapped to
// the path GET /todos/{todo_id}
func (v TodosResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": "Database connection not found"}))
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	// To find the Todo the parameter todo_id is used.
	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Render(http.StatusNotFound, r.JSON(map[string]string{"error": "Todo not found"}))
	}

	response := map[string]interface{}{
		"success": true,
		"data":    todo,
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

// Create adds a Todo to the DB. This function is mapped to the
// path POST /todos
func (v TodosResource) Create(c buffalo.Context) error {
	// Allocate an empty Todo
	todo := &models.Todo{}

	// Bind todo to the html form elements
	if err := c.Bind(todo); err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": err.Error()}))
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": "Database connection not found"}))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(todo)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": err.Error()}))
	}

	if verrs.HasAny() {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"error":  "Validation failed",
			"errors": verrs.Errors,
		}))
	}

	response := map[string]interface{}{
		"success": true,
		"data":    todo,
		"message": "Todo created successfully",
	}

	return c.Render(http.StatusCreated, r.JSON(response))
}

// Update changes a Todo in the DB. This function is mapped to
// the path PUT /todos/{todo_id}
func (v TodosResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": "Database connection not found"}))
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Render(http.StatusNotFound, r.JSON(map[string]string{"error": "Todo not found"}))
	}

	// Bind Todo to the html form elements
	if err := c.Bind(todo); err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": err.Error()}))
	}

	verrs, err := tx.ValidateAndUpdate(todo)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": err.Error()}))
	}

	if verrs.HasAny() {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"error":  "Validation failed",
			"errors": verrs.Errors,
		}))
	}

	response := map[string]interface{}{
		"success": true,
		"data":    todo,
		"message": "Todo updated successfully",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

// Destroy deletes a Todo from the DB. This function is mapped
// to the path DELETE /todos/{todo_id}
func (v TodosResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": "Database connection not found"}))
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	// To find the Todo the parameter todo_id is used.
	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Render(http.StatusNotFound, r.JSON(map[string]string{"error": "Todo not found"}))
	}

	if err := tx.Destroy(todo); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": err.Error()}))
	}

	response := map[string]interface{}{
		"success": true,
		"message": "Todo deleted successfully",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

// ToggleComplete toggles the completion status of a Todo
func (v TodosResource) ToggleComplete(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": "Database connection not found"}))
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Render(http.StatusNotFound, r.JSON(map[string]string{"error": "Todo not found"}))
	}

	// Toggle the completed status
	todo.Completed = !todo.Completed

	verrs, err := tx.ValidateAndUpdate(todo)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"error": err.Error()}))
	}

	if verrs.HasAny() {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]interface{}{
			"error":  "Validation failed",
			"errors": verrs.Errors,
		}))
	}

	response := map[string]interface{}{
		"success": true,
		"data":    todo,
		"message": "Todo completion status toggled successfully",
	}

	return c.Render(http.StatusOK, r.JSON(response))
}
