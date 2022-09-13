package methods

import (
	"bytes"
	"strings"
)

func ZhangSanFeng(pinyins []string) string {
	buffer := bytes.Buffer{}
	for _, pinyin := range pinyins {
		buffer.WriteString(strings.ToUpper(pinyin[:1]) + strings.ToLower(pinyin[1:]))
	}
	return buffer.String()
}

func ZhangSF(pinyins []string) string {
	buffers := bytes.Buffer{}
	for i, pinyin := range pinyins {
		if i == 0 {
			buffers.WriteString(strings.ToUpper(pinyin[:1]) + strings.ToLower(pinyin[1:]))
			continue
		}
		buffers.WriteString(strings.ToUpper(pinyin[:1]))
	}
	return buffers.String()
}

func ZSF(pinyins []string) string {
	buffers := bytes.Buffer{}
	for _, pinyin := range pinyins {
		buffers.WriteString(strings.ToUpper(pinyin[:1]))
	}
	return buffers.String()
}
