package faker

import (
	"github.com/crit/critical-go/faker/data"
)

func randomName() string {
	var fname string

	if randomElement([]string{"male", "female"}) == "male" {
		fname = randomElement(data.PersonFirstNameMale)
	} else {
		fname = randomElement(data.PersonFirstNameFemale)
	}

	lname := randomElement(data.PersonLastName)

	return fname + " " + lname
}
