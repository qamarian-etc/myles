package main

func iNewPInterface_AAAAAE () (*tPInterface_AAAAAE) { /* To get a portable interface of this comp, it
	is recommended that you do it with this interface. */

	return &tPInterface_AAAAAE {}
}

type tPInterface_AAAAAE struct {}

func (comp *tPInterface_AAAAAE) Decode (path string) (string, error) {
	return iDecode_AAAAAE (path)
}
