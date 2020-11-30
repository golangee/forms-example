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

package router

import (
	"github.com/golangee/forms-example/www/forms/dom"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"net/url"
	"sort"
	"strings"
)

type Route struct {
	Path        string
	Constructor func(q Query)
}

type Router struct {
	routes2Actions map[string]func(q Query)
	funcs          []dom.Releasable
	lastLocation   string
	lastFragment   string
	unmatchedRoute func(Query)
	log            log.Logger
}

func NewRouter() *Router {
	r := &Router{
		routes2Actions: make(map[string]func(Query)),
		log:            log.NewLogger(ecs.Log("router")),
	}

	r.funcs = append(r.funcs, dom.GetWindow().OnHashChange(r.checkLocation))
	r.funcs = append(r.funcs, dom.GetWindow().HashChange(r.checkLocation))
	r.funcs = append(r.funcs, dom.GetWindow().OnPopState(r.checkLocation))

	r.lastLocation = "$%&/"
	r.lastFragment = r.lastLocation
	return r
}

func (r *Router) Routes() []Route {
	var res []Route
	for k, v := range r.routes2Actions {
		res = append(res, Route{
			Path:        k,
			Constructor: v,
		})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Path < res[j].Path
	})
	return res
}

func (r *Router) AddRoute(path string, f func(Query)) *Router {
	pIdx := strings.IndexRune(path, '?')
	if pIdx >= 0 {
		path = path[:pIdx]
	}

	r.log.Print(ecs.Msg("registered route"), ecs.URLPath(path))
	r.routes2Actions[path] = f
	return r
}

func (r *Router) SetUnhandledRouteAction(f func(Query)) *Router {
	r.unmatchedRoute = f
	return r
}

func (r *Router) Start() {
	r.checkLocation()
}

func (r *Router) Reload(force bool) {
	dom.GetWindow().Location().Reload(force)
}

func (r *Router) Invalidate() error {
	f, err := url.Parse(r.lastFragment)
	if err != nil {
		return err
	}

	r.onFragmentChanged(f.Path, f.Query())
	return nil
}

func (r *Router) Release() {
	for _, f := range r.funcs {
		f.Release()
	}
}

func (r *Router) checkLocation() {
	defer dom.GlobalPanicHandler()

	location := dom.GetWindow().Location().Href()
	if r.lastLocation != location {
		u, err := url.Parse(location)
		if err != nil {
			r.log.Print(ecs.Msg("Failed to parse location"), ecs.URLPath(location), ecs.ErrMsg(err))
			return
		}
		r.onLocationChanged(u)
		r.lastLocation = location

		if u.Fragment != r.lastFragment {
			f, err := url.Parse(u.Fragment)
			if err != nil {
				r.log.Print(ecs.Msg("Failed to parse fragment as url"), ecs.URLPath(u.String()), ecs.ErrMsg(err))
				return
			}
			r.onFragmentChanged(f.Path, f.Query())
			r.lastFragment = u.Fragment
		}

	}
}

func (r *Router) onLocationChanged(location *url.URL) {

}

func (r *Router) onFragmentChanged(path string, query url.Values) {
	defer dom.GlobalPanicHandler()

	if path == "" {
		path = "/"
	}
	q := Query{values: query, path: path}
	f := r.routes2Actions[path]
	if f != nil {
		f(q)
	} else {
		if r.unmatchedRoute != nil {
			r.unmatchedRoute(q)
		} else {
			r.log.Print(ecs.Msg("unmatched route"), ecs.URLPath(path))
		}
	}
}

func (r *Router) Navigate(u *url.URL) {
	Navigate(u.String())
}

// Navigate issues a navigation request to the window element.
func Navigate(url string) {
	dom.Post(func() {
		dom.GetWindow().SetLocation(url)
	})
}
