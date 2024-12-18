package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"shop_srvs/user_srv/global"
	"shop_srvs/user_srv/model"
	"shop_srvs/user_srv/proto"
	"strings"
	"time"
)

type UserService struct{}

func ModelToResponse(user model.User) proto.UserInfoResponse {
	// 在grpc的message中字段有默认值，不能随便赋值nil进去，容易出错
	var userInfoResponse = proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoResponse.Birthday = uint64(user.Birthday.Unix())
	}

	return userInfoResponse
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

func (s *UserService) GetUserList(ctx context.Context, request *proto.PageInfo) (*proto.UserListResponse, error) {
	// 获取用户列表
	var users []model.User
	result := global.DB.Model(&users).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(request.Page), int(request.PageSize))).Find(&users)

	for _, user := range users {
		userInfoResponse := ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoResponse)
	}

	return rsp, nil
}

func (s *UserService) GetUserByMobile(ctx context.Context, request *proto.MobileRequest) (*proto.UserInfoResponse, error) {

	var user model.User

	if request.Mobile == "" {
		return nil, status.Errorf(codes.InvalidArgument, "手机号码不能为空")
	}
	result := global.DB.Where("mobile = ?", request.Mobile).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoResponse := ModelToResponse(user)

	return &userInfoResponse, nil
}

func (s *UserService) GetUserById(ctx context.Context, request *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User

	if request.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ID不能为空")
	}
	result := global.DB.Where("id = ?", request.Id).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoResponse := ModelToResponse(user)

	return &userInfoResponse, nil
}

func (s *UserService) CreateUser(ctx context.Context, request *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {

	var user model.User

	result := global.DB.Where("mobile = ?", request.Mobile).First(&user)
	if result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = request.Mobile
	user.Nickname = request.Nickname

	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodePwd := password.Encode(request.Password, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodePwd)

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	rsp := ModelToResponse(user)

	return &rsp, nil
}

func (s *UserService) UpdateUser(ctx context.Context, request *proto.UpdateUserInfo) (*emptypb.Empty, error) {
	var user model.User

	result := global.DB.Where("id = ?", request.Id).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	birthday := time.Unix(int64(request.Birthday), 0)

	user.Nickname = request.Nickname
	user.Gender = request.Gender
	user.Birthday = &birthday

	res := global.DB.Save(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &empty.Empty{}, nil
}

func (s *UserService) CheckPassword(ctx context.Context, request *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {

	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(request.EncryptedPassword, "$")
	check := password.Verify(request.Password, passwordInfo[2], passwordInfo[3], options)

	return &proto.CheckResponse{
		Success: check,
	}, nil
}
