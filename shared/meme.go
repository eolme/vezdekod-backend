package shared

import (
	errors "errors"

	config "github.com/eolme/backmemes/config"
	entities "github.com/eolme/backmemes/entities"
)

func GetMeme(id string) (*entities.Meme, error) {
	var meme entities.Meme

	result := config.Database.Preload("Author").Take(&meme, id)

	if result.RowsAffected == 0 {
		return nil, errors.New("fail")
	}

	return &meme, nil
}

//
// "Показывайте его чаще других, а конкурирующие мемы, набирающие много лайков, показывайте реже."
// то бишь самые непопулярные и приоритетные показываются чаще других
//
// сортируем случайным образом относительно их лайков как веса и приоритета к самому популярному
// и берем до 10 штук (на относительно больших стоит увеличить)
// пересортируем случайным образом и берем 1 штуку
//
const MAGIC_QUERY = "SELECT * FROM(SELECT *, (1.0 * ABS(RANDOM()) / 9223372036854775807 * (`Memes`.`Likes` - (`Memes`.`Prio` * ((SELECT MAX(`Memes`.`Likes`) FROM `Memes`) + 1)))) AS `Weight` FROM `Memes` ORDER BY `Weight` LIMIT 100) RESULTS ORDER BY RANDOM() LIMIT 1"

func GetRandomMeme() (*entities.Meme, error) {
	var meme entities.Meme

	result := config.Database.Raw(MAGIC_QUERY).Scan(&meme)

	if result.RowsAffected == 0 {
		return nil, errors.New("fail")
	}

	// Preload не работает с raw запросами
	config.Database.First(&meme.Author, meme.AuthorId)

	return &meme, nil
}

func SetMemeLike(id string) error {
	result := config.Database.Model(&entities.Meme{}).Where("ID", id).Set("Likes", "Likes + 1")

	if result.RowsAffected == 0 {
		return errors.New("fail")
	}

	return nil
}

func SetMemePrio(id string, prio string) error {
	result := config.Database.Model(&entities.Meme{}).Where("ID", id).Update("Prio", prio)

	if result.RowsAffected == 0 {
		return errors.New("fail")
	}

	return nil
}

func GetDashboard() (*[]entities.Meme, error) {
	var memes []entities.Meme

	result := config.Database.Preload("Author").Order("Likes DESC").Limit(100).Find(&memes)

	if result.RowsAffected == 0 {
		return nil, errors.New("fail")
	}

	return &memes, nil
}
