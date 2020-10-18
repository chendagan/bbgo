package store

import (
	"github.com/c9s/bbgo/pkg/types"
)

//go:generate callbackgen -type MarketDataStore
type MarketDataStore struct {
	Symbol string

	// KLineWindows stores all loaded klines per interval
	KLineWindows map[types.Interval]types.KLineWindow `json:"-"`

	LastKLine types.KLine

	kLineUpdateCallbacks []func(kline types.KLine)

	orderBook *types.StreamOrderBook

	orderBookUpdateCallbacks []func()
}

func NewMarketDataStore(symbol string) *MarketDataStore {
	return &MarketDataStore{
		Symbol: symbol,

		orderBook: types.NewStreamBook(symbol),

		// KLineWindows stores all loaded klines per interval
		KLineWindows: make(map[types.Interval]types.KLineWindow, len(types.SupportedIntervals)), // 12 interval, 1m,5m,15m,30m,1h,2h,4h,6h,12h,1d,3d,1w
	}
}

func (store *MarketDataStore) handleOrderBookUpdate(book types.OrderBook) {
	if book.Symbol != store.Symbol {
		return
	}

	store.orderBook.Update(book)
}

func (store *MarketDataStore) handleOrderBookSnapshot(book types.OrderBook) {
	if book.Symbol != store.Symbol {
		return
	}

	store.orderBook.Load(book)
}

func (store *MarketDataStore) BindStream(stream types.Stream) {
	stream.OnKLineClosed(store.handleKLineClosed)
	stream.OnBookSnapshot(store.handleOrderBookSnapshot)
	stream.OnBookUpdate(store.handleOrderBookUpdate)

	store.orderBook.BindStream(stream)
}

func (store *MarketDataStore) handleKLineClosed(kline types.KLine) {
	if kline.Symbol != store.Symbol {
		return
	}

	store.AddKLine(kline)
}

func (store *MarketDataStore) AddKLine(kline types.KLine) {
	var interval = types.Interval(kline.Interval)
	var window = store.KLineWindows[interval]
	window.Add(kline)

	store.LastKLine = kline

	store.EmitKLineUpdate(kline)
}