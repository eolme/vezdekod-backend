package api

import (
	bufio "bufio"
	"encoding/json"
	fmt "fmt"
	time "time"

	fiber "github.com/gofiber/fiber/v2"
	fasthttp "github.com/valyala/fasthttp"

	shared "github.com/eolme/backmemes/shared"
)

func GetMeme(ctx *fiber.Ctx) error {
	meme, err := shared.GetMeme(ctx.Params("id"))
	if err != nil {
		return ctx.SendStatus(404)
	}

	return ctx.Status(200).JSON(meme)
}

func GetRandomMeme(ctx *fiber.Ctx) error {
	meme, err := shared.GetRandomMeme()
	if err != nil {
		return ctx.SendStatus(428)
	}

	return ctx.Status(200).JSON(meme)
}

func SetMemeLike(ctx *fiber.Ctx) error {
	err := shared.SetMemeLike(ctx.Params("id"))
	if err != nil {
		return ctx.SendStatus(400)
	}

	return ctx.SendStatus(200)
}

func SetMemePrio(ctx *fiber.Ctx) error {
	err := shared.SetMemePrio(ctx.Params("id"), ctx.Params("prio"))
	if err != nil {
		return ctx.SendStatus(400)
	}

	return ctx.SendStatus(200)
}

func GetDashboard(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/event-stream")
	ctx.Set("Cache-Control", "no-store, no-transform")
	ctx.Set("Transfer-Encoding", "identity")
	ctx.Set("Connection", "keep-alive")
	ctx.Set("Keep-Alive", "timeout=0")

	ctx.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(writer *bufio.Writer) {
		fmt.Fprintf(writer, ":ok\n\n")

		for {
			dashboard, err := shared.GetDashboard()
			response, err := json.Marshal(dashboard)
			if err != nil {
				fmt.Fprintf(writer, "event: error\n")
			} else {
				fmt.Fprintf(writer, "data: %s\n\n", response)
			}

			writer.Flush()
			time.Sleep(time.Second)
		}
	}))

	return nil
}
