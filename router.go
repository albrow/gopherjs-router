package router

import (
	"github.com/gopherjs/gopherjs/js"
	"strings"
	"time"
)

// Router consists of handler functions which react automatically
// when there is a change in the url hash.
type Router struct {
	routes map[string]func()
}

func getHash() string {
	return js.Global.Get("location").Get("hash").String()
}

func setHash(string) {
	js.Global.Get("location").Set("hash", "/")
}

// New returns a freshly initialized Router
func New() *Router {
	return &Router{
		routes: map[string]func(){},
	}
}

// HandleFunc will cause the router to call f whenever the
// hash of the url (everything after the '#' symbol) matches path.
func (r *Router) HandleFunc(path string, f func()) {
	r.routes[path] = f
}

// Start will listen for changes in the hash of the url and
// trigger the appropriate handler function.
func (r *Router) Start() {
	r.setInitialHash()
	// TODO detect support for onhashchange
	// and use legacy version if not supported
	r.watchHash()
}

func (r *Router) setInitialHash() {
	if hash := getHash(); hash == "" {
		setHash("/")
	} else {
		r.hashChanged(hash)
	}
}

func (r *Router) watchHash() {
	js.Global.Set("onhashchange", func() {
		r.hashChanged(getHash())
	})
}

func (r *Router) legacyWatchHash() {
	t := time.NewTicker(50 * time.Millisecond)
	go func() {
		hash := getHash()
		for {
			<-t.C
			newHash := getHash()
			if hash != newHash {
				hash = newHash
				r.hashChanged(hash)
			}
		}
	}()
}

func (r *Router) hashChanged(hash string) {
	// path is everything after the '#'
	path := strings.SplitN(hash, "#", 2)[1]
	if f, found := r.routes[path]; found {
		f()
	}
}
