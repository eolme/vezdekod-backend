package cli

import (
	bufio "bufio"
	fmt "fmt"
	os "os"
	exec "os/exec"
	runtime "runtime"
	strconv "strconv"
	time "time"

	entities "github.com/eolme/backmemes/entities"
	shared "github.com/eolme/backmemes/shared"
)

func clear() {
	switch runtime.GOOS {
	case "linux":
	case "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		break
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func print(meme *entities.Meme) {
	fmt.Println(meme.PhotoUrl)
	fmt.Println("ID:", meme.ID)
	fmt.Printf("Likes: %d\r\n", meme.Likes)
	fmt.Println("Author:", meme.Author.FirstName, meme.Author.LastName)
	fmt.Printf("Link: https://vk.com/%s\r\n", meme.Author.ScreenName)
}

func Terminal() {
	var input string

	reader := bufio.NewScanner(os.Stdin)

	scan := func() {
		reader.Scan()
	}

	next := func() {
		scan()
		input = reader.Text()
	}

	reset := func() {
		clear()
		input = ""
	}

	for ok := true; ok; ok = input != "exit" {
		reset()

		fmt.Println("Menu:")
		fmt.Println("1. Get random meme")
		fmt.Println("2. Get meme")
		fmt.Println("3. Like meme")
		fmt.Println("4. Set prio meme")
		fmt.Println("5. Dashboard")

		next()

		switch input {
		case "1":
			for ok := true; ok; ok = input != "exit" {
				reset()

				meme, _ := shared.GetRandomMeme()
				print(meme)

				fmt.Println()
				fmt.Println("like or skip (default skip)")

				next()

				if input == "like" {
					shared.SetMemeLike(strconv.FormatUint(uint64(meme.ID), 10))
				}
			}
			reset()
			break
		case "2":
			reset()

			fmt.Println("Ger meme by id")
			fmt.Println("Enter ID:")

			next()

			meme, err := shared.GetMeme(input)
			if err != nil {
				fmt.Println("Not found or bad input!")
			} else {
				print(meme)
			}

			fmt.Println()
			fmt.Println("Press enter to continue...")

			next()
			reset()
			break
		case "3":
			reset()

			fmt.Println("Like meme")
			fmt.Println("Enter ID:")

			next()

			err := shared.SetMemeLike(input)
			if err != nil {
				fmt.Println("Not found or bad input!")
			} else {
				fmt.Println("Liked!")
			}

			fmt.Println()
			fmt.Println("Press enter to continue...")

			next()
			reset()
			break
		case "4":
			reset()

			fmt.Println("Set prio meme")
			fmt.Println("Enter ID:")

			next()
			id := input

			fmt.Println("Enter prio:")

			next()

			err := shared.SetMemePrio(id, input)
			if err != nil {
				fmt.Println("Not found or bad input!")
			} else {
				fmt.Println("Prioriterized!")
			}

			fmt.Println()
			fmt.Println("Press enter to continue...")

			next()
			reset()
			break
		case "5":
			for {
				reset()

				dashboard, err := shared.GetDashboard()

				if err != nil {
					fmt.Println("Failed")
				} else {
					fmt.Println("Place -- ID  -- Likes -- AuthorId")
					for index, meme := range *dashboard {
						fmt.Printf("%5d    %3d  %5d      %d\n", index+1, meme.ID, meme.Likes, meme.AuthorId)
					}
				}

				time.Sleep(time.Second)
			}
		default:
			reset()
			break
		}
	}
}
