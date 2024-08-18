package gochartjs

import (
	"regexp"

	"github.com/goccy/go-json"
)

type ChartJs_Data struct {
	Labels   []string          `json:"labels"`
	Datasets []ChartJs_Dataset `json:"datasets"`
}

type ChartJs_Dataset struct {
	Label           string  `json:"label"`
	Data            []int   `json:"data"`
	Fill            bool    `json:"fill"`
	BackgroundColor string  `json:"backgroundColor,omitempty"`
	BorderColor     string  `json:"borderColor,omitempty"`
	BorderWidth     int     `json:"borderWidth,omitempty"`
	LineTension     float64 `json:"lineTension,omitempty"`
}

type ChartJs_Options struct {
	Responsive bool                   `json:"responsive"`
	Scales     ChartJs_Scales         `json:"scales"`
	Plugins    ChartJs_PluginsOptions `json:"plugins"`
}

type ChartJs_PluginsOptions struct {
	Legend ChartJs_Legend `json:"legend,omitempty"`
	Title  ChartJs_Title  `json:"title,omitempty"`
}

// display	boolean	true	Is the legend shown?
// position	string	'top'	Position of the legend. more...
// align	string	'center'	Alignment of the legend. more...
// maxHeight	number		Maximum height of the legend, in pixels
// maxWidth	number		Maximum width of the legend, in pixels
// fullSize	boolean	true	Marks that this box should take the full width/height of the canvas (moving other boxes). This is unlikely to need to be changed in day-to-day use.
// onClick	function		A callback that is called when a click event is registered on a label item. Arguments: [event, legendItem, legend].
// onHover	function		A callback that is called when a 'mousemove' event is registered on top of a label item. Arguments: [event, legendItem, legend].
// onLeave	function		A callback that is called when a 'mousemove' event is registered outside of a previously hovered label item. Arguments: [event, legendItem, legend].
// reverse	boolean	false	Legend will show datasets in reverse order.
// labels	object		See the Legend Label Configuration section below.
// rtl	boolean		true for rendering the legends from right to left.
// textDirection	string	canvas' default	This will force the text direction 'rtl' or 'ltr' on the canvas for rendering the legend, regardless of the css specified on the canvas
// title	object		See the Legend Title Configuration section below.

type ChartJs_Legend struct {
	Display       bool                 `json:"display"`
	Position      string               `json:"position,omitempty"`
	Align         string               `json:"align,omitempty"`
	MaxHeight     int                  `json:"maxHeight,omitempty"`
	MaxWidth      int                  `json:"maxWidth,omitempty"`
	FullSize      bool                 `json:"fullSize,omitempty"`
	OnClick       jsFunc               `json:"onClick,omitempty"`
	OnHover       jsFunc               `json:"onHover,omitempty"`
	OnLeave       jsFunc               `json:"onLeave,omitempty"`
	Reverse       bool                 `json:"reverse,omitempty"`
	Labels        ChartJs_LegendLabels `json:"labels,omitempty"`
	Rtl           bool                 `json:"rtl,omitempty"`
	TextDirection string               `json:"textDirection,omitempty"`
	Title         ChartJs_Title        `json:"title,omitempty"`
}

// boxWidth	number	40	Width of coloured box.
// boxHeight	number	font.size	Height of the coloured box.
// color
// Color
// Chart.defaults.color	Color of label and the strikethrough.
// font	Font	Chart.defaults.font	See Fonts
// padding	number	10	Padding between rows of colored boxes.
// generateLabels	function		Generates legend items for each thing in the legend. Default implementation returns the text + styling for the color box. See Legend Item for details.
// filter	function	null	Filters legend items out of the legend. Receives 2 parameters, a Legend Item and the chart data.
// sort	function	null	Sorts legend items. Type is : sort(a: LegendItem, b: LegendItem, data: ChartData): number;. Receives 3 parameters, two Legend Items and the chart data. The return value of the function is a number that indicates the order of the two legend item parameters. The ordering matches the return value
//  of Array.prototype.sort()
// pointStyle
// pointStyle
// 'circle'	If specified, this style of point is used for the legend. Only used if usePointStyle is true.
// textAlign	string	'center'	Horizontal alignment of the label text. Options are: 'left', 'right' or 'center'.
// usePointStyle	boolean	false	Label style will match corresponding point style (size is based on pointStyleWidth or the minimum value between boxWidth and font.size).
// pointStyleWidth	number	null	If usePointStyle is true, the width of the point style used for the legend.
// useBorderRadius	boolean	false	Label borderRadius will match corresponding borderRadius.
// borderRadius	number	undefined	Override the borderRadius to use.

