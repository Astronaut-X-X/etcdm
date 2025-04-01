package pkg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadWriteFile(t *testing.T) {
	// 创建临时目录用于测试
	tempDir, err := os.MkdirTemp("", "file-test")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tempDir) // 测试结束后清理

	// 测试文件路径
	testFilePath := filepath.Join(tempDir, "test.txt")

	// 测试数据
	testData := []byte("这是测试数据")

	// 测试 WriteFile
	err = WriteFile(testFilePath, testData)
	if err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	// 测试 ReadFile
	readData, err := ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	// 验证读取的数据是否与写入的数据一致
	if string(readData) != string(testData) {
		t.Errorf("读取的数据与写入的数据不一致，期望: %s, 实际: %s", string(testData), string(readData))
	}

	// 测试读取不存在的文件
	_, err = ReadFile(filepath.Join(tempDir, "non-existent.txt"))
	if err == nil {
		t.Error("读取不存在的文件应该返回错误")
	}
}

func TestWriteFileOverwrite(t *testing.T) {
	// 创建临时目录用于测试
	tempDir, err := os.MkdirTemp("", "file-test")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tempDir) // 测试结束后清理

	// 测试文件路径
	testFilePath := filepath.Join(tempDir, "overwrite.txt")

	// 第一次写入
	firstData := []byte("第一次写入的数据")
	err = WriteFile(testFilePath, firstData)
	if err != nil {
		t.Fatalf("第一次写入文件失败: %v", err)
	}

	// 第二次写入（覆盖）
	secondData := []byte("第二次写入的数据")
	err = WriteFile(testFilePath, secondData)
	if err != nil {
		t.Fatalf("第二次写入文件失败: %v", err)
	}

	// 读取并验证
	readData, err := ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	// 验证读取的数据是否与第二次写入的数据一致（覆盖成功）
	if string(readData) != string(secondData) {
		t.Errorf("覆盖写入失败，期望: %s, 实际: %s", string(secondData), string(readData))
	}
}
