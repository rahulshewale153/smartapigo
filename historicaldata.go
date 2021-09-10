package smartapigo

import "net/http"

// HistoricalDataParams represents parameters for getting Order History.
type HistoricalDataParams struct {
	Exchange    string `json:"exchange"`
	SymbolToken string `json:"symboltoken"`
	Interval    string `json:"interval"`
	FromDate    string `json:"fromdate"`
	ToDate      string `json:"todate"`
}

// GetHistoricalData gets Order History according to datetime .
func (c *Client) GetHistoricalData(historicalDataParams HistoricalDataParams) ([][]interface{}, error) {
	var historicalData [][]interface{}
	var (
		params map[string]interface{}
		err    error
	)
	params = structToMap(historicalDataParams, "json")
	err = c.doEnvelope(http.MethodPost, URIHISTORYDATA, params, nil, &historicalData, true)
	return historicalData, err
}
