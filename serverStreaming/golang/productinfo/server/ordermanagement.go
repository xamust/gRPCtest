package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"log"
	pb "productinfo/server/ecommerce"
	"strings"
)

type server struct {
	orderMap map[string]*pb.Order
}

func (s *server) SearchOrders(value *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range s.orderMap {
		log.Println(key, order)
		for _, item := range order.Items {
			log.Println(item)
			if strings.Contains(item, value.Value) {
				if err := stream.Send(order); err != nil {
					return fmt.Errorf("error sending msg to stream: %v", err)
				}
				log.Print("Match! : v%", key)
				break
			}
		}
	}
	return nil
}
