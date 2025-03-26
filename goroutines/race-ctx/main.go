package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func looper(c *fiber.Ctx, times int) {
	ctx := c.Context()
	for i := 0; i < times; i++ {
		go func(c context.Context) {
			fmt.Println("loop at", i)
			getCtx(ctx)
		}(ctx)
	}
}

func getCtx(ctx context.Context) {
	if ctx == nil {
		fmt.Println("ctx is nil")
		return
	}

	fmt.Println("ctx valid")
}

func main() {
	cfg := fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			fmt.Println("error handler called", err)
			return nil
		},
	}
	app := fiber.New(cfg)

	app.Get("/", func(c *fiber.Ctx) error {
		looper(c, 20)

		return fmt.Errorf("Sample error")
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
