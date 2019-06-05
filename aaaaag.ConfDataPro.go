package main

func iNewLSIA_AAAAAG () (*tLSIA_AAAAAG) { /* To get a local standardized interface A of this comp, it
	is recommended that you do it with this remote interface. */

	return &tLSIA_AAAAAG {}
}

// Local Standardized Interface A { ...
type tLSIA_AAAAAG struct {}

func (comp *tLSIA_AAAAAG) ScalarData (dataName string) (string, error) {
	return iScalarData_AAAAAG (dataName)
}

func (comp *tLSIA_AAAAAG) SliceData (dataName string) ([]string, error) {
	return iSliceData_AAAAAG (dataName)
}

func (comp *tLSIA_AAAAAG) MapData (dataName string) (map[string]string, error) {
	return iMapData_AAAAAG (dataName)
}
// ... }
