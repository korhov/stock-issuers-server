package securities

var Headers = map[string]int{
	"COUPON_PERCENT":              32, // nolint
	"CURRENCY":                    14, // nolint
	"DATESTAMP":                   0,  // nolint
	"DECISION_DATE":               16, // nolint
	"DISCLOSURE_PART_PAGE":        37, // nolint
	"DISCLOSURE_RF_INFO_PAGE":     38, // nolint
	"EARLY_REDEMPTION":            34, // nolint
	"EARLY_REPAYMENT":             33, // nolint
	"EMITENT_FULL_NAME":           11, // nolint
	"HAS_PROSPECTUS":              21, // nolint
	"HAS_RESTRICTION_CIRCULATION": 29, // nolint
	"INCLUDED_DURING_CREATION":    24, // nolint
	"INCLUDED_WITHOUT_COMPLIANCE": 27, // nolint
	"INN":                         12, // nolint
	"INSTRUMENT_CATEGORY":         6,  // nolint
	"INSTRUMENT_ID":               1,  // nolint
	"INSTRUMENT_TYPE":             5,  // nolint
	"ISIN":                        8,  // nolint
	"ISSUE_AMOUNT":                15, // nolint
	"ISS_BOARDS":                  35, // nolint
	"IS_CONCESSION_AGREEMENT":     22, // nolint
	"IS_MORTGAGE_AGENT":           23, // nolint
	"LISTING_LEVEL_HIST":          30, // nolint
	"LIST_SECTION":                2,  // nolint
	"NOMINAL":                     13, // nolint
	"OBLIGATION_PROGRAM_RN":       31, // nolint
	"OKSM_EDR":                    17, // nolint
	"ONLY_EMITENT_FULL_NAME":      18, // nolint
	"OTHER_SECURITIES":            36, // nolint
	"QUALIFIED_INVESTOR":          20, // nolint
	"REGISTRY_DATE":               10, // nolint
	"REGISTRY_NUMBER":             9,  // nolint
	"REG_COUNTRY":                 19, // nolint
	"RETAINED_WITHOUT_COMPLIANCE": 28, // nolint
	"RN":                          3,  // nolint
	"SECURITY_HAS_DEFAULT":        25, // nolint
	"SECURITY_HAS_TECH_DEFAULT":   26, // nolint
	"SUPERTYPE":                   4,  // nolint
	"TRADE_CODE":                  7,  // nolint
}

type SecurityPaper struct {
	COUPON_PERCENT              string `json:"COUPON_PERCENT"`              // nolint
	CURRENCY                    string `json:"currency"`                    // nolint
	DATESTAMP                   string `json:"DATESTAMP"`                   // nolint
	DECISION_DATE               string `json:"DECISION_DATE"`               // nolint
	DISCLOSURE_PART_PAGE        string `json:"DISCLOSURE_PART_PAGE"`        // nolint
	DISCLOSURE_RF_INFO_PAGE     string `json:"DISCLOSURE_RF_INFO_PAGE"`     // nolint
	EARLY_REDEMPTION            string `json:"EARLY_REDEMPTION"`            // nolint
	EARLY_REPAYMENT             string `json:"EARLY_REPAYMENT"`             // nolint
	EMITENT_FULL_NAME           string `json:"EMITENT_FULL_NAME"`           // nolint
	HAS_PROSPECTUS              string `json:"HAS_PROSPECTUS"`              // nolint
	HAS_RESTRICTION_CIRCULATION string `json:"HAS_RESTRICTION_CIRCULATION"` // nolint
	INCLUDED_DURING_CREATION    string `json:"INCLUDED_DURING_CREATION"`    // nolint
	INCLUDED_WITHOUT_COMPLIANCE string `json:"INCLUDED_WITHOUT_COMPLIANCE"` // nolint
	INN                         string `json:"inn"`                         // nolint
	INSTRUMENT_CATEGORY         string `json:"INSTRUMENT_CATEGORY"`         // nolint
	INSTRUMENT_ID               string `json:"instrument_id"`               // nolint
	INSTRUMENT_TYPE             string `json:"INSTRUMENT_TYPE"`             // nolint
	ISIN                        string `json:"isin"`                        // nolint
	ISSUE_AMOUNT                string `json:"ISSUE_AMOUNT"`                // nolint
	ISS_BOARDS                  string `json:"ISS_BOARDS"`                  // nolint
	IS_CONCESSION_AGREEMENT     string `json:"IS_CONCESSION_AGREEMENT"`     // nolint
	IS_MORTGAGE_AGENT           string `json:"IS_MORTGAGE_AGENT"`           // nolint
	LISTING_LEVEL_HIST          string `json:"LISTING_LEVEL_HIST"`          // nolint
	LIST_SECTION                string `json:"LIST_SECTION"`                // nolint
	NOMINAL                     string `json:"NOMINAL"`                     // nolint
	OBLIGATION_PROGRAM_RN       string `json:"OBLIGATION_PROGRAM_RN"`       // nolint
	OKSM_EDR                    string `json:"OKSM_EDR"`                    // nolint
	ONLY_EMITENT_FULL_NAME      string `json:"ONLY_EMITENT_FULL_NAME"`      // nolint
	OTHER_SECURITIES            string `json:"OTHER_SECURITIES"`            // nolint
	QUALIFIED_INVESTOR          string `json:"QUALIFIED_INVESTOR"`          // nolint
	REGISTRY_DATE               string `json:"REGISTRY_DATE"`               // nolint
	REGISTRY_NUMBER             string `json:"REGISTRY_NUMBER"`             // nolint
	REG_COUNTRY                 string `json:"REG_COUNTRY"`                 // nolint
	RETAINED_WITHOUT_COMPLIANCE string `json:"RETAINED_WITHOUT_COMPLIANCE"` // nolint
	RN                          string `json:"RN"`                          // nolint
	SECURITY_HAS_DEFAULT        string `json:"SECURITY_HAS_DEFAULT"`        // nolint
	SECURITY_HAS_TECH_DEFAULT   string `json:"SECURITY_HAS_TECH_DEFAULT"`   // nolint
	SUPERTYPE                   string `json:"SUPERTYPE"`                   // nolint
	TRADE_CODE                  string `json:"ticker"`                      // nolint
}

type SecurityPaperList struct {
	CURRENCY          string `json:"currency"`          // nolint
	EMITENT_FULL_NAME string `json:"EMITENT_FULL_NAME"` // nolint
	INN               string `json:"inn"`               // nolint
	INSTRUMENT_ID     string `json:"instrument_id"`     // nolint
	INSTRUMENT_TYPE   string `json:"INSTRUMENT_TYPE"`   // nolint
	ISIN              string `json:"isin"`              // nolint
	ISSUE_AMOUNT      string `json:"ISSUE_AMOUNT"`      // nolint
	LIST_SECTION      string `json:"LIST_SECTION"`      // nolint
	NOMINAL           string `json:"NOMINAL"`           // nolint
	REGISTRY_DATE     string `json:"REGISTRY_DATE"`     // nolint
	REGISTRY_NUMBER   string `json:"REGISTRY_NUMBER"`   // nolint
	RN                string `json:"RN"`                // nolint
	SUPERTYPE         string `json:"SUPERTYPE"`         // nolint
	TRADE_CODE        string `json:"ticker"`            // nolint
}
