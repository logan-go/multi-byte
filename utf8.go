package multiByte

type UTF8MultiByte struct {
	Content    []byte
	StartIndex uint64
	EndIndex   uint64
}

func GetUTF8MultiBytes(str []byte) []*UTF8MultiByte {
	length := len(str)
	result := make([]*UTF8MultiByte, 0)
	for i := uint64(0); i < uint64(length); i++ {
		if str[i] >= 0 && str[i] <= 127 {
			result = append(result, &UTF8MultiByte{Content: str[i : i+1], StartIndex: i, EndIndex: i})
		} else if str[i] >= 192 && str[i] <= 223 {
			result = append(result, &UTF8MultiByte{Content: str[i : i+2], StartIndex: i, EndIndex: i + 1})
			i += 1
		} else if str[i] >= 224 && str[i] <= 239 {
			result = append(result, &UTF8MultiByte{Content: str[i : i+3], StartIndex: i, EndIndex: i + 2})
			i += 2
		} else if str[i] >= 240 && str[i] <= 247 {
			result = append(result, &UTF8MultiByte{Content: str[i : i+4], StartIndex: i, EndIndex: i + 3})
			i += 3
		} else if str[i] >= 248 && str[i] <= 251 {
			result = append(result, &UTF8MultiByte{Content: str[i : i+5], StartIndex: i, EndIndex: i + 4})
			i += 4
		} else if str[i] >= 252 && str[i] <= 253 {
			result = append(result, &UTF8MultiByte{Content: str[i : i+6], StartIndex: i, EndIndex: i + 5})
			i += 5
		}
	}
	return result
}
