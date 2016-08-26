package schalmei

import (
	"appengine"
	"appengine/blobstore"
	"encoding/json"
	"fmt"

	"github.com/mjibson/go-dsp/spectral"
	"github.com/mjibson/go-dsp/wav"
)

type Rank struct {
	Name string
}

type Note struct {
	Freq float64
	Blob appengine.BlobKey
}

func (n *Note) Wav(c appengine.Context) (*wav.Wav, error) {
	br := blobstore.NewReader(c, n.Blob)
	return wav.ReadWav(br)
}

func (n *Note) GraphUrl(key string) string {
	u, _ := router.Get("note-graph").URL("key", key)
	return u.String()
}

func (n *Note) PwelchUrl(key string) string {
	u, _ := router.Get("note-pwelch").URL("key", key)
	return u.String()
}

func (n *Note) Chart(w *wav.Wav, reqid string) string {
	cols := make([]Col, w.NumChannels+1)
	cols[0] = Col{
		Id:    "sample",
		Label: "Sample #",
		Type:  "number",
	}

	for i := 1; i <= int(w.NumChannels); i++ {
		cols[i] = Col{
			Id:    fmt.Sprintf("ch_%v", i),
			Label: fmt.Sprintf("Channel %v", i),
			Type:  "number",
		}
	}

	NS_LIM := 50
	ns := w.NumSamples
	if ns > NS_LIM {
		ns = NS_LIM
	}
	rows := make([]Row, ns)
	for i := 0; i < ns; i++ {
		cells := make([]Cell, len(cols))
		cells[0] = Cell{
			Value: i,
		}
		for j := 1; j <= int(w.NumChannels); j++ {
			cells[j] = Cell{
				Value: w.Data[j-1][i],
			}
		}
		rows[i] = Row{
			Cells: cells,
		}
	}

	dt := DataTable{
		Cols: cols,
		Rows: rows,
	}
	qr := QueryResponse{
		ReqId:   reqid,
		Status:  "ok",
		Version: "0.6",
		Table:   dt,
	}

	b, _ := json.Marshal(qr)
	return string(b)
}

func (n *Note) PwelchChart(w *wav.Wav, reqid string) string {
	cols := make([]Col, w.NumChannels+1)
	cols[0] = Col{
		Id:    "frequency",
		Label: "Frequency (Hz)",
		Type:  "number",
	}

	Pxx := make([][]float64, w.NumChannels)
	freqs := make([][]float64, w.NumChannels)
	po := &spectral.PwelchOptions{
		NFFT: 1 << 13,
	}
	for i := 0; i < int(w.NumChannels); i++ {
		x := make([]float64, w.NumSamples)
		for j := 0; j < w.NumSamples; j++ {
			x[j] = float64(w.Data[i][j])
		}
		Pxx[i], freqs[i] = spectral.Pwelch(x, float64(w.SampleRate), po)
	}

	for i := 1; i <= int(w.NumChannels); i++ {
		cols[i] = Col{
			Id:    fmt.Sprintf("ch_%v", i),
			Label: fmt.Sprintf("Channel %v", i),
			Type:  "number",
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	ns := min(len(Pxx[0]), 200)
	rows := make([]Row, ns)
	for i := 0; i < ns; i++ {
		cells := make([]Cell, len(cols))
		cells[0] = Cell{
			Value: freqs[0][i],
		}
		for j := 1; j <= int(w.NumChannels); j++ {
			cells[j] = Cell{
				Value: Pxx[j-1][i],
			}
		}
		rows[i] = Row{
			Cells: cells,
		}
	}

	dt := DataTable{
		Cols: cols,
		Rows: rows,
	}
	qr := QueryResponse{
		ReqId:   reqid,
		Status:  "ok",
		Version: "0.6",
		Table:   dt,
	}

	b, _ := json.Marshal(qr)
	return string(b)
}

type QueryResponse struct {
	ReqId   string    `json:"reqId"`
	Status  string    `json:"status"`
	Version string    `json:"version"`
	Table   DataTable `json:"table"`
}

type DataTable struct {
	Cols []Col       `json:"cols"`
	Rows []Row       `json:"rows"`
	P    interface{} `json:"p,omitempty"`
}

type Col struct {
	Id    string      `json:"id,omitempty"`
	Label string      `json:"label,omitempty"`
	Type  string      `json:"type"`
	P     interface{} `json:"p,omitempty"`
}

type Row struct {
	Cells []Cell `json:"c"`
}

type Cell struct {
	Value  interface{} `json:"v,omitempty"`
	Format string      `json:"f,omitempty"`
	P      interface{} `json:"p,omitempty"`
}
