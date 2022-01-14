package main

import (
	"bufio"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"io/fs"
	"log"
	"net/http"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
)

//go:embed assets
var assets embed.FS

func main() {
	var port int

	flag.IntVar(&port, "p", 8000, "Server port number")
	flag.Parse()

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	assetsFS, err := fs.Sub(assets, "assets")
	if err != nil {
		log.Fatal(err)
	}
	app.Use(filesystem.New(filesystem.Config{
		Root:  http.FS(assetsFS),
		Index: "index.html",
	}))

	app.Get("/list-services", func(c *fiber.Ctx) error {
		out, err := exec.Command("systemctl", "list-units", "--type=service", "--state=running", "--no-pager").Output()

		if err != nil {
			fmt.Printf("%s", err)
		}

		return c.SendString(string(out[:]))
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		var messageType int = 1
		var message []byte
		var err error

		services := []string{}
		json.Unmarshal([]byte(c.Query("services")), &services)

		args := []string{"-b"}
		for _, service := range services {
			args = append(args, "-u", service+".service")
		}
		args = append(args, "-f", "-n", "100", "-o", "json")

		cmd := exec.Command("journalctl", args...)
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			message = []byte(scanner.Text())
			if err = c.WriteMessage(messageType, message); err != nil {
				log.Println(err)
			}
		}

		cmd.Wait()
	}, websocket.Config{
		WriteBufferSize: 8192,
	}))

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
