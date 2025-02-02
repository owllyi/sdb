package server

import (
	"github.com/yemingfeng/sdb/internal/service"
	"github.com/yemingfeng/sdb/pkg/pb"
	"golang.org/x/net/context"
)

type SetServer struct {
	pb.UnimplementedSDBServer
}

func (server *SetServer) SPush(_ context.Context, request *pb.SPushRequest) (*pb.SPushResponse, error) {
	res, err := service.SPush(request.Key, request.Values, request.Sync)
	return &pb.SPushResponse{Success: res}, err
}

func (server *SetServer) SPop(_ context.Context, request *pb.SPopRequest) (*pb.SPopResponse, error) {
	res, err := service.SPop(request.Key, request.Values, request.Sync)
	return &pb.SPopResponse{Success: res}, err
}

func (server *SetServer) SExist(_ context.Context, request *pb.SExistRequest) (*pb.SExistResponse, error) {
	res, err := service.SExist(request.Key, request.Values)
	return &pb.SExistResponse{Exists: res}, err
}

func (server *SetServer) SDel(_ context.Context, request *pb.SDelRequest) (*pb.SDelResponse, error) {
	res, err := service.SDel(request.Key, request.Sync)
	return &pb.SDelResponse{Success: res}, err
}

func (server *SetServer) SCount(_ context.Context, request *pb.SCountRequest) (*pb.SCountResponse, error) {
	res, err := service.SCount(request.Key)
	return &pb.SCountResponse{Count: res}, err
}
