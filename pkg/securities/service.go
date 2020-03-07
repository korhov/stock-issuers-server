package securities

import (
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"context"
	"encoding/csv"
	"os"
)

type Service interface {
	GetList(ctx context.Context, req RequestList) (ResponseList, error)
	GetByID(id int) []string
	GetIDByTicket(ticket string) (int, bool)
}

type TypeTickers map[string]int
type service struct {
	log     *zap.SugaredLogger
	tr      opentracing.Tracer
	records [][]string
	tickers TypeTickers
}

func NewService(log *zap.SugaredLogger, tracer opentracing.Tracer) Service {
	f, err := os.Open("./data/securities-list.csv")
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	tickers := TypeTickers{}
	for i, record := range records {
		if record[Headers["TRADE_CODE"]] != "" {
			tickers[record[Headers["TRADE_CODE"]]] = i
		}
	}

	return &service{
		log:     log,
		tr:      tracer,
		records: records,
		tickers: tickers,
	}
}

func (s *service) GetByID(ticketID int) []string {
	return s.records[ticketID]
}

func (s *service) GetIDByTicket(ticket string) (int, bool) {
	ticketID, ok := s.tickers[ticket]
	return ticketID, ok
}
