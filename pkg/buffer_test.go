/**
 * @author CHUAN
 * @date 2024-04-30 3:43 pm
 */

package pkg

import "testing"

func BenchmarkBuffer_Add(b *testing.B) {
	buffer := NewBuffer(100)
	temp := []int{}
	for i := 0; i < b.N; i++ {
		temp = append(temp, i)
	}
	b.ResetTimer() // Start timing after setup is complete
	for i := 0; i < b.N; i++ {
		buffer.Add(temp[i])
	}
}

func BenchmarkBuffer_BatchAdd(b *testing.B) {
	buffer := NewBuffer(100)
	temp := []int{}
	for i := 0; i < b.N; i++ {
		temp = append(temp, i)
	}
	b.ResetTimer() // Start timing after setup is complete
	buffer.BatchAdd(temp)
}
