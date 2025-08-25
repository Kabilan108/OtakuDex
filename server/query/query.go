package query

import (
	"github.com/metafates/gache"
	"github.com/kabilan108/OtakuDex/filesystem"
	"github.com/kabilan108/OtakuDex/where"
)

type queryRecord struct {
	Rank  int    `json:"rank"`
	Query string `json:"query"`
}

var cacher = gache.New[map[string]*queryRecord](
	&gache.Options{
		Path:       where.Queries(),
		FileSystem: &filesystem.GacheFs{},
	},
)
