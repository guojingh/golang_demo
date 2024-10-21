package main

import (
	"bookstore/pb"
	"context"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestServer_ListBooks(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("connect database err: %v", err)
	}

	s := server{bs: &bookstore{db: db}}

	req := &pb.ListBooksRequest{
		Shelf: 3,
	}
	res, err := s.ListBooks(context.Background(), req)
	if err != nil {
		t.Fatalf("ListBooks err: %v", err)
	}

	t.Logf("next_page_token: %v", res.NextPageToken)

	for i, book := range res.Books {
		t.Logf("%d:%#v\n", i, book)
	}
}
