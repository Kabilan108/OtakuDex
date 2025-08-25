package main

import (
	"github.com/kabilan108/OtakuDex/cmd"
	"github.com/kabilan108/OtakuDex/config"
	"github.com/kabilan108/OtakuDex/log"
	"github.com/samber/lo"
)

func main() {
	lo.Must0(config.Setup())
	lo.Must0(log.Setup())
	cmd.Execute()
}
