package models

type Params struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int64  `json:"age,omitempty"`
}
