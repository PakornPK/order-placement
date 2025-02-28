package service_test

import (
	"testing"

	"github.com/PakornPK/order-placement/service"
	"github.com/stretchr/testify/assert"
)

func TestPlaceOrder(t *testing.T) {
	serv := service.NewOrderService()
	t.Run("case no.1", func(t *testing.T) {
		inputs := []service.InputOrder{
			{
				No:                1,
				PlatformProductId: "FG0A-CLEAR-IPHONE16PROMAX",
				Qty:               2,
				UnitPrice:         50,
				TotalPrice:        100,
			},
		}
		expects := []service.CleanedOrder{
			{
				No:         1,
				ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
				MaterialId: "FG0A-CLEAR",
				ModelId:    "IPHONE16PROMAX",
				Qty:        2,
				UnitPrice:  50,
				TotalPrice: 100,
			},
			{
				No:         2,
				ProductId:  "WIPING-CLOTH",
				Qty:        2,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         3,
				ProductId:  "CLEAR-CLEANNER",
				Qty:        2,
				UnitPrice:  0,
				TotalPrice: 0,
			},
		}

		out1, out2 := serv.PlaceOrder(inputs)

		assert.Equal(t, expects, out1)
		assert.Nil(t, out2)
	})
	t.Run("case no.2", func(t *testing.T) {
		inputs := []service.InputOrder{
			{
				No:                1,
				PlatformProductId: "x2-3&FG0A-CLEAR-IPHONE16PROMAX",
				Qty:               2,
				UnitPrice:         50,
				TotalPrice:        100,
			},
		}
		expects := []service.CleanedOrder{
			{
				No:         1,
				ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
				MaterialId: "FG0A-CLEAR",
				ModelId:    "IPHONE16PROMAX",
				Qty:        2,
				UnitPrice:  50,
				TotalPrice: 100,
			},
			{
				No:         2,
				ProductId:  "WIPING-CLOTH",
				Qty:        2,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         3,
				ProductId:  "CLEAR-CLEANNER",
				Qty:        2,
				UnitPrice:  0,
				TotalPrice: 0,
			},
		}

		out1, out2 := serv.PlaceOrder(inputs)

		assert.Equal(t, expects, out1)
		assert.Nil(t, out2)
	})
	t.Run("case no.3", func(t *testing.T) {
		inputs := []service.InputOrder{
			{
				No:                1,
				PlatformProductId: "x2-3&FG0A-MATTE-IPHONE16PROMAX*3",
				Qty:               1,
				UnitPrice:         90,
				TotalPrice:        90,
			},
		}
		expects := []service.CleanedOrder{
			{
				No:         1,
				ProductId:  "FG0A-MATTE-IPHONE16PROMAX",
				MaterialId: "FG0A-MATTE",
				ModelId:    "IPHONE16PROMAX",
				Qty:        3,
				UnitPrice:  30,
				TotalPrice: 90,
			},
			{
				No:         2,
				ProductId:  "WIPING-CLOTH",
				Qty:        3,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         3,
				ProductId:  "MATTE-CLEANNER",
				Qty:        3,
				UnitPrice:  0,
				TotalPrice: 0,
			},
		}

		out1, out2 := serv.PlaceOrder(inputs)

		assert.Equal(t, expects, out1)
		assert.Nil(t, out2)
	})
	t.Run("case no.4", func(t *testing.T) {
		inputs := []service.InputOrder{
			{
				No:                1,
				PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B",
				Qty:               1,
				UnitPrice:         80,
				TotalPrice:        80,
			},
		}
		expects := []service.CleanedOrder{
			{
				No:         1,
				ProductId:  "FG0A-CLEAR-OPPOA3",
				MaterialId: "FG0A-CLEAR",
				ModelId:    "OPPOA3",
				Qty:        1,
				UnitPrice:  40,
				TotalPrice: 40,
			},
			{
				No:         2,
				ProductId:  "FG0A-CLEAR-OPPOA3-B",
				MaterialId: "FG0A-CLEAR",
				ModelId:    "OPPOA3-B",
				Qty:        1,
				UnitPrice:  40,
				TotalPrice: 40,
			},
			{
				No:         3,
				ProductId:  "WIPING-CLOTH",
				Qty:        2,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         4,
				ProductId:  "CLEAR-CLEANNER",
				Qty:        2,
				UnitPrice:  0,
				TotalPrice: 0,
			},
		}

		out1, out2 := serv.PlaceOrder(inputs)

		assert.Equal(t, expects, out1)
		assert.Nil(t, out2)
	})

	t.Run("case no.5", func(t *testing.T) {
		inputs := []service.InputOrder{
			{
				No:                1,
				PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B/FG0A-MATTE-OPPOA3",
				Qty:               1,
				UnitPrice:         120,
				TotalPrice:        120,
			},
		}
		expects := []service.CleanedOrder{
			{
				No:         1,
				ProductId:  "FG0A-CLEAR-OPPOA3",
				MaterialId: "FG0A-CLEAR",
				ModelId:    "OPPOA3",
				Qty:        1,
				UnitPrice:  40,
				TotalPrice: 40,
			},
			{
				No:         2,
				ProductId:  "FG0A-CLEAR-OPPOA3-B",
				MaterialId: "FG0A-CLEAR",
				ModelId:    "OPPOA3-B",
				Qty:        1,
				UnitPrice:  40,
				TotalPrice: 40,
			},
			{
				No:         3,
				ProductId:  "FG0A-MATTE-OPPOA3",
				MaterialId: "FG0A-MATTE",
				ModelId:    "OPPOA3",
				Qty:        1,
				UnitPrice:  40,
				TotalPrice: 40,
			},
			{
				No:         4,
				ProductId:  "WIPING-CLOTH",
				Qty:        3,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         5,
				ProductId:  "CLEAR-CLEANNER",
				Qty:        2,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         6,
				ProductId:  "MATTE-CLEANNER",
				Qty:        1,
				UnitPrice:  0,
				TotalPrice: 0,
			},
		}

		out1, out2 := serv.PlaceOrder(inputs)

		assert.Equal(t, expects, out1)
		assert.Nil(t, out2)
	})

	t.Run("case no.6", func(t *testing.T) {
		inputs := []service.InputOrder{
			{
				No:                1,
				PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3",
				Qty:               1,
				UnitPrice:         120,
				TotalPrice:        120,
			},
		}
		expects := []service.CleanedOrder{
			{
				No:         1,
				ProductId:  "FG0A-CLEAR-OPPOA3",
				MaterialId: "FG0A-CLEAR",
				ModelId:    "OPPOA3",
				Qty:        2,
				UnitPrice:  40,
				TotalPrice: 80,
			},
			{
				No:         2,
				ProductId:  "FG0A-MATTE-OPPOA3",
				MaterialId: "FG0A-MATTE",
				ModelId:    "OPPOA3",
				Qty:        1,
				UnitPrice:  40,
				TotalPrice: 40,
			},
			{
				No:         3,
				ProductId:  "WIPING-CLOTH",
				Qty:        3,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         4,
				ProductId:  "CLEAR-CLEANNER",
				Qty:        2,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         5,
				ProductId:  "MATTE-CLEANNER",
				Qty:        1,
				UnitPrice:  0,
				TotalPrice: 0,
			},
		}

		out1, out2 := serv.PlaceOrder(inputs)

		assert.Equal(t, expects, out1)
		assert.Nil(t, out2)
	})

	t.Run("case no.7", func(t *testing.T) {
		inputs := []service.InputOrder{
			{
				No:                1,
				PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2",
				Qty:               1,
				UnitPrice:         160,
				TotalPrice:        160,
			},
			{
				No:                2,
				PlatformProductId: "FG0A-PRIVACY-IPHONE16PROMAX",
				Qty:               1,
				UnitPrice:         50,
				TotalPrice:        50,
			},
		}
		expects := []service.CleanedOrder{
			{
				No:         1,
				ProductId:  "FG0A-CLEAR-OPPOA3",
				MaterialId: "FG0A-CLEAR",
				ModelId:    "OPPOA3",
				Qty:        2,
				UnitPrice:  40,
				TotalPrice: 80,
			},
			{
				No:         2,
				ProductId:  "FG0A-MATTE-OPPOA3",
				MaterialId: "FG0A-MATTE",
				ModelId:    "OPPOA3",
				Qty:        2,
				UnitPrice:  40,
				TotalPrice: 80,
			},
			{
				No:         3,
				ProductId:  "FG0A-PRIVACY-IPHONE16PROMAX",
				MaterialId: "FG0A-PRIVACY",
				ModelId:    "IPHONE16PROMAX",
				Qty:        1,
				UnitPrice:  50,
				TotalPrice: 50,
			},
			{
				No:         4,
				ProductId:  "WIPING-CLOTH",
				Qty:        5,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         5,
				ProductId:  "CLEAR-CLEANNER",
				Qty:        2,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         6,
				ProductId:  "MATTE-CLEANNER",
				Qty:        2,
				UnitPrice:  0,
				TotalPrice: 0,
			},
			{
				No:         7,
				ProductId:  "PRIVACY-CLEANNER",
				Qty:        1,
				UnitPrice:  0,
				TotalPrice: 0,
			},
		}

		out1, out2 := serv.PlaceOrder(inputs)

		assert.Equal(t, expects, out1)
		assert.Nil(t, out2)
	})
}
