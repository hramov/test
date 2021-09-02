package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"

	"testing/domains/dto"
	entity "testing/domains/entities"
	"testing/modules/cache"
	pb "testing/modules/grpc/proto"
	"testing/modules/logger"
	"testing/modules/queue"

	container "github.com/golobby/container/v3"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTestingServer
}

func (s *Server) GetUsers(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	var userEntity entity.UserEntity
	container.NamedResolve(&userEntity, "UserEntity")
	var cache cache.Cache
	container.NamedResolve(&cache, "Cache")
	var logger logger.Logger
	container.NamedResolve(&logger, "Logger")

	var pbUsers []*pb.GetSingleReply

	cachedUsers, err := cache.Get("users")
	if err == nil {
		err := json.Unmarshal([]byte(cachedUsers), &pbUsers)
		if err != nil {
			return nil, err
		}
		return &pb.GetReply{
			User: pbUsers,
		}, nil
	}

	users, err := userEntity.Get()
	for i := 0; i < len(users); i++ {
		pbUsers = append(pbUsers, &pb.GetSingleReply{
			Email:    &users[i].Email,
			Password: &users[i].Password,
		})
	}

	data, err := json.Marshal(pbUsers)
	if err != nil {
		return nil, err
	}
	cache.Set("users", string(data))

	return &pb.GetReply{
		User: pbUsers,
	}, nil
}

func (s *Server) AddUser(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	var userEntity entity.UserEntity
	container.NamedResolve(&userEntity, "UserEntity")

	var queue queue.Queue
	container.NamedResolve(&queue, "Queue")

	queue.SendMessage("logs", fmt.Sprintf("Adding new user with email %s", *in.Email))

	id, err := userEntity.Add(&dto.UserDto{
		Email:    *in.Email,
		Password: *in.Password,
	})

	retId := fmt.Sprintf("%d", id)
	return &pb.AddReply{
		Id: &retId,
	}, err
}

func (s *Server) DeleteUser(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteReply, error) {
	var userEntity entity.UserEntity
	container.NamedResolve(&userEntity, "UserEntity")
	reqId, err := strconv.ParseUint(*in.Id, 10, 64)
	id, err := userEntity.Delete(reqId)

	retId := fmt.Sprintf("%d", id)
	return &pb.DeleteReply{
		Id: &retId,
	}, err
}

func (s *Server) Start() {
	log.Println("Server has beed started")
	lis, err := net.Listen("tcp", ":5005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterTestingServer(server, &Server{})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
