package dto

type NewMusicDto struct {
	Title string `json:"title" validate:"required"`
	Genre string `json:"genre" validate:"required"`

	// MusicFile is checked inside the controller
}

type DeleteMusicDto struct {
	MusicID string `json:"music_id" validate:"required"`
}
