package smartapigo

import (
	"testing"
)

func (ts *TestSuite) TestGetHistoricalData(t *testing.T) {
	t.Parallel()
	params := HistoricalDataParams{
		Exchange:    "NSE",
		SymbolToken: "15083",
		Interval:    "FIVE_MINUTE",
		FromDate:    "2021-09-09 09:15",
		ToDate:      "2021-09-09 12:45",
	}
	historicalData, err := ts.TestConnect.GetHistoricalData(params)
	if err != nil {
		t.Errorf("Error while fetching historicalData. %v", err)
	}
	if len(historicalData) < 0 {
		t.Errorf("Error while getting historicalData. %v", err)
	}

}
