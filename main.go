package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"casai.com/posada/lib"
	pb "casai.com/posada/proto"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/gogo/protobuf/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var (
	prod = flag.Bool("prod", false, "Whether we'll run with prod config.")
)

const (
	devFile  = "dev.cfg"
	prodFile = "prod.cfg"
)

func sendToMessageQueue() {

}

func loadConfig(prod bool) *pb.Config {
	fname := devFile
	if prod {
		log.Print("Reading configuration file for production")
		fname = prodFile
	} else {
		log.Print("Reading configuration file for dev")
	}
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err, " at ", fname)
	}
	conf := &pb.Config{}
	if err := proto.UnmarshalText(string(in), conf); err != nil {
		log.Fatalln("Failed to parse Config at protobuffer:", err, " contents ", string(in))
	}

	return conf
}

type posadaServer struct {
	db *pg.DB
}

// GetListing returns the feature at the given point.
func (s *posadaServer) GetListing(c context.Context, r *pb.GetListingRequest) (*pb.GetListingResponse, error) {
	return &pb.GetListingResponse{}, nil
}

// UpdateListing returns the feature at the given point.
func (s *posadaServer) UpdateListing(c context.Context, r *pb.UpdateListingRequest) (*pb.UpdateListingResponse, error) {
	return &pb.UpdateListingResponse{}, nil
}

// CreateListing returns the feature at the given point.
func (s *posadaServer) CreateListing(c context.Context, r *pb.CreateListingRequest) (*pb.CreateListingResponse, error) {
	return &pb.CreateListingResponse{}, nil
}

// DeleteListing returns the feature at the given point.
func (s *posadaServer) DeleteListing(c context.Context, r *pb.DeleteListingRequest) (*pb.DeleteListingResponse, error) {
	return &pb.DeleteListingResponse{}, nil
}

// GetReservation returns the feature at the given point.
func (s *posadaServer) GetReservation(c context.Context, r *pb.GetReservationRequest) (*pb.GetReservationResponse, error) {
	return &pb.GetReservationResponse{}, nil
}

// UpdateReservation returns the feature at the given point.
func (s *posadaServer) UpdateReservation(c context.Context, r *pb.UpdateReservationRequest) (*pb.UpdateReservationResponse, error) {
	return &pb.UpdateReservationResponse{}, nil
}

// CreateReservation returns the feature at the given point.
func (s *posadaServer) CreateReservation(c context.Context, r *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	return &pb.CreateReservationResponse{}, nil
}

// DeleteReservation returns the feature at the given point.
func (s *posadaServer) DeleteReservation(c context.Context, r *pb.DeleteReservationRequest) (*pb.DeleteReservationResponse, error) {
	return &pb.DeleteReservationResponse{}, nil
}

// GetGuest returns the feature at the given point.
func (s *posadaServer) GetGuest(c context.Context, r *pb.GetGuestRequest) (*pb.GetGuestResponse, error) {
	return &pb.GetGuestResponse{}, nil
}

// UpdateGuest returns the feature at the given point.
func (s *posadaServer) UpdateGuest(c context.Context, r *pb.UpdateGuestRequest) (*pb.UpdateGuestResponse, error) {
	return &pb.UpdateGuestResponse{}, nil
}

// CreateGuest returns the feature at the given point.
func (s *posadaServer) CreateGuest(c context.Context, r *pb.CreateGuestRequest) (*pb.CreateGuestResponse, error) {
	return &pb.CreateGuestResponse{}, nil
}

// DeleteGuest returns the feature at the given point.
func (s *posadaServer) DeleteGuest(c context.Context, r *pb.DeleteGuestRequest) (*pb.DeleteGuestResponse, error) {
	return &pb.DeleteGuestResponse{}, nil
}

// GetCalendar returns the feature at the given point.
func (s *posadaServer) GetCalendar(c context.Context, r *pb.GetCalendarRequest) (*pb.GetCalendarResponse, error) {
	return &pb.GetCalendarResponse{}, nil
}

// UpdateCalendar returns the feature at the given point.
func (s *posadaServer) UpdateCalendar(c context.Context, r *pb.UpdateCalendarRequest) (*pb.UpdateCalendarResponse, error) {
	return &pb.UpdateCalendarResponse{}, nil
}

// CreateCalendar returns the feature at the given point.
func (s *posadaServer) CreateCalendar(c context.Context, r *pb.CreateCalendarRequest) (*pb.CreateCalendarResponse, error) {
	return &pb.CreateCalendarResponse{}, nil
}

// DeleteCalendar returns the feature at the given point.
func (s *posadaServer) DeleteCalendar(c context.Context, r *pb.DeleteCalendarRequest) (*pb.DeleteCalendarResponse, error) {
	return &pb.DeleteCalendarResponse{}, nil
}

// GetConversation returns the feature at the given point.
func (s *posadaServer) GetConversation(c context.Context, r *pb.GetConversationRequest) (*pb.GetConversationResponse, error) {
	return &pb.GetConversationResponse{}, nil
}

// UpdateConversation returns the feature at the given point.
func (s *posadaServer) UpdateConversation(c context.Context, r *pb.UpdateConversationRequest) (*pb.UpdateConversationResponse, error) {
	return &pb.UpdateConversationResponse{}, nil
}

// CreateConversation returns the feature at the given point.
func (s *posadaServer) CreateConversation(c context.Context, r *pb.CreateConversationRequest) (*pb.CreateConversationResponse, error) {
	return &pb.CreateConversationResponse{}, nil
}

// DeleteConversation returns the feature at the given point.
func (s *posadaServer) DeleteConversation(c context.Context, r *pb.DeleteConversationRequest) (*pb.DeleteConversationResponse, error) {
	return &pb.DeleteConversationResponse{}, nil
}

// For testing purposes.
type IPgDb interface {
	CreateTable(model interface{}, opt *orm.CreateTableOptions) error
}

// maybeCreateSchemas creates tables if they don't exist.
func maybeCreateSchemas(db IPgDb) {
	for _, model := range lib.Tables {
		err := db.CreateTable(model, &orm.CreateTableOptions{})
		if err != nil {
			if strings.Contains(err.Error(), "FATAL") {
				log.Fatal(err)
			}
			log.Print(strings.Replace(err.Error(), "ERROR", "INFO", -1))
		}
	}
}

func main() {
	flag.Parse()
	conf := loadConfig(*prod)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.GetPort()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if conf.GetTls() != "" {
		creds, err := credentials.NewServerTLSFromFile(conf.GetCertFile(), conf.GetKeyFile())
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	db := pg.Connect(&pg.Options{
		User:     conf.GetDbUser(),
		Password: conf.GetDbPass(),
		Database: conf.GetDbName(),
		Addr:     conf.GetDbAddr(),
	})
	defer db.Close()

	pb.RegisterPosadaServer(grpcServer, &posadaServer{db: db})
	maybeCreateSchemas(db)
	grpcServer.Serve(lis)
}
