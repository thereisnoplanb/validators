package Status

type Enum int

// EU VAT is valid.
const Valid Enum = 0

// EU VAT length is incorrect. Length is shorter then 3 signs including country code.
const IncorrectLength Enum = 1

// EU VAT country code is unrecognized. Country code is case sensitive.
const UnrecognizedCountryCode Enum = 2

// EU VAT number does not comply with the number pattern for the given country.
const InvalidFormat Enum = 3

// EU VAT number the checksum is incorrect.
const IncorrectChecksum Enum = 4

var names = map[Enum]string{
	Valid:                   "Valid",
	IncorrectLength:         "Incorrect length",
	UnrecognizedCountryCode: "Unrecognized country code",
	InvalidFormat:           "Invalid format",
	IncorrectChecksum:       "Incorrect checksum",
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
		IncorrectLength,
		UnrecognizedCountryCode,
		InvalidFormat,
		IncorrectChecksum,
	}
}
