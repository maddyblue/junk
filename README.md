# appstats

A port of the python appstats implementation to the Go runtime on Google App Engine.

Currently in early development: useful but possibly dangerous or wrong. **Does not run in production** due to appengine's implementation. This is being worked on.

Compatible with [`github.com/gorilla/mux`](http://www.gorillatoolkit.org/pkg/mux) and [`github.com/mjibson/goon`](https://github.com/mjibson/goon).

## installation

In your main `.go` file:

```import "github.com/mjibson/appstats"```

Add to the handler section in `init()`:

```http.HandleFunc("/_ah/stats/", appstats.AppstatsHandler)```

Change all handler functions to the following signature:

```func(appengine.Context, http.ResponseWriter, *http.Request)```

Wrap all calls to those functions in the `appstats.NewHandler` wrapper:

```http.Handle("/", appstats.NewHandler(Main))```

## example

```
import "appengine"
import "github.com/mjibson/appstats"
import "net/http"

func init() {
	http.Handle("/", appstats.NewHandler(Main))
	http.HandleFunc("/_ah/stats/", appstats.AppstatsHandler)
}

func Main (c appengine.Context, w http.ResponseWriter, r *http.Request) {
	// do stuff with c: datastore.Get(c, key, entity)
	w.Write([]byte("success"))
}
```

## usage

Do things and view at [http://localhost:8080/_ah/stats/](http://localhost:8080/_ah/stats/) like normal.

## optional configuration

Refer to the [variables section](http://godoc.org/github.com/mjibson/appstats#_variables) of the documentation.

## todo

Cost calculation is experimental. Currently it only includes write ops (read and small ops are TODO).
