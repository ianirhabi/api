package barang

type Requestbarang struct {
	Itemcategory string `json:"item_category"`
	CodeITEM     string `json:"code_item"`
	Created      string `json:"created"`
	UserId       int64  `json:"user_id"`
}

type Respons struct {
	Status string      `json:"status, omitempty"`
	Data   interface{} `json:"data"`
	Total  int         `json:"total, omitempty"`
}
