package grpcserver

import (
	"context"
	proto "github.com/a1ekaeyVorobyev/otus_go_hw/hw21/pkg/calendar"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"net"

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
	server   *grpc.Server
}

type CalendarServerGrpc struct {
	log      *logrus.Logger
	calendar *calendar.Calendar
}


func (s *Server) Run() {
	s.Logger.Info("Start GRPC server:", s.Config.GrpcServer)

	listener, err := net.Listen("tcp", s.Config.GrpcServer)
	if err != nil {
		s.Logger.Fatalf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer()
	proto.RegisterCalendarServer(s.server,s)

	err = s.server.Serve(listener)
	if err != nil {
		s.Logger.Fatalf("grps server startup error: %v", err)
	}
}

func (s *Server) Shutdown() {
	s.Logger.Info("grps server shutdown ...")
	s.server.GracefulStop()
}


func (s *Server) AddEvent(ctx context.Context, e *proto.Event) (*empty.Empty, error) {
	s.Logger.Debug("gRPC event AddEvent(): ", e)
	start,err := ptypes.Timestamp(e.StartTime)
	if err!=nil{
		return &empty.Empty{}, err
	}
	finish,err := ptypes.Timestamp(e.StartTime)
	if err!=nil{
		return &empty.Empty{}, err
	}
	err = s.Calendar.AddEvent(event.Event{
		StartTime:		start,
		EndTime:		finish,
		Duration: 		int(e.Duration),
		TypeDuration:	int(e.Typeduration),
		Title:			e.Title,
		Note: 			e.Note,
	})

	return &empty.Empty{}, err
}

func (s *Server) GetEvent(ctx context.Context, id *proto.Id) (*proto.Event, error) {
	s.Logger.Debug("Income gRPC GetEvent() id:", id)
	calendarEvent, err := s.Calendar.GetEvent(int(id.Id))
	start,err := ptypes.TimestampProto(calendarEvent.StartTime)
	if err!=nil{
		return nil, err
	}
	finish,err := ptypes.TimestampProto(calendarEvent.EndTime)
	if err!=nil{
		return nil, err
	}
	return &proto.Event{
		Id:          	int32(calendarEvent.Id),
		StartTime:   	start,
		EndTime:     	finish,
		Duration: 		int32(calendarEvent.Duration),
		Typeduration:	int32(calendarEvent.TypeDuration),
		Title:       	calendarEvent.Title,
		Note: 			calendarEvent.Note,
	}, err
}

func (s *Server) DeleteEvent(ctx context.Context, e *proto.Id) (*empty.Empty, error) {
	s.Logger.Debug("gRPC  event DeleteEvent() id:", e)

	return &empty.Empty{}, s.Calendar.DeleteEvent(int(e.Id))
}

func (s *Server) CountRecord(ctx context.Context, e *empty.Empty) (*proto.Count,error) {
	s.Logger.Debug("gRPC  event CountRecord()")
	var cnt = *new(proto.Count)
	cnt.Count=int32(s.Calendar.CountRecord())
	return &cnt,nil
}


func (s *Server) EditEvent(ctx context.Context, e *proto.Event) (*empty.Empty, error) {
	s.Logger.Debug("Income gRPC EditEvent() event:", e)
	start,err := ptypes.Timestamp(e.StartTime)
	if err!=nil{
		return &empty.Empty{}, err
	}
	finish,err := ptypes.Timestamp(e.StartTime)
	if err!=nil{
		return &empty.Empty{}, err
	}
	err = s.Calendar.EditEvent(event.Event{
		StartTime:	start,
		EndTime:    finish,
		Duration: 		int(e.Duration),
		TypeDuration:	int(e.Typeduration),
		Title:			e.Title,
		Note: 			e.Note,
	})

	return &empty.Empty{}, err
}

func (s *Server) GetAllEvents(ctx context.Context, e *empty.Empty) (*proto.Events, error) {
	s.Logger.Debug("gRPC event GetAllEvents()")
	calendarEvents, err := s.Calendar.GetAllEvents()
	protobufEvents := make([]*proto.Event, 0, s.Calendar.CountRecord())

	for _, calendarEvent := range calendarEvents {
		start,err := ptypes.TimestampProto(calendarEvent.StartTime)
		if err!=nil{
			return nil, err
		}
		finish,err := ptypes.TimestampProto(calendarEvent.EndTime)
		if err!=nil{
			return nil, err
		}
		protobufEvent := proto.Event{
			Id:          	int32(calendarEvent.Id),
			StartTime:   	start,
			EndTime:     	finish,
			Duration: 		int32(calendarEvent.Duration),
			Typeduration:	int32(calendarEvent.TypeDuration),
			Title:       	calendarEvent.Title,
			Note: 			calendarEvent.Note,
		}
		protobufEvents = append(protobufEvents, &protobufEvent)
	}

	return &proto.Events{Events: protobufEvents}, err
}
