package handler

import (
	"gorm.io/gorm"
	"shop_srvs/goods_srv/proto"
)

type GoodsService struct {
	proto.UnimplementedGoodsServer
}

//// 商品接口
//func (s *GoodsService) GoodsList(ctx context.Context, request *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
//	return nil, nil
//}
//
//func (s *GoodsService) BatchGetGoods(ctx context.Context, request *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
//	return nil, nil
//}
//func (s *GoodsService) CreateGoods(ctx context.Context, request *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
//	return nil, nil
//}
//func (s *GoodsService) DeleteGoods(ctx context.Context, request *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
//	return nil, nil
//}
//func (s *GoodsService) UpdateGoods(ctx context.Context, request *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
//	return nil, nil
//}
//func (s *GoodsService) GetGoodsDetail(ctx context.Context, request *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
//	return nil, nil
//}

// 分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
