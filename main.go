package reverse

// I want this to be as simple as possible for api creating.
// apis might need a new user creation setup function, because
// our api creators might have databases or something that they connect to.
// but maybe that should be our database instances that we just setup when we see they're trying to connect to it?
// that could be pretty cool... So all of their stuff could just be written in goose migrations
// and sqlc syntax and we'll use that to populate database instances for them? This is getting kinda weird though.

// and my lack of experience is showing.

// dbs should just be external. Just allow people to send along whatever their packages receive and
// have environment variables that point to their dbs and what not. I don't care what they have externally, they'll have to price those in to their decision.

import (
	"strings"
)

//Reverse does some freaking awesome things
// lets go boys
func Reverse(Orig string) (a string, b string, d string, err error) {
	var c []string = strings.Split(Orig, "")

	// Look at all this stuff it does
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}

	return strings.Join(c, ""), "", "", nil
}

// and look at this crap
func AnotherThing(random string, second string) (testName string, err error) {
	return "asdf", nil
}
