package input

import "strings"

func SplitAsMap(vList []string, sep string) map[string]string {
	valueData := make(map[string]string)
	for _, s := range vList {
		s = strings.TrimSpace(s)
		idx := strings.Index(s, sep)
		if idx == -1 {
			continue
		}
		valueData[s[:idx]] = s[idx+1:]
	}
	return valueData
}

func SplitAsMapFunc(vList []string, sep string, f func(string, string) string) map[string]string {
	valueData := make(map[string]string)
	for _, s := range vList {
		s = strings.ReplaceAll(s, " ", "")
		s = strings.TrimSpace(s)
		idx := strings.Index(s, sep)
		if idx == -1 {
			continue
		}
		key := s[:idx]
		value := s[idx+1:]
		if f != nil {
			value = f(key, value)
		}
		valueData[key] = value
	}
	return valueData
}

// SplitAsMapConvFunc use func as callback to map value for any
func SplitAsMapConvFunc[V any](vList []string, sep string, f func(string, string) (V, bool)) map[string]V {
	if f == nil {
		panic("SplitAsMapConvFunc: must set func as callback")
	}
	valueData := make(map[string]V)
	SplitAsMapFunc(vList, sep, func(key string, value string) string {
		conv, ok := f(key, value)
		if !ok {
			return ""
		}
		valueData[key] = conv
		return ""
	})
	return valueData
}

func SplitMapBasic(vList []string) map[string]string {
	return SplitAsMap(vList, ".")
}

func SplitAsMapBasicFunc[V any](vList []string, f func(string, string) (V, bool)) map[string]V {
	return SplitAsMapConvFunc(vList, ".", f)
}
