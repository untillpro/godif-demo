/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"flag"
	"log"
	"path"

	"github.com/untillpro/godif-demo/answerer"
	"github.com/untillpro/godif-demo/iconfigfile"
	"github.com/untillpro/godif-demo/iconfigmem"
	"github.com/untillpro/godif-demo/ikvdbbbolt"
	"github.com/untillpro/godif-demo/ikvdbmem"
	"github.com/untillpro/godif-demo/ui"
	"github.com/untillpro/godif/services"
)

func main() {

	// Parse command line

	pInMem := flag.Bool("m", false, "Use in-memory key-value database")
	pVerbose := flag.Bool("v", false, "Use verbose output")
	flag.Parse()

	// Declare ikvdb implementation depending on `-m` option
	if *pInMem {
		ikvdbmem.Declare()
		iconfigmem.Declare()
	} else {
		iconfigfile.Declare(path.Join(".", ".data"))
		ikvdbbbolt.Declare(path.Join(".data", "answers.db"))
	}

	// Declare answerer service
	answerer.Declare()

	// Declare ui service
	ui.Declare(*pInMem, "ui")

	// Run
	services.SetVerbose(*pVerbose)
	err := services.Run()

	if err != nil {
		log.Println("Error :", err)
	}

}
