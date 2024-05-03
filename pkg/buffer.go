/**
 * @author CHUAN
 * @date 2024-04-30 3:29 pm
 */

package pkg

// Buffer
// @Description: a buffer that stores int
type Buffer struct {
	Data []int
	Sum  int
	Size int // Size of the buffer
}

// NewBuffer
//
//	@Description: creates a new Buffer
//	@param Size
//	@return *Buffer
func NewBuffer(size int) *Buffer {
	if size == 0 {
		panic("Size cannot be 0")
	}
	return &Buffer{
		Data: make([]int, size),
		Size: size,
	}
}

// AddByAppend
//
//	@Description: add a value to the buffer
//	@receiver b
//	@param value
func (b *Buffer) AddByAppend(value int) {
	b.Data = append(b.Data, value)
	// add the new value
	b.Sum += value
	if len(b.Data) > b.Size {
		// subtract the value to be removed first
		b.Sum -= b.Data[0]
		b.Data = b.Data[1:]
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
	// reset sum
	b.Sum = 0
	// take the last n elements
	if newSize >= b.Size {
		for i := 0; i < b.Size; i++ {
			b.Data[i] = values[i+newSize-b.Size]
			b.Sum += values[i+newSize-b.Size]
		}
		return
	}
	// shift existing Data
	shift := b.Size - newSize
	for i := 0; i < newSize; i++ {
		b.Sum += b.Data[i+shift]
		b.Data[i] = b.Data[i+shift]
	}
	// set new values
	for i := 0; i < newSize; i++ {
		b.Sum += values[i]
		b.Data[i+shift] = values[i]
	}
}

// AddByShift
//
//	@Description: add a value to the buffer by shifting
//	@receiver b
//	@param value
func (b *Buffer) AddByShift(value int) {
	size := len(b.Data)
	if size == 0 {
		// do nothing
		return
	}
	// subtract the value to be removed
	b.Sum -= b.Data[0]
	for i := 0; i < b.Size-1; i++ {
		b.Data[i] = b.Data[i+1]
	}
	b.Data[b.Size-1] = value
	// add the new value
	b.Sum += value
}

func (b *Buffer) GetData() []int {
	return b.Data
}
