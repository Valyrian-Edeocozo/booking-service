package entries

import (
	"strings"
)

func FormatNames(bookings []string, firstName string, lastName string) []string {
	bookings = append(bookings, firstName+" "+lastName)

	firstNames := []string{}

	for _, name := range bookings {
		var names = strings.Fields(name)
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}
