package nodes

type TagItem struct {
	Name string `json:"name"`
	Color string `json:"color"`
}

type NodeItem struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Tags []TagItem `json:"tags"`
}

type GetNodeForm struct {
	Start	int64 `form:"start" binding:"number"`
	Size	int64 `form:"size" binding:"number"`
}

type Ret struct {
	Amount int64 `json:"amount"`
	Nodes []NodeItem `json:"nodes"`
}

type AddNodeForm struct {
	Name string `json:"name" binding:"required"`
	Tags []string `json:"tags"`
	Url string `json:"url" binding:"required"`
	Port uint `json:"port" binding:"required"`
	CpuNum uint `json:"cpunum" binding:"required"`
	GpuNum uint `json:"gpunum" binding:"required"`
	CpuMem uint `json:"cpumem" binding:"required"`
	GpuMem uint `json:"gpumem" binding:"required"`
}
