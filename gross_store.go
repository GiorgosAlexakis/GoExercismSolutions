package gross
// Units stores the Gross Store unit measurements.
func Units() map[string]int {
   return map[string]int{
        "quarter_of_a_dozen":3,
        "half_of_a_dozen":6,
        "dozen":12,
        "small_gross":120,
        "gross":144,
        "great_gross":1728,
    }
}
// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := make(map[string]int)
	return bill
}
// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value,exists := units[unit]
    if !exists {
        return false
    } else {
    	bill[item] = bill[item] + value
        return true
    }
}
// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_,itemExists := bill[item]
    if !itemExists {
        return false
    }
	_,unitExists := units[unit]
    if !unitExists {
        return false
    }
	newValue := bill[item] - units[unit]
    if newValue < 0 {
        return false
    }
	if newValue == 0 {
		delete(bill,item)
        return true
    }
	if newValue > 0 {
        bill[item] = newValue
        return true
    }
	return false
}
// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity,exists := bill[item]
    if !exists {
        return 0,false
    } else {
    	return quantity,true
    }
}
