package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if (kind == "car" || kind == "truck") {
        return true 
    } else {
    	return false
        }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    message := " is clearly the better choice."
	if option1 < option2 {
        return option1 + message
    } else {
    	return option2 + message
        }
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    switch {
        case age < 3.0: 
    		return originalPrice * 0.8
        case age >= 3.0 && age < 10.0:
    		return originalPrice * 0.7
        case age >= 10.0:
    		return originalPrice * 0.5
        default: 
    		return originalPrice * 0.8
    }
}
