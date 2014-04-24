// Web servers
package main

import (
	"fmt"
	"net/http"
)

type String string

func (self String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, self)
}

type Struct struct {
	Greeting	string
	Punct		string
	Who			string
}

func (self Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, self.Greeting, self.Punct, self.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"A", ":", "B"})
	http.ListenAndServe("localhost:4000", nil)
}
