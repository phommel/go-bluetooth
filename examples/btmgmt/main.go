// Example use of the btmgmt wrapper
package main

import (
	"os"

	"git.enexoma.de/r/smartcontrol/libraries/go-bluetooth/linux"
	log "github.com/Sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel)

	list, err := linux.GetAdapters()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	for i, a := range list {
		log.Infof("%d) %s (%v)", i+1, a.Name, a.Addr)
	}

}
