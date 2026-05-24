// main.go
package main

import (
	_ "embed"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"

	"github.com/webview/webview_go"
)

//go:embed index.html
var htmlContent string

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	procGetParent    = user32.NewProc("GetParent")
	procGetWindowLong = user32.NewProc("GetWindowLongW")
	procSetWindowLong = user32.NewProc("SetWindowLongW")
	procSetWindowPos  = user32.NewProc("SetWindowPos")
)

const (
	GWL_STYLE   = -16
	WS_CAPTION  = 0x00C00000
	WS_THICKFRAME = 0x00040000
	WS_MINIMIZEBOX = 0x00020000
	WS_MAXIMIZEBOX = 0x00010000
	SWP_FRAMECHANGED = 0x0020
	SWP_NOMOVE       = 0x0002
	SWP_NOSIZE       = 0x0001
	SWP_NOZORDER     = 0x0004
)

func removeWindowFrame(hwnd uintptr) {
	parent, _, _ := procGetParent.Call(hwnd)
	for parent != 0 {
		hwnd = parent
		parent, _, _ = procGetParent.Call(hwnd)
	}

	style, _, _ := procGetWindowLong.Call(hwnd, uintptr(GWL_STYLE))
	if style != 0 {
		newStyle := style &^ WS_CAPTION &^ WS_THICKFRAME &^ WS_MINIMIZEBOX &^ WS_MAXIMIZEBOX
		procSetWindowLong.Call(hwnd, uintptr(GWL_STYLE), newStyle)
		procSetWindowPos.Call(hwnd, 0, 0, 0, 0, 0, SWP_FRAMECHANGED|SWP_NOMOVE|SWP_NOSIZE|SWP_NOZORDER)
	}
}

func main() {
	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("Local Workspace Engine")
	w.SetSize(1200, 800, webview.HintNone)

	w.Bind("nativeOpenDialog", func() string {
		return ""
	})

	w.Bind("nativeSaveDialog", func(currentPath, currentName, content string) string {
		if currentPath == "" {
			currentPath = filepath.Join(".", currentName)
		}
		err := os.WriteFile(currentPath, []byte(content), 0644)
		if err != nil {
			return ""
		}
		return currentPath
	})

	w.Bind("closeApp", func() {
		w.Terminate()
	})

	w.Bind("minimizeApp", func() {
		w.SetSize(0, 0, webview.HintMin)
	})

	w.Bind("maximizeApp", func() {
		w.SetSize(1200, 800, webview.HintMax)
	})

	go func() {
		hwnd := w.Window()
		if hwnd != nil {
			removeWindowFrame(uintptr(hwnd))
		}
	}()

	w.SetHtml(htmlContent)
	w.Run()
}
