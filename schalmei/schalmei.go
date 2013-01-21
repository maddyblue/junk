package schalmei

import (
	"appengine"
	"appengine/blobstore"
	"appengine/datastore"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mjibson/goon"
)

var router = new(mux.Router)
var templates *template.Template

type rankData struct {
	Id int64
	Rank *Rank
	Url string
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

	router.HandleFunc("/", Main).Name("main")
	router.HandleFunc("/rank/create", RankCreate).Name("create-rank")
	router.HandleFunc("/rank/list", RankList).Name("list-ranks")
	router.HandleFunc("/rank/get/{id:[0-9]+}", RankGet).Name("get-rank")
	router.HandleFunc("/upload-url/{id:[0-9]+}", UploadUrl).Name("upload-url")
	router.HandleFunc("/upload-success/{id:[0-9]+}", UploadSuccess).Name("upload-success")
	http.Handle("/", router)
}

func serveError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	fmt.Println(err)
}

func Main(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "base.html", nil)

	if err != nil {
		serveError(w, err)
	}
}

func RankCreate(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	var g Rank

	n := goon.NewGoon(r)
	e, err := n.NewEntity(nil, &g)

	g.Name = string(b)

	err = n.Put(e)
	if err != nil {
		serveError(w, err)
		return
	}

	b, _ = json.Marshal(e.Key.IntID())
	w.Write(b)
}

func RankList(w http.ResponseWriter, r *http.Request) {
	g := goon.NewGoon(r)
	q := datastore.NewQuery("Rank")
	var gg []*Rank
	es, _ := g.GetAll(q, &gg)

	rs := make([]rankData, len(es))
	for i, e := range es {
		url, _ := router.Get("get-rank").URL("id", strconv.FormatInt(e.Key.IntID(), 10))
		rs[i] = rankData{
			Id: e.Key.IntID(),
			Rank: e.Src.(*Rank),
			Url: url.String(),
		}
	}

	b, _ := json.Marshal(rs)
	w.Write(b)
}

func RankGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s := vars["id"]
	id, _ := strconv.ParseInt(s, 10, 64)

	n := goon.NewGoon(r)
	g := &Rank{}
	e, _ := n.GetById(g, "", id, nil)
	u, _ := router.Get("upload-url").URL("id", s)
	b, _ := json.Marshal(rankData{
		Id: e.Key.IntID(),
		Rank: g,
		Url: u.String(),
	})
	w.Write(b)
}

func UploadUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c := appengine.NewContext(r)
	url, _ := router.Get("upload-success").URL("id", vars["id"])
	url, _ = blobstore.UploadURL(c, url.String(), nil)
	b, _ := json.Marshal(url.String())
	w.Write(b)
}

func UploadSuccess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	blobs, values, err := blobstore.ParseUpload(r)
	if err != nil {
		serveError(w, err)
		return
	}

	n := goon.NewGoon(r)
	var rp Rank
	rid, _ := strconv.ParseInt(vars["id"], 10, 64)
	rank, err := n.GetById(&rp, "", rid, nil)
	if err != nil || rank.NotFound {
		serveError(w, nil)
		return
	}

	file := blobs["file"]
	if len(file) == 0 {
		serveError(w, nil)
		return
	}

	freq, err := strconv.ParseFloat(values.Get("freq"), 64)
	if err != nil {
		serveError(w, err)
		return
	}

	note := Note{
		Freq: freq,
		Blob: file[0].BlobKey,
	}

	e, _ := n.NewEntity(rank.Key, &note)
	_ = n.Put(e)
	url, _ := router.Get("get-rank").URL("id", vars["id"])
	http.Redirect(w, r, url.String(), http.StatusFound)
}
