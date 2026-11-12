package temu

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// sign 仅排序外层字段；对象和数组保留原有顺序，以紧凑 JSON 参与签名。
func sign(params map[string]json.RawMessage, secret string) (string, error) {
	keys := make([]string, 0, len(params))
	for key := range params {
		if key != "sign" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	var raw strings.Builder
	raw.WriteString(secret)
	for _, key := range keys {
		raw.WriteString(key)
		value := params[key]
		if len(value) > 0 && value[0] == '"' {
			var text string
			if err := json.Unmarshal(value, &text); err != nil {
				return "", err
			}
			raw.WriteString(text)
		} else {
			var compact bytes.Buffer
			if err := json.Compact(&compact, value); err != nil {
				return "", err
			}
			raw.Write(compact.Bytes())
		}
	}
	raw.WriteString(secret)
	return fmt.Sprintf("%X", md5.Sum([]byte(raw.String()))), nil
}
