package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	quantityToValue := map[string]int{}
    
    quantityToValue["quarter_of_a_dozen"] = 3
    quantityToValue["half_of_a_dozen"] = 6
    quantityToValue["dozen"] = 12
    quantityToValue["small_gross"] = 120
    quantityToValue["gross"] = 144
    quantityToValue["great_gross"] = 1728

    return quantityToValue
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, exists := units[unit]

    if !exists {
        return false
    }

    already, inCart := bill[item]

    if inCart {
        bill[item] = already + quantity
    } else {
        bill[item] = quantity
    }

    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantity, exists := units[unit]

    if !exists {
        return false
    }

    already, inCart := bill[item]

    if !inCart || already < quantity {
        return false
    }

    if already == quantity {
        delete(bill, item)
    } else {
    	bill[item] = already - quantity
    }

    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	already, inCart := bill[item]

    if !inCart {
        return 0, false
    }

    return already, true
}
