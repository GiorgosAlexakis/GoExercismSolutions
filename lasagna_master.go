package lasagna
// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string,numOfLayers int) int {
    if numOfLayers == 0 {
        numOfLayers = 2
    }
	return len(layers) * numOfLayers
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int,sauce float64) {
    for _,layer := range layers {
        if layer == "noodles" {
            noodles += 50
        }
    	if layer == "sauce" {
            sauce += 0.2
        }
    }
	return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList,myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64,portions int) []float64{
    scaledQuantities := make([]float64,len(quantities))
	for i,_ := range scaledQuantities {
        scaledQuantities[i] = (quantities[i] / 2.0) * float64(portions)
    }    
    return scaledQuantities
}
