package main

func iNewLSIA_AAAAAK () (*tLSIA_AAAAAK) { /* To get a local standardized interface A of this comp, it
	is recommended that you do it with this remote interface. */

	return &tLSIA_AAAAAK {}
}

// Local Standardized Interface A { ...
type tLSIA_AAAAAK struct {}

func (comp *tLSIA_AAAAAK) Record (log string) (error) {
	return iRecord_AAAAAK (log)
}

func (comp *tLSIA_AAAAAK) Shutdown () {
	iShutdown_AAAAAK ()
}
// ... }
