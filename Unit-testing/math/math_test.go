package math

import "testing"

type AddData struct {
	x, y   int
	result int
}

func TestAdd(t *testing.T) {

	// result := Add(1, 3)

	// if result != 4 {
	// 	t.Errorf("Add(1, 3) FAILED. Expected %d, got %d\n", 4, result)
	// } else {
	// 	t.Logf("Add(1, 3) PASSED. Expected %d, got %d\n", 4, result)
	// }
	testdemo := []AddData{
		{1, 3, 4},
		{-1, 6, 5},
		{3, 5, 8},
	}

	for _, datum := range testdemo {
		result := Add(datum.x, datum.y)

		if result != datum.result {
			t.Errorf("Add(%d, %d) FAILED. Expected %d got %d\n",
				datum.x, datum.y, datum.result, result)
		}
	}
}
func TestDivide(t *testing.T) {

	result := Divide(5, 0)

	if result != 0 {
		t.Errorf("Divide(5,0) FAILED. Expected %d, got %f\n", 0, result)
	} else {
		t.Logf("Divdie(5,0) PASSED. Expected %d, got %f\n", 0, result)
	}
}
