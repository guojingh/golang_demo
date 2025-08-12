package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

func ExtractRawXML(data []byte, tagName string) (string, error) {
	d := xml.NewDecoder(bytes.NewReader(data))
	var buf bytes.Buffer
	inTag := false

	for {
		tok, err := d.RawToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local == tagName {
				inTag = true
				buf.WriteString(fmt.Sprintf("<%s>", tagName))
			}
		case xml.EndElement:
			if t.Name.Local == tagName && inTag {
				buf.WriteString(fmt.Sprintf("</%s>", tagName))
				return buf.String(), nil
			}
		case xml.CharData:
			if inTag {
				buf.Write([]byte(t))
			}
		case xml.Comment:
			if inTag {
				buf.WriteString(fmt.Sprintf("<!--%s-->", string(t)))
			}
		case xml.Directive:
			if inTag {
				buf.WriteString(fmt.Sprintf("<!%s>", string(t)))
			}
		case xml.ProcInst:
			if inTag {
				buf.WriteString(fmt.Sprintf("<?%s %s?>", t.Target, string(t.Inst)))
			}
		}
	}
	return "", fmt.Errorf("tag <%s> not found", tagName)
}

func main() {
	xmlData := []byte(`<status>1</status>\n<tool_call>[{'name': 'mcp-server-search-image.edit_image', 'arguments': {'image': 'https://yawen-gips.baidu-int.com/it/u=4005592140,581189832&fm=3042&app=3042&f=PNG?w=837&h=1358b', 'query': '变清晰', 'type': '3'}}]</tool_call>`)
	raw, err := ExtractRawXML(xmlData, "tool_call")
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("原始XML片段:", raw)
	}
}
