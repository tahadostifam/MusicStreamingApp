package music

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tahadostifam/MusicStreamingApp/api/models"
	"github.com/tahadostifam/MusicStreamingApp/database"
	"gorm.io/gorm"
)

type MyTestSuite struct {
	suite.Suite
	musicRepo    *Repository
	db           *gorm.DB
	sampleMusic  models.Music
	sampleArtist string
}

func (s *MyTestSuite) SetupSuite() {
	s.db = database.CreateTestDatabase()
	s.musicRepo = NewRepository(s.db)

	s.sampleArtist = uuid.NewString()
	s.sampleMusic = models.Music{
		ArtistID: s.sampleArtist,
		Title:    "Sample Music",
		Genre:    "Sample genre",
		Plays:    0,
		FileName: "sample-file-name",
		Duration: time.Minute * 3,
	}
}

func (s *MyTestSuite) TestA_Create() {
	music, err := s.musicRepo.Create(
		s.sampleMusic.ArtistID,
		s.sampleMusic.Title,
		s.sampleMusic.Genre,
		s.sampleMusic.FileName,
		s.sampleMusic.Plays,
		s.sampleMusic.Duration,
	)
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), music)

	// Update info of sampleMusic variable
	s.sampleMusic.MusicID = music.MusicID

	assert.Equal(s.T(), s.sampleMusic.ArtistID, music.ArtistID)
	assert.Equal(s.T(), s.sampleMusic.Title, music.Title)
	assert.Equal(s.T(), s.sampleMusic.Genre, music.Genre)
	assert.Equal(s.T(), s.sampleMusic.FileName, music.FileName)
	assert.Equal(s.T(), s.sampleMusic.Plays, music.Plays)
	assert.Equal(s.T(), s.sampleMusic.Duration, music.Duration)
}

func (s *MyTestSuite) TestB_FindByMusicID() {
	music, err := s.musicRepo.FindByMusicID(s.sampleMusic.MusicID)

	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), music)
	assert.Equal(s.T(), s.sampleMusic.MusicID, music.MusicID)
}

func (s *MyTestSuite) TestC_Update() {
	// TODO
}

func (s *MyTestSuite) TestD_Delete() {
	err := s.musicRepo.Delete(s.sampleMusic.MusicID)

	assert.NoError(s.T(), err)
}

func TestMySuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}
