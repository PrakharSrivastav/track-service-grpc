package model

import (
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/renstrom/shortuuid"
)

type Artist struct {
	Id   string
	Name string
}

func NewArtist() *Artist {
	return &Artist{
		Id:   shortuuid.New(),
		Name: "Some Name",
	}
}

func (a *Artist) ToProto() *pb.Artist {
	return &pb.Artist{
		Id:          a.Id,
		Name:        a.Name,
		Description: "some desc",
	}
}
