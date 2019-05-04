package server

import (
	"context"
	model "model"
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

func (s *Server) CreateAddress(ctx context.Context, req *pb.CreateAddressRequest) (res *pb.CreateAddressResponse, err error) {
	if isValidAddress(req) {
		address := model.Address{}
		address.Email = req.Email
		s.model.CreateAddress(&address)
		return &pb.CreateAddressResponse{}, err
	}
	return &pb.CreateAddressResponse{}, nil
}

func isValidAddress(req *pb.CreateAddressRequest) bool {
	if len(req.GetEmail()) > 0 {
		return true
	}
	return false
}
