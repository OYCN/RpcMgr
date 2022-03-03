package rpc

type AssetType struct {
	CpuNum uint `json:"cpunum" binding:"number"`
	GpuNum uint `json:"gpunum" binding:"number"`
	CpuMem uint `json:"cpumem" binding:"number"`
	GpuMem uint `json:"gpumem" binding:"number"`
}

type CreatCtxForm struct {
	NodeID uint  `json:"nodeid" binding:"number"`
	Asset AssetType  `json:"asset"`
}

type CtxRet struct {
	CtxId uint `json:"ctxid"`
	NodeId uint `json:"nodeid"`
	CpuNum uint `json:"cpunum"`
	GpuNum uint `json:"gpunum"`
	CpuMem uint `json:"cpumem"`
	GpuMem uint `json:"gpumem"`
}
