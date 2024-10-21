package main

import (
	"bookstore/pb"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

const (
	defaultCursor   = "0" // 默认游标
	defaultPageSize = 2   // 默认每一页显示数量
)

// bookstore grpc 服务

type server struct {
	pb.UnimplementedBookstoreServer

	bs *bookstore //data.go
}

// ListShelves 列出所有书架的 RPC 方法
func (s *server) ListShelves(ctx context.Context, in *emptypb.Empty) (*pb.ListShelvesResponse, error) {
	// 调用 orm 操作的那些方法
	sl, err := s.bs.GetShelfList(ctx)
	if errors.Is(err, gorm.ErrEmptySlice) {
		// 没有数据
		return &pb.ListShelvesResponse{}, nil
	}

	if err != nil { // 查询数据库失败
		return nil, status.Error(codes.Internal, "query failed")
	}

	// 封装返回数据
	nsl := make([]*pb.Shelf, 0, len(sl))
	for _, s := range sl {
		nsl = append(nsl, &pb.Shelf{
			Id:    s.ID,
			Theme: s.Theme,
			Size:  s.Size,
		})
	}

	return &pb.ListShelvesResponse{Shelves: nsl}, nil
}

// CreateShelf 创建书架
func (s *server) CreateShelf(ctx context.Context, in *pb.CreateShelfRequest) (*pb.Shelf, error) {
	// 参数检查
	if len(in.GetShelf().GetTheme()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid theme")
	}

	data := Shelf{
		Theme: in.GetShelf().Theme,
		Size:  in.GetShelf().Size,
	}

	// 数据库创建
	ns, err := s.bs.CreateShelf(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, "create failed")
	}

	return &pb.Shelf{Id: ns.ID, Theme: ns.Theme, Size: ns.Size}, nil
}

func (s *server) GetShelf(ctx context.Context, in *pb.GetShelfRequest) (*pb.Shelf, error) {
	// 参数 check
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}

	// 查询数据库
	shelf, err := s.bs.GetShelf(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "query failed")
	}

	// 封装返回数据
	return &pb.Shelf{Id: shelf.ID, Theme: shelf.Theme, Size: shelf.Size}, nil
}

func (s server) DeleteShelf(ctx context.Context, in *pb.DeleteShelfRequest) (*emptypb.Empty, error) {
	// 参数 check
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}

	err := s.bs.DeleteShelf(ctx, in.Shelf)
	if err != nil {
		return nil, status.Error(codes.Internal, "delete failed")
	}

	return &emptypb.Empty{}, nil
}

// ListBooks 列出书架的所有图书
func (s *server) ListBooks(ctx context.Context, in *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	// 参数 check
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	var (
		cursor   = defaultCursor
		pageSize = defaultPageSize
	)

	if len(in.GetPageToken()) > 0 {
		// 有分页的话就先解析分页数据
		pageInfo := Token(in.GetPageToken()).Decode()
		// 再判断解析结果是否有效
		if pageInfo.InValid() {
			return nil, status.Error(codes.InvalidArgument, "invalid page_token")
		}
		cursor = pageInfo.NextID
		pageSize = int(pageInfo.PageSize)
	}

	// 查询数据库,基于游标实现分页
	bookList, err := s.bs.GetBookListByShelfID(ctx, in.GetShelf(), cursor, pageSize+1)
	if err != nil {
		fmt.Printf("GetBookListByShelfID failed, err:%v\n", err)
		return nil, status.Error(codes.Internal, "query failed")
	}

	var (
		hasNextPage   bool
		nextPageToken string
		realSize      = len(bookList)
	)

	// 当查询数据库的结果数大于pageSize
	if len(bookList) > pageSize {
		hasNextPage = true  //1.有下一页
		realSize = pageSize //2.下面格式化数据没必要把所有数据都返回，只需要返回pageSize的数据
	}
	// 封装返回的数据
	// 将 []*book ---> []*pb.Books
	res := make([]*pb.Book, 0, realSize)
	for i := 0; i < realSize; i++ {
		res = append(res, &pb.Book{
			Id:     bookList[i].ID,
			Author: bookList[i].Author,
			Title:  bookList[i].Title,
		})
	}

	// 如果有下一页，就要生成下一页的 page_token
	if hasNextPage {
		nextPageInfo := Page{
			NextID:        strconv.FormatInt(res[realSize-1].Id, 10), //最后一个返回结果的ID
			NextTimeAtUTC: time.Now().Unix(),
			PageSize:      int64(pageSize),
		}
		nextPageToken = string(nextPageInfo.Encode())
	}
	return &pb.ListBooksResponse{Books: res, NextPageToken: nextPageToken}, nil
}
