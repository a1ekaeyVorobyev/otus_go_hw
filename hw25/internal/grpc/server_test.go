package grpcserver

import (
	"context"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw22/internal/calendar/calendar"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw22/internal/calendar/event"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw22/internal/config"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw22/internal/storage"
	proto "github.com/a1ekaeyVorobyev/otus_go_hw/hw22/pkg/calendar"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestAddEventGetEvent(t *testing.T) {
	inFile := storage.InFile{}
	inFile.Init()
	inFile.Clear()

	cal := calendar.Calendar{Storage: &inFile, Logger: &logrus.Logger{}}
	grpcServer := Server{Config: config.Config{GrpcServer: "127.0.0.1:55051"}, Calendar: &cal, Logger: &logrus.Logger{}}
	ctx := context.Background()
	go grpcServer.Run()
	time.Sleep(time.Second)
	defer grpcServer.Shutdown()
	conn, err := grpc.Dial("127.0.0.1:55051", []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		t.Error("Fail connect to GRPC server")
	}
	client := proto.NewCalendarClient(conn)
	//Add 10 record
	source := make(map[int]event.Event)
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := "Event " + string(i)
		note := "This envets" + string(i) + " start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		source[len(source)] = event
		start, err := ptypes.TimestampProto(event.StartTime)
		if err != nil {
			t.Error(err.Error())
		}
		finish, err := ptypes.TimestampProto(event.EndTime)
		if err != nil {
			t.Error(err.Error())
		}
		sendGrpcEvent := proto.Event{StartTime: start, EndTime: finish, Duration: int32(event.Duration),
			Typeduration: int32(event.TypeDuration), Title: event.Title, Note: event.Note}
		_, err = client.AddEvent(ctx, &sendGrpcEvent)
		if err != nil {
			t.Error("Fail AddEvent")
		}
	}

	getGrpcEvents, err := client.GetAllEvents(ctx, &empty.Empty{})
	if err != nil {
		t.Error("Fail GetEvent() with id=0")
	}

	cnt, err := client.CountRecord(ctx, &empty.Empty{})
	if err != nil {
		t.Error("Fail CountRecord()")
	}
	if cnt.Count != 10 {
		t.Error("Fail GetEvent() with id=0")
	}
	//check record event
	for _, v := range getGrpcEvents.Events {
		s := source[int(v.Id)]
		start, err := ptypes.Timestamp(v.StartTime)
		if err != nil {
			t.Error(err.Error())
		}
		finish, err := ptypes.Timestamp(v.EndTime)
		if err != nil {
			t.Error(err.Error())
		}
		if s.Title != v.Title || s.Note != v.Note || s.StartTime != start || s.EndTime != finish ||
			s.Duration != int(v.Duration) || s.TypeDuration != int(v.Typeduration) {
			t.Error("Get Event's not same", start, "-", s.StartTime)
		}

	}
	//delete all event
	for _, v := range getGrpcEvents.Events {
		_, err = client.DeleteEvent(ctx, &proto.Id{Id: v.Id})
		if err != nil {
			t.Error("Fail DelEvent() with id=0")
		}
	}
	cnt, err = client.CountRecord(ctx, &empty.Empty{})
	if err != nil {
		t.Error("Fail DeleteRecord()")
	}
	if cnt.Count != 0 {
		t.Error("Don't delete all record")
	}
}
