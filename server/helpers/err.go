package helpers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ThrowErr(c *fiber.Ctx, err error, status string) {
	errResp, _ := json.Marshal(Result{
		Status: status,
		Data:   make([]Book, 0),
	})
	log.Fatal(err)
	c.Status(500).Send(errResp)
}
