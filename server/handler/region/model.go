package region

type params struct {
	ParentID uint `form:"parent_id"`
}

type Ack struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Child []Ack  `json:"child"`
}
