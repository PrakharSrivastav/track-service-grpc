package service

import (
	"context"

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

func (f *impl) GetAll(ctx context.Context) ([]*pb.Track, error) {
	panic("not implemented")
}

func (f *impl) GetArtistByAlbum(ctx context.Context, req *pb.SimpleTrackRequest) ([]*pb.Track, error) {
	panic("not implemented")
}

func (f *impl) GetArtistByTrack(ctx context.Context, req *pb.SimpleTrackRequest) ([]*pb.Track, error) {
	panic("not implemented")
}

func (f *impl) CleanupAndInit() error {
	return f.repo.setupDatabase()
}
