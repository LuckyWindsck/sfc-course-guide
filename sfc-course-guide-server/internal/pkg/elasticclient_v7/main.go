package elasticclient_v7

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)

const (
	esURL = "http://elastic:changeme@elasticsearch:9200"
)

// Search Result
type Hit interface {
}

type HitStat struct {
	Total   int64
	Latency int64 // Unit: milliseconds
}

type ClientSearchResult struct {
	Query string `json:",omitempty"`
	Stat  HitStat
	Hits  []Hit
}

func initClient() (client *elastic.Client, err error) {
	client, err = elastic.NewClient(elastic.SetURL(esURL))
	if err != nil {
		fmt.Println(err)
		fmt.Println("connect es error")
		return
	}
	fmt.Println("connected successfully")
	return
}

var client, err = initClient()
var ctx = context.Background()

func sourceQueryString(query elastic.Query) (queryString interface{}, err error) {
	queryString, err = query.Source()
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func countDocument(index string) (count int64, err error) {
	count, err = client.Count(index).Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func GetAllDocumentCounts() (counts map[string]int64, err error) {
	indexNames, err := client.IndexNames()
	if err != nil {
		fmt.Println(err)
		return
	}

	counts = make(map[string]int64)
	for _, indexName := range indexNames {
		// Skip system index
		if strings.HasPrefix(indexName, ".") {
			continue
		}

		count, err := client.Count(indexName).Do(ctx)
		if err != nil {
			fmt.Println(err)
			continue
		}

		counts[indexName] = count
	}

	return
}

func GetAllCourse() (clientSearchResult ClientSearchResult, err error) {
	// func GetAllCourse() (searchResult interface{}, err error) {
	query := elastic.NewBoolQuery()
	query = query.Filter(
		elastic.NewTermQuery("year", 2019),
		elastic.NewTermQuery("summary_flag", true),
	)

	searchResult, err := client.
		Search("class_info").
		Query(query).
		From(0).Size(10).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	clientSearchResult.Stat.Latency = searchResult.TookInMillis
	clientSearchResult.Stat.Total = searchResult.TotalHits()

	for _, hit := range searchResult.Hits.Hits {
		clientSearchResult.Hits = append(clientSearchResult.Hits, hit.Source)
	}

	return
}
