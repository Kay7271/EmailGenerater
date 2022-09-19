package util

import (
	"bytes"
	"converter/conf"
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var commands = map[string]string{
	"windows": "explorer",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func OpenBrowser(platform, port string) error {
	run, ok := commands[platform]
	uri := "http://127.0.0.1:" + port
	if !ok {
		return fmt.Errorf("unknown platform : %s", platform)
	}
	cmd := exec.Command(run, uri)
	return cmd.Start()
}

func GetPinYin(name string) []string {
	a := pinyin.NewArgs()
	var res []string
	for i, ch := range name {
		if py, ok := conf.DuoYinZiMap[string(ch)]; i == 0 && ok {
			res = append(res, py)
		} else {
			res = append(res, pinyin.Pinyin(string(ch), a)[0][0])
		}
	}
	return res
}

func DoConvert(method, suffix, names string) (string, error) {
	convertor := conf.MethodMap[method]
	nameSlic := SplitText(names)
	bfs := bytes.Buffer{}
	for _, name := range nameSlic {
		if name == "" {
			continue
		}
		pinYin := GetPinYin(name)
		bfs.WriteString(convertor(pinYin) + "@" + suffix + "\n")
	}
	return bfs.String(), nil
}

func readXlsx(bts []byte) string {
	buffer := bytes.Buffer{}
	xlsx, err := xlsx.OpenBinary(bts)
	if err != nil {
		return ""
	}
	sheet1 := xlsx.Sheets[0]
	column := findColumn(sheet1)
	if column == -1 {
		return ""
	}
	for _, row := range sheet1.Rows[1:] {
		buffer.WriteString(row.Cells[column].String() + "\n")
	}
	return buffer.String()
}

func findColumn(sheet *xlsx.Sheet) int {
	firstRow := sheet.Rows[0]
	for i, cell := range firstRow.Cells {
		if strings.Contains(strings.ToUpper(cell.String()), "NAME") || strings.Contains(cell.String(), "姓名") {
			return i
		}
	}
	return -1
}

func readTxt(bts []byte) string {
	return string(bts)
}

func GetNamesFromFile(fileType string, bts []byte) string {
	if fileType == "txt" {
		return readTxt(bts)
	} else if fileType == "xlsx" {
		return readXlsx(bts)
	}
	return ""
}

func GenerateTxtFile(text string) []byte {
	return []byte(text)
}

func GenerateExcelFile(text string) []byte {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("email_sheet")
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "邮箱地址"

	emails := SplitText(text)
	for _, email := range emails {
		if email == "" {
			continue
		}
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = email
	}
	err := file.Save("./config/tmp.xlsx")
	if err != nil {
		return nil
	}
	bts, _ := ioutil.ReadFile("./config/tmp.xlsx")
	os.Remove("./config/tmp.xlsx")
	return bts
}

func SplitText(text string) []string {
	var textSlic []string
	textSlic = strings.Split(text, "\n")
	for i, text := range textSlic {
		textSlic[i] = strings.TrimSpace(text)
	}
	return textSlic
}
