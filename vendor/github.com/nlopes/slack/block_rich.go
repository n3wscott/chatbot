package slack

// RichTextBlock
type RichTextBlock struct {
	Type      MessageBlockType              `json:"type"`
	BlockID   string                        `json:"block_id,omitempty"`
	Elements  []*RichTextElementBlockObject `json:"elements,omitempty"`
	Accessory *Accessory                    `json:"accessory,omitempty"`
}

// BlockType returns the type of the block
func (s RichTextBlock) BlockType() MessageBlockType {
	return s.Type
}

type RichTextElementBlockObject struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
	//Style *RichTextElementStyle `json:"style,omitempty"`
}

//type RichTextElementStyle struct {
//	Code   *bool `json:"code,omitempty"`
//	Strike *bool `json:"strike,omitempty"`
//	Bold   *bool `json:"bold,omitempty"`
//	Italic *bool `json:"italic,omitempty"`
//}

// Slack decided to make this style element an object or a string. so ... gonna nope out of this for now.
