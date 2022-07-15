package slice

import (
	"testing"
)

//
// TODO
//
func TestAll(t *testing.T) {

	t.Run("Slice Test", func(t *testing.T) {
		slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		slc = Remove(slc, 3)
		if len(slc) != 8 {
			t.Error("Funciton not working")
		}
	})
}
