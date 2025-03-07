package handler

import (
	"context"
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"shop_srvs/order_srv/global"
	"shop_srvs/order_srv/model"
	"shop_srvs/order_srv/proto"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (s *InventoryServer) SetInv(ctx context.Context, req *proto.GoodsInvInfo) (*emptypb.Empty, error) {
	//设置库存， 如果我要更新库存
	var inv model.Inventory
	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv)
	inv.Goods = req.GoodsId
	inv.Stocks = req.Num

	global.DB.Save(&inv)
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) InvDetail(ctx context.Context, req *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	var inv model.Inventory
	if result := global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有库存信息")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num:     inv.Stocks,
	}, nil
}

func (s *InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {

	// 扣减库存，本地事务
	tx := global.DB.Begin()

	client := goredislib.NewClient(&goredislib.Options{
		Addr: "192.168.15.21:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	rs := redsync.New(pool)

	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory

		mutex := rs.NewMutex(fmt.Sprintf("goods_id_%d", goodInfo.GoodsId))

		if err := mutex.Lock(); err != nil {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "获取锁异常")
		}

		if result := global.DB.Where("goods = ?", goodInfo.GoodsId).First(&inv); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
		}
		// 判断库存是否充足
		if inv.Stocks < goodInfo.Num {
			tx.Rollback()
			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
		}
		// 扣减
		inv.Stocks -= goodInfo.Num
		tx.Save(&inv)

		if ok, err := mutex.Unlock(); !ok || err != nil {
			return nil, status.Errorf(codes.Internal, "释放锁异常")
		}
	}

	tx.Commit()

	return nil, nil
}

//func (s *InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
//	//扣减库存， 本地事务 [1:10,  2:5, 3: 20]
//	//数据库基本的一个应用场景：数据库事务
//	//并发情况之下 可能会出现超卖 1
//	client := goredislib.NewClient(&goredislib.Options{
//		Addr: "192.168.0.104:6379",
//	})
//	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
//	rs := redsync.New(pool)
//
//	mutex := rs.NewMutex(fmt.Sprintf("order_%d", req.OrderSn))
//	if err := mutex.Lock(); err != nil {
//		return nil, status.Errorf(codes.Internal, "获取redis分布式锁异常")
//	}
//	tx := global.DB.Begin()
//	//m.Lock() //获取锁 这把锁有问题吗？  假设有10w的并发， 这里并不是请求的同一件商品  这个锁就没有问题了吗？
//
//	//这个时候应该先查询表，然后确定这个订单是否已经扣减过库存了，已经扣减过了就别扣减了
//	//并发时候会有漏洞， 同一个时刻发送了重复了多次， 使用锁，分布式锁
//	sellDetail := model.StockSellDetail{
//		OrderSn: req.OrderSn,
//		Status:  1,
//	}
//	var details []model.GoodsDetail
//	for _, goodInfo := range req.GoodsInfo {
//		details = append(details, model.GoodsDetail{
//			Goods: goodInfo.GoodsId,
//			Num:   goodInfo.Num,
//		})
//
//		var inv model.Inventory
//		//if result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods:goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
//		//	tx.Rollback() //回滚之前的操作
//		//	return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
//		//}
//
//		//for {
//
//		if result := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
//			tx.Rollback() //回滚之前的操作
//			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
//		}
//		//判断库存是否充足
//		if inv.Stocks < goodInfo.Num {
//			tx.Rollback() //回滚之前的操作
//			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
//		}
//		//扣减， 会出现数据不一致的问题 - 锁，分布式锁
//		inv.Stocks -= goodInfo.Num
//		tx.Save(&inv)
//
//		//update inventory set stocks = stocks-1, version=version+1 where goods=goods and version=version
//		//这种写法有瑕疵，为什么？
//		//零值 对于int类型来说 默认值是0 这种会被gorm给忽略掉
//		//if result := tx.Model(&model.Inventory{}).Select("Stocks", "Version").Where("goods = ? and version= ?", goodInfo.GoodsId, inv.Version).Updates(model.Inventory{Stocks: inv.Stocks, Version: inv.Version+1}); result.RowsAffected == 0 {
//		//	zap.S().Info("库存扣减失败")
//		//}else{
//		//	break
//		//}
//		//}
//		//tx.Save(&inv)
//	}
//	sellDetail.Detail = details
//	//写selldetail表
//	if result := tx.Create(&sellDetail); result.RowsAffected == 0 {
//		tx.Rollback()
//		return nil, status.Errorf(codes.Internal, "保存库存扣减历史失败")
//	}
//	tx.Commit() // 需要自己手动提交操作
//
//	if ok, err := mutex.Unlock(); !ok || err != nil {
//		return nil, status.Errorf(codes.Internal, "释放redis分布式锁异常")
//	}
//	//m.Unlock() //释放锁
//	return &emptypb.Empty{}, nil
//}

func (*InventoryServer) Reback(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	//库存归还： 1：订单超时归还 2. 订单创建失败，归还之前扣减的库存 3. 手动归还
	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory
		if result := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
			tx.Rollback() //回滚之前的操作
			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
		}

		//扣减， 会出现数据不一致的问题 - 锁，分布式锁
		inv.Stocks += goodInfo.Num
		tx.Save(&inv)
	}
	tx.Commit() // 需要自己手动提交操作
	return &emptypb.Empty{}, nil
}

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
