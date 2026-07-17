package main

import (
	"embed"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

//go:embed all:site
var siteFS embed.FS

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	}
	if cmd != nil {
		go func() {
			time.Sleep(800 * time.Millisecond)
			_ = cmd.Run()
		}()
	}
}

func main() {
	port := "7890"
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(siteFS)))

	srv := &http.Server{Addr: ":" + port, Handler: mux}
	url := fmt.Sprintf("http://127.0.0.1:%s/", port)
	fmt.Printf("🎮 Six to Survive 启动中...\n")
	fmt.Printf("   浏览器将自动打开，如未打开请访问: %s\n", url)
	fmt.Printf("   关闭此窗口即退出游戏\n")
	openBrowser(url)
	_ = srv.ListenAndServe()
}
