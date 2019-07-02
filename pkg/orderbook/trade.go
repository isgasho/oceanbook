package orderbook

import (
	"github.com/draveness/oceanbook/api/protobuf-spec/oceanbookpb"
	"github.com/shopspring/decimal"
)

// Trade .
type Trade struct {
	ID       uint64
	Pair     string
	Price    decimal.Decimal
	Quantity decimal.Decimal
	TakerID  uint64
	MakerID  uint64
}

// Serialize returns protobuf encoded trade.
func (t *Trade) Serialize() *oceanbookpb.Trade {
	return &oceanbookpb.Trade{
		Id:       t.ID,
		Pair:     t.Pair,
		Price:    t.Price.String(),
		Quantity: t.Quantity.String(),
		TakerId:  t.TakerID,
		MakerId:  t.MakerID,
	}
}
