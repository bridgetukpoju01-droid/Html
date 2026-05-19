package main
import(
	"fmt"
)
func ValidateBanner(banner map[rune][] string) error {
	if len (banner) != 95 {
		return fmt.Errorf("incomplete banner")
	}
	for key, value := range banner {
		if key < 32 || key > 126 {
			return fmt.Errorf("invalid character")
		}
		if len(value) != 8{
			return fmt.Errorf ("invalid length")
		}
	} return nil 
}