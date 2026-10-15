// Package config — код под тестом для упражнения 03.
package config

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Save пишет cfg в файл строками "key=value", отсортированными по ключу.
func Save(path string, cfg map[string]string) error {
	keys := make([]string, 0, len(cfg))
	for k := range cfg {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, k := range keys {
		fmt.Fprintf(&b, "%s=%s\n", k, cfg[k])
	}
	return os.WriteFile(path, []byte(b.String()), 0o600)
}

// Load читает файл, записанный Save.
func Load(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	cfg := make(map[string]string)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if k, v, ok := strings.Cut(sc.Text(), "="); ok {
			cfg[k] = v
		}
	}
	return cfg, sc.Err()
}
