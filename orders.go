package ibrestapi

//LiveOrders get a list of the current live orders
func LiveOrders() (LiveOrdersResponse, error) {
	var response LiveOrdersResponse
	err := get(&response, "/iserver/account/orders")
	return response, err
}

//PlaceOrder Place a order
func PlaceOrder(payloadStruct PlaceOrderPayload, accountID string) (PlaceOrderResponses, error) {
	var response PlaceOrderResponses
	err := post(payloadStruct, &response, "/iserver/account/"+accountID+"/order")
	return response, err
}

//PlaceOrderReply Reply to the response from PlaceOrder replyID should be the id recived from PlaceOrder
func PlaceOrderReply(payloadStruct PlaceOrderReplyPayload, replyID string) (PlaceOrderReplyResponses, error) {
	var response PlaceOrderReplyResponses
	err := post(payloadStruct, &response, "/iserver/account/"+replyID+"/order")
	return response, err
}

//PreviewOrder Preview a order
func PreviewOrder(payloadStruct PlaceOrderPayload, accountID string) (PreviewOrderResponse, error) {
	var response PreviewOrderResponse
	err := post(payloadStruct, &response, "/iserver/account/"+accountID+"/order/whatif")
	return response, err
}

//ModifyOrder Modify a order
func ModifyOrder(payloadStruct ModifyOrderPayload, accountID string, orderID string) (ModifyOrderResponses, error) {
	var response ModifyOrderResponses
	err := post(payloadStruct, &response, "/iserver/account/"+accountID+"/order/"+orderID)
	return response, err
}

//DeleteOrder Modify a order
func DeleteOrder(accountID string, orderID string) (ModifyOrderResponses, error) {
	var response ModifyOrderResponses
	err := get(&response, "/iserver/account/"+accountID+"/order/"+orderID)
	return response, err
}
