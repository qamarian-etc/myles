package main

func iNewLSIA_AAAAAL () (*tLSIA_AAAAAL) { /* To get a local standardized interface A of this comp, it
	is recommended that you do it with this remote interface. */

	return &tLSIA_AAAAAL {}
}

// Local Standardized Interface A { ...
type tLSIA_AAAAAL struct {}

func (comp *tLSIA_AAAAAL) Shutdown () {
	iShutdown_AAAAAL ()
}
// ... }
