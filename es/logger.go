package es

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/olivere/elastic"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"doc":{
			"properties":{
				"package":{
					"type":"keyword"
				},
				"method":{
					"type":"keyword"
				}
				"message":{
					"type":"text"
				},
				"timestamp":{
					"type":"date"
				}
			}
		}
	}
}`

type Logger struct {
	enabled bool
	client  *elastic.Client
}

type Log struct {
	Package   string
	Method    string
	Message   string
	Timestamp time.Time
}

func NewLogger() *Logger {
	// Check if tests are being run
	if flag.Lookup("test.v") != nil {
		return &Logger{false, nil}
	}

	ctx := context.Background()
	esURL := "http" + os.Getenv("ES_HOST") + ":" + os.Getenv("ES_PORT")

	client, err := elastic.NewClient(elastic.SetURL(esURL))
	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(esURL).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(os.Getenv("ES_LOG_INDEX")).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex(os.Getenv("ES_LOG_INDEX")).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	return &Logger{true, client}
}

func (l *Logger) Error(log *Log) {
	if !l.enabled {
		return
	}

}
