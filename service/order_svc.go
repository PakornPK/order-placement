package service

import (
	"strconv"
	"strings"
)

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

func (i *InputOrder) extractProduct() {
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

func (i *InputOrder) getMaterialId() string {
	return i.materialId
}

func (i *InputOrder) getModelId() string {
	return i.modelId
}

func (i *InputOrder) calculate() int {
	return i.calPrice
}

func (i *InputOrder) getTextureId() string {
	return i.textureId
}

type CleanedOrder struct {
	No         int    `json:"no"`
	ProductId  string `json:"productId"`
	MaterialId string `json:"materialId"`
	ModelId    string `json:"modelId"`
	Qty        int    `json:"qty"`
	UnitPrice  int    `json:"unitPrice"`
	TotalPrice int    `json:"totalPrice"`
}

type OrderService interface {
	PlaceOrder(orders []InputOrder) ([]CleanedOrder, error)
}

type orderService struct{}

func NewOrderService() OrderService {
	return orderService{}
}

func (s orderService) PlaceOrder(orders []InputOrder) ([]CleanedOrder, error) {
	var res []CleanedOrder
	currentNo := 1
	cleaner := make(map[string]int, 0)
	wiping := make(map[string]int, 0)
	inputs := prepareInput(orders)
	for _, order := range inputs {
		order.extractProduct()
		res = append(res, CleanedOrder{
			No:         currentNo,
			ProductId:  order.PlatformProductId,
			MaterialId: order.getMaterialId(),
			ModelId:    order.getModelId(),
			Qty:        order.Qty,
			UnitPrice:  order.UnitPrice,
			TotalPrice: order.calculate(),
		})
		currentNo += 1
		wiping["WIPING-CLOTH"] += order.Qty
		cleaner[order.getTextureId()+"-CLEANNER"] += order.Qty
	}

	for k, v := range wiping {
		res = append(res, CleanedOrder{
			No:         currentNo,
			ProductId:  k,
			Qty:        v,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		currentNo += 1
	}

	for k, v := range cleaner {
		res = append(res, CleanedOrder{
			No:         currentNo,
			ProductId:  k,
			Qty:        v,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		currentNo += 1
	}

	return res, nil
}

func prepareInput(inputs []InputOrder) []InputOrder {
	var newInputs []InputOrder
	mapItems := make(map[string]int)
	for _, v := range inputs {
		cleanProductId(&v)
		pdIds := strings.Split(v.PlatformProductId, "/")
		size := len(pdIds)
		if size == 1 {
			break
		}
		for _, val := range pdIds {
			for index, c := range val {
				if c >= 65 && c <= 90 {
					val = val[index:]
					break
				}
			}
			mapItems[val] += v.Qty
		}
	}

	for k := range mapItems {
		pd := strings.Split(k, "-")
		model := strings.Split(pd[2], "*")
		if len(model) != 1 {
			qty, _ := strconv.Atoi(model[1])
			mapItems[k] = qty
		}
	}

	var qty int
	for _, v := range mapItems {
		qty += v
	}

	for _, v := range inputs {
		pdIds := strings.Split(v.PlatformProductId, "/")
		size := len(pdIds)
		if size == 1 {
			starAt := strings.Index(v.PlatformProductId, "*")
			if starAt > 0 {
				qty, _ := strconv.Atoi(v.PlatformProductId[starAt+1:])
				v.Qty = qty
				unitPrice := v.TotalPrice / v.Qty
				v.UnitPrice = unitPrice
				v.PlatformProductId = v.PlatformProductId[:starAt]
			}
			cleanProductId(&v)
			newInputs = append(newInputs, v)
			continue
		}
		for i, p := range pdIds {
			for index, c := range p {
				if c >= 65 && c <= 90 {
					p = p[index:]
					break
				}
			}
			before := p
			starAt := strings.Index(p, "*")
			if starAt > 0 {
				p = p[:starAt]
			}

			unitPrice := v.UnitPrice / qty
			totalPrice := unitPrice * mapItems[before]
			tmp := InputOrder{
				No:                i + 1,
				PlatformProductId: p,
				Qty:               mapItems[before],
				UnitPrice:         unitPrice,
				TotalPrice:        totalPrice,
			}

			cleanProductId(&tmp)
			newInputs = append(newInputs, tmp)
		}
	}

	return newInputs
}

func cleanProductId(v *InputOrder) {
	for index, c := range v.PlatformProductId {
		if c >= 65 && c <= 90 {
			v.PlatformProductId = v.PlatformProductId[index:]
			break
		}
	}
}
