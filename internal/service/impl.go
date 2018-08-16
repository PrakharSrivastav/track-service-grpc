package service

import (
	"fmt"

	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/jmoiron/sqlx"
)

type impl struct {
	repo *repository
}

func New(db *sqlx.DB) Service {
	return &impl{repo: &repository{db: db}}
}

func (f *impl) Get(id string) (*pb.Track, error) {
	track, err := f.repo.get(id)
	if err != nil {
		return nil, err
	}
	return track.ToProto(), nil
}

func (f *impl) GetAll() ([]*pb.Track, error) {
	tracks, err := f.repo.getAll()
	if err != nil {
		return nil, err
	}

	var proto []*pb.Track
	for _, t := range tracks {
		proto = append(proto, t.ToProto())
	}
	return proto, nil
}

//artist,album
func (f *impl) GetTracksByAlbum(albumID string) ([]*pb.Track, error) {
	tracks, err := f.repo.getTracksByAlbum(albumID)
	if err != nil {
		return nil, err
	}

	var proto []*pb.Track
	for _, t := range tracks {
		proto = append(proto, t.ToProto())
	}
	return proto, nil
}

func (f *impl) GetTracksByArtist(artistID string) ([]*pb.Track, error) {
	tracks, err := f.repo.getTracksByArtist(artistID)
	if err != nil {
		return nil, err
	}

	var proto []*pb.Track
	for _, t := range tracks {
		proto = append(proto, t.ToProto())
	}
	return proto, nil
}

func (f *impl) CleanupAndInit() error {
	fmt.Println("Inside CleanupAndInit")
	return f.repo.setupDatabase()
}
