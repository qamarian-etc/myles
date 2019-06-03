package main

func iNewPInterface_AAAAAC () (*tPInterface_AAAAAC) { /* To get a portable interface of this comp, it
	is recommended that you do it with this interface. */

	return &tPInterface_AAAAAC {}
}

type tPInterface_AAAAAC struct {}

func (comp *tPInterface_AAAAAC) Output (compID, outputType, output string, moreOutput ... string) {
	iOutputT2_AAAAAC (compID, outputType, output, moreOutput ...)
}
