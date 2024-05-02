/**
 * @author CHUAN
 * @date 2024-04-30 3:29 pm
 */

package pkg

// Buffer
// @Description: a buffer that stores int
type Buffer struct {
	data []int
	size int // size of the buffer
}

// NewBuffer
//
//	@Description: creates a new Buffer
//	@param size
//	@return *Buffer
func NewBuffer(size int) *Buffer {
	if size == 0 {
		panic("size cannot be 0")
	}
	return &Buffer{
		data: make([]int, size),
		size: size,
	}
}

// AddByAppend
//
//	@Description: add a value to the buffer
//	@receiver b
//	@param value
func (b *Buffer) AddByAppend(value int) {
	b.data = append(b.data, value)
	if len(b.data) > b.size {
		b.data = b.data[1:]
	}
}

// BatchAdd
//
//	@Description: add multiple values to the buffer
//	@receiver b
//	@param values
func (b *Buffer) BatchAdd(values []int) {
	newSize := len(values)
	if newSize == 0 {
		// nothing to add
		return
	}
	// take the last n elements
	if newSize >= b.size {
		for i := 0; i < b.size; i++ {
			b.data[i] = values[i+newSize-b.size]
		}
		return
	}
	// shift existing data
	shift := b.size - newSize
	for i := 0; i < shift; i++ {
		b.data[i] = b.data[i+newSize]
	}
	// set new values
	for i := 0; i < newSize; i++ {
		b.data[newSize+i] = values[i]
	}
}

func (b *Buffer) AddByShift(value int) {
	size := len(b.data)
	if size == 0 {
		// do nothing
		return
	}
	for i := 0; i < b.size-1; i++ {
		b.data[i] = b.data[i+1]
	}
	b.data[b.size-1] = value
}
