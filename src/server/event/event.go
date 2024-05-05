package event

import (
	"log"
)

type Event struct {
	Canceled bool
}

func (e *Event) IsCanceled() bool {
	return e.Canceled
}

func (e *Event) Cancel() {
	e.Canceled = true
}

func (e *Event) Call() bool {
	// todo
	// var pluginManager = Grasscutter.getPluginManager();
	// if (pluginManager != null) pluginManager.invokeEvent(this);
	log.Println("Pushing event to all listeners...")

	return !e.IsCanceled()
}
