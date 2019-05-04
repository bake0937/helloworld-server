package server

import (
	"context"
	pb "proton"
)

func (s *Server) ListAddress(ctx context.Context, req *pb.ListAddressRequest) (res *pb.ListAddressResponse, err error) {
	address_list, err := s.model.AllFindAddress()

	if address_list == nil {
		return &pb.ListAddressResponse{}, nil
	}
	var addresses []*pb.Address

	for _, val := range address_list {
		addresses = append(addresses, &pb.Address{
			Id:    uint32(val.ID),
			Email: val.Email},
		)
	}

	res = &pb.ListAddressResponse{Address: addresses}
	return res, nil
}
