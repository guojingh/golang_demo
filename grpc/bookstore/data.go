package main

import (
	"context"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	SHELF_DEFAULT_SIZE = 100
)

// 使用 GORM

func NewDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	err = db.AutoMigrate(&Shelf{}, &Book{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 定义Model

// Shelf 书架
type Shelf struct {
	ID       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}

// Book 图书
type Book struct {
	ID       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}

// 数据库操作
type bookstore struct {
	db *gorm.DB
}

// GetShelfList 查询书架列表
func (b *bookstore) GetShelfList(ctx context.Context) ([]*Shelf, error) {
	var vl []*Shelf
	err := b.db.Debug().WithContext(ctx).Find(&vl).Error
	if errors.Is(err, gorm.ErrEmptySlice) {
		return vl, errors.New("查询数据为空")
	}

	if err != nil {
		return vl, err
	}

	return vl, nil
}

// CreateShelf 创建书架
func (b *bookstore) CreateShelf(ctx context.Context, data Shelf) (*Shelf, error) {
	// 参数校验
	if len(data.Theme) <= 0 {
		return nil, errors.New("shelf theme param error")
	}

	var size int64
	if data.Size <= 0 {
		size = int64(SHELF_DEFAULT_SIZE)
	}

	// 创建书架
	v := Shelf{Theme: data.Theme, Size: size, CreateAt: time.Now(), UpdateAt: time.Now()}
	err := b.db.Debug().WithContext(ctx).Create(&v).Error

	if err != nil {
		return &v, err
	}

	return &v, nil
}

// GetShelf 获取书架
func (b *bookstore) GetShelf(ctx context.Context, id int64) (*Shelf, error) {
	v := Shelf{}
	err := b.db.Debug().WithContext(ctx).First(&v, id).Error
	return &v, err
}

// DeleteShelf 删除书架
func (b *bookstore) DeleteShelf(ctx context.Context, id int64) error {
	return b.db.Debug().WithContext(ctx).Delete(&Shelf{}, id).Error
}

// GetBookListByShelfID 根据书架id查询图书
func (b *bookstore) GetBookListByShelfID(ctx context.Context, shelfId int64, cursor string, pageSize int) ([]*Book, error) {
	var vl []*Book

	err := b.db.Debug().WithContext(ctx).Where("shelf_id = ? and id > ?", shelfId, cursor).Order("id asc").Limit(pageSize).Find(&vl).Error
	return vl, err
}
