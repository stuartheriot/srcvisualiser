package charts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/gin-gonic/gin"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func createBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(400)})
	}
	return items
}

func BuildChart(ctx *gin.Context) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My first Bar Chart with go-echarts ",
		Subtitle: "Seems easy. Basic charts all good.",
	}))

	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Cat A", createBarItems()).
		AddSeries("Cat B", createBarItems())

	bar.Render(ctx.Writer)
}

// ------------------------- 3D BAR CHART FOR FILESTATS ---------------------------------
var (
	bar3DRangeColor = []string{
		"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
		"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
	}

	bar3DHrs = [...]string{
		"12a", "1a", "2a", "3a", "4a", "5a", "6a", "7a", "8a", "9a", "10a", "11a",
		"12p", "1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p", "9p", "10p", "11p",
	}

	bar3DFiletypes = [...]string{"java", "javascript", "html", "yaml", "json", "other"}

	bar3DFiletypeIndex = map[string]int{"java": 0, "javascript": 1, "html": 2, "yaml": 3, "json": 4, "other": 5}
)

type FileStat struct {
	Repo              string `json:"repo"`
	Tier              string `json:"tier"`
	Name              string `json:"name"`
	Path              string `json:"path"`
	Filetype          string `json:"filetype"`
	Category          string `json:"category"`
	NumCommits        int    `json:"numCommits"`
	NumLines          int    `json:"locTotal"`
	NumAdded          int    `json:"locAdded"`
	NumRemoved        int    `json:"locRemoved"`
	NumAuthors        int    `json:"numAuthors"`
	FirstCommitAuthor string `json:"firstCommitAuthor"`
	MostCommitsAuthor string `json:"mostCommitsAuthor"`
}

func ChartFileData(ctx *gin.Context, filename string) {
	bar3d := charts.NewBar3D()
	bar3d.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "DATAFILE:" + filename,
		Subtitle: "3D Bar Chart",
	}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        3000,
			Range:      []float32{0, 300},
			InRange:    &opts.VisualMapInRange{Color: bar3DRangeColor},
		}),
		charts.WithGrid3DOpts(opts.Grid3D{
			BoxWidth: 300,
			BoxDepth: 150,
		}),
	)

	bar3dData, bar3Dscales := buildFilestats3dData(filename)

	bar3d.SetGlobalOptions(
		charts.WithXAxis3DOpts(opts.XAxis3D{Data: bar3Dscales["X"]}),
		charts.WithYAxis3DOpts(opts.YAxis3D{Data: bar3DFiletypes, Name: "Type"}),
		charts.WithZAxis3DOpts(opts.ZAxis3D{Name: "LOC", Data: bar3Dscales["Z"]}),
	)

	bar3d.AddSeries("Filestats", bar3dData)
	bar3d.Render(ctx.Writer)
}

// type Chart3DData struct {
// 	// Name of the data item.
// 	Name string `json:"name,omitempty"`

// 	// Value of the data item.
// 	// []interface{}{1, 2, 3}
// 	Value []interface{} `json:"value,omitempty"`

// 	// ItemStyle settings in this series data.
// 	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

// 	// The style setting of the text label in a single bar.
// 	Label *Label `json:"label,omitempty"`
// }

type DataPoint struct {
	Commits int
	Type    int
	Count   int
}

const CommitIncremens = 10

func buildFilestats3dData(fname string) ([]opts.Chart3DData, map[string][]string) {
	data := []opts.Chart3DData{}
	scales := map[string][]string{"X": {}, "Y": {}, "Z": {}}

	// Read in JSON file
	jsondata := []FileStat{}

	file, err := ioutil.ReadFile("./data/" + fname)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &jsondata)
	if err != nil {
		log.Fatal(err)
	}

	// for _, d := range bar3DDays {
	// 	ret = append(ret, opts.Chart3DData{
	// 		Value: []interface{}{d[0], d[1], d[2]},
	// 	})
	// }

	maxloc := 0
	maxCommits := 0
	for _, v := range jsondata {
		if v.NumLines > maxloc {
			maxloc = v.NumLines
		}
		if v.NumCommits > maxCommits {
			maxCommits = v.NumCommits
		}
	}
	// Determine increment size for total commits with max of 10
	scales["X"] = buildIntScale(CommitIncremens, maxCommits)
	scales["Z"] = buildIntScale(CommitIncremens, maxloc)

	datapoints := []map[string]int{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}
	for _, v := range jsondata {
		idx := v.NumCommits / CommitIncremens
		datapoints[idx][v.Filetype] += 1
	}
	for k, v := range datapoints {
		for fk, cnt := range v {
			data = append(data, opts.Chart3DData{
				Value: []interface{}{k, fk, cnt}})
		}
	}
	return data, scales
}

func buildIntScale(numInc int, maxVal int) []string {
	bar3Dscale := []string{}

	size := maxVal / numInc

	for i := 0; i < numInc+1; i++ {
		bar3Dscale = append(bar3Dscale, fmt.Sprintf("%d", i*size))
	}

	return bar3Dscale
}
