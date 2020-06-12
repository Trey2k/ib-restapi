package ibrestapi

//SearchPayload the payload struct for Search
type SearchPayload struct {
	Symbol  string `json:"symbol"`
	Name    bool   `json:"name"`
	SecType string `json:"secType"`
}

//Start Order Structs

//PlaceOrderPayload is the payload struct for PlaceOrders
type PlaceOrderPayload struct {
	AcctID          string `json:"acctId"`
	Conid           int    `json:"conid"`
	SecType         string `json:"secType"`
	COID            string `json:"cOID"`
	ParentID        string `json:"parentId"`
	OrderType       string `json:"orderType"`
	ListingExchange string `json:"listingExchange"`
	OutsideRTH      bool   `json:"outsideRTH"`
	Price           int    `json:"price"`
	Side            string `json:"side"`
	Ticker          string `json:"ticker"`
	Tif             string `json:"tif"`
	Referrer        string `json:"referrer"`
	Quantity        int    `json:"quantity"`
	UseAdaptive     bool   `json:"useAdaptive"`
}

//PlaceOrderReplyPayload is the payload struct for PlaceOrderReply
type PlaceOrderReplyPayload struct {
	Confirmed bool `json:"confirmed"`
}

//ModifyOrderPayload is the payload struct for ModifyOrder
type ModifyOrderPayload struct {
	AcctID          string `json:"acctId"`
	Conid           int    `json:"conid"`
	OrderType       string `json:"orderType"`
	OutsideRTH      bool   `json:"outsideRTH"`
	Price           int    `json:"price"`
	AuxPrice        int    `json:"auxPrice"`
	Side            string `json:"side"`
	ListingExchange string `json:"listingExchange"`
	Ticker          string `json:"ticker"`
	Tif             string `json:"tif"`
	Quantity        int    `json:"quantity"`
}
