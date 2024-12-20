package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	add := Add

	res := add(2, 2)
	exp := 4
	if res != exp {
		t.Fatalf("Expected:\n%v\nFound\n%v\n", exp, res)
	}
}

func TestListAdd10(t *testing.T) {
	res := ListAdd10(Add, []int{1, 2, 3, 4, 5})
	expc := []int{11, 12, 13, 14, 15}

	if len(res) != len(expc) {
		t.Fatal("==============Falid======================")
	}

	for i := range len(res) {
		if res[i] != expc[i] {
			t.Errorf("Expected:\n%v", expc)
			t.Errorf("Found:\n%v", res)
		}
	}

}

func TestEntity(t *testing.T) {
	fn := Entity()

	e1 := fn()

	e2 := fn()
	e3 := fn()

	if e1 != 1 || e2 != 2 || e3 != 3 {
		t.Error("==========Faild=================")
	}
}

func TestCurrySum(t *testing.T) {
	if CurrySum(2)(2) != 4 {
		fmt.Println("Faild curry sum")
	} else {
		fmt.Println("Curry sum PASS")
	}
}
