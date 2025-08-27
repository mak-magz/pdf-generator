// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a simple hello, world demonstration web server.
//
// It serves version information on /version and answers
// any other request like /name by saying "Hello, name!".
//
// See golang.org/x/example/outyet for a more sophisticated server.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mak-magz/pdf-generator/internal/router"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: helloserver [options]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	addr = flag.String("addr", "localhost:8080", "address to serve")
)

func main() {
	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	// Parse and validate arguments (none).
	args := flag.Args()
	if len(args) != 0 {
		usage()
	}

	router := router.SetupRouter()

	router.LoadHTMLGlob("pkg/templates/*")
	router.Static("/pkg/templates", "./pkg/templates")

	srv := &http.Server{
		Addr:    *addr,
		Handler: router,
	}

	log.Printf("serving http://%s\n", *addr)
	log.Fatal(srv.ListenAndServe())
}
