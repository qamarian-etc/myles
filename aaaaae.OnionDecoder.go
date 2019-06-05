package main

func iNewLSIA_AAAAAE () (*tLSIA_AAAAAE) { /* To get a local standardized interface A of this comp, it
	is recommended that you do it with this remote interface. */

	return &tLSIA_AAAAAE {}
}

// Local Standardized Interface A { ...
type tLSIA_AAAAAE struct {}

func (comp *tLSIA_AAAAAE) Decode (path string) (string, error) {
	return iDecode_AAAAAE (path)
}
// ... }
