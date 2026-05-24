// main.go
package main

import (
	_ "embed"
	"github.com/webview/webview_go"
)

//go:embed index.html
var htmlContent string

func main() {
	// Debug set to false for pure clean production rendering environment
	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("Local Workspace Engine")
	w.SetSize(1200, 800, webview.HintNone)

	// Expose native window control bindings directly to HTML JavaScript sandbox safely
	w.Bind("closeApp", func() {
		w.Terminate()
	})

	w.Bind("minimizeApp", func() {
		w.SetSize(0, 0, webview.HintMin)
	})

	w.Bind("maximizeApp", func() {
		w.SetSize(1200, 800, webview.HintMax)
	})

	// Stream embedded standalone web package
	w.SetHtml(htmlContent)
	w.Run()
}
