package editorconfig

import "io/ioutil"

func CreateEditorconfig() {
	content := []byte(`
root = true
[*]
indent_style = space
indent_size = 2
charset = utf-8
trim_trailing_whitespace = true
insert_final_newline = true
end_of_line = lf
max_line_length = off
[Makefile]
indent_style = tab
[*.go]
indent_style = tab
`)
	err := ioutil.WriteFile(".editorconfig", content, 0644)
	if err != nil {
		panic(err)
	}
}
