package main

import (
	"image/color"
	"image/png"
	"io"
	"math/rand"
	"os"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/psykhi/wordclouds"
)

var wcData = map[string]interface{}{
	"恩n":                  10000,
	"岁":                   6181,
	"Ywsmer":              4386,
	"Yws World":           4055,
	"YwsCommunications":   2467,
	"Ywsl A":              2244,
	"Ywsitness":           1898,
	"Ywsrfect":            1484,
	"Yws:":                1689,
	"Yws ":                1112,
	"Ywsepp":              985,
	"Ywsham":              847,
	"Ywsmilton":           582,
	"Yws     ":            555,
	"Ywsen Mark":          550,
	"Ywsbraham":           462,
	"Yws":                 366,
	"Ywsilliams":          282,
	"Ywseball tournament": 273,
	"Ywseak":              265,
}

func generateWCData(data map[string]interface{}) (items []opts.WordCloudData) {
	items = make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

func wcBase() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "basic WordCloud example",
		}))

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
				}),
		)
	return wc
}

func wcCardioid() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "cardioid shape"}),
	)

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "cardioid",
				}),
		)
	return wc
}

func wcStar() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "star shape",
		}))

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "star",
				}),
		)
	return wc
}

type WordcloudExamples struct{}

func (WordcloudExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		wcBase(),
		wcCardioid(),
		wcStar(),
	)

	f, err := os.Create("wordcloud.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

func main() {

	// ec-chart
	// wc := WordcloudExamples{}
	// wc.Examples()

	// wordclouds
	wordCloud()

}

// psykhi/wordclouds"
func wordCloud() {

	// word list
	inputWords := make(map[string]int, 0)
	for i := 1; i <= 100; i++ {
		key := RandomString(5)
		inputWords[key] = i
	}

	// word color
	co := []color.RGBA{
		{0x1b, 0x1b, 0x1b, 0xff},
		{0x48, 0x48, 0x4B, 0xff},
		{0x59, 0x3a, 0xee, 0xff},
		{0x65, 0xCD, 0xFA, 0xff},
		{0x70, 0xD6, 0xBF, 0xff},
		{194, 69, 39, 255},
		{38, 103, 118, 255},
		{173, 210, 224, 255},
	}
	colors := make([]color.Color, 0)
	for _, c := range co {
		colors = append(colors, c)
	}

	// display option
	oarr := []wordclouds.Option{
		wordclouds.FontMaxSize(200),
		wordclouds.FontMinSize(100),
		// wordclouds.MaskBoxes(boxes),
		wordclouds.RandomPlacement(true),
		// wordclouds.BackgroundColor(color.Black),
		wordclouds.FontFile("fonts/Roboto-Regular.ttf"),
		wordclouds.Height(2048),
		wordclouds.Width(2048),
		wordclouds.Colors(colors),
	}

	// wordcloud instance
	w := wordclouds.NewWordcloud(inputWords,
		oarr...,
	)

	img := w.Draw()
	f, err := os.Create("img.png")
	if err != nil {
		panic(err)
	}
	// write image
	png.Encode(f, img)

	// Don't forget to close files
	f.Close()
}

// random string
func RandomString(n int) string {
	const letterRunes string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(letterRunes[rand.Int63()%int64(len(letterRunes))])
	}
	return sb.String()
}
