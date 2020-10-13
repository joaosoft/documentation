package printer

import "fmt"

var (
	// ErrorUndefinedBrand is an error that is returned if the brand isn't
	//defined.
	ErrorUndefinedBrand = fmt.Errorf("undefined brand")

	// ErrorUndefinedModel is an error that is returned if the model isn't
	//defined.
	ErrorUndefinedModel = fmt.Errorf("undefined model")
)
