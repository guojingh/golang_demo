package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
)

type ClientManager struct {
	client *elasticsearch.Client
	once   sync.Once
}

var defaultManager = &ClientManager{}

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func (cm *ClientManager) GetClient() *elasticsearch.Client {
	cm.once.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: []string{"http://172.16.56.129:9200"},
		}
		es, err := elasticsearch.NewClient(cfg)
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}
		cm.client = es
	})

	return cm.client
}

func GetDefaultClient() *elasticsearch.Client {
	return defaultManager.GetClient()
}

// CreateIndex creates an index with the given name and mapping.
func CreateIndex(ctx context.Context, es *elasticsearch.Client, indexName string) error {
	body := map[string]interface{}{
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"id":   map[string]interface{}{"type": "integer"},
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

// DeleteIndex deletes the specified index.
func DeleteIndex(ctx context.Context, es *elasticsearch.Client, indexName string) error {
	res, err := es.Indices.Delete(
		[]string{indexName},
		es.Indices.Delete.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("error deleting the index: %w", err)
	}
	defer res.Body.Close()

	return nil
}

// InsertData inserts the given users into the specified index.
func InsertData(ctx context.Context, es *elasticsearch.Client, indexName string, users []User) error {
	var buf strings.Builder
	for _, user := range users {
		meta := []byte(fmt.Sprintf(`{"index":{"_id":"%d"}}%s`, user.ID, "\n"))
		data, err := json.Marshal(user)
		if err != nil {
			return fmt.Errorf("error encoding the user: %w", err)
		}
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
		buf.WriteByte('\n')
	}

	res, err := es.Bulk(strings.NewReader(buf.String()), es.Bulk.WithIndex(indexName))
	if err != nil {
		return fmt.Errorf("error indexing the documents: %w", err)
	}
	defer res.Body.Close()

	return nil
}

// SearchData performs a search on the specified index using the given query and returns the results.
func SearchData(ctx context.Context, es *elasticsearch.Client, indexName string, query string) ([]SearchResult, error) {
	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting response: %w", err)
	}
	defer res.Body.Close()

	var r struct {
		Hits struct {
			Hits []struct {
				ID     string          `json:"_id"`
				Source json.RawMessage `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %w", err)
	}

	results := make([]SearchResult, len(r.Hits.Hits))
	for i, hit := range r.Hits.Hits {
		if err := json.Unmarshal(hit.Source, &results[i]); err != nil {
			return nil, fmt.Errorf("error unmarshaling search result %d: %w", i, err)
		}
	}

	return results, nil
}

// UpdateData updates the specified index with the provided update data.
/*func UpdateData(ctx contextdemo.Context, es *elasticsearch.Client, indexName string, updateData []UpdateResult) error {
	var buf strings.Builder
	for _, update := range updateData {
		meta := []byte(fmt.Sprintf(`{"update":{"_id":"%s"}}%s`, update.ID, "\n"))
		data, err := json.Marshal(map[string]map[string]string{"doc": {"name": update.Name}})
		if err != nil {
			return fmt.Errorf("error encoding the update data: %w", err)
		}
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
		buf.WriteByte('\n')
	}

	res, err := esutil.Bulk(
		ctx,
		es.Bulk.WithIndex(indexName),
		es.Bulk.WithBody(strings.NewReader(buf.String())),
		es.Bulk.WithRefresh("true"),
	)
	if err != nil {
		return fmt.Errorf("error updating the data: %w", err)
	}
	defer res.Body.Close()

	return nil
}*/

// SearchResult represents a search result from Elasticsearch.
type SearchResult struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	client := GetDefaultClient()

	if err := CreateIndex(context.Background(), client, "web2"); err != nil {
		fmt.Printf("Error creating index: %s", err)
		return
	}

	users := []User{
		{
			ID:      0,
			Name:    "小張",
			Age:     19,
			Married: false,
		},
		{
			ID:      1,
			Name:    "小李",
			Age:     29,
			Married: true,
		},
	}
	err := InsertData(context.Background(), client, "test", users)
	if err != nil {
		fmt.Printf("error inserting data: %v", err)
	}
}
