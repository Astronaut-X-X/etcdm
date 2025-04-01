package pkg

import (
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	// 测试生成的UUID长度是否正确
	uuid := GenerateUUID()
	if len(uuid) != 36 {
		t.Errorf("生成的UUID长度错误: 期望36, 实际%d", len(uuid))
	}

	// 测试多次生成的UUID是否唯一
	uuidMap := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		uuid := GenerateUUID()
		if uuidMap[uuid] {
			t.Error("生成了重复的UUID")
		}
		uuidMap[uuid] = true
	}

	// 测试UUID格式是否正确
	uuid = GenerateUUID()
	if matched := isValidUUIDFormat(uuid); !matched {
		t.Errorf("生成的UUID格式错误: %s", uuid)
	}
}

// 检查UUID格式是否符合标准
func isValidUUIDFormat(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	// 检查分隔符位置
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	// 检查是否只包含有效的十六进制字符和分隔符
	for i, c := range uuid {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			if c != '-' {
				return false
			}
			continue
		}
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}

	return true
}
