package usecase

import (
	"context"
	"github.com/ct-core-standard/internal/model"
	"github.com/ct-core-standard/internal/repository/mongo_storage"
	"github.com/ct-core-standard/internal/repository/service"
	svr "github.com/ct-core-standard/pkg/protobuf"
)

type AdListingUC interface {
	GetByListID(ctx context.Context, adID int64) (svr.GetAdInfoResponse, error)
}

type Repos struct {
	AdListingSrv   service.AdListingService     `inject:""`
	//EsAdListing    es_storage.ESAdListing       `inject:""`
	MongoAdListing mongo_storage.MongoAdListing `inject:""`
}

type adListingImpl struct {
	*Repos
}

func NewAdListingUC(repos *Repos) AdListingUC {
	return &adListingImpl{
		repos,
	}
}

func (al *adListingImpl) GetByListID(ctx context.Context, adID int64) (res svr.GetAdInfoResponse, err error) {
	ad, err := al.AdListingSrv.GetByListID(ctx, adID)
	res = al.formatAdListingResp(ad)
	return res, err
}

func (al *adListingImpl) formatAdListingResp(ad model.AdListing) svr.GetAdInfoResponse {
	resp := svr.GetAdInfoResponse{
		Ad:         &svr.AdInfo{
			AdId:    ad.AdID,
			ListId:  ad.ListID,
			Subject: ad.Subject,
			Body:    ad.Body,
		},
		Parameters: nil,
	}
	return resp
}
