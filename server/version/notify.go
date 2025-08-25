package version

import (
	"fmt"
	"github.com/kabilan108/OtakuDex/color"
	"github.com/kabilan108/OtakuDex/constant"
	"github.com/kabilan108/OtakuDex/icon"
	"github.com/kabilan108/OtakuDex/key"
	"github.com/kabilan108/OtakuDex/style"
	"github.com/kabilan108/OtakuDex/util"
	"github.com/spf13/viper"
)

func Notify() {
	if !viper.GetBool(key.CliVersionCheck) {
		return
	}

	erase := util.PrintErasable(fmt.Sprintf("%s Checking if new version is available...", icon.Get(icon.Progress)))
	version, err := Latest()
	erase()
	if err == nil {
		comp, err := Compare(version, constant.Version)
		if err == nil && comp <= 0 {
			return
		}
	}

	fmt.Printf(`
%s New version is available %s %s
%s

`,
		style.Fg(color.Green)("▇▇▇"),
		style.Bold(version),
		style.Faint(fmt.Sprintf("(You're on %s)", constant.Version)),
		style.Faint("https://github.com/kabilan108/OtakuDex/releases/tag/v"+version),
	)

}
