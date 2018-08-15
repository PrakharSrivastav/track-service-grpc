package rpc

import (
	"context"
	"fmt"

	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/PrakharSrivastav/track-service-grpc/internal/model"
	"github.com/PrakharSrivastav/track-service-grpc/internal/service"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type TrackService struct {
	service service.Service
}

func (f *TrackService) GetAll(_ *empty.Empty, stream pb.TrackService_GetAllServer) error {
	fmt.Println("Inside the function")
	var tracks []*model.Track
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	fmt.Println("Added to the list")
	for _, a := range tracks {
		fmt.Println("Iterating")
		if err := stream.Send(a.ToProto()); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

func (f *TrackService) GetTrackByAlbum(req *pb.SimpleTrackRequest, stream pb.TrackService_GetTrackByAlbumServer) error {
	fmt.Println("Inside the function")
	var tracks []*model.Track
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	fmt.Println("Added to the list")
	for _, a := range tracks {
		fmt.Println("Iterating")
		if err := stream.Send(a.ToProto()); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

func (f *TrackService) GetTrackByArtist(req *pb.SimpleTrackRequest, stream pb.TrackService_GetTrackByArtistServer) error {
	fmt.Println("Inside the function")
	var tracks []*model.Track
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	tracks = append(tracks, model.NewTrack())
	fmt.Println("Added to the list")
	for _, a := range tracks {
		fmt.Println("Iterating")
		if err := stream.Send(a.ToProto()); err != nil {
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
