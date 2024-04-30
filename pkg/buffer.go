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
	return &Buffer{
		data: []int{},
	}
}

// Add
//
//	@Description: add a value to the buffer
//	@receiver b
//	@param value
func (b *Buffer) Add(value int) {
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
	// take the last n elements
	if len(values) > b.size {
		values = values[len(values)-b.size:]
	}
	// reset data with predefined size
	b.data = make([]int, 0, b.size)
	// assign values to data
	for i := 0; i < len(values); i++ {
		b.data[i] = values[i]
	}
}
