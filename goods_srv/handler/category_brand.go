package handler

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"shop_srvs/goods_srv/proto"
)

// 品牌分类
func (s *GoodsService) CategoryBrandList(ctx context.Context, in *proto.CategoryBrandFilterRequest, opts ...grpc.CallOption) (*proto.CategoryBrandListResponse, error) {
	return nil, nil
}

// 通过category获取brands
func (s *GoodsService) GetCategoryBrandList(ctx context.Context, in *proto.CategoryInfoRequest, opts ...grpc.CallOption) (*proto.BrandListResponse, error) {
	return nil, nil
}
func (s *GoodsService) CreateCategoryBrand(ctx context.Context, in *proto.CategoryBrandRequest, opts ...grpc.CallOption) (*proto.CategoryBrandResponse, error) {
	return nil, nil
}
func (s *GoodsService) DeleteCategoryBrand(ctx context.Context, in *proto.CategoryBrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
func (s *GoodsService) UpdateCategoryBrand(ctx context.Context, in *proto.CategoryBrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
