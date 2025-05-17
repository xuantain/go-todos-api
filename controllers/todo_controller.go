package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "go-todos-api/docs"
	"go-todos-api/models"
	"go-todos-api/pkg/helpers"
	"go-todos-api/pkg/types"
	"go-todos-api/repositories"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	TodoRepo repositories.TodoRepository
	UserRepo repositories.UserRepository
}

func (ctl TodoController) Index(c *gin.Context) {
	message, err := c.Cookie("message")
	if err != nil {
		c.HTML(http.StatusOK, "Index", gin.H{
			"title":   "Login",
			"message": message,
		})
		c.SetCookie("message", "", -1, "/", c.Request.Host, false, true) // Clear the cookie
	} else {
		c.HTML(http.StatusOK, "Index", gin.H{
			"title": "Home Page",
		})
	}
}

func (ctl TodoController) Welcome(c *gin.Context) {

	user := helpers.Auth(c)
	if user.Username == "" {
		c.Redirect(http.StatusFound, "/page/login")
	}
	username := user.Username
	// username, err := c.Cookie("username")
	// if err != nil {
	// 	c.SetCookie("message", "Unauthentication", 180, "/", c.Request.Host, false, true)
	// 	c.Redirect(http.StatusFound, "/page/")
	// }

	msg := fmt.Sprintf("Welcome back %v!", username)
	c.HTML(http.StatusOK, "Welcome", gin.H{
		"welcome":   msg,
		"linkTitle": "Click here to access API docs",
		"link":      "/api/docs/index.html",
	})
}

func (ctl TodoController) ListTodos(c *gin.Context) {

	user := helpers.Auth(c)
	userId := user.Id
	// userId, err := c.Cookie("userId")
	// if err != nil {
	// 	c.SetCookie("message", "Unauthentication", 180, "/", c.Request.Host, false, true)
	// 	c.Redirect(http.StatusFound, "/page/")
	// }

	todoList, err := ctl.TodoRepo.ListByUserId(c.Request.Context(), userId, 0, 10)
	if err != nil {
		// todo: db connection error or ?
	}

	message, err := c.Cookie("message")

	if err != nil {
		c.HTML(http.StatusOK, "TodoList", gin.H{
			"title":   "Todos Page",
			"todos":   todoList,
			"message": message,
		})
		c.SetCookie("message", "", -1, "/", c.Request.Host, false, true) // Clear the cookie
	} else {
		c.HTML(http.StatusOK, "TodoList", gin.H{
			"title": "Todos Page",
			"todos": todoList,
		})
	}
}

func (ctl TodoController) CreateTodo(c *gin.Context) {

	newTodo := models.Todo{}

	if c.Request.Method == http.MethodPost {

		if err := c.ShouldBind(&newTodo); err != nil {
			helpers.Flash(c, "error", err.Error())
			return
		}

		if err := ctl.TodoRepo.Create(c.Request.Context(), &newTodo); err != nil {
			helpers.Flash(c, "error", err.Error())
			return
		}

		c.Redirect(http.StatusFound, "/page/todos")
		return
	}

	newTodo.TargetDate = types.DateType{Time: time.Now()}
	c.HTML(http.StatusOK, "Todo", gin.H{
		"title": "Todo",
		"todo":  newTodo,
	})
	c.Abort()
}

func (ctl TodoController) UpdateTodo(c *gin.Context) {

	u64, err := strconv.ParseUint(c.Param("todoId"), 10, 32)
	if err != nil {
		helpers.Flash(c, "error", err.Error())
		c.Abort()
	}
	todoId := uint(u64)

	todo, err := ctl.TodoRepo.FindByID(c, todoId)
	if err != nil {
		helpers.Flash(c, "error", err.Error())
		c.Abort()
	}

	if c.Request.Method == http.MethodPost {

		updatedTodo := models.Todo{}

		if err := c.ShouldBind(&updatedTodo); err != nil {
			helpers.Flash(c, "error", err.Error())
			c.Abort()
		}

		// Todo: The field should be serialized by c.ShouldBind instead of manually updating like following
		targetDate, ok := c.GetPostForm("targetDate")
		if ok {
			fmt.Println("Update ############################## " + targetDate)
		}

		parsedTime, err := time.Parse(types.DateFormat, targetDate)
		if err != nil {
			fmt.Println("Error parsing time string:", err)
			return
		}

		updatedTodo.ID = todoId
		updatedTodo.UserID = todo.UserID
		updatedTodo.TargetDate = types.DateType{Time: parsedTime}

		if err := ctl.TodoRepo.Update(c.Request.Context(), &updatedTodo); err != nil {
			helpers.Flash(c, "error", err.Error())
			c.Abort()
		}

		c.Redirect(http.StatusFound, "/page/todos")
		return
	}

	c.HTML(http.StatusOK, "Todo", gin.H{
		"title": "Todo",
		"todo":  todo,
	})
	c.Abort()
}

func (ctl TodoController) DeleteTodo(c *gin.Context) {

	todoId := ""
	if err := c.Request.ParseForm(); err == nil {
		todoId = c.Request.PostForm.Get("id")
	}

	if todoId != "" {

		u64, err := strconv.ParseUint(todoId, 10, 32)
		if err != nil {
			message := "Missing Todo Id"
			c.SetCookie("message", message, 180, "/", c.Request.Host, false, true)
			c.Redirect(http.StatusFound, "/page/todos")
		}
		todoId := uint(u64)

		if err := ctl.TodoRepo.Delete(c.Request.Context(), todoId); err != nil {
			message := fmt.Sprintf("Cannot delete Todo {%d}", todoId)
			c.SetCookie("message", message, 180, "/", c.Request.Host, false, true)
			c.Redirect(http.StatusFound, "/page/todos")
			return
		}

		message := fmt.Sprintf("Deleted todo id:{%d}", todoId)
		c.SetCookie("message", message, 180, "/", c.Request.Host, false, true)
		// Note: To redirect > Must use http.StatusFound; Or http.StatusTemporaryRedirect with method
		c.Redirect(http.StatusFound, "/page/todos")
		return
	}

	message := "Missing Todo Id"
	c.SetCookie("message", message, 180, "/", c.Request.Host, false, true)
	c.Redirect(http.StatusFound, "/page/todos")
	c.Abort()
}
