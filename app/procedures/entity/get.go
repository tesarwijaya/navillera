package entity

import (
  "fmt"

  pb "github.com/tesarwijaya/navillera/proto"
)

// Get ...
func Get(request *pb.Request) (*pb.Response, error) {
  fmt.Printf("\nEntity Get: %v\n", request)
  return &pb.Response{
		Status: 200,
		Data: map[string]string{
			"message": "success",
		},
	}, nil
}
