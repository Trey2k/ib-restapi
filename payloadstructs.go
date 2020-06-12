package ibrestapi

//SearchPayload the payload struct for Search
type SearchPayload struct {
	Symbol  string
	Name    bool
	SecType string
}

//Start Order Structs

//PlaceOrderPayload is the payload struct for PlaceOrders
type PlaceOrderPayload struct {
	AcctID          string
	Conid           int
	SecType         string
	COID            string
	ParentID        string
	OrderType       string
	ListingExchange string
	OutsideRTH      bool
	Price           int
	Side            string
	Ticker          string
	Tif             string
	Referrer        string
	Quantity        int
	UseAdaptive     bool
}

//PlaceOrderReplyPayload is the payload struct for PlaceOrderReply
type PlaceOrderReplyPayload struct {
	Confirmed bool
}

//ModifyOrderPayload is the payload struct for ModifyOrder
type ModifyOrderPayload struct {
	AcctID          string
	Conid           int
	OrderType       string
	OutsideRTH      bool
	Price           int
	AuxPrice        int
	Side            string
	ListingExchange string
	Ticker          string
	Tif             string
	Quantity        int
}
