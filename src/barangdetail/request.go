package barangdetail

type Respons struct {
	Status string      `json:"status, omitempty"`
	Data   interface{} `json:"data_detail"`
	Total  string      `json:"total, omitempty"`
}

type RequestBarangDetail struct {
	Code          string `json:"code"`
	Stock         string `json:"stock"`
	Hp            int64  `json:"harga_pokok"`
	Hj            int64  `json:"harga_jual"`
	Code_kategory int64  `json:"category_code"`
	Name          string `json:"deskripsi"`
	Created       string `json:"created"`
	UserID        int64  `json:"userid"`
}

type UpdateBarangDetail struct {
	Code          string `json:"code"`
	Stock         string `json:"stock"`
	Hp            int64  `json:"harga_pokok"`
	Hj            int64  `json:"harga_jual"`
	Code_kategory int64  `json:"category_code"`
	Name          string `json:"deskripsi"`
	Created       string `json:"created"`
	UserID        int64  `json:"userid"`
}
