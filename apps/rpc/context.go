package rpc

import (
	"webbk/config"
	"webbk/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func checkAsset(free uint, req uint) bool {
	if free == 0 {
		return false
	}
	return free >= req
}

func HandleCreatCtx(c *gin.Context) {
	var form CreatCtxForm
	err := c.ShouldBindJSON(&form);
	if err != nil {
		c.JSON(400, config.Ret {
			Status: false,
			Data: err.Error(),
		})
		return
	}

	var count int64
	sql := db.Db.Model(&db.Node{}).Where(&db.Node{Model: gorm.Model{ID: form.NodeID}})
	sql.Count(&count)
	if count == 0 {
		c.JSON(400, config.Ret {
			Status: false,
			Data: "Node Not found: NodeID",
		})
		return
	} else if count == 1 {
		var node db.Node
		{
			result := sql.First(&node)
			if result.Error != nil {
				panic(result.Error.Error())
			}
		}
		// TODO: atomic operation
		cond := checkAsset(node.FreeCpuNum, form.Asset.CpuNum)
		cond = cond && checkAsset(node.FreeGpuNum, form.Asset.GpuNum)
		cond = cond && checkAsset(node.FreeCpuMem, form.Asset.CpuMem)
		cond = cond && checkAsset(node.FreeGpuMem, form.Asset.GpuMem)
		if !cond {
			c.JSON(400, config.Ret {
				Status: false,
				Data: "Node Not found: Asset",
			})
			return
		}
		node.FreeCpuNum -= form.Asset.CpuNum
		node.FreeGpuNum -= form.Asset.GpuNum
		node.FreeCpuMem -= form.Asset.CpuMem
		node.FreeGpuMem -= form.Asset.GpuMem
		db.Db.Save(&node)
		ctx := db.RpcCtx{
			NodeID: node.ID,
			CpuNum: form.Asset.CpuNum,
			GpuNum: form.Asset.GpuNum,
			CpuMem: form.Asset.CpuMem,
			GpuMem: form.Asset.GpuMem,
		}
		result := db.Db.Create(&ctx)
		if result.Error != nil {
			panic(result.Error.Error())
		}
		c.JSON(200, config.Ret {
			Status: true,
			Data: CtxRet {
				CtxId: ctx.ID,
				NodeId: node.ID,
				CpuNum: form.Asset.CpuNum,
				GpuNum: form.Asset.GpuNum,
				CpuMem: form.Asset.CpuMem,
				GpuMem: form.Asset.GpuMem,
			},
		})
		return

	} else {
		panic("Repeated Node ID")
	}
}