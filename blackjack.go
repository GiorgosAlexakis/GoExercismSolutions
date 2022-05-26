package blackjack
type set map[string]struct{}
var contains = struct{}{}
// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	ruleset :=  map[string]int{
        "ace" :11,
        "two":2,
        "three":3,
        "four":4,
        "five":5,
        "six":6,
        "seven":7,
        "eight":8,
        "nine":9,
        "ten":10,
        "jack":10,
        "queen":10,
        "king":10,
        "other":0,
    }
	return ruleset[card]
}
func fillStandCardSet() set {
    dealerCardsList := []string{"ten","jack","queen","king","ace"}
	dealerCards := make(set,len(dealerCardsList))
    for _,v := range dealerCardsList {
        dealerCards[v] = contains
    }
    return dealerCards
}
func stand(dealerCard string,standCardSet set) bool {
    _,contains := standCardSet[dealerCard]
    return contains
}
// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    standCardSet := fillStandCardSet()
    sum := ParseCard(card1)+ParseCard(card2)
    switch {
    case card1=="ace" && card2=="ace":
    	return "P"
    case sum==21:
        if stand(dealerCard,standCardSet) {
            return "S"
        } else {
        	return "W"
        }
    case sum >=17 && sum <=20:
    	return "S"
    case sum >= 12 && sum <=16:
    	if ParseCard(dealerCard) >= 7 {
            return "H"
        } else {
        return "S"
        }
    case sum <= 11:
		return "H"
    default:
    	return "H"
    }
}
