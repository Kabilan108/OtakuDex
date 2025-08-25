package integration

import (
	"github.com/kabilan108/OtakuDex/integration/anilist"
	"github.com/kabilan108/OtakuDex/source"
)

// Integrator is the interface that wraps the basic integration methods.
type Integrator interface {
	// MarkRead marks a chapter as read
	MarkRead(chapter *source.Chapter) error
}

var (
	Anilist Integrator = anilist.New()
)
