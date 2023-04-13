// Package coser images
package query_material

import (
	"encoding/json"
	"fmt"
	bz "github.com/FloatTech/AnimeAPI/bilibili"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var imgDir = "D:\\video\\固定素材\\背景图\\竖屏"
var imgDir2 = "D:\\video\\固定素材\\杂\\群相册"

var videoWorkDirs = []string{"D:\\video\\搞笑图", "D:\\video\\搞笑图_4016", "D:\\video\\搞笑图_2796"}
var videoUpMids = map[string]struct{}{
	"481895777":  struct{}{},
	"1207346390": struct{}{},
	"432832299":  struct{}{},
}

var filesMap = map[string][]string{}

var (
	limit            = ctxext.NewLimiterManager(time.Second*10, 1)
	searchVideo      = `bilibili.com\\?/video\\?/(?:av(\d+)|([bB][vV][0-9a-zA-Z]+))`
	searchDynamic    = `(t.bilibili.com|m.bilibili.com\\?/dynamic)\\?/(\d+)`
	searchArticle    = `bilibili.com\\?/read\\?/(?:cv|mobile\\?/)(\d+)`
	searchLiveRoom   = `live.bilibili.com\\?/(\d+)`
	searchVideoRe    = regexp.MustCompile(searchVideo)
	searchDynamicRe  = regexp.MustCompile(searchDynamic)
	searchArticleRe  = regexp.MustCompile(searchArticle)
	searchLiveRoomRe = regexp.MustCompile(searchLiveRoom)
)

func init() {

	engine := control.Register("获取背景图", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "获取背景图",
		Help:             "获取背景图 [第几期]",
	})
	engine.ApplySingle(ctxext.DefaultSingle).OnPrefix("获取背景图").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("少女祈祷中......"))
			searchKey, ok := ctx.State["args"].(string)
			if !ok {
				ctx.SendChain(message.Text("参数错误：格式为-> '获取背景图 第几期', 注意中间有空格"))
				return
			}

			searchKey = strings.ReplaceAll(searchKey, " ", "")
			var sourceMaterials []string
			if sourceMaterials, ok = filesMap[searchKey]; !ok {
				sourceMaterials = findVideoMaterial(ctx, searchKey)
			}
			if len(sourceMaterials) == 0 {
				return
			}
			ctx.SendChain(message.Text(fmt.Sprintf("已找到：%s", searchKey)))

			var images []message.MessageSegment
			for _, ma := range sourceMaterials {
				images = append(images, message.Image("file:///"+ma))
			}
			if id := ctx.Send(message.Message{ctxext.FakeSenderForwardNode(ctx, images...)}).ID(); id == 0 {
				ctx.SendChain(message.Text("ERROR: 可能被风控或下载图片用时过长，请耐心等待"))
			}
		})

	engine.OnKeyword("哔哩哔哩", match()).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			url := ctx.State["regex_matched"].([]string)[0]
			realurl, err := bz.GetRealURL("https://" + url)
			if err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
				return
			}
			//searchVideoRe.MatchString(realurl)
			ctx.State["regex_matched"] = searchVideoRe.FindStringSubmatch(realurl)
			handleVideo(ctx)
		})

}

func findVideoMaterial(ctx *zero.Ctx, searchKey string) (sourceMaterials []string) {
	//找配置文件
	configFile := ""
	for _, dir := range videoWorkDirs {
		f := ""
		_ = filepath.Walk(filepath.Join(dir, "已使用素材", "目标视频目录"), func(path string, info fs.FileInfo, err error) error {
			if strings.ReplaceAll(GetFilenameNotSuffix(info.Name()), " ", "") == searchKey {
				f = path
				return nil
			}
			return nil
		})

		exist, _ := PathExists(f)
		if exist {
			configFile = f
			break
		}
	}
	if configFile == "" {
		ctx.SendChain(message.Text("未找到"))
		return
	}
	data, err := os.ReadFile(configFile)
	if err != nil {
		ctx.SendChain(message.Text("未找到"))
		return
	}
	var record WorkflowRecord
	err = json.Unmarshal(data, &record)
	if err != nil {
		ctx.SendChain(message.Text("未找到"))
		return
	}

	materials := make(map[string]struct{})
	for _, w := range record.Workflow {
		if w.Work == "图片设置背景图" {
			p := w.Params.(map[string]any)
			materials[filepath.Base(AnyToString(p["dir"]))] = struct{}{}
		}
		if w.Work == "插入" {
			p := w.Params.(map[string]any)
			materials[filepath.Base(AnyToString(p["dir"]))] = struct{}{}
		}
	}
	if len(materials) == 0 {
		ctx.SendChain(message.Text("未找到"))
		return
	}

	bgs, _ := ReadDir(imgDir)
	imgs, _ := ReadDir(imgDir2)
	bgs = append(bgs, imgs...)
	// 获取源文件
	for _, path := range bgs {
		fileName := filepath.Base(path)
		if _, ok := materials[fileName]; ok {
			sourceMaterials = append(sourceMaterials, path)
			continue
		}
	}

	if len(sourceMaterials) == 0 {
		ctx.SendChain(message.Text(fmt.Sprintf("未找到：%s", searchKey)))
		return
	}
	filesMap[searchKey] = sourceMaterials
	return
}

