package video_tools

import (
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
)

var DownloadDir = ""

func match() zero.Rule {
	return func(ctx *zero.Ctx) bool {
		if len(ctx.Event.Message) == 0 {
			return false
		}
		for _, message := range ctx.Event.Message {
			if message.Type != "json" {
				continue
			}
			data := gjson.Parse(message.Data["data"])
			prompt := data.Get("prompt").String()
			if prompt != "[QQ小程序]哔哩哔哩" {
				continue
			}
			bvShortUrl := data.Get("meta.detail_1.qqdocurl").String()
			if bvShortUrl == "" {
				continue
			}
			ctx.State["bv_short_url"] = bvShortUrl
			return true
		}
		return false
	}
}

/*func init() {
	zero.OnKeyword("哔哩哔哩", match()).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			//ctx.SendChain(message.Text("少女祈祷中......"))
			urls, ok := ctx.State["bv_short_url"]
			if !ok && len(urls.(string)) == 0 {
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
}*/
