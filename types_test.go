package ghost_test

import (
	"encoding/json"
	"fmt"
	"testing"

	ghost "github.com/acim/go-ghost"
)

func TestMobiledoc(t *testing.T) {
	html := "<iframe type=text/html width=640 height=390 src=http://www.youtube.com/embed/Mq99gFKztaI frameborder=0></iframe>"

	m := ghost.Mobiledoc{
		Version: "0.3.1",
		Cards: [][]interface{}{
			[]interface{}{
				"html",
				ghost.Card{
					HTML: &html,
				},
			},
		},
		Sections: [][]interface{}{
			[]interface{}{10, 0},
			[]interface{}{1, "p", []interface{}{}},
		},
	}
	j, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	}

	js, err := json.Marshal(string(j))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(js))
}
