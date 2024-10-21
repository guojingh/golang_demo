package main

import (
	"fmt"
	"oneofdemo/api"

	"github.com/golang/protobuf/protoc-gen-go/generator"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// oneof 实例
func oneofDemo() {
	// client
	/*	req := &api.NoticeReaderRequest{
		Msg: "李文周的博客更新了",
		NoticeWay: &api.NoticeReaderRequest_Email{
			Email: "123@xxx.com",
		},
	}*/

	req := &api.NoticeReaderRequest{
		Msg: "李文周的博客更新了",
		NoticeWay: &api.NoticeReaderRequest_Phone{
			Phone: "15536983501",
		},
	}

	// server
	// 类型断言
	switch v := req.NoticeWay.(type) {
	case *api.NoticeReaderRequest_Email:
		noticeWithEmail(v)
	case *api.NoticeReaderRequest_Phone:
		noticeWithPhone(v)
	}
}

// 使用 google/protobuf/wrappers.proto
/*func wrapValueDemo() {
	//client
	book := api.Book{
		Title: "《学习go语言》",
		Price: &wrapperspb.Int64Value{Value: 9900},
		Memo:  &wrapperspb.StringValue{Value: "学就完事了"},
	}

	if book.Price == nil { // 没有给 price 赋值
		fmt.Println("没有设置price")
	} else {
		// 赋值了放心大胆的去用
		fmt.Println(book.GetPrice().GetValue())
	}

	if book.GetMemo() != nil {
		fmt.Println(book.GetMemo().GetValue())
	}
}*/

func optionalDemo() {
	//client
	book := api.Book{
		Title: "《学习微服务》",
		Price: proto.Int64(9900),
	}

	// server
	// 如何判断 book.Price 有没有被赋值呢？
	if book.Price == nil {
		fmt.Println("no price")
	} else {
		fmt.Printf("book with price:%v\n", book.GetPrice())
	}
}

// field 使用 field_mask 实现部分更新实例
func fieldMaskDemo() {
	// client
	paths := []string{"price", "info.b", "author"} // 更新的字段信息
	req := &api.UpdateBookRequest{
		Op: "q1mi",
		Book: &api.Book{
			Author: "七米2号",
			Price:  proto.Int64(8800),
			Info: &api.Book_Info{
				B: "bbbbb",
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths}, // 提供情报（哪些字段更新了）
	}

	// Server
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(req.UpdateMask, generator.CamelCase)
	var bookDst = make(map[string]interface{})
	// 将数据读取到 map[string]interface{}
	// fieldmask-utils 支持读取到结构体，更多用法可查看文档
	fieldmask_utils.StructToMap(mask, req.Book, bookDst)
	fmt.Printf("bookDst:%#v\n", bookDst)
}

func noticeWithEmail(in *api.NoticeReaderRequest_Email) {
	fmt.Printf("notice reader by email:%v\n", in.Email)
}

func noticeWithPhone(in *api.NoticeReaderRequest_Phone) {
	fmt.Printf("notice reader by phone:%v\n", in.Phone)
}

func main() {
	//oneofDemo()
	//wrapValueDemo()
	//optionalDemo()
	fieldMaskDemo()
}
