package model

import (
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/renstrom/shortuuid"
)

type Track struct {
	Id          string
	Name        string
	ArtistId    string
	AlbumId     string
	Description string
	Duration    float64
}

func NewTrack() *Track {
	return &Track{
		Id:          shortuuid.New(),
		Name:        "Some Name",
		AlbumId:     shortuuid.New(),
		ArtistId:    shortuuid.New(),
		Description: "Some description",
		Duration:    float64(123.32),
	}
}

func (a *Track) ToProto() *pb.Track {
	return &pb.Track{
		Id:          a.Id,
		Name:        a.Name,
		AlbumId:     a.AlbumId,
		ArtistId:    a.ArtistId,
		Duration:    a.Duration,
		Description: "some desc",
	}
}
