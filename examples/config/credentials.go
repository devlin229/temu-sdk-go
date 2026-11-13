// Package config 提供所有调用示例共用的本地配置读取，不改变 SDK 客户端配置方式。
package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Credentials 是调用示例所需凭据，加载过程不输出其中的值。
type Credentials struct {
	AppKey      string
	AppSecret   string
	AccessToken string
}

// LoadCredentials 优先读取系统环境变量，缺失字段从项目 examples/.env 补齐。
// 支持从项目根目录或其子目录运行；不覆盖系统环境变量，不修改进程环境。
func LoadCredentials() (Credentials, error) {
	keys := []string{"TEMU_APP_KEY", "TEMU_APP_SECRET", "TEMU_ACCESS_TOKEN"}
	values := make(map[string]string, len(keys))
	complete := true
	for _, key := range keys {
		values[key] = strings.TrimSpace(os.Getenv(key))
		if values[key] == "" {
			complete = false
		}
	}
	if !complete {
		path, err := envPath()
		if err != nil {
			return Credentials{}, err
		}
		raw, err := os.ReadFile(path)
		if err != nil && !os.IsNotExist(err) {
			return Credentials{}, fmt.Errorf("读取 examples/.env: %w", err)
		}
		if err == nil {
			local, err := parseEnv(string(raw))
			if err != nil {
				return Credentials{}, err
			}
			for _, key := range keys {
				if values[key] == "" {
					values[key] = local[key]
				}
			}
		}
	}
	var missing []string
	for _, key := range keys {
		if strings.TrimSpace(values[key]) == "" {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		return Credentials{}, fmt.Errorf("缺少凭据 %s，请设置系统环境变量，或复制 examples/.env.example 为 examples/.env 后填写", strings.Join(missing, "、"))
	}
	return Credentials{AppKey: values[keys[0]], AppSecret: values[keys[1]], AccessToken: values[keys[2]]}, nil
}

func envPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("定位 examples/.env: %w", err)
	}
	for {
		if info, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil && !info.IsDir() {
			// examples 是独立模块；继续向上找包含 examples 目录的仓库根目录。
			if examples, err := os.Stat(filepath.Join(dir, "examples")); err == nil && examples.IsDir() {
				return filepath.Join(dir, "examples", ".env"), nil
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("找不到项目 examples/.env，请从项目目录运行或设置完整的 TEMU_APP_KEY、TEMU_APP_SECRET、TEMU_ACCESS_TOKEN 环境变量")
		}
		dir = parent
	}
}

// 支持 KEY=value、export、单/双引号及注释；不执行 shell 或展开变量。
func parseEnv(raw string) (map[string]string, error) {
	values := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimPrefix(raw, "\uFEFF")))
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "export ") || strings.HasPrefix(line, "export\t") {
			line = strings.TrimSpace(line[len("export"):])
		}
		key, value, ok := strings.Cut(line, "=")
		key = strings.TrimSpace(key)
		if !ok || key == "" || strings.ContainsAny(key, " \t\"'") {
			return nil, fmt.Errorf("examples/.env 第 %d 行格式错误，应为 KEY=value", lineNumber)
		}
		value, ok = envValue(strings.TrimSpace(value))
		if !ok {
			return nil, fmt.Errorf("examples/.env 第 %d 行值格式错误，请检查引号", lineNumber)
		}
		values[key] = value
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("读取 examples/.env 行失败: %w", err)
	}
	return values, nil
}

func envValue(raw string) (string, bool) {
	if raw == "" {
		return "", true
	}
	quote := raw[0]
	if quote == '\'' || quote == '"' {
		for i := 1; i < len(raw); i++ {
			if quote == '"' && raw[i] == '\\' {
				i++
				continue
			}
			if raw[i] != quote {
				continue
			}
			tail := strings.TrimSpace(raw[i+1:])
			if tail != "" && !strings.HasPrefix(tail, "#") {
				return "", false
			}
			if quote == '\'' {
				return raw[1:i], true
			}
			value, err := strconv.Unquote(raw[:i+1])
			return value, err == nil
		}
		return "", false
	}
	for i := 0; i < len(raw); i++ {
		if raw[i] == '#' && (i == 0 || raw[i-1] == ' ' || raw[i-1] == '\t') {
			return strings.TrimSpace(raw[:i]), true
		}
	}
	return strings.TrimSpace(raw), true
}
