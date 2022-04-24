package seed

import (
	os "os"
	strconv "strconv"
	strings "strings"
	time "time"

	api "github.com/SevereCloud/vksdk/v2/api"
	gorm "gorm.io/gorm"

	entities "github.com/eolme/backmemes/entities"
)

func Seed(db *gorm.DB) {
	token := os.Getenv("TOKEN")

	if token == "" {
		return
	}

	vk := api.NewVK(os.Getenv("TOKEN"))

	// 430 фоток, надеюсь достаточно много
	seeds := map[string][]string{
		"-197700721": {
			"283939598",
			"281940823",
			"274262016",
		},
		"-96775998": {
			"240611621",
			"276666593",
		},
	}

	for owner, ids := range seeds {
		// Prevent too many requests
		time.Sleep(time.Second)

		albums, err := vk.PhotosGetAlbums(api.Params{
			"owner_id":  owner,
			"album_ids": strings.Join(ids[:], ","),
		})
		if err != nil {
			panic(err)
		}

		for _, album := range albums.Items {
			// Prevent too many requests
			time.Sleep(time.Second)

			photos, err := vk.PhotosGetExtended(api.Params{
				"owner_id": owner,
				"album_id": strconv.Itoa(album.ID),
				"count":    "1000",
			})
			if err != nil {
				panic(err)
			}

			for _, photo := range photos.Items {
				// Prevent too many requests
				time.Sleep(time.Second)

				users, err := vk.UsersGet(api.Params{
					"user_ids": strconv.Itoa(photo.UserID),
					"fields":   "screen_name",
				})
				if err != nil {
					panic(err)
				}

				user := users[0]

				var author entities.Author
				var meme entities.Meme

				db.Debug().Where(&entities.Author{
					UserId: user.ID,
				}).Assign(&entities.Author{
					FirstName:  user.FirstName,
					LastName:   user.LastName,
					ScreenName: user.ScreenName,
				}).FirstOrCreate(&author)

				db.Debug().Where(&entities.Meme{
					PhotoId:    photo.ID,
					PhotoOwner: photo.OwnerID,
				}).Assign(&entities.Meme{
					PhotoUrl: photo.MaxSize().URL,
					Likes:    photo.Likes.Count,
					Author:   author,
				}).FirstOrCreate(&meme)
			}
		}
	}
}
