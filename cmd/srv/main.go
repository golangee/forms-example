// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"github.com/lpar/gzipped/v2"
	"github.com/golangee/forms-example/build"
	"github.com/golangee/forms/theme/material"
	"log"
	"net/http"
	"path"
	"strings"
)

func withIndexHTML(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			fmt.Println(r.URL.Path)
			newpath := path.Join(r.URL.Path, "index.html")
			r.URL.Path = newpath
		}
		h.ServeHTTP(w, r)
	})
}

func version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(build.Env().String()))
}

func main() {
	log.Println("build time", build.Time)
	log.Println("build commit", build.Commit)
	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", ".", "the directory of static files to host")
	flag.Parse()

	material.Resources(http.DefaultServeMux)
	http.HandleFunc("/version", version)
	http.Handle("/", withIndexHTML(gzipped.FileServer(gzipped.Dir(*directory))))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}
