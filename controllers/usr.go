package usr

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kev-tsx/go-sql-server/db"
)

func GetAll(c *fiber.Ctx) error {
	res, err := db.DB.Query("SELECT * FROM user")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	return c.JSON(&fiber.Map{
		"data": "users from  backend",
	})
}
