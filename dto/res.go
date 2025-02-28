package dto

type CleanedOrder struct {
	No         int    `json:"no"`
	ProductId  string `json:"productId"`
	MaterialId string `json:"materialId"`
	ModelId    string `json:"modelId"`
	Qty        int    `json:"qty"`
	UnitPrice  int    `json:"unitPrice"`
	TotalPrice int    `json:"totalPrice"`
}
