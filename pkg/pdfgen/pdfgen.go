package pdfgen

// TODO
import (
	"github.com/jung-kurt/gofpdf"
	"log"
	"regexp"
	"strings"
)

// 正则，把<body>里面的内容取出来，没有的话直接返回原字符串
func extractBodyContent(html string) string {
	re := regexp.MustCompile(`<body>([\s\S]*?)</body>`)
	match := re.FindStringSubmatch(html)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}
	return html
}

// 生成PDF
// 注：该代码只实现生成pdf，你记得remove
func CreatePdf(htmlStr string, uid string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	//加入三个字体文件，是不是应该来个文件夹放字体这些文件什么的……
	pdf.AddUTF8Font("deng", "", "Deng.ttf")
	pdf.AddUTF8Font("deng", "B", "Dengb.ttf")
	pdf.AddUTF8Font("deng", "I", "Dengl.ttf")
	pdf.SetFont("deng", "", 20)

	pdf.SetLeftMargin(45)
	pdf.SetFontSize(8)
	htmlStr = extractBodyContent(htmlStr)
	html := pdf.HTMLBasicNew()
	html.Write(5.0, htmlStr)       //也可以手动指定 line height
	pdfName := "./" + uid + ".pdf" //生成pdf的路径，前面的可以加相对路径
	err := pdf.OutputFileAndClose(pdfName)
	if err != nil {
		//panic(err.Error())//一个低级的panic
		log.Println(err) //一个低级的打印
	}
	return err

}
