package handler

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"shop_srvs/goods_srv/proto"
)

// 品牌和轮播图
func (s *GoodsService) BrandList(ctx context.Context, in *proto.BrandFilterRequest, opts ...grpc.CallOption) (*proto.BrandListResponse, error) {
	return nil, nil
}
func (s *GoodsService) CreateBrand(ctx context.Context, in *proto.BrandRequest, opts ...grpc.CallOption) (*proto.BrandInfoResponse, error) {
	return nil, nil
}
func (s *GoodsService) DeleteBrand(ctx context.Context, in *proto.BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
func (s *GoodsService) UpdateBrand(ctx context.Context, in *proto.BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
