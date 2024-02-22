package data

import (
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/chippolot/haunted-limo/api/_pkg/common"
)

const (
	StoryDatPath = "stories.json"
)

type StoriesData struct {
	Stories []*StoryData
}

type StoryData struct {
	Key                string
	Title              string
	StoryType          string
	Story              string
	BackgroundColor    string
	LogoFontLink       template.URL
	LogoFontFamilyName string
	LogoFontStyle      string
	LogoFontWeight     int
	LogoFontSerif      string
}

func LoadStoryData() ([]*StoryData, error) {
	filePath := common.GetDataFilePath(StoryDatPath)
	bytes, err := common.LoadFileBytes(filePath)
	if err != nil {
		return nil, err
	}

	var stories StoriesData
	err = json.Unmarshal(bytes, &stories)
	return stories.Stories, err
}

func FindStoryData(dataList []*StoryData, storyType string) (*StoryData, error) {
	for _, data := range dataList {
		if data.StoryType == storyType {
			return data, nil
		}
	}
	return nil, fmt.Errorf("could not find story data with type %s", storyType)
}
