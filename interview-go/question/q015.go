package question

// 使用位图实现0～100之内数字的增删改查功能

import "errors"

type Bitmap struct {
	data []uint64
	size int
}

func NewBitmap(size int) *Bitmap {
	return &Bitmap{
		data: make([]uint64, (size+63)/64), // 分配足够的uint64来存储所有位
		size: size,
	}
}

// Add 添加数字到位图
func (b *Bitmap) Add(num int) error {
	if num < 0 || num >= b.size {
		return errors.New("number out of range")
	}
	b.data[num/64] |= 1 << (num % 64) // 设置相应的位
	return nil
}

// Delete 从位图中删除数字
func (b *Bitmap) Delete(num int) error {
	if num < 0 || num >= b.size {
		return errors.New("number out of range")
	}
	b.data[num/64] &^= 1 << (num % 64) // 清除相应的位
	return nil
}

// Exists 检查数字是否存在于位图中
func (b *Bitmap) Exists(num int) (bool, error) {
	if num < 0 || num >= b.size {
		return false, errors.New("number out of range")
	}
	return (b.data[num/64] & (1 << (num % 64))) != 0, nil // 检查相应的位
}

// List 返回位图中所有设置了的数字
func (b *Bitmap) List() []int {
	var result []int
	for i := 0; i < b.size; i++ {
		if (b.data[i/64] & (1 << (i % 64))) != 0 {
			result = append(result, i)
		}
	}
	return result
}
