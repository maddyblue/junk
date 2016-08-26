package schalmei

import (
	"appengine/blobstore"
	"appengine/datastore"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mjibson/MiniProfiler/go/miniprofiler"
	mpg "github.com/mjibson/MiniProfiler/go/miniprofiler_gae"
	"github.com/mjibson/goon"
)

var router = new(mux.Router)
var templates *template.Template

type rankData struct {
	Id    int64
	Rank  *Rank
	Url   string
	Notes []*noteData
}

type noteData struct {
	Note      *Note
	GraphUrl  string
	PwelchUrl string
}

func init() {
	var err error

	templates, err = template.New("").Funcs(funcs).
		ParseFiles(
		"templates/base.html",
	)

	if err != nil {
		log.Print(err)
		return
	}

	router.Handle("/", mpg.NewHandler(Main)).Name("main")
	router.Handle("/rank/create", mpg.NewHandler(RankCreate)).Name("create-rank")
	router.Handle("/rank/list", mpg.NewHandler(RankList)).Name("list-ranks")
	router.Handle("/rank/get/{id:[0-9]+}", mpg.NewHandler(RankGet)).Name("get-rank")
	router.Handle("/upload-url/{id:[0-9]+}", mpg.NewHandler(UploadUrl)).Name("upload-url")
	router.Handle("/upload-success/{id:[0-9]+}", mpg.NewHandler(UploadSuccess)).Name("upload-success")
	router.Handle("/note/graph/{key}", mpg.NewHandler(NoteGraph)).Name("note-graph")
	router.Handle("/note/pwelch/{key}", mpg.NewHandler(NotePwelch)).Name("note-pwelch")
	http.Handle("/", router)

	miniprofiler.Position = "right"
	miniprofiler.ShowControls = false
	miniprofiler.Enable = func(r *http.Request) bool { return true }
}

func serveError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	fmt.Println("serve error:", err)
}

func Main(c mpg.Context, w http.ResponseWriter, r *http.Request) {
	/*
		datastore.AllocateIDs(c, "raank", nil, 5)
		n := goon.FromContext(c)
		q := datastore.NewQuery("raank")
		q = q.KeysOnly()
		q = q.Limit(10)
		var gg []Rank
		n.GetAll(q, &gg)
	*/

	err := templates.ExecuteTemplate(w, "base.html", miniprofiler.Includes(r, c.P))
	if err != nil {
		serveError(w, err)
	}
}

func RankCreate(c mpg.Context, w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	var g Rank

	n := goon.FromContext(c)
	e, err := n.NewEntity(nil, &g)

	g.Name = string(b)

	err = n.Put(e)
	if err != nil {
		serveError(w, err)
		return
	}

	url, err := router.Get("get-rank").URL("id", strconv.FormatInt(e.Key.IntID(), 10))
	b, err = json.Marshal(url.String())
	w.Write(b)
}

func RankList(c mpg.Context, w http.ResponseWriter, r *http.Request) {
	g := goon.FromContext(c)
	q := datastore.NewQuery(goon.Kind(Rank{}))
	var gg []*Rank
	es, _ := g.GetAll(q, &gg)

	rs := make([]rankData, len(es))
	for i, e := range es {
		url, _ := router.Get("get-rank").URL("id", strconv.FormatInt(e.Key.IntID(), 10))
		rs[i] = rankData{
			Id:   e.Key.IntID(),
			Rank: gg[i],
			Url:  url.String(),
		}
	}

	b, _ := json.Marshal(rs)
	w.Write(b)
}

func RankGet(c mpg.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := vars["id"]
	id, _ := strconv.ParseInt(s, 10, 64)

	n := goon.FromContext(c)

	c.P.Step("get by id", func() {
		g := &Rank{}
		e, err := n.GetById(g, "", id, nil)
		if err != nil {
			serveError(w, err)
			return
		}

		c.P.Step("query note", func() {
			q := datastore.NewQuery("Note").Ancestor(e.Key)
			ndata := []*Note{}
			ns, _ := n.GetAll(q, &ndata)
			notes := make([]*noteData, len(ndata))
			for i, ndat := range ndata {
				notes[i] = &noteData{
					ndat,
					ndat.GraphUrl(ns[i].Key.Encode()),
					ndat.PwelchUrl(ns[i].Key.Encode()),
				}
			}

			u, _ := router.Get("upload-url").URL("id", s)
			b, _ := json.Marshal(rankData{
				Id:    e.Key.IntID(),
				Rank:  g,
				Url:   u.String(),
				Notes: notes,
			})
			w.Write(b)
		})
	})
}

func UploadUrl(c mpg.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url, _ := router.Get("upload-success").URL("id", vars["id"])
	url, _ = blobstore.UploadURL(c, url.String(), nil)
	b, _ := json.Marshal(url.String())
	w.Write(b)
}

func UploadSuccess(c mpg.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	blobs, values, err := blobstore.ParseUpload(r)

	del_blobs := func() {
		for _, binfos := range blobs {
			for _, binfo := range binfos {
				blobstore.Delete(c, binfo.BlobKey)
			}
		}
	}

	if err != nil {
		serveError(w, err)
		del_blobs()
		return
	}

	n := goon.FromContext(c)
	var rp Rank
	rid, _ := strconv.ParseInt(vars["id"], 10, 64)
	rank, err := n.GetById(&rp, "", rid, nil)
	if err != nil || rank.NotFound {
		serveError(w, nil)
		del_blobs()
		return
	}

	file := blobs["file"]
	if len(file) == 0 {
		serveError(w, nil)
		del_blobs()
		return
	}

	freq, err := strconv.ParseFloat(values.Get("freq"), 64)
	if err != nil {
		serveError(w, err)
		del_blobs()
		return
	}

	note := Note{
		Freq: freq,
		Blob: file[0].BlobKey,
	}

	wav, err := note.Wav(c)
	_ = wav
	if err != nil {
		serveError(w, err)
		del_blobs()
		return
	}

	e, _ := n.NewEntity(rank.Key, &note)
	_ = n.Put(e)
	url, _ := router.Get("get-rank").URL("id", vars["id"])
	http.Redirect(w, r, url.String(), http.StatusFound)
}

func NoteGraph(c mpg.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := datastore.DecodeKey(vars["key"])
	n := &Note{}
	g := goon.FromContext(c)
	_, _ = g.Get(n, key)
	wv, _ := n.Wav(c)

	tqx := r.URL.Query().Get("tqx")
	reqId := strings.Split(tqx, ":")[1]
	fmt.Fprintf(w, n.Chart(wv, reqId))
}

func NotePwelch(c mpg.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := datastore.DecodeKey(vars["key"])
	n := &Note{}
	g := goon.FromContext(c)
	_, _ = g.Get(n, key)
	wv, _ := n.Wav(c)

	tqx := r.URL.Query().Get("tqx")
	reqId := strings.Split(tqx, ":")[1]
	c.P.Step("pwelch", func() {
		fmt.Fprintf(w, n.PwelchChart(wv, reqId))
	})
}
