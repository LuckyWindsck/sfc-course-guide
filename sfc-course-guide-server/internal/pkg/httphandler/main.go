package httphandler

import (
	// "github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/elasticclient"
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/elasticclient_v7"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO: For development only
func Pretty(c *gin.Context, data interface{}) {
	_, pretty := c.Request.URL.Query()["pretty"]

	if pretty {
		c.IndentedJSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Hello, Gin",
		"message": "SFC COURSE GUIDE SERVER",
	})
}

func GetSearch(c *gin.Context) {
	courses, err := elasticclient_v7.GetAllCourse()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	Pretty(c, courses)

	return
	// query := c.Query("query")
	// pretty := c.Query("pretty")
	// var clientSearchResult elasticclient.ClientSearchResult
	// var err error
	// if query == "" {
	// 	clientSearchResult, err = elasticclient.GetAllCourse()
	// } else {
	// 	clientSearchResult, err = elasticclient.SearchCourse(query)
	// }

	// if err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// }

	// if pretty == "true" {
	// 	c.IndentedJSON(http.StatusOK, clientSearchResult)
	// } else {
	// 	c.JSON(http.StatusOK, clientSearchResult)
	// }
}

func GetCount(c *gin.Context) {
	counts, err := elasticclient_v7.GetAllDocumentCounts()

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	Pretty(c, counts)
	return
}

func GetTest(c *gin.Context) {
	c.String(http.StatusOK, "test")
}
