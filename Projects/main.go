package main
import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)	
	type Todo struct {
		ID int `json:"id"`
		Completed bool `json:"completed"`
		Body string `json:"body"`
	}
func main(){
	fmt.Println("Hello, World!")
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello, World",
		})
	})

 // Create a Todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
	todo := &Todo{} // {id: 0, completed: false, body: ""}
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.Body == "" {	
		return c.Status(400).JSON(fiber.Map{
			"error": "Body is required",
		})
	}
	n := len(todos)
	if n==0 {todo.ID = 1}else {todo.ID = todos[n-1].ID + 1 ;}
	todos = append(todos, *todo)
	return c.Status(201).JSON(todo)

})
	// Update a Todo
	app.Patch("/api/todos/:id",func(c *fiber.Ctx) error {
	 id := c.Params("id")
	 for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos[i].Completed = true
			return c.Status(200).JSON(todos[i])
		}
	}
	return c.Status(404).JSON(fiber.Map{
			"error": "Todo not found",})
	})
	// Delete a Todo 
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error{
		id := c.Params("id")
		for i,todo :=range todos {
			if fmt.Sprint(todo.ID) == id {
				DeletedTodo := todos[i];
				todos = append(todos[:i],todos[i+1:]...)
				return c.Status(200).JSON(DeletedTodo)
			}		
		}
		return c.Status(404).JSON(fiber.Map{
			"error" : "Todo not found",
		})
	})

	// get all Todo 
	app.Get("/api/todos",func(c *fiber.Ctx) error{
		return c.Status(200).JSON(todos)
	})


	log.Fatal(app.Listen("localhost:"+PORT))
}