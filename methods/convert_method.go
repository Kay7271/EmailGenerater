package methods

import (
	"bytes"
	"strings"
)

func SanFengZhang(pinyins []string) string {
	if len(pinyins) <= 2 {
		return ZhangSanFeng(pinyins)
	}
	buffer := bytes.Buffer{}
	var firstName string
	for i, pinyin := range pinyins {
		if i == 0 {
			firstName = pinyin
			continue
		}
		buffer.WriteString(strings.ToLower(pinyin))
	}
	buffer.WriteString(".")
	buffer.WriteString(firstName)
	return buffer.String()
}

func ZhangSanFeng(pinyins []string) string {
	buffer := bytes.Buffer{}
	for _, pinyin := range pinyins {
		buffer.WriteString(strings.ToLower(pinyin))
	}
	return buffer.String()
}

func ZhangSF(pinyins []string) string {
	buffers := bytes.Buffer{}
	for i, pinyin := range pinyins {
		if i == 0 {
			buffers.WriteString(strings.ToLower(pinyin))
			continue
		}
		buffers.WriteString(strings.ToLower(pinyin[:1]))
	}
	return buffers.String()
}

func ZSF(pinyins []string) string {
	buffers := bytes.Buffer{}
	for _, pinyin := range pinyins {
		buffers.WriteString(strings.ToLower(pinyin[:1]))
	}
	return buffers.String()
}
