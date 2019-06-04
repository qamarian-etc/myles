package main

func iNewLSIA_AAAAAA () (*tLSIA_AAAAAA) { /* To get a local standardized interface A of this comp, it
	is recommended that you do it with this remote interface. */

	return &tLSIA_AAAAAA {}
}

// Local Standardized Interface A { ...
type tLSIA_AAAAAA struct {}

func (comp *tLSIA_AAAAAA) RegComp (compID string, initFunc, dnitFunc func (), dep []string) error {
	// This interface can be used to register other comps with comp AAAAAA.

	return iRegComp_AAAAAA (compID, initFunc, dnitFunc, dep)
}

func (comp *tLSIA_AAAAAA) Shutdown () { /* This interface can be used to gracefully shutdown
	your mylware. */

	iShutdown_AAAAAA ()
}
// ... }
