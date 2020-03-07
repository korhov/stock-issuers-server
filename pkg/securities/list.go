package securities

import (
	"github.com/korhov/stock-issuers-server/internal/tracing"

	"context"
)

type RequestList struct {
	InstrumentType string
}

type ResponseList struct {
	List []SecurityPaperList `json:"list"`
}

const (
	MaxList int = 10
)

func (s *service) GetList(ctx context.Context, req RequestList) (ResponseList, error) {
	span, ctx := tracing.StartSpan(ctx, s.tr, "securities.GetList") // nolint
	defer span.Finish()

	response := ResponseList{}
	for _, record := range s.records[1:] {
		if !(req.InstrumentType != "" && req.InstrumentType == record[Headers["INSTRUMENT_TYPE"]]) {
			continue
		}
		response.List = append(response.List, SecurityPaperList{
			CURRENCY:          record[Headers["CURRENCY"]],
			EMITENT_FULL_NAME: record[Headers["EMITENT_FULL_NAME"]],
			INN:               record[Headers["INN"]],
			INSTRUMENT_ID:     record[Headers["INSTRUMENT_ID"]],
			INSTRUMENT_TYPE:   record[Headers["INSTRUMENT_TYPE"]],
			ISIN:              record[Headers["ISIN"]],
			ISSUE_AMOUNT:      record[Headers["ISSUE_AMOUNT"]],
			LIST_SECTION:      record[Headers["LIST_SECTION"]],
			NOMINAL:           record[Headers["NOMINAL"]],
			REGISTRY_DATE:     record[Headers["REGISTRY_DATE"]],
			REGISTRY_NUMBER:   record[Headers["REGISTRY_NUMBER"]],
			RN:                record[Headers["RN"]],
			SUPERTYPE:         record[Headers["SUPERTYPE"]],
			TRADE_CODE:        record[Headers["TRADE_CODE"]],
		})
		if len(response.List) >= MaxList {
			break
		}
	}

	return response, nil
}
