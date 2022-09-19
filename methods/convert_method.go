package methods

import (
	"bytes"
	"strings"
)

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

func SanFengZhang(pinyins []string) string {
	buffer := bytes.Buffer{}
	var firstName string
	for i, pinyin := range pinyins {
		if i == 0 {
			firstName = pinyin
			continue
		}
		buffer.WriteString(pinyin)
	}
	buffer.WriteString(firstName)
	return buffer.String()
}

func SFZhang(pinyins []string) string {
	buffer := bytes.Buffer{}
	var firstName string
	for i, pinyin := range pinyins {
		if i == 0 {
			firstName = pinyin
			continue
		}
		buffer.WriteString(pinyin[:1])
	}
	buffer.WriteString(firstName)
	return buffer.String()
}

func SanFengDotZhang(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffer := bytes.Buffer{}
	var firstName string
	for i, pinyin := range pinyins {
		if i == 0 {
			firstName = pinyin
			continue
		}
		buffer.WriteString(pinyin)
	}
	buffer.WriteString(".")
	buffer.WriteString(firstName)
	return buffer.String()
}

func ZhangDotSanFeng(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffers := bytes.Buffer{}
	for i, pinyin := range pinyins {
		if i == 0 {
			buffers.WriteString(pinyin)
			buffers.WriteString(".")
			continue
		}
		buffers.WriteString(pinyin)
	}
	return buffers.String()
}

func SFDotZhang(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffer := bytes.Buffer{}
	var firstName string
	for i, pinyin := range pinyins {
		if i == 0 {
			firstName = pinyin
			continue
		}
		buffer.WriteString(pinyin[:1])
	}
	buffer.WriteString(".")
	buffer.WriteString(firstName)
	return buffer.String()
}

func ZhangDotSF(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffers := bytes.Buffer{}
	for i, pinyin := range pinyins {
		if i == 0 {
			buffers.WriteString(pinyin)
			buffers.WriteString(".")
			continue
		}
		buffers.WriteString(pinyin[:1])
	}
	return buffers.String()
}

func SDotFZhang(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffer := bytes.Buffer{}
	var firstName string
	for i, pinyin := range pinyins {
		if i == 0 {
			firstName = pinyin
			continue
		}
		buffer.WriteString(pinyin[:1])
		if i <= len(pinyins)-2 {
			buffer.WriteString(".")
		}
	}
	buffer.WriteString(firstName)
	return buffer.String()
}

func SanHYFengDotZhang(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffer := bytes.Buffer{}
	var firstName string
	for i, pinyin := range pinyins {
		if i == 0 {
			firstName = pinyin
			continue
		}
		buffer.WriteString(pinyin)
		if i <= len(pinyins)-2 {
			buffer.WriteString("-")
		}
	}
	buffer.WriteString(".")
	buffer.WriteString(firstName)
	return buffer.String()
}

func SanFeng_Zhang(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffer := bytes.Buffer{}
	var firstName string
	for i, pinyin := range pinyins {
		if i == 0 {
			firstName = pinyin
			continue
		}
		buffer.WriteString(pinyin)
	}
	buffer.WriteString("_")
	buffer.WriteString(firstName)
	return buffer.String()
}

func SF_Zhang(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffer := bytes.Buffer{}
	var firstName string
	for i, pinyin := range pinyins {
		if i == 0 {
			firstName = pinyin
			continue
		}
		buffer.WriteString(pinyin[:1])
	}
	buffer.WriteString("_")
	buffer.WriteString(firstName)
	return buffer.String()
}

func Zhang_SanFeng(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffers := bytes.Buffer{}
	for i, pinyin := range pinyins {
		if i == 0 {
			buffers.WriteString(pinyin)
			buffers.WriteString("_")
			continue
		}
		buffers.WriteString(pinyin)
	}
	return buffers.String()
}

func Zhang_SF(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffers := bytes.Buffer{}
	for i, pinyin := range pinyins {
		if i == 0 {
			buffers.WriteString(pinyin)
			buffers.WriteString("_")
			continue
		}
		buffers.WriteString(pinyin[:1])
	}
	return buffers.String()
}

func San_Feng_Zhang(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	pys := []string{}
	pys = append(pys, pinyins[1:]...)
	pys = append(pys, pinyins[:1]...)
	return Zhang_San_Feng(pys)
}

func Zhang_San_Feng(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffer := bytes.Buffer{}
	for i, pinyin := range pinyins {
		buffer.WriteString(pinyin)
		if i != len(pinyins)-1 {
			buffer.WriteString("_")
		}
	}
	return buffer.String()
}

func ZhangHYSanFeng(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffer := bytes.Buffer{}
	for i, pinyin := range pinyins {
		if i == 0 {
			buffer.WriteString(pinyin)
			buffer.WriteString("-")
			continue
		}
		buffer.WriteString(pinyin)
	}
	return buffer.String()
}

func SanFengHYZhang(pinyins []string) string {
	if len(pinyins) == 1 {
		return pinyins[0]
	}
	buffer := bytes.Buffer{}
	var firstName string
	for i, pinyin := range pinyins {
		if i == 0 {
			firstName = pinyin
			continue
		}
		buffer.WriteString(pinyin)
	}
	buffer.WriteString("-")
	buffer.WriteString(firstName)
	return buffer.String()
}
