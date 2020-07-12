package elastic

import (
	"employee/config"
	"github.com/elastic/go-elasticsearch/v8"
)

func generateElasticClient(conf config.Configuration) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			conf.Elasticsearch.Host,
		},
		Username:      conf.Elasticsearch.Username,
		Password:      conf.Elasticsearch.Password,
		MaxRetries:    15,
		RetryOnStatus: []int{502, 503, 504, 429},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return es, err
	}
	return es, nil
}
