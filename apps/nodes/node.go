package nodes

import (
	"webbk/config"
	"webbk/db"

	"github.com/gin-gonic/gin"
)

func HandleAddNode(c *gin.Context) {
	var form AddNodeForm
	err := c.ShouldBindJSON(&form);
	if err != nil {
		c.JSON(400, config.Ret {
			Status: false,
			Data: err.Error(),
		})
		return
	}
	var count int64
	db.Db.Model(&db.Node{}).Where(&db.Node{Name: form.Name}).Count(&count)
	if count != 0 {
		c.JSON(400, config.Ret {
			Status: false,
			Data: "This Node named: `" + form.Name + "` is already exist",
		})
		return
	}
	tagSet := make(map[string] db.Tag)
	for _, tag := range(form.Tags) {
		tagSet[tag] = db.Tag{
			Name: tag,
		}
	}
	tagList := make([]db.Tag, len(tagSet))
	idx := 0
	for _, t := range(tagSet) {
		sql := db.Db.Model(&db.Tag{}).Where(&t)
		var count int64
		sql.Count(&count)
		if count == 0 {
			db.Db.Create(&t)
		} else {
			result := sql.First(&t)
			if result.Error != nil {
				c.JSON(400, config.Ret {
					Status: false,
					Data: result.Error.Error(),
				})
				return
			}
		}
		tagList[idx] = t
		idx += 1
	}
	node := db.Node {
		Name: form.Name,
		Tags: tagList,
		Url: form.Url,
		Port: form.Port,
		TotalCpuNum: form.CpuNum,
		TotalGpuNum: form.GpuNum,
		TotalCpuMem: form.CpuMem,
		TotalGpuMem: form.GpuMem,
		FreeCpuNum: form.CpuNum,
		FreeGpuNum: form.GpuNum,
		FreeCpuMem: form.CpuMem,
		FreeGpuMem: form.GpuMem,
	}
	result := db.Db.Create(&node)
	if result.Error != nil {
		c.JSON(400, config.Ret {
			Status: false,
			Data: result.Error.Error(),
		})
		return
	}
	if result.RowsAffected != 1 {
		panic("Affected rows is not 1, maybe an error")
	}
	ret := config.Ret {
		Status: true,
		Data: node.ID,
	}
	c.JSON(200, ret)
}