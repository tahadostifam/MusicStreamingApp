package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tahadostifam/MusicStreamingApp/api/models"
	"github.com/tahadostifam/MusicStreamingApp/database"
	"gorm.io/gorm"
	"testing"
)

type MyTestSuite struct {
	suite.Suite
	authRepo   *Repository
	db         *gorm.DB
	sampleUser models.User
}

func (s *MyTestSuite) SetupSuite() {
	s.db = database.CreateTestDatabase()
	s.authRepo = NewRepository(s.db)

	// define a sample information for our user!
	s.sampleUser = models.User{
		Email:    "sample@mail.com",
		Password: "1234",
	}
}

func (s *MyTestSuite) TestA_Create() {
	user, err := s.authRepo.Create(s.sampleUser.Email, s.sampleUser.Password)

	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), user)
}

func (s *MyTestSuite) TestB_FindBy() {
	user, err := s.authRepo.FindBy(s.sampleUser.Email)

	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), user)
	assert.Equal(s.T(), user.Email, s.sampleUser.Email)
}

func (s *MyTestSuite) TestC_Update() {
	newName := "Sample"
	newPassword := "a-stronger@password"

	updatedUser, err := s.authRepo.Update(s.sampleUser.Email, newName, newPassword)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), updatedUser.Email, s.sampleUser.Email)
	assert.Equal(s.T(), updatedUser.Name, newName)
}

func (s *MyTestSuite) TestD_Delete() {
	err := s.authRepo.Delete(s.sampleUser.Email)

	assert.NoError(s.T(), err)
}

func TestMySuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}
