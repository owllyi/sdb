package server

import (
	"context"
	"github.com/yemingfeng/sdb/internal/conf"
	"github.com/yemingfeng/sdb/pkg/pb"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type SDBGrpcServer struct {
	grpcServer *grpc.Server
	StringServer
	ListServer
	SetServer
	SortedSetServer
	BloomFilterServer
	HyperLogLogServer
	PubSubServer
}

func NewSDBGrpcServer() *SDBGrpcServer {
	grpcprometheus.EnableHandlingTimeHistogram()

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			grpcrecovery.StreamServerInterceptor(),
			grpcprometheus.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpcmiddleware.ChainUnaryServer(
				ratelimit.UnaryServerInterceptor(CreateRateLimit(conf.Conf.Server.Rate))),
			grpcrecovery.UnaryServerInterceptor(),
			grpcprometheus.UnaryServerInterceptor,
			grpcmiddleware.ChainUnaryServer(func(ctx context.Context, req interface{},
				info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				start := time.Now()

				defer func() {
					cost := time.Since(start)
					if cost.Milliseconds() > conf.Conf.Server.SlowQueryThreshold {
						log.Printf("Slow query: %s, cost: %d", req, cost.Milliseconds())
					}
				}()

				return handler(ctx, req)
			}),
		)),
	)
	sdbGrpcServer := SDBGrpcServer{grpcServer: grpcServer}
	pb.RegisterSDBServer(grpcServer, &sdbGrpcServer)
	return &sdbGrpcServer
}

func (sdbGrpcServer *SDBGrpcServer) Start() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(conf.Conf.Server.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}
	if err := sdbGrpcServer.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
	log.Printf("serve: %d", conf.Conf.Server.GRPCPort)
}
