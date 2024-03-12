package music

import (
	"errors"
	"time"

	"github.com/tahadostifam/MusicStreamingApp/api/models"
	"gorm.io/gorm"
)

var ErrMusicNotFound = errors.New("music not found")

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(artistID, title, genre, fileName string, plays int, duration time.Duration) (*models.Music, error) {
	obj := models.Music{
		ArtistID: artistID,
		Title:    title,
		Genre:    genre,
		Duration: duration,
		FileName: fileName,
		Plays:    plays,
	}
	result := r.db.Create(&obj)

	if result.Error != nil {
		return nil, result.Error
	}

	return &obj, nil
}

func (r *Repository) FindByMusicID(musicID string) (*models.Music, error) {
	music := &models.Music{}
	tx := r.db.Where("music_id = ?", musicID).Find(music)
	if tx.Error != nil || music.MusicID == "" {
		return nil, ErrMusicNotFound
	}

	return music, nil
}

func (r *Repository) FindByMusicFileName(fileName string) (*models.Music, error) {
	music := &models.Music{}
	tx := r.db.Where("file_name = ?", fileName).Find(&music)
	if tx.Error != nil {
		return nil, ErrMusicNotFound
	}

	return music, nil
}

func (r *Repository) Delete(musicID string) error {
	music := &models.Music{}
	tx := r.db.Where("music_id = ?", musicID).Unscoped().Delete(&music)
	if tx.Error != nil {
		return ErrMusicNotFound
	}

	return nil
}

func (r *Repository) Update(email string, newName, newPassword string) (*models.User, error) {
	// user, err := r.FindByEmail(email)
	// if err != nil {
	// 	return nil, ErrUserNotFound
	// }

	// user.Name = newName
	// user.Password = newPassword

	// tx := r.db.Save(&user)
	// if tx.Error != nil {
	// 	return nil, tx.Error
	// }

	// return user, nil
	return nil, nil
}
