package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")
	// fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test2"
	u.Email = "test2@example.com"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "yukihiro"
	// u.Email = "yukihiro@gmail.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(3)
	// fmt.Println(user)
	// user.CreateTodo("Third Todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := models.GetTodos()

	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	// user2, _ := models.GetUser(2)
	// todos, _ := user2.GetTodosByUser()
	// todos, _ := user.GetTodosByUser()

	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	t, _ := models.GetTodo(3)

	// t.Content = "Update Todo"
	// t.UpdateTodo()
	t.DeleteTodo()
}
