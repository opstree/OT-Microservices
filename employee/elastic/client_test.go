package elastic

import (
	"employee/config"
	"reflect"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

func Test_generateElasticClient(t *testing.T) {
	type args struct {
		conf config.Configuration
	}
	tests := []struct {
		name    string
		args    args
		want    *elasticsearch.Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateElasticClient(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateElasticClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateElasticClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
