package grpcserver

import (
	"context"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/calendar"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
	proto "github.com/a1ekaeyVorobyev/otus_go_hw/hw22/pkg/calendar"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)
type Config struct{
	Server 			string `yaml:"Server"`
}

type server struct {
	config   Config
	logger   *logrus.Logger
	calendar *calendar.Calendar
	server   *grpc.Server
}

type CalendarServerGrpc struct {
	log      *logrus.Logger
	calendar *calendar.Calendar
}

func NewGRPSDerver(conf Config, logger *logrus.Logger, calendar *calendar.Calendar) (s *server, err error) {
	s = &server{}
	s.config = conf
	s.logger = logger
	s.calendar = calendar
	return s,nil
}

func (s *server) Run() {
	s.logger.Info("Start GRPC server:", s.config.Server)

	listener, err := net.Listen("tcp", s.config.Server)
	if err != nil {
		s.logger.Fatalf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer()
	proto.RegisterCalendarServer(s.server, s)

	err = s.server.Serve(listener)
	if err != nil {
		s.logger.Fatalf("grps server startup error: %v", err)
	}
}

func (s *server) Shutdown() {
	s.logger.Info("grps server shutdown ...")
	s.server.GracefulStop()
}

func (s *server) AddEvent(ctx context.Context, e *proto.Event) (*empty.Empty, error) {
	s.logger.Debug("gRPC event AddEvent(): ", e)
	start, err := ptypes.Timestamp(e.StartTime)
	if err != nil {
		return &empty.Empty{}, err
	}
	finish, err := ptypes.Timestamp(e.EndTime)
	if err != nil {
		return &empty.Empty{}, err
	}
	err = s.calendar.AddEvent(event.Event{
		StartTime:    start,
		EndTime:      finish,
		Duration:     int(e.Duration),
		TypeDuration: int(e.Typeduration),
		Title:        e.Title,
		Note:         e.Note,
	})

	return &empty.Empty{}, err
}

func (s *server) GetEvent(ctx context.Context, id *proto.Id) (*proto.Event, error) {
	s.logger.Debug("Income gRPC GetEvent() id:", id)
	calendarEvent, err := s.calendar.GetEvent(int(id.Id))
	start, err := ptypes.TimestampProto(calendarEvent.StartTime)
	if err != nil {
		return nil, err
	}
	finish, err := ptypes.TimestampProto(calendarEvent.EndTime)
	if err != nil {
		return nil, err
	}
	return &proto.Event{
		Id:           int32(calendarEvent.Id),
		StartTime:    start,
		EndTime:      finish,
		Duration:     int32(calendarEvent.Duration),
		Typeduration: int32(calendarEvent.TypeDuration),
		Title:        calendarEvent.Title,
		Note:         calendarEvent.Note,
	}, err
}

func (s *server) DeleteEvent(ctx context.Context, e *proto.Id) (*empty.Empty, error) {
	s.logger.Debug("gRPC  event DeleteEvent() id:", e)

	return &empty.Empty{}, s.calendar.DeleteEvent(int(e.Id))
}

func (s *server) CountRecord(ctx context.Context, e *empty.Empty) (*proto.Count, error) {
	s.logger.Debug("gRPC  event CountRecord()")
	var cnt = *new(proto.Count)
	cnt.Count = int32(s.calendar.CountRecord())
	return &cnt, nil
}

func (s *server) EditEvent(ctx context.Context, e *proto.Event) (*empty.Empty, error) {
	s.logger.Debug("Income gRPC EditEvent() event:", e)
	start, err := ptypes.Timestamp(e.StartTime)
	if err != nil {
		return &empty.Empty{}, err
	}
	finish, err := ptypes.Timestamp(e.EndTime)
	if err != nil {
		return &empty.Empty{}, err
	}
	err = s.calendar.EditEvent(event.Event{
		StartTime:    start,
		EndTime:      finish,
		Duration:     int(e.Duration),
		TypeDuration: int(e.Typeduration),
		Title:        e.Title,
		Note:         e.Note,
	})

	return &empty.Empty{}, err
}

func (s *server) GetAllEvents(ctx context.Context, e *empty.Empty) (*proto.Events, error) {
	s.logger.Debug("gRPC event GetAllEvents()")
	calendarEvents, err := s.calendar.GetAllEvents()
	protobufEvents := make([]*proto.Event, 0, s.calendar.CountRecord())

	for _, calendarEvent := range calendarEvents {
		start, err := ptypes.TimestampProto(calendarEvent.StartTime)
		if err != nil {
			return nil, err
		}
		finish, err := ptypes.TimestampProto(calendarEvent.EndTime)
		if err != nil {
			return nil, err
		}
		protobufEvent := proto.Event{
			Id:           int32(calendarEvent.Id),
			StartTime:    start,
			EndTime:      finish,
			Duration:     int32(calendarEvent.Duration),
			Typeduration: int32(calendarEvent.TypeDuration),
			Title:        calendarEvent.Title,
			Note:         calendarEvent.Note,
		}
		protobufEvents = append(protobufEvents, &protobufEvent)
	}

	return &proto.Events{Events: protobufEvents}, err
}
