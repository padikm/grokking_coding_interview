package leet

func IsValid(s string) bool {
	//1. Map of paranthesis
	var stack []string
	var braces = map[string]string{")": "(", "}": "{", "]": "["}
	for i, _ := range s {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, string(s[i]))
		} else if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			if len(stack) == 0 {
				return false
			} else if braces[string(s[i])] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
