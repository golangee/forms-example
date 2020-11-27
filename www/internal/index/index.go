package index

import (
	"encoding/base64"
	"encoding/json"
)

type Tutorial struct {
	Path string
	Doc  string
	Code string
}

type TutorialIndex struct {
	Tutorials []Tutorial
}

func (t TutorialIndex) Find(path string) Tutorial {
	for _, tutorial := range t.Tutorials {
		if tutorial.Path == path {
			return tutorial
		}
	}

	return Tutorial{}
}

var Tutorials TutorialIndex

func init() {
	buf, err := base64.StdEncoding.DecodeString(tutorials)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(buf, &Tutorials)
	if err != nil {
		panic(err)
	}
}
