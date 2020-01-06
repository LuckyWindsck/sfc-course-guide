package elasticclient_v7

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/olivere/elastic/v7"
	"os"
	"strings"
)

type Course struct { // Example Course
	Title   string
	Faculty string
}

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
	err = godotenv.Load()
	if err != nil {
		panic("loading env file failed")
	}

	elasticUser := os.Getenv("ELASTIC_USER")
	elasticPassword := os.Getenv("ELASTIC_PASSWORD")

	esURL := "http://" + elasticUser + ":" + elasticPassword + "@elasticsearch:9200"

	client, err = elastic.NewClient(elastic.SetURL(esURL))
	if err != nil {
		fmt.Println(err)
		panic("connect elasticsearch error")
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

func GetTest() interface{} {
	boolQuery := elastic.NewBoolQuery()
	boolQuery = boolQuery.Filter(
		elastic.NewTermQuery("year", 2019),
		elastic.NewTermQuery("summary_flag", true),
	)

	classInfos, err := client.
		Search("class_info").
		Query(boolQuery).
		From(0).Size(10).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var ids []string
	for _, classInfo := range classInfos.Hits.Hits {
		var class struct {
			FacultyID string `json:"faculty_id"`
		}
		err = json.Unmarshal(classInfo.Source, &class)
		if err != nil {
			fmt.Println("error:", err)
		}

		if class.FacultyID != "" {
			ids = append(ids, class.FacultyID)
		}
	}

	var names []string
	for _, id := range ids {
		matchQuery := elastic.NewMatchQuery("id", id)
		faculties, err := client.
			Search("faculty").
			Query(matchQuery).
			Do(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		var teacher struct {
			Name string `json:"name"`
		}

		numTeacher := len(faculties.Hits.Hits)
		if numTeacher > 1 {
			panic("Many teachers can be found with this ID")
		}
		if numTeacher == 0 {
			panic("No teacher is found with this ID")
		}

		err = json.Unmarshal(faculties.Hits.Hits[0].Source, &teacher)
		if err != nil {
			fmt.Println("error:", err)
		}

		names = append(names, teacher.Name)
	}

	dict := make(map[string]string)
	for i, id := range ids {
		dict[id] = names[i]
	}

	var courseList []Course

	for _, classInfo := range classInfos.Hits.Hits {
		cInfo := classInfo.Source
		var temp struct {
			Title     string `json:"title"`
			FacultyID string `json:"faculty_id"`
		}
		err = json.Unmarshal(cInfo, &temp)
		if err != nil {
			fmt.Println("error:", err)
		}

		var c Course
		c.Title = temp.Title
		c.Faculty = dict[temp.FacultyID]
		courseList = append(courseList, c)
	}

	return courseList
}
