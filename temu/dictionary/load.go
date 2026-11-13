package dictionary

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
)

//go:embed data/common/*.json data/category/*.json data/product/*.json
var bundled embed.FS

// Load 从 SDK 自带的快照读取数据，不访问网络、凭据、temp 或 docs。
// 每次返回独立实例，修改实例不会污染后续加载或其他调用方。
func Load() (*Snapshot, error) {
	root, err := fs.Sub(bundled, "data")
	if err != nil {
		return nil, fmt.Errorf("open bundled dictionaries: %w", err)
	}
	return LoadFS(root)
}

// LoadFS 从完整的外部快照读取，目录结构与 data 下相同。
// 可传入 os.DirFS，调用方自行维护更新版本；不自动混合不同版本的字典。
func LoadFS(root fs.FS) (*Snapshot, error) {
	if root == nil {
		return nil, errors.New("dictionary filesystem must not be nil")
	}
	snapshot := &Snapshot{}
	entries := []struct {
		path   string
		target any
	}{
		{"common/provinces.json", &snapshot.Provinces},
		{"category/size_chart_classes.json", &snapshot.SizeChartClasses},
		{"category/base_plate_supported.json", &snapshot.BasePlateSupported},
		{"category/model_required.json", &snapshot.ModelRequired},
		{"category/vehicle_required.json", &snapshot.VehicleRequired},
		{"category/net_weight_required.json", &snapshot.NetWeightRequired},
		{"category/total_pieces_required.json", &snapshot.TotalPiecesRequired},
		{"category/size_spec_groups.json", &snapshot.SizeSpecGroups},
		{"product/name_length_rules.json", &snapshot.NameLengthRules},
		{"product/customized_technology.json", &snapshot.CustomizedTechnology},
	}
	for _, entry := range entries {
		raw, err := fs.ReadFile(root, entry.path)
		if err != nil {
			return nil, fmt.Errorf("read dictionary %s: %w", entry.path, err)
		}
		var envelope struct {
			Data json.RawMessage `json:"data"`
		}
		if err := json.Unmarshal(raw, &envelope); err != nil {
			return nil, fmt.Errorf("decode dictionary %s: %w", entry.path, err)
		}
		if len(envelope.Data) == 0 || bytes.Equal(bytes.TrimSpace(envelope.Data), []byte("null")) {
			return nil, fmt.Errorf("dictionary %s: data must not be missing or null", entry.path)
		}
		if err := json.Unmarshal(raw, entry.target); err != nil {
			return nil, fmt.Errorf("decode dictionary %s: %w", entry.path, err)
		}
	}
	if snapshot.NameLengthRules.Data.DefaultLimit <= 0 {
		return nil, errors.New("dictionary product/name_length_rules.json: default_limit must be positive")
	}
	for _, rule := range snapshot.NameLengthRules.Data.Overrides {
		if rule.Limit <= 0 {
			return nil, errors.New("dictionary product/name_length_rules.json: override limit must be positive")
		}
	}
	return snapshot, nil
}
