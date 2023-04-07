package video_tools

import (
	"encoding/binary"
	"fmt"
	"github.com/imroc/req/v3"
	zero "github.com/wdvxdr1123/ZeroBot"
	"io"
	"os"
	"path/filepath"
	"time"

	"strings"
)

var DownloadDir = ""

func init() {
	zero.OnMessage(zero.OnlyGroup).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			//ctx.SendChain(message.Text("少女祈祷中......"))
			urls, ok := ctx.State["image_url"]
			if !ok && len(urls.([]string)) == 0 {
				return
			}

			for _, url := range urls.([]string) {
				resp, err := req.R().Get(url)
				if err != nil {
					continue
				}
				contentType := resp.GetHeader("Content-Type")
				ext := ""
				if contentType == "image/jpeg" {
					ext = ".jpg"
				} else if contentType == "image/png" {
					ext = ".png"
				} else {
					ext = "." + strings.ReplaceAll(contentType, "image/", "")
				}
				size := binary.Size(resp.Bytes())
				if size < 1024*300 {
					//s.log.Debugf("图片下载大小不足 %d", size)
					continue
				}
				file, err := os.Create(filepath.Join(DownloadDir,
					fmt.Sprintf("%s_%d%s", time.Now().Format("01_02_15_04_05"), ctx.Event.Time, ext)))
				if err != nil {
					continue
				}

				_, err = io.Copy(file, resp.Body)
				file.Close()
				if err != nil {
					continue
				}
			}
		})
}
