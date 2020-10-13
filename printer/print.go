package printer

import "fmt"

// Creates a new printer service.
// Setting the brand and mode arguments on the printer service
//
//     // Create a new printer with a given brand, mode
//     myPrinter := printer.NewPrinter("hp", "404")
func NewPrinter(brand, mode string) *Printer {
	return &Printer{
		Brand: brand,
		Model: mode,
	}
}

// Prints with the default settings
func (p *Printer) Print(quality Quality) error {
	fmt.Println("do a print!")
	return nil
}

// Prints in black & white colors
func (p *Printer) PrintBlackWhite(quality Quality) error {
	fmt.Println("do a black & white print!")
	return nil
}
