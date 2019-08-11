package main

import (
	"os"
	"strings"

	"github.com/demandbase/gitleaks/src"
	log "github.com/sirupsen/logrus"
)

func main() {
	_, err := gitleaks.Run(gitleaks.ParseOpts())
	if err != nil {
		if strings.Contains(err.Error(), "whitelisted") {
			log.Info(err.Error())
			os.Exit(0)
		}
		log.Error(err)
		os.Exit(gitleaks.ErrExit)
	}
}
