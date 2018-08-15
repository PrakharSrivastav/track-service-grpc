package model

import (
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/renstrom/shortuuid"
)

type Track struct {
	ID          string  `db:"id"`
	Name        string  `db:"name"`
	ArtistID    string  `db:"artistId"`
	AlbumID     string  `db:"albumId"`
	Description string  `db:"description"`
	Duration    float64 `db:"duration"`
}

func NewTrack() *Track {
	return &Track{
		ID:          shortuuid.New(),
		Name:        "Some Name",
		AlbumID:     shortuuid.New(),
		ArtistID:    shortuuid.New(),
		Description: "Some description",
		Duration:    float64(123.32),
	}
}

func (a *Track) ToProto() *pb.Track {
	return &pb.Track{
		Id:          a.ID,
		Name:        a.Name,
		AlbumId:     a.AlbumID,
		ArtistId:    a.ArtistID,
		Duration:    a.Duration,
		Description: a.Description,
	}
}
