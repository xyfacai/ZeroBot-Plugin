package meizitu

import (
	"fmt"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/message"
	"math/rand"
	"strings"
	"time"
)

// https://www.meinv.im/search/?page=2&s=%E5%A4%A7%E5%B0%BA%E5%BA%A6
var searchApi = "https://www.meinv.im/search/"
var ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
var poke = rate.NewManager[int64](time.Minute*5, 5) // 戳

func init() { // 插件主体
	engine := control.Register("meizitu", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "妹子图",
		Help:             "- 来点黑丝妹子\n- 来点白丝妹子\n- 来点jk妹子\n- 来点巨乳妹子\n- 来点足控妹子\n- 来点网红妹子",
	})

	engine.OnRegex(`^来点\s*(.+)妹子$`, zero.OnlyGroup).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch {
			case poke.Load(ctx.Event.GroupID).Acquire():
				// 5分钟共8块命令牌 一次消耗1块命令牌
				ctx.SendChain(message.Text("少女祈祷中......"))
				key := ctx.State["regex_matched"].([]string)[1]
				key = strings.TrimSpace(key)
				if len(key) == 0 {
					ctx.SendChain(message.Text("类型不能为空： 命令提示： 来点黑丝妹子\n 来点jk妹子"))
					return
				}
				img, err := getOneMeizi(key)
				if err != nil {
					ctx.SendChain(message.Text("ERROR: ", err))
					return
				}
				if img == nil || len(img) == 0 {
					ctx.SendChain(message.Text("未找到: ", key))
					return
				}

				m := message.Message{ctxext.FakeSenderForwardNode(ctx, message.ImageBytes(img))}
				if id := ctx.Send(m).ID(); id == 0 {
					ctx.SendChain(message.Text("ERROR: 可能被风控或下载图片用时过长，请耐心等待"))
				}
			default:
				// 频繁触发，不回复
				time.Sleep(time.Second * 1)
				ctx.SendChain(message.Text("要冒烟了，过会再来吧"))
			}
		})
}

func getOneMeizi(key string) (imgData []byte, err error) {
	var articles []string
	queryUrl := fmt.Sprintf("%s?s=%s", searchApi, key)
	client := req.C().SetCommonHeaders(map[string]string{
		"user-agent": ua,
		"referer":    "https://www.meinv.im/type/1/",
	})

	client.SetProxyURL("http://127.0.0.1:7890")
	for i := 1; i <= 3; i++ {
		queryUrl = fmt.Sprintf("%s&page=%d", queryUrl, i)
		resp, err := client.R().
			SetRetryCount(2).
			AddRetryHook(func(resp *req.Response, err error) {
				//log.Logger().Warn("重试发送封面确认请求")
				time.Sleep(time.Second)
				return
			}).
			AddRetryCondition(func(resp *req.Response, err error) bool {
				return err != nil || resp.StatusCode != 200
			}).Get(queryUrl)
		if err != nil {
			return nil, err
		}

		dom, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return nil, err
		}
		dom.Find("article > div > h2 > a").Each(func(i int, articleA *goquery.Selection) {
			//articleA := selection.Find(")
			//<h2><a href="/article/68302/">清纯白色丝袜动漫美女许岚露内内图大尺度超污</a></h2>
			if u, ok := articleA.Attr("href"); ok {
				articles = append(articles, u)
			}
		})
	}

	if len(articles) == 0 {
		return
	}

	for i := 1; i <= 3; i++ {
		article := articles[rand.Intn(len(articles))]
		var imgs []string

		resp, err := client.R().
			SetHeader("refer", "https://www.meinv.im").
			SetRetryCount(2).
			AddRetryHook(func(resp *req.Response, err error) {
				//log.Logger().Warn("重试发送封面确认请求")
				time.Sleep(time.Second)
				return
			}).
			AddRetryCondition(func(resp *req.Response, err error) bool {
				return err != nil || resp.StatusCode != 200
			}).Get(fmt.Sprintf("https://www.meinv.im%s", article))
		if err != nil {
			return nil, err
		}

		dom, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return nil, err
		}
		dom.Find("article > p").Each(func(i int, selection *goquery.Selection) {
			img := selection.Find("img")
			//<p><img class="alignnone size-medium wp-image-42" src="/static/zde/timg.gif" data-src="/static/images/20230312/6004/1678467221vsiS.jpg" alt="清纯白色丝袜动漫美女许岚露内内图大尺度超污" style="width:800px;hight:auto" />
			if u, ok := img.Attr("data-src"); ok {
				imgs = append(imgs, u)
			}
		})
		if len(imgs) == 0 {
			continue
		}
		img := imgs[rand.Intn(len(imgs))]

		reps, err := client.R().Get("https://www.meinv.im" + img)
		if err != nil {
			return nil, err
		}
		return reps.Bytes(), nil
	}
	return nil, nil
}
