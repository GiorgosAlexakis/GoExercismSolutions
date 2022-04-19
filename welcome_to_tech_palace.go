package techpalace
import "strings"
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	defaultMessage := "Welcome to the Tech Palace, "
    nameInUpperCase := strings.ToUpper(customer)
    return defaultMessage + nameInUpperCase
}
// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    stars := strings.Repeat("*",numStarsPerLine)
	return stars + "\n"+ welcomeMsg + "\n" + stars
}
// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    MessageNoStars := strings.ReplaceAll(oldMsg,"*","")
    MessageNoSpaces := strings.TrimSpace(MessageNoStars)
    MessageNoNewLines := strings.ReplaceAll(MessageNoSpaces,"\n","")
    return MessageNoNewLines
}
