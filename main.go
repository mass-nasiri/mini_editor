// main.go
package main

import (
	_ "embed"
	"github.com/webview/webview_go"
)

//go:embed index.html
var htmlContent string

func main() {
	// Create a new webview window with a safe 1200x800 resolution
	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("Go Ultra-Lightweight Live Editor")
	w.SetSize(1200, 800, webview.HintNone)

	// Inject the embedded HTML bundle directly into the webview component
	w.SetHtml(htmlContent)

	// Block and run the operating system native window loop
	w.Run()
}
