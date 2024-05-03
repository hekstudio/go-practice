/**
 * @author CHUAN
 * @date 2024-04-30 3:43 pm
 */

package pkg

import (
	"fmt"
	"testing"
)

func BenchmarkBuffer_AddByAppend(b *testing.B) {
	buffer := NewBuffer(100)
	for i := 0; i < b.N; i++ {
		buffer.AddByAppend(i)
	}
}

func BenchmarkBuffer_AddByShift(b *testing.B) {
	buffer := NewBuffer(100)
	for i := 0; i < b.N; i++ {
		buffer.AddByShift(i)
	}
}

func BenchmarkBuffer_BatchAdd(b *testing.B) {
	buffer := NewBuffer(100)
	var temp []int
	for i := 0; i < b.N; i++ {
		temp = append(temp, i)
	}
	b.ResetTimer() // Start timing after setup is complete
	buffer.BatchAdd(temp)
}

func TestBuffer_AddByAppend1(t *testing.T) {
	bufferSize := 335
	newSize := 334
	buffer := NewBuffer(bufferSize)
	// test AddByAppend
	for i := 0; i < newSize; i++ {
		buffer.AddByAppend(i)
	}
	for i := 0; i < bufferSize; i++ {
		fmt.Printf("%d ", buffer.Data[i])
	}
	fmt.Printf("sum: %d\n", buffer.Sum)
	fmt.Println()
	// test AddByShift
	buffer = NewBuffer(bufferSize)
	for i := 0; i < newSize; i++ {
		buffer.AddByShift(i)
	}
	for i := 0; i < bufferSize; i++ {
		fmt.Printf("%d ", buffer.Data[i])
	}
	fmt.Printf("sum: %d\n", buffer.Sum)
	fmt.Println()
	// test BatchAdd
	buffer = NewBuffer(bufferSize)
	var temp []int
	for i := 0; i < newSize; i++ {
		temp = append(temp, i)
	}
	buffer.BatchAdd(temp)
	for i := 0; i < bufferSize; i++ {
		fmt.Printf("%d ", buffer.Data[i])
	}
	fmt.Printf("sum: %d\n", buffer.Sum)
	fmt.Println()
}

func TestBuffer_AddByAppend2(t *testing.T) {
	bufferSize := 10
	for newSize := 0; newSize < 30; newSize++ {
		buffer := NewBuffer(bufferSize)
		// test AddByAppend
		for i := 0; i < newSize; i++ {
			buffer.AddByAppend(i)
		}
		for i := 0; i < bufferSize; i++ {
			fmt.Printf("%d ", buffer.Data[i])
		}
		fmt.Printf("sum: %d\n", buffer.Sum)
		fmt.Println()
		// test AddByShift
		buffer = NewBuffer(bufferSize)
		for i := 0; i < newSize; i++ {
			buffer.AddByShift(i)
		}
		for i := 0; i < bufferSize; i++ {
			fmt.Printf("%d ", buffer.Data[i])
		}
		fmt.Printf("sum: %d\n", buffer.Sum)
		fmt.Println()
		// test BatchAdd
		buffer = NewBuffer(bufferSize)
		var temp []int
		for i := 0; i < newSize; i++ {
			temp = append(temp, i)
		}
		buffer.BatchAdd(temp)
		for i := 0; i < bufferSize; i++ {
			fmt.Printf("%d ", buffer.Data[i])
		}
		fmt.Printf("sum: %d\n", buffer.Sum)
		fmt.Println("----------------------------------------")
	}
}
