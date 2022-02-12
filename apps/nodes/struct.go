package nodes

type TagItem struct {
	Name string `json:"name"`
	Color string `json:"color"`
}

type NodeItem struct {
	Key uint64 `json:"key"`
	Name string `json:"name"`
	Tags []TagItem `json:"tags"`
}
