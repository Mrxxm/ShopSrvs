package handler

import (
	"context"
	"shop_srvs/goods_srv/global"
	"shop_srvs/goods_srv/model"
	"shop_srvs/goods_srv/proto"
)

// 品牌和轮播图
func (s *GoodsService) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	var brandListResponse proto.BrandListResponse

	var brands []model.Brands
	result := global.DB.Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	var count int64
	global.DB.Model(&model.Brands{}).Count(&count)
	global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)

	var brandResponses []*proto.BrandInfoResponse
	for _, brand := range brands {
		brandResponse := proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		}
		brandResponses = append(brandResponses, &brandResponse)
	}

	brandListResponse.Data = brandResponses
	brandListResponse.Total = int32(count)

	return &brandListResponse, nil
}

//
//func (s *GoodsService) CreateBrand(context.Context, *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
//	return nil, nil
//}
//func (s *GoodsService) DeleteBrand(context.Context, *proto.BrandRequest) (*emptypb.Empty, error) {
//	return nil, nil
//}
//func (s *GoodsService) UpdateBrand(context.Context, *proto.BrandRequest) (*emptypb.Empty, error) {
//	return nil, nil
//}
