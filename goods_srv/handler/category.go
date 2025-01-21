package handler

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"shop_srvs/goods_srv/proto"
)

// 商品分类
func (s *GoodsService) GetAllCategorysList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*proto.CategoryListResponse, error) {
	return nil, nil
}

func (s *GoodsService) GetSubCategory(ctx context.Context, in *proto.CategoryListRequest, opts ...grpc.CallOption) (*proto.SubCategoryListResponse, error) {
	return nil, nil
}
func (s *GoodsService) CreateCategory(ctx context.Context, in *proto.CategoryInfoRequest, opts ...grpc.CallOption) (*proto.CategoryInfoResponse, error) {
	return nil, nil
}
func (s *GoodsService) DeleteCategory(ctx context.Context, in *proto.DeleteCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
func (s *GoodsService) UpdateCategory(ctx context.Context, in *proto.CategoryInfoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
