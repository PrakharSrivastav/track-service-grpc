package service

import pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"

// Service defines a contract for the impl
type Service interface {
	CleanupAndInit() error

	Get(id string) (*pb.Track, error)
	GetAll() ([]*pb.Track, error)
	GetTracksByAlbum(albumID string) ([]*pb.Track, error)
	GetTracksByArtist(artistID string) ([]*pb.Track, error)
}
