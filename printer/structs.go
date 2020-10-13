package printer

// The printer quality is a numeric value where the bigger number, the better quality
type Quality int

type Printer struct {
	// The printer brand
	Brand string

	// The printer brand
	Model string

	// Defines the quality level of the printer
	Quality Quality
}
