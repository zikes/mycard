package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/zikes/box"
	"github.com/zikes/card"
)

func main() {
	// The color library uses tty detection to automatically disable colors if
	// stdout is not a terminal. If the process is started by e.g. systemd
	// that would result in un-colored output, so this disables detection.
	color.NoColor = false

	banner := ` ╺━┓╻╻┏ ┏━╸┏━┓
 ┏━┛┃┣┻┓┣╸ ┗━┓
 ┗━╸╹╹ ╹┗━╸┗━┛
Jason Hutchinson`

	b := &box.Box{
		BoxStyle:    &box.DoubleStyle,
		Margin:      1,
		Padding:     1,
		BorderColor: color.New(color.FgGreen),
	}

	sections := []interface{}{
		box.Section{
			Content:      banner,
			BorderColor:  color.New(color.FgGreen),
			ContentColor: color.New(color.FgGreen, color.Bold),
			Padding:      1,
			Alignment:    box.CenterAlign,
		},
		box.Section{
			Padding:     2,
			LinePadding: 1,
			Content: card.Section{
				Items: []card.Item{
					{Label: "Work", Data: "WEHCO Media, Inc"},
					{"", ""},
					{Label: "Email", Data: "zikes@zikes.me"},
					{Label: "Web", Data: "https://zikes.me"},
					{Label: "Twitter", Data: "https://twitter.com/zikes"},
					{Label: "GitHub", Data: "https://github.com/zikes"},
				},
			},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, b.Sprint(sections...))
	})

	http.ListenAndServe("localhost:8000", nil)
}
