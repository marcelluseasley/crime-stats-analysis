package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/marcelluseasley/crime-stats-analysis/server/bing"

	"github.com/marcelluseasley/crime-stats-analysis/crimestats"
)

type server struct{
	crimestats.UnimplementedCrimeStatsServiceServer
}

func(s *server) CrimeStats(ctx context.Context, req *crimestats.CrimeStatsRequest) (*crimestats.CrimeStatsResponse, error){
	var address string
	var latitude float64
	var longitude float64

	log.Printf("CrimeStats function was invoked with %v", req)

	location, err := bing.GetLocations(req.GetStreet(),req.GetCity(), req.GetState(), req.GetZipcode())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not find location: %v", err)
	}

	if location != nil && len(location.ResourceSets) > 0 && len(location.ResourceSets[0].Resources) > 0 {
		mainResource := location.ResourceSets[0].Resources[0]
		address = mainResource.Name

		if len(mainResource.Point.Coordinates) != 2 {
			return nil, status.Errorf(codes.Internal, "Wrong number of coordinates: %v", err)
		}
			latitude = mainResource.Point.Coordinates[0]
			longitude = mainResource.Point.Coordinates[1]
		
		
	}

	return &crimestats.CrimeStatsResponse{
		Items: []*crimestats.CrimeStatsResponse_Crime{
			{
				Description: address,
				Datetime:    "",
				Location:    []float64{
					latitude,
					longitude,
				},
			}, // Add a comma here
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	crimestats.RegisterCrimeStatsServiceServer(s, &server{})

	// Enable gRPC server reflection
	reflection.Register(s)

	log.Println("Server started. Listening on port 50051.")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}