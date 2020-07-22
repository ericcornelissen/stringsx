package stringsx

import (
	"fmt"
	"testing"
)

func ExampleAll() {
	elems := []string{"Hello", "world", "!"}
	all := All(elems, func(s string) bool { return len(s) > 0 })
	fmt.Print(all)
	// Output: true
}

func ExampleAny() {
	elems := []string{"Hello", "world", "!"}
	any := Any(elems, func(s string) bool { return s == "world" })
	fmt.Print(any)
	// Output: true
}

func ExampleIsEmpty() {
	fmt.Println(IsEmpty(""))
	fmt.Println(IsEmpty("foobar"))
	// Output: true
	// false
}

func ExampleMapAll() {
	elems := []string{" foo  ", "  bar "}
	mapped := MapAll(elems, TrimSpace)
	fmt.Print(mapped)
	// Output: [foo bar]
}

func TestAll(t *testing.T) {
	condition := func(s string) bool {
		return s == ""
	}

	t.Run("empty list", func(t *testing.T) {
		elems := []string{}
		result := All(elems, condition)

		if result == false {
			t.Error("Unexpected result for empty list")
		}
	})
	t.Run("condition never holds", func(t *testing.T) {
		elems := []string{"a", "b", "c"}
		result := All(elems, condition)

		if result == true {
			t.Error("Unexpected result for list")
		}
	})
	t.Run("condition holds once", func(t *testing.T) {
		elems := []string{"a", "", "c"}
		result := All(elems, condition)

		if result == true {
			t.Error("Unexpected result for list")
		}
	})
	t.Run("condition holds most of the time", func(t *testing.T) {
		elems := []string{"", "", "c"}
		result := All(elems, condition)

		if result == true {
			t.Error("Unexpected result for list")
		}
	})
	t.Run("condition always holds", func(t *testing.T) {
		elems := []string{"", "", ""}
		result := All(elems, condition)

		if result == false {
			t.Error("Unexpected result for list")
		}
	})
}

func TestAny(t *testing.T) {
	condition := func(s string) bool {
		return s == ""
	}

	t.Run("empty list", func(t *testing.T) {
		elems := []string{}
		result := Any(elems, condition)

		if result == true {
			t.Error("Unexpected result for empty list")
		}
	})
	t.Run("condition never holds", func(t *testing.T) {
		elems := []string{"a", "b", "c"}
		result := Any(elems, condition)

		if result == true {
			t.Error("Unexpected result for list")
		}
	})
	t.Run("condition holds once", func(t *testing.T) {
		elems := []string{"a", "", "c"}
		result := Any(elems, condition)

		if result == false {
			t.Error("Unexpected result for list")
		}
	})
	t.Run("condition holds most of the time", func(t *testing.T) {
		elems := []string{"", "", "c"}
		result := Any(elems, condition)

		if result == false {
			t.Error("Unexpected result for list")
		}
	})
	t.Run("condition always holds", func(t *testing.T) {
		elems := []string{"", "", ""}
		result := Any(elems, condition)

		if result == false {
			t.Error("Unexpected result for list")
		}
	})
}

func TestIsEmpty(t *testing.T) {
	if isEmpty := IsEmpty(""); isEmpty == false {
		t.Error("Unexpected result for empty string")
	}

	if isEmpty := IsEmpty("Hello world!"); isEmpty == true {
		t.Error("Unexpected result for non-empty string")
	}
}

func TestMapAll(t *testing.T) {
	elems := []string{
		"a",
		" b",
		"c ",
		" d ",
	}

	mapped := MapAll(elems, TrimSpace)

	if mapped[0] != TrimSpace(elems[0]) {
		t.Errorf("Incorrect first value (got '%s')", mapped[0])
	}

	if mapped[1] != TrimSpace(elems[1]) {
		t.Errorf("Incorrect second value (got '%s')", mapped[1])
	}

	if mapped[2] != TrimSpace(elems[2]) {
		t.Errorf("Incorrect third value (got '%s')", mapped[2])
	}

	if mapped[3] != TrimSpace(elems[3]) {
		t.Errorf("Incorrect fourth value (got '%s')", mapped[3])
	}
}
