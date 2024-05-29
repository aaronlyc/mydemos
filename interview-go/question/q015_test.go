package question

import "testing"

func TestBitmap_Exists(t *testing.T) {
	bitmap := NewBitmap(101) // 创建一个新的位图，大小为101

	// 添加数字
	bitmap.Add(10)
	bitmap.Add(20)
	bitmap.Add(11)
	bitmap.Add(21)
	bitmap.Add(12)
	bitmap.Add(22)
	bitmap.Add(13)
	bitmap.Add(23)

	// 检查数字是否存在
	exists, _ := bitmap.Exists(10)
	t.Logf("10 exists: %t", exists)

	// 删除数字
	bitmap.Delete(10)
	exists, _ = bitmap.Exists(10)
	t.Logf("10 exists after deletion: %t", exists)

	// List所有数字
	list := bitmap.List()
	t.Logf("list all: %v", list)
}
