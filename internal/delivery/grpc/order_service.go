package grpc

import (
    "context"
    "project/internal/service"
    
    pb "project/proto"
)

type OrderService struct {
    Service *service.OrderService
    pb.UnimplementedOrderServiceServer
}

func (s *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
    orders, err := s.Service.ListOrders()
    if err != nil {
        return nil, err
    }

    var grpcOrders []*pb.Order
    for _, order := range orders {
        grpcOrders = append(grpcOrders, &pb.Order{
            Id:           int32(order.ID),
            CustomerName: order.CustomerName,
            TotalAmount:  float32(order.TotalAmount),
            CreatedAt:    order.CreatedAt.String(),
        })
    }

    return &pb.ListOrdersResponse{Orders: grpcOrders}, nil
}