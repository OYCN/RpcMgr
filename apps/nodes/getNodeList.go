package nodes

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"

	"webbk/config"
)

var tagList = [...]TagItem{
	{"cpu", "#2db7f5"},
	{"x86", "#2db7f5"},
	{"gpu", "#87d068"},
	{"sm75", "#87d068"},
}

const maxGetNum int = 100
const maxGetSize int = config.NodesMaxSizeInOnce

type argsTemp struct {
	Start	int `form:"start" binding:"required"`
	Size	int `form:"size" binding:"required"`
}

func HandleNodesGetNodeList(c *gin.Context) {
	var args argsTemp
	if err := c.ShouldBindQuery(&args); err != nil {
		ret := config.Ret {
			Status: false,
			Data: err.Error(),
		}
		c.JSON(400, ret)
		return
	}
	start := args.Start
	size := args.Size
	if start < 0 {
		start = maxGetNum + start
	}
	if size < 0 {
		ret := config.Ret {
			Status: false,
			Data: "size must be positive",
		}
		c.JSON(400, ret)
		return
	}
	if size > maxGetSize {
		ret := config.Ret {
			Status: false,
			Data: "Exceeded maximum step",
		}
		c.JSON(400, ret)
		return
	}
	rets := make([]NodeItem, size)
	for retsIdx, _ := range rets {
		tags := make([]TagItem, 1)
		for tagsIdx, _ := range tags {
			tags[tagsIdx] = tagList[rand.Intn(len(tagList))]
		}
		rets[retsIdx] = NodeItem {
			Key: uint64(start + retsIdx),
			Name: strconv.Itoa(int(start + retsIdx)),
			Tags: tags,
		}
	}
	fmt.Println(rets)
	ret := config.Ret {
		Status: true,
		Data: rets,
	}
	c.JSON(200, ret)
}
