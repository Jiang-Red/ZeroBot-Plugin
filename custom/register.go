// Package custom registers the local custom plugins and their priorities.
package custom

import (
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/antirecall"    // 反撤回
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/antirepeat"    // 限制复读
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/cloudmusic"    // 网易云热评
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/custom"        // 自订工具
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/delreply"      // 回复撤回
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/dice"          // 骰子
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/guessname"     // 音游开字符
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/managerplugin" // 自定义群管插件
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/moehu"         // 图库
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/phigros"       // phigros b19
	_ "github.com/FloatTech/ZeroBot-Plugin/custom/plugin/sdwebui"       // sdwebui
)
