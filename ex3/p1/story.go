package p1

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strings"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

var htmlTemplateDef = `<!DOCTYPE html>
<html>
	<head>
		<title>Choose your own adventure</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
{{range .Story}}
		<p>{{.}}</p>
{{end}}
		<ul>
{{range .Options}}
			<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
{{end}}
		</ul>
	</body>
</html>
`

var htmlTemplate = template.Must(template.New("htmlStory").Parse(htmlTemplateDef))

func ParseStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)

	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

func getChapterName(requestURI string) string {
	n := strings.TrimLeft(requestURI, "/")
	return n
}

func StoryHandler(s *Story) http.Handler {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chapterName := getChapterName(r.RequestURI)
		if chapterName == "" {
			http.Redirect(w, r, "/intro", 307)
			return
		}

		chapter, ok := (*s)[chapterName]
		if !ok {
			w.WriteHeader(404)
			return
		}

		htmlTemplate.Execute(w, chapter)
	})
	return h
}
