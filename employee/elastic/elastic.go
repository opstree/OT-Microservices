package elastic

import (
	"bytes"
	"context"
	conf "employee/config"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/sirupsen/logrus"
	"sync/atomic"
	"time"
)

var (
	res             *esapi.Response
	countSuccessful uint64
	indexName       = "employee-management"
	r               map[string]interface{}
)

// PostDataInSearch method will push data to elasticsearch
func PostDataInSearch(c conf.Configuration, id string, data interface{}) {
	esClient, err := generateElasticClient(c)
	if err != nil {
		logrus.Errorf("Unable to create client connection with elasticsearch: %v", err)
	}
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         indexName,
		Client:        esClient,
		NumWorkers:    1,
		FlushBytes:    int(5e+6),
		FlushInterval: 30 * time.Second,
	})

	if err != nil {
		logrus.Errorf("Error creating the bulk indexer: %v", err)
	}
	checkIndexState := indexExists(c, indexName)

	if checkIndexState != 200 {
		res, err = esClient.Indices.Create(indexName)
		res.Body.Close()
		if err != nil {
			logrus.Errorf("Cannot create the employee index: %v, %v", err, res)
		}
	}

	putDataInSearch(data, bulkIndexer, id)
	if err := bulkIndexer.Close(context.Background()); err != nil {
		logrus.Errorf("Unexpected error: While closing bulk indexing: %v", err)
	}
	logrus.Infof("Successfully pushed employee's information in elasticsearch")
}

func putDataInSearch(jsonData interface{}, bulkIndexer esutil.BulkIndexer, id string) {
	data, err := json.Marshal(jsonData)
	if err != nil {
		logrus.Errorf("Cannot marshal data into JSON: %v", err)
	}
	err = bulkIndexer.Add(
		context.Background(),
		esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: id,
			Body:       bytes.NewReader(data),
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
				atomic.AddUint64(&countSuccessful, 1)
			},
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
				if err != nil {
					logrus.Errorf("ERROR: Performing bulk indexing: %v", err)
				} else {
					logrus.Errorf("ERROR: Performing bulk indexing: %v", err)
				}
			},
		},
	)
	if err != nil {
		logrus.Errorf("Unexpected error: Bulk operation over index: %v", err)
	}
}

func indexExists(c conf.Configuration, index string) int {
	esClient, err := generateElasticClient(c)
	if err != nil {
		logrus.Errorf("Unable to create client connection with elastic: %v", err)
	}
	resp, err := esClient.Indices.Exists([]string{indexName})
	if err != nil {
		logrus.Errorf("Unexpected error: while checking index: %v", err)
		return 404
	}
	return resp.StatusCode
}

// SearchDataInElastic will search data in elasticsearch
func SearchDataInElastic(c conf.Configuration, Id string) map[string]interface{} {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": Id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		logrus.Errorf("Error encoding query: %v", err)
	}
	es, err := generateElasticClient(c)

	if err != nil {
		logrus.Errorf("Unable to create client connection with elastic: %v", err)
	}
	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			logrus.Errorf("Error parsing the response body: %v", err)
		} else {
			// Print the response status and error information.
			logrus.Errorf("[%v] %v: %v",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logrus.Errorf("Error parsing the response body: %v", err)
	}
	return r
}

// SearchALLDataInElastic will search all data in elasticsearch
func SearchALLDataInElastic(c conf.Configuration) map[string]interface{} {
	var buf bytes.Buffer
	es, err := generateElasticClient(c)

	if err != nil {
		logrus.Errorf("Unable to create client connection with elastic: %v", err)
	}
	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			logrus.Errorf("Error parsing the response body: %v", err)
		} else {
			// Print the response status and error information.
			logrus.Errorf("[%v] %v: %v",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logrus.Errorf("Error parsing the response body: %v", err)
	}
	return r
}

// CheckElasticHealth is a method to check elasticsearch health
func CheckElasticHealth(c conf.Configuration) (bool, error) {
	es, err := generateElasticClient(c)
	if err != nil {
		logrus.Errorf("Unable to create client connection with elastic: %v", err)
	}

	_, err = es.Info()
	if err != nil {
		return false, err
	}

	return true, nil
}
