package Status

type Enum int

const (
	// Postal code is valid.
	Valid Enum = 0

	// Country code is unrecognized. Country code is case sensitive.
	UnrecognizedCountryCode Enum = 1

	// Postal code does not comply with the code pattern for the given country.
	InvalidFormat Enum = 2
)

var names = map[Enum]string{
	Valid:                   "Valid",
	UnrecognizedCountryCode: "Unrecognized country code",
	InvalidFormat:           "Invalid format",
}

// Converts the value of this instance to its equivalent string representation.
//
// # Returns:
//
//	name string
//
// The string representation of the value of this instance.
func (enum Enum) String() (name string) {
	if name, isDefined := names[enum]; isDefined {
		return name
	}
	return "Unknown"
}

// Returns a bool telling whether a current instance enum exists in a specified enumeration.
//
// # Returns:
//
//	isDefined bool
//
// True if a constant in enumeration has a defined value; false otherwise.
func (enum Enum) IsDefined() (isDefined bool) {
	_, isDefined = names[enum]
	return isDefined
}

// Returns a bool telling whether a given enum exists in a specified enumeration.
//
// # Parameters:
//
//	enum Enum
//
// The value of a constant in enumeration.
//
// # Returns:
//
//	isDefined bool
//
// True if a constant in enumeration has a value equal to enum; false otherwise.
func IsDefined(enum Enum) (isDefined bool) {
	return enum.IsDefined()
}

// Retrieves a slice of the values of the constants in a specified enumeration.
//
// # Returns:
//
//	enums []Enum
//
// A slice that contains the values of the constants in enumeration.
func GetValues() (enums []Enum) {
	return []Enum{
		Valid,
		UnrecognizedCountryCode,
		InvalidFormat,
	}
}
