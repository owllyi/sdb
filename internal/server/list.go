package server

import (
	"github.com/yemingfeng/sdb/internal/service"
	"github.com/yemingfeng/sdb/pkg/pb"
	"golang.org/x/net/context"
)

type ListServer struct {
	pb.UnimplementedSDBServer
}

func (server *ListServer) LPush(_ context.Context, request *pb.LPushRequest) (*pb.LPushResponse, error) {
	res, err := service.LPush(request.Key, request.Values, request.Sync)
	return &pb.LPushResponse{Success: res}, err
}

func (server *ListServer) LPop(_ context.Context, request *pb.LPopRequest) (*pb.LPopResponse, error) {
	res, err := service.LPop(request.Key, request.Values, request.Sync)
	return &pb.LPopResponse{Success: res}, err
}

func (server *ListServer) LRange(_ context.Context, request *pb.LRangeRequest) (*pb.LRangeResponse, error) {
	res, err := service.LRange(request.Key, request.Offset, request.Limit)
	return &pb.LRangeResponse{Values: res}, err
}

func (server *ListServer) LExist(_ context.Context, request *pb.LExistRequest) (*pb.LExistResponse, error) {
	res, err := service.LExist(request.Key, request.Values)
	return &pb.LExistResponse{Exists: res}, err
}

func (server *ListServer) LDel(_ context.Context, request *pb.LDelRequest) (*pb.LDelResponse, error) {
	res, err := service.LDel(request.Key, request.Sync)
	return &pb.LDelResponse{Success: res}, err
}

func (server *ListServer) LCount(_ context.Context, request *pb.LCountRequest) (*pb.LCountResponse, error) {
	res, err := service.LCount(request.Key)
	return &pb.LCountResponse{Count: res}, err
}
