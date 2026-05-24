// main.go
package main

import (
	_ "embed"
	"github.com/webview/webview_go"
)

//go:embed index.html
var htmlContent string

func main() {
	// Initialize context with hardware rendering optimization
	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("ArchFlow WorkSpace Engine")
	w.SetSize(1200, 800, webview.HintNone)

	// Stream embedded layout fully offline
	w.SetHtml(htmlContent)
	w.Run()
}
