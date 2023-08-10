package schedules

import (
	"github.com/jmoiron/sqlx"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func scheduleService() (client ScheduleClient, err error) {
	port := ":8000"
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client = NewScheduleClient(conn)
	return
}

func RegisterRouteGRPC(server *grpc.Server, db *sqlx.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	handler := NewScheduleHandlerGPRC(svc)

	RegisterScheduleServer(server, handler)
}
