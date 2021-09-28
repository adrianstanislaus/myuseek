package playlists

import (
	"myuseek/business/playlists"
	"time"

	"gorm.io/gorm"
)

type Playlist struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	Description string
	Creator_id  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (playlist *Playlist) ToDomain() playlists.Domain {
	return playlists.Domain{
		Id:          playlist.Id,
		Name:        playlist.Name,
		Description: playlist.Description,
		Creator_id:  playlist.Creator_id,
		CreatedAt:   playlist.CreatedAt,
		UpdatedAt:   playlist.UpdatedAt,
	}
}

func FromDomain(domain playlists.Domain) Playlist {
	return Playlist{
		Id:          domain.Id,
		Name:        domain.Name,
		Description: domain.Description,
		Creator_id:  domain.Creator_id,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt}
}

func ToListDomain(data []Playlist) (result []playlists.Domain) {
	result = []playlists.Domain{}
	for _, playlist := range data {
		result = append(result, playlist.ToDomain())
	}
	return result
}
