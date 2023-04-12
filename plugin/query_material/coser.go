// Package coser images
package query_material

import (
	"encoding/json"
	"fmt"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var imgDir = "D:\\video\\固定素材\\背景图\\竖屏\\美女"

var videoWorkDirs = []string{"D:\\video\\搞笑图", "D:\\video\\搞笑图_4016", "D:\\video\\搞笑图_2796"}

var filesMap = map[string]string{}

func init() {
	control.Register("获取背景图", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "获取背景图",
		Help:             "获取背景图 [第几期]",
	}).ApplySingle(ctxext.DefaultSingle).OnPrefix("获取背景图").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("少女祈祷中......"))
			searchKey, ok := ctx.State["args"].(string)
			if !ok {
				ctx.SendChain(message.Text("参数错误：格式为-> '获取背景图 第几期', 注意中间有空格"))
				return
			}
			searchKey = strings.ReplaceAll(searchKey, " ", "")
			if bgPath, ok := filesMap[searchKey]; ok {
				ctx.SendChain(message.Text(fmt.Sprintf("已找到：%s", searchKey)))
				if id := ctx.Send(message.Message{ctxext.FakeSenderForwardNode(ctx,
					message.Image("file:///"+bgPath))}).ID(); id == 0 {
					ctx.SendChain(message.Text("ERROR: 可能被风控或下载图片用时过长，请耐心等待"))
				}
				return
			}
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
			bgPath := ""
			bgPath2 := ""
			for _, w := range record.Workflow {
				if w.Work == "图片设置背景图" {
					p := w.Params.(map[string]any)
					bgPath = AnyToString(p["dir"])
					break
				}
			}
			if bgPath == "" {
				ctx.SendChain(message.Text("未找到"))
				return
			}
			fileName := filepath.Base(bgPath)
			// 获取源文件
			for _, md := range record.MaterialDurations {
				if filepath.Base(md.Material) == fileName {
					bgPath2 = md.Material
					break
				}
			}

			exist, _ := PathExists(bgPath2)
			if exist {
				bgPath = bgPath2
			}
			exist, _ = PathExists(bgPath)
			if !exist {
				ctx.SendChain(message.Text(fmt.Sprintf("未找到：%s", searchKey)))
				return
			}
			filesMap[searchKey] = bgPath
			ctx.SendChain(message.Text(fmt.Sprintf("已找到：%s", searchKey)))
			if id := ctx.Send(message.Message{ctxext.FakeSenderForwardNode(ctx,
				message.Image("file:///"+bgPath))}).ID(); id == 0 {
				ctx.SendChain(message.Text("ERROR: 可能被风控或下载图片用时过长，请耐心等待"))
			}
		})
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
