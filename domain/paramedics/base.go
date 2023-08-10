package paramedics

import (
	"github.com/jmoiron/sqlx"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func paramedicService() (client ParamedicClient, err error) {
	port := ":7000"
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client = NewParamedicClient(conn)
	return
}

func RegisterRouteGRPC(server *grpc.Server, db *sqlx.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	handler := NewParamedicHandlerGPRC(svc)

	RegisterParamedicServer(server, handler)
}
