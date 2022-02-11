package nodes

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

type tagItem struct {
	Name string `json:"name"`
	Color string `json:"color"`
}

type nodeItem struct {
	Key uint64 `json:"key"`
	Name string `json:"name"`
	Tags []tagItem `json:"tags"`
}

var tagList = [...]tagItem{
	{"cpu", "#2db7f5"},
	{"x86", "#2db7f5"},
	{"gpu", "#87d068"},
	{"sm75", "#87d068"},
}

func HandleNodesGetNodeList(c *gin.Context) {
	start, startErr := strconv.Atoi(c.Query("start"))
	step, stepErr := strconv.Atoi(c.Query("step"))
	if startErr != nil && stepErr != nil {
		c.String(503,
			"Err to parse input to Int:\n" +
			"start: " + c.Query("start") + "\n" +
			"step: "  + c.Query("start"))
		return
	}
	if start + step > 17 {
		step = 0
	}
	rets := make([]nodeItem, step)
	for retsIdx, _ := range rets {
		tags := make([]tagItem, 1)
		for tagsIdx, _ := range tags {
			tags[tagsIdx] = tagList[rand.Intn(len(tagList))]
		}
		rets[retsIdx] = nodeItem {
			Key: uint64(start + retsIdx),
			Name: strconv.Itoa(int(start + retsIdx)),
			Tags: tags,
		}
	}
	fmt.Println(rets)
	c.JSON(200, rets)
}
