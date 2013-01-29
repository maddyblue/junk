# appstats

A port of the python appstats implementation to the Go runtime on Google App Engine.

Currently in early development: useful but possibly dangerous or wrong.

## installation

From your app engine project's directory, run:

```git clone git://github.com/mjibson/appstats.git```

Add to `app.yaml`:

```
- url: /_ah/stats/static
  static_dir: appstats/static
```

In your main `.go` file:

```import "appstats"```

Add to the handler section in `init()`:

```http.HandleFunc("/_ah/stats/", appstats.Handler)```

Replace all instances of `appengine.NewContext` with `appstats.NewContext`.

On the line after each `appstats.NewContext`, add the line (assuming `c` is the result from `NewContext`):

```defer c.Save()```

So you should end up with:

```
c := appstats.NewContext(r)
defer c.Save()
```

## usage

Do things and view at [http://localhost:8080/_ah/stats/](http://localhost:8080/_ah/stats/) like normal.
