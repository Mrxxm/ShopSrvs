package handler

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"shop_srvs/goods_srv/proto"
)

// 轮播图
func (s *GoodsService) BannerList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*proto.BannerListResponse, error) {
	return nil, nil
}
func (s *GoodsService) CreateBanner(ctx context.Context, in *proto.BannerRequest, opts ...grpc.CallOption) (*proto.BannerResponse, error) {
	return nil, nil
}
func (s *GoodsService) DeleteBanner(ctx context.Context, in *proto.BannerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
func (s *GoodsService) UpdateBanner(ctx context.Context, in *proto.BannerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
