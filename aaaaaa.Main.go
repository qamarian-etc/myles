package main

func iNewPInterface_AAAAAA () (*tPInterface_AAAAAA) { /* To get a portable interface of this comp, it
	is recommended that you do it with this interface. */

	return &tPInterface_AAAAAA {}
}

type tPInterface_AAAAAA struct {}

func (comp *tPInterface_AAAAAA) RegComp (compID string, initFunc, dnitFunc func (),
	dep []string) (error) { // This interface can be used to register other comp.

	return iRegComp_AAAAAA (compID, initFunc, dnitFunc, dep)
}

func (comp *tPInterface_AAAAAA) Shutdown () { /* This interface can be used to gracefully shutdown
	your mylware. */

	iShutdown_AAAAAA ()
}
