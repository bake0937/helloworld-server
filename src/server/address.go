package server

import (
	"context"
	pb "proton"
)

func (s *Server) ListAddress(ctx context.Context, req *pb.ListAddressRequest) (res *pb.ListAddressResponse, err error) {

	// company := s.model.NewClinotCompany()

	// if err := company.Find(sq.Eq{"uuid": req.Uuid}); err != nil {
	// 	return nil, status.Error(codes.NotFound, "The clinot_company not found")
	// }

	// ClinotAddresses, err := s.model.FindClinotAddressUuid(company.ID)

	// var addresses []*proton.ClinotAddress

	// if ClinotAddresses == nil {
	// 	return &proton.ListClinotAddressResponse{}, nil
	// }

	// for _, ca := range ClinotAddresses {
	// 	addresses = append(addresses, &proton.ClinotAddress{
	// 		Id:              int32(ca.ID),
	// 		ClinotCompanyId: int32(ca.ClinotCompanyID),
	// 		Email:           ca.Email},
	// 	)
	// }

	//res = &proton.ListClinotAddressResponse{Name: company.Name, ClinotAddresses: addresses}
	return nil, nil
}
