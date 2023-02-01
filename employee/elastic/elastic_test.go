package elastic

import (
	"context"
	conf "employee/config"
	"reflect"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func TestPostDataInSearch(t *testing.T) {
	type args struct {
		c      conf.Configuration
		id     string
		data   interface{}
		ctxReq context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostDataInSearch(tt.args.c, tt.args.id, tt.args.data, tt.args.ctxReq)
		})
	}
}

func Test_putDataInSearch(t *testing.T) {
	type args struct {
		jsonData    interface{}
		bulkIndexer esutil.BulkIndexer
		id          string
		ctxReq      context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			putDataInSearch(tt.args.jsonData, tt.args.bulkIndexer, tt.args.id, tt.args.ctxReq)
		})
	}
}

func Test_indexExists(t *testing.T) {
	type args struct {
		c     conf.Configuration
		index string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexExists(tt.args.c, tt.args.index); got != tt.want {
				t.Errorf("indexExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchDataInElastic(t *testing.T) {
	type args struct {
		c      conf.Configuration
		Id     string
		ctxReq context.Context
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchDataInElastic(tt.args.c, tt.args.Id, tt.args.ctxReq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchDataInElastic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchALLDataInElastic(t *testing.T) {
	type args struct {
		c      conf.Configuration
		ctxReq context.Context
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchALLDataInElastic(tt.args.c, tt.args.ctxReq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchALLDataInElastic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckElasticHealth(t *testing.T) {
	type args struct {
		c      conf.Configuration
		ctxReq context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckElasticHealth(tt.args.c, tt.args.ctxReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckElasticHealth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckElasticHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
