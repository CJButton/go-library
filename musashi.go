package musashi

import (
	"fmt"
	"musashi/miyamoto"
)

/*
Musashi returns Musashi Myamoto!
*/
func Musashi() string {
	fmt.Println(miyamoto.Miyamoto())
	lastName := miyamoto.Miyamoto()
	return "Musashi" + " " + lastName
}
