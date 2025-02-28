package service

import (
	"strconv"
	"strings"

	dto "github.com/PakornPK/order-placement/Dto"
)

type OrderService interface {
	PlaceOrder(orders []dto.InputOrder) ([]dto.CleanedOrder, error)
}

type orderService struct{}

func NewOrderService() OrderService {
	return orderService{}
}

func (s orderService) PlaceOrder(orders []dto.InputOrder) ([]dto.CleanedOrder, error) {
	var res []dto.CleanedOrder
	currentNo := 1
	cleaner := make(map[string]int, 0)
	wiping := make(map[string]int, 0)
	textureStr := ""
	inputs := prepareInput(orders)
	for _, order := range inputs {
		order.ExtractProduct()
		res = append(res, dto.CleanedOrder{
			No:         currentNo,
			ProductId:  order.PlatformProductId,
			MaterialId: order.GetMaterialId(),
			ModelId:    order.GetModelId(),
			Qty:        order.Qty,
			UnitPrice:  order.UnitPrice,
			TotalPrice: order.Calculate(),
		})
		currentNo += 1
		wiping["WIPING-CLOTH"] += order.Qty
		cleaner[order.GetTextureId()+"-CLEANNER"] += order.Qty
		textureStr += (order.GetTextureId() + "-CLEANNER,")
	}
	textureStr = textureStr[:len(textureStr)-1]
	for k, v := range wiping {
		res = append(res, dto.CleanedOrder{
			No:         currentNo,
			ProductId:  k,
			Qty:        v,
			UnitPrice:  0,
			TotalPrice: 0,
		})
		currentNo += 1
	}
	var merged []string
	mapTmp := make(map[string]string, 0)
	for _, v := range strings.Split(textureStr, ",") {
		if _, ok := mapTmp[v]; !ok {
			mapTmp[v] = v
			merged = append(merged, v)
		}
	}

	for _, v := range merged {
		res = append(res, dto.CleanedOrder{
			No:         currentNo,
			ProductId:  v,
			Qty:        cleaner[v],
			UnitPrice:  0,
			TotalPrice: 0,
		})
		currentNo += 1
	}

	return res, nil
}

func prepareInput(inputs []dto.InputOrder) []dto.InputOrder {
	var newInputs []dto.InputOrder
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
			tmp := dto.InputOrder{
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

func cleanProductId(v *dto.InputOrder) {
	for index, c := range v.PlatformProductId {
		if c >= 65 && c <= 90 {
			v.PlatformProductId = v.PlatformProductId[index:]
			break
		}
	}
}
