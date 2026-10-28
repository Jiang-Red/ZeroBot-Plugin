// Package dice dice！golang部分移植版
package dice

import (
	"time"

	fcext "github.com/FloatTech/floatbox/ctxext"
	sql "github.com/FloatTech/sqlite"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	db     sql.Sqlite
	initDB zero.Rule
	engine = control.AutoRegister(&ctrl.Options[*zero.Ctx]{
		DisableOnDefault: true,
		Brief:            "骰子?",
		Help:             "试图移植的dice\n-.jrrp\n-.ra\n-.rd",
		PublicDataFolder: "Dice",
	})
)

func init() {
	initDB = fcext.DoOnceOnSuccess(func(ctx *zero.Ctx) bool {
		db = sql.New(engine.DataFolder() + "dice.db")
		if err := db.Open(time.Hour * 24); err != nil {
			ctx.SendChain(message.Text("ERROR: ", err))
			return false
		}
		for name, model := range map[string]any{
			"strjrrp": &strjrrp{},
			"rsl":     &rsl{},
			"set":     &set{},
		} {
			if err := db.Create(name, model); err != nil {
				ctx.SendChain(message.Text("ERROR: ", err))
				return false
			}
		}
		return true
	})
}
