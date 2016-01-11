package fuzzy_test

import "selvan.github.com/fuzzy"
import "testing"
import "strings"

func assertEq(t *testing.T, first float32, second float32) {
	if first != second {
		t.Errorf("Expected %f, got %f", first, second)
	}

	t.Logf("assertEq :: First %f , second %f", first, second)	
}

func assertGt(t *testing.T, first float32, second float32) {
	if first <= second {
		t.Errorf("Expected %f >= %f", first, second)
	}

	t.Logf("assertGt :: First %f , second %f", first, second)
}


func TestComputeScore(t *testing.T) {
	assertEq(t, 1, fuzzy.ComputeScore("hello world", "hello world"))
	assertEq(t, 0, fuzzy.ComputeScore("hello world", ""))

	assertEq(t, 0, fuzzy.ComputeScore("hello world", "Hello World"))
	assertEq(t, 1, fuzzy.ComputeScore("hello world", strings.ToLower("Hello World")))

	assertEq(t, 0, fuzzy.ComputeScore("he", "hel"))
	assertEq(t, 0, fuzzy.ComputeScore("hetz tight", "world"))

	assertGt(t, fuzzy.ComputeScore("hello world", "hello"), fuzzy.ComputeScore("hello world", "lo"))
	assertGt(t, fuzzy.ComputeScore("hello world", "hello"), fuzzy.ComputeScore("hello world", "World"))
	assertGt(t, fuzzy.ComputeScore("hello world", "h"), fuzzy.ComputeScore("hello world", "w"))
	assertGt(t, fuzzy.ComputeScore("hello world", "h"), fuzzy.ComputeScore("he", "hel"))
	assertGt(t, fuzzy.ComputeScore("hello world", "world"), fuzzy.ComputeScore("helloworld", "world"))
	assertGt(t, fuzzy.ComputeScore("hello-world", "world"), fuzzy.ComputeScore("helloworld", "world"))
	assertGt(t, fuzzy.ComputeScore("hello_world", "world"), fuzzy.ComputeScore("helloworld", "world"))
	assertEq(t, fuzzy.ComputeScore("hello world", "world"), fuzzy.ComputeScore("hello_world", "world"))
	assertGt(t, fuzzy.ComputeScore("hello world", "ed"), fuzzy.ComputeScore("hello world", "zz"))
	
	assertEq(t, 1, fuzzy.ComputeScore("世界", "世界"))
	assertGt(t, fuzzy.ComputeScore("世界", "世"), 0)

	
}