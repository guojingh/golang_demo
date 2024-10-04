package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
)

type ClientManager struct {
	client      *elasticsearch.Client
	once        sync.Once
	logDataChan chan interface{}
	index       string
}

var (
	esClient = &ClientManager{}
)

var defaultManager = &ClientManager{}

func (cm *ClientManager) GetClient(addr []string) (*elasticsearch.Client, error) {
	var err error
	var es *elasticsearch.Client
	cm.once.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: addr,
		}
		es, err = elasticsearch.NewClient(cfg)
		if err != nil {
			fmt.Printf("Error creating the client: %s", err)
		}
		cm.client = es
	})
	return cm.client, err
}

func GetDefaultClient(addr []string) (*elasticsearch.Client, error) {
	return defaultManager.GetClient(addr)
}

// Init 将日志数据写入 ElasticSearch
func Init(addr []string, index string, maxSize, goroutineNum int) (err error) {
	client, err := GetDefaultClient(addr)

	esClient.client = client
	esClient.logDataChan = make(chan interface{}, maxSize)

	// 从通道中取出数据，写入到 ES 中
	for i := 0; i < goroutineNum; i++ {
		go sendToEs(index)
	}
	return
}

func sendToEs(index string) {
	for m1 := range esClient.logDataChan {
		fmt.Println(m1)
		b, err := json.Marshal(m1)
		if err != nil {
			fmt.Printf("Error marshalling data: %s", err)
			continue
		}

		err = InsertData(context.Background(), esClient.client, index, b)
		if err != nil {
			fmt.Printf("Error inserting data: %s", err)
			return
		}
	}
}

// PutLogData 通过一个首字母大写的函数从包外接收msg，发送到chan中
func PutLogData(msg interface{}) {
	esClient.logDataChan <- msg
}

// CreateIndex creates an index with the given name and mapping.
func CreateIndex(ctx context.Context, es *elasticsearch.Client, indexName string) error {
	body := map[string]interface{}{
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"name": map[string]interface{}{"type": "text"},
			},
		},
	}

	jsonBody, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling index mapping: %w", err)
	}

	res, err := es.Indices.Create(
		indexName,
		es.Indices.Create.WithContext(ctx),
		es.Indices.Create.WithBody(bytes.NewReader(jsonBody)),
	)
	if err != nil {
		return fmt.Errorf("error creating the index: %w", err)
	}
	defer res.Body.Close()

	return nil
}

// InsertData inserts the given users into the specified index.
func InsertData(ctx context.Context, es *elasticsearch.Client, indexName string, msg []byte) error {
	err := CreateIndex(ctx, es, indexName)
	if err != nil {
		fmt.Printf("Error creating the index: %s", err)
		return err
	}
	res, err := es.Bulk(bytes.NewReader(msg), es.Bulk.WithIndex(indexName))
	fmt.Println("发送数据成功")
	if err != nil {
		return fmt.Errorf("error indexing the documents: %w", err)
	}
	defer res.Body.Close()

	return nil
}
