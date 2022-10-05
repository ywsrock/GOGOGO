package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"unicode"

	"github.com/xuri/excelize/v2"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var WriteStartRow = 9

func main() {

	// parse filename(ファイル名のみ指定）
	frmFile := ""
	flag.StringVar(&frmFile, "f", "", "Formファイル指定")
	flag.Parse()

	frmFile = path.Base(frmFile)
	pwd, _ := os.Getwd()
	frmFile = fmt.Sprintf("%s/%s", pwd, frmFile)

	// Check the file status, if case of error, Processing exit
	if _, err := os.Stat(frmFile); err != nil {
		log.Fatal(err)
	}
	// ファイル名読込
	f, err := os.Open(frmFile)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// 　書き込みファイル
	exf := OpenTemplateExcel("template.xlsx")
	defer func() {
		if err := exf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Print string buffer
	sb := bytes.Buffer{}
	// Print Attributes map
	res := make(map[string]string)
	// Print stdout Header
	writeHeader(sb)

	// 日本語　ShiftJIS
	r := transform.NewReader(f, japanese.ShiftJIS.NewDecoder())
	scan := bufio.NewScanner(r)

	for scan.Scan() {
		// read line
		line := scan.Text()
		// trim space
		replaceLine(line, &sb, res, exf)
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	if err := exf.Save(); err != nil {
		fmt.Println("Excelファイル保存失敗！")
	}
}

func writeHeader(sb bytes.Buffer) {
	sb.WriteString("Name")
	sb.WriteString("\t")
	sb.WriteString("Type")
	sb.WriteString("Caption")
	sb.WriteString("\t")
	sb.WriteString("Text")
	sb.WriteString("\t")
	sb.WriteString("IMEMode")
	sb.WriteString("\t")
	sb.WriteString("Enabled")
	sb.WriteString("\t")
	sb.WriteString("Object.Width")
	sb.WriteString("\t")
	sb.WriteString("Locked")

	fmt.Print(sb.String())
	sb.Reset()
}

func replaceLine(str string, sb *bytes.Buffer, res map[string]string, f *excelize.File) {
	keys := [...]string{"Enabled ", "Caption ", "IMEMode ", "Text ", "Object.Width ", "Locked "}
	// if start with begin out a new line
	if b := pickUpWord(str, sb, res, f); b {
		return
	}
	for _, v := range keys {
		if strings.Contains(str, v) {
			kv := strings.Split(str, "=")
			v = strings.TrimSpace(v)
			if len(kv) > 1 {
				res[v] = kv[1]
			}
			return
		}
	}
}

func pickUpWord(w string, sb *bytes.Buffer, res map[string]string, excel *excelize.File) bool {
	// start with "Begin" $1= type $2 = Name
	r := regexp.MustCompile(`(?im)\s*Begin|BeginProperty\s+(.*)`)
	r1 := regexp.MustCompile(`\s`)

	if r.MatchString(w) {
		// print current line
		// rests the buffer
		PrintStdout(sb, res)
		// WriteExcel
		WriteExcel(res, excel)

		sb.Reset()
		for k := range res {
			delete(res, k)
		}

		rs := strings.TrimSpace(r.ReplaceAllString(w, "$1"))
		// split by space
		rms := r1.Split(rs, -1)
		// remove space
		n := r1.ReplaceAllString(rms[1], "")
		// remove space
		t := r1.ReplaceAllString(rms[0], "")
		res["Name"] = n
		res["Type"] = t

		return true
	}
	return false
}

func WriteExcel(res map[string]string, excel *excelize.File) {
	// 項目定義
	sn := excel.GetSheetName(0)

	// Cellに値を設定（サンプルエラー処理無視）
	// 項目名	D9
	dn := fmt.Sprintf("%s%d", "D", WriteStartRow)
	if v, b := res["Caption"]; b {
		err := excel.SetCellValue(sn, dn, strings.TrimFunc(v, TrimF))
		if err != nil {
			return
		}
	} else {
		err := excel.SetCellValue(sn, dn, strings.TrimFunc(res["Text"], TrimF))
		if err != nil {
			return
		}
	}

	// 論理名	L9
	ln := fmt.Sprintf("%s%d", "L", WriteStartRow)
	err := excel.SetCellValue(sn, ln, res["Name"])
	if err != nil {
		return
	}

	// 種別		Q9
	qn := fmt.Sprintf("%s%d", "Q", WriteStartRow)
	err = excel.SetCellValue(sn, qn, res["Type"])
	if err != nil {
		return
	}

	// IME 		X9
	xn := fmt.Sprintf("%s%d", "X", WriteStartRow)
	err = excel.SetCellValue(sn, xn, res["IMEMode"])
	if err != nil {
		return
	}

	// Enable 	BC9
	bcn := fmt.Sprintf("%s%d", "BC", WriteStartRow)
	err = excel.SetCellValue(sn, bcn, res["Enabled"])
	if err != nil {
		return
	}
	// ListItem Size BE9
	if res["Object.Width"] != "" {
		// Object.width
		ben := fmt.Sprintf("%s%d", "BE", WriteStartRow)
		err = excel.SetCellValue(sn, ben, "幅:"+res["Object.Width"])
		if err != nil {
			return
		}
	}

	// Lock　BA9
	ban := fmt.Sprintf("%s%d", "BA", WriteStartRow)
	err = excel.SetCellValue(sn, ban, res["Locked"])
	if err != nil {
		return
	}
	WriteStartRow += 1
}

func PrintStdout(sb *bytes.Buffer, res map[string]string) {
	sb.WriteString(res["Name"])
	sb.WriteString("\t")
	sb.WriteString(res["Type"])
	sb.WriteString("\t")
	sb.WriteString(res["Caption"])
	sb.WriteString("\t")
	sb.WriteString("\t")
	sb.WriteString(res["Text"])
	sb.WriteString("\t")
	sb.WriteString(res["IMEMode"])
	sb.WriteString("\t")
	sb.WriteString(res["Enabled"])
	sb.WriteString("\t")
	if res["Object.Width"] != "" {
		sb.WriteString("幅:")
		sb.WriteString(res["Object.Width"])
	}
	sb.WriteString("\t")
	sb.WriteString(res["Locked"])
	fmt.Println(sb.String())
}

func OpenTemplateExcel(excel string) *excelize.File {

	f, err := excelize.OpenFile(excel)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func TrimF(r rune) bool {
	switch {
	case unicode.IsSpace(r): // スペース削除
		return true
	case r == 34: // ""を削除
		return true
	default:
		return false
	}
}
