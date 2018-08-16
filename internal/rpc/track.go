package rpc

import (
	"context"
	"fmt"

	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/PrakharSrivastav/track-service-grpc/internal/service"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type TrackService struct {
	service service.Service
}

func (f *TrackService) GetAll(_ *empty.Empty, stream pb.TrackService_GetAllServer) error {
	tracks, err := f.service.GetAll()
	if err != nil {
		return err
	}
	for _, a := range tracks {
		fmt.Println("Iterating")
		if err := stream.Send(a); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

func (f *TrackService) GetTrackByAlbum(req *pb.SimpleTrackRequest, stream pb.TrackService_GetTrackByAlbumServer) error {
	tracks, err := f.service.GetTracksByAlbum(req.GetId())
	if err != nil {
		return err
	}
	for _, a := range tracks {
		fmt.Println("Iterating")
		if err := stream.Send(a); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

func (f *TrackService) GetTrackByArtist(req *pb.SimpleTrackRequest, stream pb.TrackService_GetTrackByArtistServer) error {
	tracks, err := f.service.GetTracksByArtist(req.GetId())
	if err != nil {
		return err
	}
	for _, a := range tracks {
		fmt.Println("Iterating")
		if err := stream.Send(a); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

func (f *TrackService) Get(_ context.Context, req *pb.SimpleTrackRequest) (*pb.Track, error) {
	return f.service.Get(req.GetId())
}

func (f *TrackService) Register(server *grpc.Server) {
	pb.RegisterTrackServiceServer(server, f)
}

func New(service service.Service) *TrackService {
	return &TrackService{
		service: service,
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
