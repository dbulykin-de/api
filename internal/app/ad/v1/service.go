package ad

import adV1 "ad-api/internal/pkg/pb/ad/v1"

type Implementation struct {
	adV1.UnimplementedAdServiceServer
}

func NewAdService() *Implementation {
	return &Implementation{}
}
