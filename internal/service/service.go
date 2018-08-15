package service

import "context"
import pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"

type Service interface {
	Get(id string) (*pb.Track, error)
	GetAll(ctx context.Context) ([]*pb.Track, error)
	GetArtistByAlbum(ctx context.Context, req *pb.SimpleTrackRequest) ([]*pb.Track, error)
	GetArtistByTrack(ctx context.Context, req *pb.SimpleTrackRequest) ([]*pb.Track, error)
	CleanupAndInit() error
}
