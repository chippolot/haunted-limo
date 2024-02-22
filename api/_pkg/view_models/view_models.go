package viewmodels

import (
	"github.com/chippolot/haunted-limo/api/_pkg/data"
)

type IndexModel struct {
	Stories []*data.StoryData
}

type StoryModel struct {
	Cfg   *data.StoryData
	Story string
}
