package workflow

const OrderFoodTaskQueue = "order-food"

var Queries = struct {
	GET_STATUS string
}{
	GET_STATUS: "GET_STATUS",
}

var Signals = struct {
	CHECKOUT string // Submits order for payment
}{
	CHECKOUT: "CHECKOUT",
}

var Updates = struct {
	ADD_ITEM      string // Adds an item to the order
	REMOVE_ITEM   string // Remove an item from the order
	UPDATE_STATUS string // Restaurant updates status of order
}{
	ADD_ITEM:      "ADD_ITEM",
	REMOVE_ITEM:   "REMOVE_ITEM",
	UPDATE_STATUS: "UPDATE_STATUS",
}