func handleVideo(ctx *zero.Ctx) {
	id := ctx.State["regex_matched"].([]string)[1]
	if id == "" {
		id = ctx.State["regex_matched"].([]string)[2]
	}
	card, err := bz.GetVideoInfo(id)
	if err != nil {
		ctx.SendChain(message.Text("ERROR: ", err))
		return
	}
	if _, ok := videoUpMids[fmt.Sprintf("%d", card.Owner.Mid)]; !ok {
		ctx.SendChain(message.Text("不是 Up 主的视频， 不能获取素材，告辞..."))
		return
	}
	var ok bool
	searchKey := card.Title
	searchKey = strings.ReplaceAll(searchKey, " ", "")
	var sourceMaterials []string
	if sourceMaterials, ok = filesMap[searchKey]; !ok {
		sourceMaterials = findVideoMaterial(ctx, searchKey)
	}
	if len(sourceMaterials) == 0 {
		return
	}
	ctx.SendChain(message.Text(fmt.Sprintf("已找到：%s", searchKey)))

	var images []message.MessageSegment
	for _, ma := range sourceMaterials {
		images = append(images, message.Image("file:///"+ma))
	}
	if id := ctx.Send(message.Message{ctxext.FakeSenderForwardNode(ctx, images...)}).ID(); id == 0 {
		ctx.SendChain(message.Text("ERROR: 可能被风控或下载图片用时过长，请耐心等待"))
	}
}

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

			regex := regexp.MustCompile(`((b23|acg).tv|bili2233.cn)/[0-9a-zA-Z]+`)
			if matched := regex.FindStringSubmatch(bvShortUrl); matched != nil {
				ctx.State["bv_short_url"] = bvShortUrl
				ctx.State["regex_matched"] = matched
				return true
			}

			return false
		}
		return false
	}
}

type WorkflowRecord struct {
	InVideos          []string `json:"in_videos"`
	Out               string   `json:"out"`
	Title             string   `json:"title"`
	CoverSource       []string `json:"cover_source"`
	Category          string
	Generator         VideoGenerator     `yaml:"generator"`
	Paths             VideoPath          `yaml:"paths"`
	Workflow          []Workflow         `json:"workflow"`
	MaterialDurations []MaterialDuration `json:"material_durations"`
}

type VideoGenerator struct {
	Execute  bool       `yaml:"execute"`
	Workflow []Workflow `yaml:"workflow"`
}

type Workflow struct {
	Work   string      `yaml:"work"`
	Params interface{} `yaml:"params"`
}

type VideoPath struct {
	WorkDir string `yaml:"work_dir"`
}

type MaterialDuration struct {
	Dur      time.Duration `json:"-"`
	Duration string        `json:"duration"`
	Material string        `json:"material"`
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func AnyToString(data any) string {
	if data == nil {
		return ""
	}
	switch data.(type) {
	case string:
		return data.(string)
	case int:
		return strconv.Itoa(data.(int))
	}
	return data.(string)
}

func GetFilenameNotSuffix(filePath string) string {
	baseName := filepath.Base(filePath)
	ext := path.Ext(filePath)                // 输出 .html
	return strings.TrimSuffix(baseName, ext) // 输出 name
}

func ReadDir(root string) ([]string, error) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	var allFiles []string
	for _, file := range files {
		if file.IsDir() {
			subFiles, err := ReadDir(filepath.Join(root, file.Name()))
			if err != nil {
				return nil, err
			}
			if len(subFiles) > 0 {
				allFiles = append(allFiles, subFiles...)
			}
		} else {
			allFiles = append(allFiles, filepath.Join(root, file.Name()))
		}
	}
	return allFiles, nil
}
