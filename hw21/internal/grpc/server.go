package grps

import (
	"context"
	proto "github.com/a1ekaeyVorobyev/otus_go_hw/hw21/pkg/calendar"


	"time"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/calendar/calendar"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/calendar/event"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw21/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/golang/protobuf/ptypes/empty"
)

type Server struct {
	Config   config.Config
	Logger   *logrus.Logger
	Calendar *calendar.Calendar
	//server   *grpc.Server
}

type CalendarServerGrpc struct {
	log      *logrus.Logger
	calendar *calendar.Calendar
}


func (s *Server) Run() {
	s.Logger.Info("Start GRPC server:", s.Config.GrpcServer)

	//listener, err := net.Listen("tcp", s.Config.GrpcServer)
	//if err != nil {
	//	s.Logger.Fatalf("failed to listen: %v", err)
	//}

	//s.server = grpc.NewServer()
	//protobuf.RegisterCalendarServer(s.server, s)

	//err = s.server.Serve(listener)
	//if err != nil {
	//	s.Logger.Fatalf("failed to run grpc server: %v", err)
	//}
}

func (s *Server) Shutdown() {
	s.Logger.Info("Graceful shutdown GRPC server...")
	//s.server.GracefulStop()
}


func (s *Server) AddEvent(ctx context.Context, e *proto.Event) (*empty.Empty, error) {
	s.Logger.Debug("Income gRPC AddEvent() event: ", e)
	err := s.Calendar.AddEvent(event.Event{
		StartTime:		Timestamp(e.StartTime),
		EndTime:		Timestamp(e.EndTime),
		Duration: 		int(e.Duration),
		TypeDuration:	int(e.Typeduration),
		Title:			e.Title,
		Note: 			e.Note,
	})

	return &empty.Empty{}, err
}

func (s *Server) GetEvent(ctx context.Context, grpcId *proto.Id) (*proto.Event, error) {
	s.Logger.Debug("Income gRPC GetEvent() id:", grpcId)
	calendarEvent, err := s.Calendar.GetEvent(int(grpcId.Id))

	return &protobuf.Event{
		Id:          int32(calendarEvent.Id),
		StartTime:   calendarEvent.StartTime.Unix(),
		EndTime:     calendarEvent.EndTime.Unix(),
		Title:       calendarEvent.Title,
		Description: calendarEvent.Description,
	}, err
}

func (s *Server) DeleteEvent(ctx context.Context, grpcId *proto.Id) (*empty.Empty, error) {
	s.Logger.Debug("Income gRPC DelEvent() id:", grpcId)

	return &empty.Empty{}, s.Calendar.DelEvent(int(grpcId.Id))
}

func (s *Server) EditEvent(ctx context.Context, grpcE *proto.Event) (*empty.Empty, error) {
	s.Logger.Debug("Income gRPC EditEvent() event:", grpcE)

	err := s.Calendar.EditEvent(event.Event{
		StartTime:   time.Unix(grpcE.StartTime, 0),
		EndTime:     time.Unix(grpcE.EndTime, 0),
		Title:       grpcE.Title,
		Description: grpcE.Description,
	})

	return &empty.Empty{}, err
}

func (s *Server) GetAllEvents(ctx context.Context, ev *empty.Empty) (*proto.Events, error) {
	s.Logger.Debug("Income gRPC GetAllEvents()")
	calendarEvents, err := s.Calendar.GetAllEvents()
	l := len(calendarEvents)

	protobufEvents := make([]*proto.Event, 0, l)

	for _, calendarEvent := range calendarEvents {
		protobufEvent := proto.Event{
			Id:          int32(calendarEvent.Id),
			StartTime:   calendarEvent.StartTime.Unix(),
			EndTime:     calendarEvent.EndTime.Unix(),
			Title:       calendarEvent.Title,
			Description: calendarEvent.Description,
		}
		protobufEvents = append(protobufEvents, &protobufEvent)
	}

	return &proto.Events{Events: protobufEvents}, err
}
