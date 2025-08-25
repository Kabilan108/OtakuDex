package provider

import (
	"github.com/kabilan108/OtakuDex/provider/generic"
	"github.com/kabilan108/OtakuDex/provider/mangadex"
	"github.com/kabilan108/OtakuDex/provider/manganato"
	"github.com/kabilan108/OtakuDex/provider/manganelo"
	"github.com/kabilan108/OtakuDex/provider/mangapill"
	"github.com/kabilan108/OtakuDex/source"
)

const CustomProviderExtension = ".lua"

var builtinProviders = []*Provider{
	{
		ID:   mangadex.ID,
		Name: mangadex.Name,
		CreateSource: func() (source.Source, error) {
			return mangadex.New(), nil
		},
	},
}

func init() {
	for _, conf := range []*generic.Configuration{
		manganelo.Config,
		manganato.Config,
		mangapill.Config,
	} {
		conf := conf
		builtinProviders = append(builtinProviders, &Provider{
			ID:   conf.ID(),
			Name: conf.Name,
			CreateSource: func() (source.Source, error) {
				return generic.New(conf), nil
			},
		})
	}
}
