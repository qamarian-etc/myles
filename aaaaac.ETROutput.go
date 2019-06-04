package main

func iNewLSIA_AAAAAC () (*tLSI_AAAAAC) { /* To get a local standardized interface A of this comp, it
	is recommended that you do it with this remote interface. */

	return &tLSI_AAAAAC {}
}

// Local Standardized Interface A { ...
type tLSI_AAAAAC struct {}

func (comp *tLSI_AAAAAC) Output (compID, outputType, output string, moreOutput ... string) {
	iOutputT2_AAAAAC (compID, outputType, output, moreOutput ...)
}
// ... }
