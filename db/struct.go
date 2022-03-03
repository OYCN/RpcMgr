package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"unique"`
	Passwd string
	Auth uint32
}

type Tag struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"unique"`
	Color string
}

type Node struct {
	gorm.Model
	Name string `gorm:"unique"`
	Tags []Tag `gorm:"foreignKey:Name"`
	Url string
	Port uint
	TotalCpuNum uint
	TotalGpuNum uint
	TotalCpuMem uint
	TotalGpuMem uint
	FreeCpuNum uint
	FreeGpuNum uint
	FreeCpuMem uint
	FreeGpuMem uint
	Priority int
}

type RpcCtx struct {
	gorm.Model
	NodeID uint
	CpuNum uint
	GpuNum uint
	CpuMem uint
	GpuMem uint
	Method string
}
