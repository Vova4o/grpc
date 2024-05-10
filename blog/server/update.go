package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vova4o/grpc-go-study/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog function was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in MongoDB: %v", err),
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return &emptypb.Empty{}, nil
}
