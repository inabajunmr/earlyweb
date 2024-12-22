package html

import (
	"fmt"
	"strings"
)

type Tag struct {
	Name       string
	Children   []Tag
	Text       string
	Attributes []Attribute
}

type Attribute struct {
	Name  string
	Value string
}

// Parse recursively parses the HTML string and extracts tags with their children
func ParseHtml(html string) []Tag {
	tags := []Tag{}
	stack := []Tag{}

	for len(html) > 0 {
		if strings.HasPrefix(html, "<") {
			end := strings.Index(html, ">")
			if end == -1 {
				break // Malformed HTML
			}

			tag := html[1:end]
			html = html[end+1:]

			if tag == "ISINDEX" {
				if len(stack) == 0 {
					tags = append(tags, Tag{Name: tag})
				} else {
					stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, Tag{Name: tag})
				}
			} else if strings.HasPrefix(tag, "/") {
				// Closing tag
				if len(stack) == 0 || stack[len(stack)-1].Name != tag[1:] {
					fmt.Println("Error: mismatched tags")
					fmt.Println(stack)
					fmt.Println(tag[1:])
					return nil
				}
				completedTag := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if len(stack) == 0 {
					tags = append(tags, completedTag)
				} else {
					stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, completedTag)
				}
			} else {
				// Opening tag
				stack = append(stack, Tag{Name: tag})
			}
		} else {
			// Text content
			end := strings.Index(html, "<")
			if end == -1 {
				end = len(html)
			}
			text := strings.TrimSpace(html[:end])
			html = html[end:]

			if len(stack) > 0 && text != "" {
				stack[len(stack)-1].Text = text
			}
		}
	}

	return tags
}
