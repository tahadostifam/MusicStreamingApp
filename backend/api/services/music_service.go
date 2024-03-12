package services

import (
	"errors"

	"github.com/tahadostifam/MusicStreamingApp/api/models"
	AuthRepository "github.com/tahadostifam/MusicStreamingApp/api/repository/auth"
	"github.com/tahadostifam/MusicStreamingApp/api/repository/music"
	"github.com/tahadostifam/MusicStreamingApp/pkg/minio_client"
	"github.com/tahadostifam/MusicStreamingApp/pkg/music_duration"
)

const (
	MaxFileUploadSizeMegabyte = 10
	MaxFileUploadSize         = MaxFileUploadSizeMegabyte << 20
)
const TmpDirectory = "tmp/uploads"
const MusicsBucketName = "musics-bucket"

var ErrMusicFileNotFound = errors.New("music file not found")
var ErrOnlyUserCan = errors.New("only the user himself can do this")

type MusicService struct {
	musicRepo music.Repository
	authRepo  AuthRepository.Repository
	mc        minio_client.MinioClient
}

func NewMusicService(musicRepo music.Repository, authRepo AuthRepository.Repository, mc minio_client.MinioClient) *MusicService {
	return &MusicService{musicRepo, authRepo, mc}
}

func (c MusicService) NewMusic(filePath string, artistID, title, genre string) (*models.Music, error) {
	// Upload the file
	uploadedFile, uploadErr := c.mc.UploadFile(MusicsBucketName, filePath)
	if uploadErr != nil {
		return nil, uploadErr
	}

	// Get music duration time
	musicDuration, durationErr := music_duration.MusicDuration(filePath)
	if durationErr != nil {
		return nil, durationErr
	}

	// Insert music to database
	music, createErr := c.musicRepo.Create(artistID, title, genre, uploadedFile.Name, 0, *musicDuration)
	if createErr != nil {
		return nil, createErr
	}

	return music, nil
}

func (c MusicService) GetMusic(fileName string) (savedFilePath string, err error) {
	music, findErr := c.musicRepo.FindByMusicFileName(fileName)
	if findErr != nil || music == nil {
		return "", ErrMusicFileNotFound
	}

	savedFilePath, getErr := c.mc.DownloadFile(TmpDirectory, MusicsBucketName, fileName)
	if getErr != nil {
		return "", getErr
	}

	return savedFilePath, nil
}

func (c MusicService) DeleteMusic(artistID, musicID string) error {
	user, findUserErr := c.authRepo.FindByUserID(artistID)
	if findUserErr != nil || user == nil {
		return findUserErr
	}

	music, findMusicErr := c.musicRepo.FindByMusicID(musicID)
	if findMusicErr != nil {
		return ErrMusicFileNotFound
	}

	if user.UserID != music.ArtistID {
		return ErrOnlyUserCan
	}

	err := c.mc.DeleteFile(MusicsBucketName, music.FileName)
	if err != nil {
		return err
	}

	err = c.musicRepo.Delete(music.MusicID)
	if err != nil {
		return err
	}

	return nil
}
