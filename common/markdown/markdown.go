package markdown

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"os"
)

func Markdown() template.HTML {
	file, err := file2Bytes("storage/uploads/markdown/go-blog.md")
	html := bluemonday.UGCPolicy().SanitizeBytes(file)
	if err != nil {
		fmt.Println(err)
	}

	s := blackfriday.Run(html)
	return template.HTML(s)

}

func file2Bytes(filename string) ([]byte, error) {

	// File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// []byte
	data := make([]byte, stats.Size())
	count, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("read file %s len: %d \n", filename, count)
	return data, nil
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 渲染markdown为html
func Render(data string) string {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	md := []byte(data)
	content := markdown.ToHTML(md, nil, renderer)
	return string(content)
}
