/*
Copyright 2012 Graham King <graham@gkg.org>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

http://www.gnu.org/licenses/agpl.html
*/

/* Minimalist URL routing for Go web apps.
 Usage:
  1. Add your routes:
      route.AddRoute("a regexp here", funcName)
      funcName should be a function implementing RouteTarget
  2. Look up a route:
      route.FindRoute(url)
     Returns the Route match, and the arguments. You then call Route.Target
     giving it response, request and the args.
*/

package route

import (
    "regexp"
    "net/http"
)

var (
    URLS []Route
)

type Route struct {
    Re *regexp.Regexp
    Target RouteTarget
}
type RouteTarget func (http.ResponseWriter, *http.Request, map[string] string)

func AddRoute(url string, target RouteTarget) {
    if URLS == nil {
        URLS = make([]Route, 0)
    }

    re := regexp.MustCompile(url)
    URLS = append(URLS, Route{re, target})
}

func FindRoute(url string) (Route, map[string] string) {

    var match []string
    args := make(map[string] string)

    for _, route := range URLS {

        match = route.Re.FindStringSubmatch(url)
        if len(match) == 0 {
            continue
        }

        for idx, val := range match[1:] {
            args[route.Re.SubexpNames()[1:][idx]] = val
        }

        return route, args
    }

    return Route{}, nil
}
