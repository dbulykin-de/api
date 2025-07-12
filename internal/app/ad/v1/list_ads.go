package ad

import (
	"context"
	"time"

	adV1 "ad-api/internal/pkg/pb/ad/v1"

	"google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) ListAds(ctx context.Context, req *adV1.ListAdsRequest) (*adV1.ListAdsResponse, error) {
	return &adV1.ListAdsResponse{
		Ads: []*adV1.ListAdsResponse_Ad{
			{
				Id:          "ad-1",
				Title:       "Настройка серверного оборудования",
				Description: "Настраиваем сервачки",
				Category:    "ad-category-it",
				AuthorId:    "user_123",
				Price: &money.Money{
					CurrencyCode: "RUB",
					Units:        100,
				},
				Status:    "active",
				CreatedAt: timestamppb.New(time.Now()),
			},
		},
	}, nil
}
