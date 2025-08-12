package main

import (
	"encoding/xml"
	"fmt"
)

type ModelOutputV4 struct {
	XMLName  xml.Name `xml:"model"`
	Status   string   `json:"status" xml:"status"`
	ToolCall string   `json:"tool_call" xml:"tool_call"`
}

func main() {
	modelOutput1 := &ModelOutputV4{}
	modelOutput2 := &ModelOutputV4{}
	// str1 := `<status>1</status>\n<tool_call>[{'name': 'mcp-server-search-image.edit_image', 'arguments': {'image': 'https://yawen-gips.baidu-int.com/it/u=4005592140,581189832&amp;fm=3042&amp;app=3042&amp;f=PNG?w=837&amp;h=1358b', 'query': '变清晰', 'type': '3'}}]</tool_call>`
	str3 := `<status>1</status>\n<tool_call><![CDATA[[{'name': 'mcp-server-search-image.edit_image', 'arguments': {'image': 'https://yawen-gips.baidu-int.com/it/u=4005592140,581189832&amp;fm=3042&amp;app=3042&amp;f=PNG?w=837&amp;h=1358b', 'query': '变清晰', 'type': '3'}}]]]></tool_call>`
	str2 := `<status>1</status>\n<tool_call>[{'name': 'mcp-server-search-image.text_to_image', 'arguments': {'query': '画一只狗'}}]</tool_call>`

	err := xml.Unmarshal([]byte("<model>"+str2+"</model>"), modelOutput2)
	if err != nil {
		fmt.Printf("xml unmarshal error: %v\n", err)
	}

	err = xml.Unmarshal([]byte("<model>"+str3+"</model>"), modelOutput1)
	if err != nil {
		fmt.Printf("xml unmarshal error: %v\n", err)
	}

	fmt.Printf("modelOutput2 = %v\n", modelOutput2)
	fmt.Printf("modelOutput1 = %v\n", modelOutput1)
}
