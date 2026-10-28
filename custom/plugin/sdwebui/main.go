package sdwebui

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"

	"github.com/FloatTech/floatbox/binary"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type arg struct {
	Prompt         string `flag:"s"`
	NegativePrompt string `flag:"n"`
	Seed           int    `flag:"seed"`
	Lenth          bool   `flag:"l"`
	Width          bool   `flag:"w"`
}

func init() {
	en := control.AutoRegister(&ctrl.Options[*zero.Ctx]{
		DisableOnDefault: true,
		Brief:            "",
		Help: "/ai -s [prompt] -n [negativeprompt] -seed [num] -[l|w]\n" +
			"/getmodels\n" +
			"/usemodel [modelname]\n" +
			"/setapi [url]\n" +
			"Tips: -l is 512*768, -w is 768*512",
		PrivateDataFolder: "sdwebui",
	}).ApplySingle(ctxext.DefaultSingle)
	en.OnShell("ai", arg{}).SetBlock(true).Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			arg := ctx.State["flag"].(*arg)
			np := arg.NegativePrompt
			c := txt2img{}

			if arg.NegativePrompt == "" {
				np = "lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry"
			}
			c.NegativePrompt = np
			c.Prompt = arg.Prompt

			c.Height = 512
			c.Width = 512

			if arg.Width {
				c.Width = 768
			}
			if arg.Lenth {
				c.Height = 768
				c.Width = 512
			}

			c.BatchSize = 1
			c.NIter = 1
			c.Tiling = true
			c.CfgScale = 7.0
			c.Steps = 20
			c.Seed = rand.Intn(2147483646) + 1
			c.SamplerIndex = "Euler a"
			c.SendImages = true
			c.ClipSkip = 2

			c.OverrideSettings.CLIPStopAtLastLayers = 2

			ctx.SendChain(message.Text("少女折寿中..."))

			imginfo, err := posttxt2img(c)
			if err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
				return
			}

			info := info{}
			err = json.Unmarshal(binary.StringToBytes(imginfo.Info), &info)
			if err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
				return
			}
			ctx.SendChain(message.Image("base64://"+imginfo.Images[0]), message.Text(info))
		})
	en.OnFullMatch("/getmodels").SetBlock(true).Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			models, err := getmodels()
			if err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
				return
			}
			sb := strings.Builder{}
			sb.WriteString("1.")
			sb.WriteString(models[0].ModelName)
			for i, v := range models[1:] {
				sb.WriteString("\n")
				sb.WriteString(strconv.Itoa(i + 2))
				sb.WriteString(".")
				sb.WriteString(v.ModelName)
			}
			ctx.SendChain(message.Text(sb.String()))
		})
	en.OnRegex(`/usemodel (.*)`, zero.AdminPermission).SetBlock(true).Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			modelname := ctx.State["regex_matched"].([]string)[1]
			err := changemodel(modelname)
			if err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
				return
			}
			ctx.SendChain(message.Text("设置成功"))
		})
	en.OnRegex(`/setapi (.*)`, zero.SuperUserPermission).SetBlock(true).Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			api := ctx.State["regex_matched"].([]string)[1]
			if !strings.Contains(api, "http") {
				ctx.SendChain(message.Text("api地址不正确"))
				return
			}
			baseurl = api
			ctx.SendChain(message.Text("设置成功"))
		})
}
