package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/calendar/event"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/config"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/logger"
	proto "github.com/a1ekaeyVorobyev/otus_go_hw/hw22/pkg/calendar"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config/config.yaml", "Config file")
	flag.Parse()
	if configFile == "" {
		_, _ = fmt.Fprint(os.Stderr, "Don't config file")
		os.Exit(2)
	}

	conf, err := config.ReadFromFile(configFile)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
	logger, f := logger.GetLogger(conf.Log)
	if f != nil {
		defer f.Close()
	}
	ctxConn, _ := context.WithTimeout(context.Background(), time.Second*10)
	conn, err := grpc.DialContext(ctxConn, conf.Grps.Server, []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
	if err != nil {
		log.Fatalln("Can't connect to grpc server, error: ", err)
	}
	client := proto.NewCalendarClient(conn)
	ctx := context.Background()
	//add 10 record
	dateStart := time.Date(2020, 1, 1, 11, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		dateStart.AddDate(0, 0, 1)
		dateEnd.AddDate(0, 0, 1)
		title := "Event " + string(i)
		note := "This envets" + string(i) + " start:" + dateStart.Format(time.RFC3339) + " finish:" + dateEnd.Format(time.RFC3339)
		event, _ := event.CreateEvent(dateStart.Format(time.RFC3339), dateEnd.Format(time.RFC3339), title, note, 0, 0)
		start, err := ptypes.TimestampProto(event.StartTime)
		if err != nil {
			LogOnError(&logger, "Fail on add Event", err)
		}
		finish, err := ptypes.TimestampProto(event.EndTime)
		if err != nil {
			LogOnError(&logger, "Fail on add Event", err)
		}
		sendGrpcEvent := proto.Event{StartTime: start, EndTime: finish, Duration: int32(event.Duration),
			Typeduration: int32(event.TypeDuration), Title: event.Title, Note: event.Note}
		_, err = client.AddEvent(ctx, &sendGrpcEvent)
		if err != nil {
			LogOnError(&logger, "Fail on add Event", err)
		}
	}

	getGrpcEvents, err := client.GetAllEvents(ctx, &empty.Empty{})
	if err != nil {
		LogOnError(&logger, "Fail on GetAllEvents", err)
	}

	cnt, err := client.CountRecord(ctx, &empty.Empty{})
	if err != nil {
		LogOnError(&logger, "Fail on CountRecord", err)
	}
	logger.Info("Count events = ", cnt.Count)
	for _, v := range getGrpcEvents.Events {
		_, err = client.DeleteEvent(ctx, &proto.Id{Id: v.Id})
		if err != nil {
			LogOnError(&logger, "Fail on DeleteEvent", err)
		}
	}
	cnt, err = client.CountRecord(ctx, &empty.Empty{})
	if err != nil {
		LogOnError(&logger, "Fail on CountRecord", err)
	}
	logger.Info("Count events = ", cnt.Count)
}

func LogOnError(log *logrus.Logger, mes string, err error) {
	if err != nil {
		log.Errorln(fmt.Sprintf("%s, error: %v", mes, err))
	}
}
