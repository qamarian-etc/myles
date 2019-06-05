package main

func iNewLSIA_AAAAAI () (*tLSIA_AAAAAI) { /* To get a local standardized interface A of this comp, it
	is recommended that you do it with this remote interface. */

	return &tLSIA_AAAAAI {}
}

// Local Standardized Interface A { ...
type tLSIA_AAAAAI struct {}

func (comp *tLSIA_AAAAAI) BeInformed (eventDescp string) {
	iBeInformed_AAAAAI (eventDescp)
}
// ... }
