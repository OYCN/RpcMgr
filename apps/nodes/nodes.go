package nodes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"

	"webbk/config"
	"webbk/db"
)

var tagList = [...]TagItem{
	{"cpu", "#2db7f5"},
	{"x86", "#2db7f5"},
	{"gpu", "#87d068"},
	{"sm75", "#87d068"},
}

const maxGetSize int64 = config.NodesMaxSizeInOnce

func HandleNodesGetNodeList(c *gin.Context) {
	var args GetNodeForm
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
		ret := config.Ret {
			Status: false,
			Data: "size must be positive",
		}
		c.JSON(400, ret)
		return
	}
	if size < 0 {
		ret := config.Ret {
			Status: false,
			Data: "size must be positive",
		}
		c.JSON(400, ret)
		return
	}
	var ret Ret
	db.Db.Model(&db.Node{}).Count(&ret.Amount)
	if size != 0 {
		if size > maxGetSize {
			size = maxGetSize
		}
		if size > ret.Amount {
			size = ret.Amount
		}
		var nodes []db.Node
		db.Db.Offset(int(start)).Limit(int(size)).Preload(clause.Associations).Find(&nodes)
		ret.Nodes = make([]NodeItem, len(nodes))
		for retsIdx, node := range nodes {
			tags := make([]TagItem, len(node.Tags))
			for tagsIdx, tag := range node.Tags {
				tags[tagsIdx] = TagItem {
					Name: tag.Name,
					Color: tag.Color,
				}
			}
			ret.Nodes[retsIdx] = NodeItem {
				ID: node.ID,
				Name: node.Name,
				Tags: tags,
			}
		}
	}
	fmt.Println(ret)
	c.JSON(200, config.Ret {
		Status: true,
		Data: ret,
	})
}
