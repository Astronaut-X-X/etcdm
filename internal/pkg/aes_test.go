package pkg

import "testing"

func TestEncryptDecrypt(t *testing.T) {
	tests := []struct {
		name      string
		plaintext string
		wantErr   bool
	}{
		{
			name:      "空字符串",
			plaintext: "",
			wantErr:   false,
		},
		{
			name:      "普通字符串",
			plaintext: "hello world",
			wantErr:   false,
		},
		{
			name:      "中文字符串",
			plaintext: "你好，世界",
			wantErr:   false,
		},
		{
			name:      "特殊字符",
			plaintext: "!@#$%^&*()_+",
			wantErr:   false,
		},
		{
			name:      "长字符串",
			plaintext: "这是一个很长的字符串，包含中文、English、数字123和特殊字符!@#",
			wantErr:   false,
		},
		{
			name:      "超长字符串",
			plaintext: "这是一个500字符长度的测试字符串，包含了常见的各类字符：中文汉字测试、English letters and words、数字123456789、特殊字符!@#$%^&*()_+-=[]{}|;:'\",.<>?/、重复内容测试测试、Unicode表情符号😀🎉🌟、换行符\n\t制表符。这里继续添加一些内容来达到指定长度：今天天气真好，阳光明媚，适合出去散步。The quick brown fox jumps over the lazy dog. 床前明月光，疑是地上霜。举头望明月，低头思故乡。春眠不觉晓，处处闻啼鸟。夜来风雨声，花落知多少。枯藤老树昏鸦，小桥流水人家。古道西风瘦马，夕阳西下，断肠人在天涯。",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试加密
			encrypted, err := Encrypt(tt.plaintext)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 如果加密成功，测试解密
			if err == nil {
				decrypted, err := Decrypt(encrypted)
				if (err != nil) != tt.wantErr {
					t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				// 验证解密后的结果是否与原文相同
				if decrypted != tt.plaintext {
					t.Errorf("解密结果与原文不匹配，got = %v, want %v", decrypted, tt.plaintext)
				}
			}
		})
	}
}

func TestDecrypt_InvalidInput(t *testing.T) {
	tests := []struct {
		name       string
		ciphertext string
		wantErr    bool
	}{
		{
			name:       "非法Base64字符串",
			ciphertext: "这不是一个有效的Base64字符串",
			wantErr:    true,
		},
		{
			name:       "空字符串",
			ciphertext: "",
			wantErr:    true,
		},
		{
			name:       "无效的加密数据",
			ciphertext: "SGVsbG8=", // 有效的Base64但不是有效的加密数据
			wantErr:    true,
		},
		{
			name:       "有效加密数据",
			ciphertext: "7LAWiUufBHJdwuHAfPQi8Q==", // 有效的Base64但不是有效的加密数据
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Decrypt(tt.ciphertext)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
