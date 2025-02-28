package dto

import "strings"

type InputOrder struct {
	No                int    `json:"no"`
	PlatformProductId string `json:"platformProductId"`
	Qty               int    `json:"qty"`
	UnitPrice         int    `json:"unitPrice"`
	TotalPrice        int    `json:"totalPrice"`
	materialId        string
	modelId           string
	textureId         string
	calPrice          int
}

func (i *InputOrder) ExtractProduct() {
	pd := strings.Split(i.PlatformProductId, "-")
	i.materialId = pd[0] + "-" + pd[1]
	if len(pd) > 3 {
		i.modelId = pd[2] + "-" + pd[3]
	} else {
		i.modelId = pd[2]
	}
	i.calPrice = i.Qty * i.UnitPrice
	i.textureId = pd[1]
}

func (i *InputOrder) GetMaterialId() string {
	return i.materialId
}

func (i *InputOrder) GetModelId() string {
	return i.modelId
}

func (i *InputOrder) Calculate() int {
	return i.calPrice
}

func (i *InputOrder) GetTextureId() string {
	return i.textureId
}