type ChartJs_LegendLabels struct {
	BoxWidth        int    `json:"boxWidth,omitempty"`
	BoxHeight       int    `json:"boxHeight,omitempty"`
	Color           string `json:"color,omitempty"`
	Font            string `json:"font,omitempty"`
	Padding         int    `json:"padding,omitempty"`
	GenerateLabels  jsFunc `json:"generateLabels,omitempty"`
	Filter          jsFunc `json:"filter,omitempty"`
	Sort            jsFunc `json:"sort,omitempty"`
	PointStyle      string `json:"pointStyle,omitempty"`
	TextAlign       string `json:"textAlign,omitempty"`
	UsePointStyle   bool   `json:"usePointStyle,omitempty"`
	PointStyleWidth int    `json:"pointStyleWidth,omitempty"`
	UseBorderRadius bool   `json:"useBorderRadius,omitempty"`
	BorderRadius    int    `json:"borderRadius,omitempty"`
}

type ChartJs_Scales struct {
	XAxes ChartJs_Axes `json:"x,omitempty"`
	YAxes ChartJs_Axes `json:"y,omitempty"`
}

type ChartJs_Axes struct {
	Display bool           `json:"display"`
	Ticks   ChartJs_Ticks  `json:"ticks"`
	Border  ChartJs_Border `json:"border"`
	Grid    ChartJs_Grid   `json:"grid"`
}

type ChartJs_Ticks struct {
	BeginAtZero bool `json:"beginAtZero"`
}

type ChartJs_Border struct {
	Display bool `json:"display"`
}

type ChartJs_Grid struct {
	Display bool `json:"display"`
}

// color
// Color
// Chart.defaults.color	Color of text.
// display	boolean	false	Is the legend title displayed.
// font	Font	Chart.defaults.font	See Fonts
// padding
// Padding
// 0	Padding around the title.
// text	string		The string title.

type ChartJs_Title struct {
	Display bool   `json:"display"`
	Font    string `json:"font"`
	Padding int    `json:"padding"`
	Text    string `json:"text"`
}

type ChartJs struct {
	Type    string           `json:"type"`
	Data    ChartJs_Data     `json:"data"`
	Options ChartJs_Options  `json:"options"`
	Plugins []ChartJs_Plugin `json:"plugins,omitempty"`
}

type ChartJs_Plugin struct {
	Id         string `json:"id"`
	BeforeDraw jsFunc `json:"beforeDraw"`
}

type jsFunc struct {
	Func       string `json:"func"`       // js function
	Entry      string `json:"entry"`      // js entry point
	ReturnType string `json:"returnType"` // js return type (string, int, float, bool, object, array)
}

// New creates a new ChartJs object
func New(chartType uint8) ChartJs {
	return ChartJs{
		Type: chartTypes[chartType],
		Data: ChartJs_Data{
			Labels:   []string{},
			Datasets: []ChartJs_Dataset{},
		},
		Options: ChartJs_Options{
			Responsive: true,
		},
	}
}

// output as json
func (c ChartJs) Json() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// Chart types
const (
	ChartTypeBar uint8 = iota
	ChartTypeBubble
	ChartTypeDoughnut
	ChartTypeLine
	ChartTypePie
	ChartTypePolarArea
	ChartTypeRadar
	ChartTypeScatter
)

var chartTypes = map[uint8]string{
	ChartTypeBar:       "bar",
	ChartTypeBubble:    "bubble",
	ChartTypeDoughnut:  "doughnut",
	ChartTypeLine:      "line",
	ChartTypePie:       "pie",
	ChartTypePolarArea: "polarArea",
	ChartTypeRadar:     "radar",
	ChartTypeScatter:   "scatter",
}

// output as js object
func (c ChartJs) Js(elementId string) string {
	return "new Chart(document.getElementById('myChart'), " + c.JsObj() + ")"
}

func (c ChartJs) JsObj() string {
	inJson := c.Json()
	// replace double quotes with no quotes on specific values to expose js functions
	// `"(\((?P<expression>[\w]+)\)\?\s?(\'[\w, #\(\)]+\')\s?:\s?(\'[\w, #\(\)]+\');?)"`gm
	// use this regex to find the strings
	var re = regexp.MustCompile(`(?m)"(\((?P<expression>[\w]+)\)\?\s?(\'[\w, #\(\)]+\')\s?:\s?(\'[\w, #\(\)]+\'));?"`)

	// replace the strings
	inJson = re.ReplaceAllString(inJson, "$1")

	return inJson

}
