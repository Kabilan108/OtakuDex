package converter

import (
	"fmt"
	"github.com/kabilan108/OtakuDex/constant"
	"github.com/kabilan108/OtakuDex/converter/cbz"
	"github.com/kabilan108/OtakuDex/converter/pdf"
	"github.com/kabilan108/OtakuDex/converter/plain"
	"github.com/kabilan108/OtakuDex/converter/zip"
	"github.com/kabilan108/OtakuDex/source"
	"github.com/samber/lo"
	"strings"
)

// Converter is the interface that all converters must implement.
type Converter interface {
	Save(chapter *source.Chapter) (string, error)
	SaveTemp(chapter *source.Chapter) (string, error)
}

var converters = map[string]Converter{
	constant.FormatPlain: plain.New(),
	constant.FormatCBZ:   cbz.New(),
	constant.FormatPDF:   pdf.New(),
	constant.FormatZIP:   zip.New(),
}

// Available returns a list of available converters.
func Available() []string {
	return lo.Keys(converters)
}

// Get returns a converter by name.
// If the converter is not available, an error is returned.
func Get(name string) (Converter, error) {
	if converter, ok := converters[name]; ok {
		return converter, nil
	}

	return nil, fmt.Errorf("unkown format \"%s\", available options are %s", name, strings.Join(Available(), ", "))
}
