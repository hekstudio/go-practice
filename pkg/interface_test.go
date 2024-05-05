/**
 * @author CHUAN
 * @date 2024-05-05 9:39 am
 */

package pkg

import "testing"

func BenchmarkReadIntDirectly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReadIntDirectly(i)
	}
}

func BenchmarkReadIntFromInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReadIntFromInterface(i)
	}
}
