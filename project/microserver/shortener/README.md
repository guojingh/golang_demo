# 短链接项目


## 搭建项目的骨架

1. 建库建表

新建发号器表
```sql
CREATE TABLE `sequence` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `stub` varchar(1) NOT NULL ,
    `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_uniq_stub` (`stub`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT = '序号表';

```

新建长链接短链接映射表：
```sql
CREATE TABLE `short_url_map` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建者',
    `is_del` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除: 0正常1删除',

    `lurl` varchar(2048) DEFAULT NULL COMMENT '长链接',
    `md5` char(32) DEFAULT NULL COMMENT '长链接MD5',
    `surl` varchar(32) DEFAULT NULL COMMENT '短链接',
    PRIMARY KEY (`id`),
    INDEX (`is_del`),
    UNIQUE (`md5`),
    UNIQUE (`surl`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT = '长短链映射表';

```

2. 搭建 go-zero框架的骨架

2.1 编写 `api` 文件，使用 goctl 命令生成代码

```go
syntax = "v1"

/* 短链接项目
 * author: guojinghu   
*/

type ConvertRequest {
    LongUrl string `json:"longUrl"`
}

type ConvertResponse {
    ShortUrl string `json:"shortUrl"`
}

type ShowRequest {
    ShortUrl string `shortUrl`
}

type ShowResponse {
    LongUrl string `json:"longUrl"`
}

service shortener-api {

    @handler ConvertHandler
    post /convert(ConvertRequest) returns(ConvertResponse)
    
    @handler ShowHandler
    get /:shortUrl(ShowRequest) returns(ShowResponse)
    
}


```

2.2 根据 api 文件生成代码

```bash
goctl api go -api .\shortener.api -dir .
```

3. 根据数据表生成model层代码
```bash
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/test" -table="short_url_map" -dir="./model"

goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/test" -table="sequence" -dir="./model"

```

4. 下载项目依赖
```bash
go mod tidy
```

5. 运行项目
```bash
go run shortener.go
```


6. 修改配置结构体和配置文件
注意：两边一定一定要对齐

## 参数校验

1. go_zero 使用 validator
https://pkg.go.dev/github.com/go-playground/validator/v10

下载依赖
```bash
go get github.com/go-playground/validator/v10
```

导入依赖
```bash
import "github.com/go-playground/validator/v10"
```

在api中为结构体添加 validate tag 并添加校验规则
```go
type ConvertRequest {
	LongUrl string `json:"longUrl" validate:"required"`
}

type ConvertResponse {
	ShortUrl string `json:"shortUrl"`
}

type ShowRequest {
	ShortUrl string `json:"shortUrl" validate:"required"`
}

type ShowResponse {
	LongUrl string `json:"longUrl"`
}
```

## go 单元测试
```bash
go test -run TestGetBasePath .\pkg\urltool\ -v


## 跑项目底下全部测试用例
go test ./...
```

第三方测试库推荐使用
go get github.com/smartystreets/goconvey@latest

使用示例
```go
func TestGet(t *testing.T) {
	convey.Convey("基础用例", t, func() {
		url := "https://www.liwenzhou.com/posts/Go/golang-menu/"
		get := Get(url)

		// 断言
		convey.So(get, convey.ShouldEqual, true)
		//convey.ShouldBeTrue(get)
	})

	convey.Convey("url请求不通示例", t, func() {
		url := "posts/go/unit-test-s"
		get := Get(url)

		// 断言
		convey.So(get, convey.ShouldEqual, false)
		//convey.ShouldBeTrue(get)
	})
}
```


## 查看短链接

### 缓存版

有两种方式
1. 使用自己实现的缓存， surl -> lurl，能够节省缓存空间，缓存的数据量小
2. 使用go-zero自带的缓存，surl -> 数据行，不需要自己实现，开发量小

这里使用第二种方式：
1. 添加缓存配置
    - 配置文件
    - 配置结构体

2. 删除旧的model层代码（新版本不需要）
    - 删除 shorturlmapmodel.go
3. 重新生成model层代码
```bash
goc
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/test" -table="short_url_map" -dir="./model" -c
```

4. 修改svccontext代码