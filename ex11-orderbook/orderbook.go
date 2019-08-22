package orderbook

type Orderbook struct {
	Bids   []*Order
	Asks   []*Order
	Trades []*Trade
	///
}

func New() *Orderbook {
	var OrdBook Orderbook = &Orderbook{}
	OrdBook.Bids = &*Order{}
	OrdBook.Asks = &*Order{}
	OrdBook.Trade = &*Trade{}
	return OrdBook
}

func BidMarketOrder(OrdBook *Orderbook, order *Order) ([]*Trade, *Order) {

}

func BidLimitOrder(OrdBook *Orderbook, order *Order) ([]*Trade, *Order) {

}

func BidLimitOrder(OrdBook *Orderbook, order *Order) ([]*Trade, *Order) {

}

func BidLimitOrder(OrdBook *Orderbook, order *Order) ([]*Trade, *Order) {

}

func (OrdBook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	switch Order.Side {
	case SideBid:
		switch Order.Kind {
		case KindMarket:
			BidMarketOrder(OrdBook, order)
		case KindLimit:
			BidLimitOrder(OrdBook, order)
		}
	case SideAsk:
		switch Order.Kind {
		case KindMarket:
			AskMarketOrder(OrdBook, order)
		case KindLimit:
			AskLimitOrder(OrdBook, order)
		}
	}
	return nil, nil
}
