package search

import (
	"github.com/korhov/stock-issuers-server/internal/tracing"
	pkgSecurities "github.com/korhov/stock-issuers-server/pkg/securities"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"

	"context"
	"errors"
)

type Service interface {
	SearchRequest(ctx context.Context, req Request) (pkgSecurities.SecurityPaper, error)
}

type Request struct {
	Ticket string
}

type service struct {
	log        *zap.SugaredLogger
	tr         opentracing.Tracer
	securities pkgSecurities.Service
}

func NewService(log *zap.SugaredLogger, tracer opentracing.Tracer, svcSecurities pkgSecurities.Service) Service {
	return &service{
		log:        log,
		tr:         tracer,
		securities: svcSecurities,
	}
}

func (s *service) SearchRequest(ctx context.Context, req Request) (pkgSecurities.SecurityPaper, error) {
	span, ctx := tracing.StartSpan(ctx, s.tr, "search.SearchRequest") // nolint
	defer span.Finish()

	if req.Ticket != "" {
		if ticketID, ok := s.securities.GetIDByTicket(req.Ticket); ok {
			securityPaper := s.securities.GetByID(ticketID)
			span.LogFields(log.Int("ticket.id", ticketID))
			return pkgSecurities.SecurityPaper{
				COUPON_PERCENT:              securityPaper[pkgSecurities.Headers["COUPON_PERCENT"]],
				CURRENCY:                    securityPaper[pkgSecurities.Headers["CURRENCY"]],
				DATESTAMP:                   securityPaper[pkgSecurities.Headers["DATESTAMP"]],
				DECISION_DATE:               securityPaper[pkgSecurities.Headers["DECISION_DATE"]],
				DISCLOSURE_PART_PAGE:        securityPaper[pkgSecurities.Headers["DISCLOSURE_PART_PAGE"]],
				DISCLOSURE_RF_INFO_PAGE:     securityPaper[pkgSecurities.Headers["DISCLOSURE_RF_INFO_PAGE"]],
				EARLY_REDEMPTION:            securityPaper[pkgSecurities.Headers["EARLY_REDEMPTION"]],
				EARLY_REPAYMENT:             securityPaper[pkgSecurities.Headers["EARLY_REPAYMENT"]],
				EMITENT_FULL_NAME:           securityPaper[pkgSecurities.Headers["EMITENT_FULL_NAME"]],
				HAS_PROSPECTUS:              securityPaper[pkgSecurities.Headers["HAS_PROSPECTUS"]],
				HAS_RESTRICTION_CIRCULATION: securityPaper[pkgSecurities.Headers["HAS_RESTRICTION_CIRCULATION"]],
				INCLUDED_DURING_CREATION:    securityPaper[pkgSecurities.Headers["INCLUDED_DURING_CREATION"]],
				INCLUDED_WITHOUT_COMPLIANCE: securityPaper[pkgSecurities.Headers["INCLUDED_WITHOUT_COMPLIANCE"]],
				INN:                         securityPaper[pkgSecurities.Headers["INN"]],
				INSTRUMENT_CATEGORY:         securityPaper[pkgSecurities.Headers["INSTRUMENT_CATEGORY"]],
				INSTRUMENT_ID:               securityPaper[pkgSecurities.Headers["INSTRUMENT_ID"]],
				INSTRUMENT_TYPE:             securityPaper[pkgSecurities.Headers["INSTRUMENT_TYPE"]],
				ISIN:                        securityPaper[pkgSecurities.Headers["ISIN"]],
				ISSUE_AMOUNT:                securityPaper[pkgSecurities.Headers["ISSUE_AMOUNT"]],
				ISS_BOARDS:                  securityPaper[pkgSecurities.Headers["ISS_BOARDS"]],
				IS_CONCESSION_AGREEMENT:     securityPaper[pkgSecurities.Headers["IS_CONCESSION_AGREEMENT"]],
				IS_MORTGAGE_AGENT:           securityPaper[pkgSecurities.Headers["IS_MORTGAGE_AGENT"]],
				LISTING_LEVEL_HIST:          securityPaper[pkgSecurities.Headers["LISTING_LEVEL_HIST"]],
				LIST_SECTION:                securityPaper[pkgSecurities.Headers["LIST_SECTION"]],
				NOMINAL:                     securityPaper[pkgSecurities.Headers["NOMINAL"]],
				OBLIGATION_PROGRAM_RN:       securityPaper[pkgSecurities.Headers["OBLIGATION_PROGRAM_RN"]],
				OKSM_EDR:                    securityPaper[pkgSecurities.Headers["OKSM_EDR"]],
				ONLY_EMITENT_FULL_NAME:      securityPaper[pkgSecurities.Headers["ONLY_EMITENT_FULL_NAME"]],
				OTHER_SECURITIES:            securityPaper[pkgSecurities.Headers["OTHER_SECURITIES"]],
				QUALIFIED_INVESTOR:          securityPaper[pkgSecurities.Headers["QUALIFIED_INVESTOR"]],
				REGISTRY_DATE:               securityPaper[pkgSecurities.Headers["REGISTRY_DATE"]],
				REGISTRY_NUMBER:             securityPaper[pkgSecurities.Headers["REGISTRY_NUMBER"]],
				REG_COUNTRY:                 securityPaper[pkgSecurities.Headers["REG_COUNTRY"]],
				RETAINED_WITHOUT_COMPLIANCE: securityPaper[pkgSecurities.Headers["RETAINED_WITHOUT_COMPLIANCE"]],
				RN:                          securityPaper[pkgSecurities.Headers["RN"]],
				SECURITY_HAS_DEFAULT:        securityPaper[pkgSecurities.Headers["SECURITY_HAS_DEFAULT"]],
				SECURITY_HAS_TECH_DEFAULT:   securityPaper[pkgSecurities.Headers["SECURITY_HAS_TECH_DEFAULT"]],
				SUPERTYPE:                   securityPaper[pkgSecurities.Headers["SUPERTYPE"]],
				TRADE_CODE:                  securityPaper[pkgSecurities.Headers["TRADE_CODE"]],
			}, nil
		}

		return pkgSecurities.SecurityPaper{}, errors.New("the Issuer is not found")
	}

	return pkgSecurities.SecurityPaper{}, errors.New("search parameters are not specified")
}
