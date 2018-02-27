package vm

import (
	"testing"
)

func TestVM_ICONST(t *testing.T) {
	program := []int{
		ICONST, 42,
		HALT,
	}

	vm := NewVM(program)
	vm.Run()
	if vm.stack.Peek() != 42 {
		t.Fatalf("incorrect value on the stack. got=%d", vm.stack.Peek())
	}
}

func TestVM_PRINT(t *testing.T) {
	program := []int{
		ICONST, 42,
		PRINT,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 0 {
		t.Fatalf("value is still present on stack. stack size is %d", vm.stack.Size())
	}
}

func TestVM_PRINT_TwoIntegers(t *testing.T) {
	program := []int{
		ICONST, 42,
		ICONST, 43,
		PRINT,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}
}

func TestVM_IADD(t *testing.T) {
	program := []int{
		ICONST, 42,
		ICONST, 43,
		IADD,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 85 {
		t.Fatalf("incorrect result. result is %d but it should be 85", vm.stack.Peek())
	}
}

func TestVM_ISUB(t *testing.T) {
	program := []int{
		ICONST, 2,
		ICONST, 4,
		ISUB,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 2 {
		t.Fatalf("incorrect result. result is %d but it should be 2", vm.stack.Peek())
	}
}

func TestVM_IMUL(t *testing.T) {
	program := []int{
		ICONST, 2,
		ICONST, 4,
		IMUL,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 8 {
		t.Fatalf("incorrect result. result is %d but it should be 8", vm.stack.Peek())
	}
}

func TestVM_ILT(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 2,
		ILT,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 1 {
		t.Fatalf("incorrect result. result is %d but it should be 1", vm.stack.Peek())
	}
}

func TestVM_ILT_NotLessThan(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 7,
		ILT,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 0 {
		t.Fatalf("incorrect result. result is %d but it should be 0", vm.stack.Peek())
	}
}

func TestVM_IEQ(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 4,
		IEQ,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 1 {
		t.Fatalf("incorrect result. result is %d but it should be 1", vm.stack.Peek())
	}
}

func TestVM_IEQ_NotEqual(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 6,
		IEQ,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 0 {
		t.Fatalf("incorrect result. result is %d but it should be 0", vm.stack.Peek())
	}
}

func TestVM_JMP(t *testing.T) {
	// Expect that stack size will be 2 since ICONST 7 and ICONST 8 will be skipped
	program := []int{
		ICONST, 4,
		ICONST, 6,
		JMP, 10,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 2 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 2", vm.stack.Size())
	}
}

func TestVM_JMPT(t *testing.T) {
	// Expect that stack size will be 0 since the ICONST 7 and ICONST 8 will be skipped
	program := []int{
		ICONST, 4,
		ICONST, 4,
		IEQ,
		JMPT, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 0 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 0", vm.stack.Size())
	}
}

func TestVM_JMPT_NotTrue(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 3,
		IEQ,
		JMPT, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 2 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 2", vm.stack.Size())
	}
}

func TestVM_JMPF(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 3,
		IEQ,
		JMPF, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 0 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 0", vm.stack.Size())
	}
}

func TestVM_JMPF_NotFalse(t *testing.T) {
	program := []int{
		ICONST, 4,
		ICONST, 4,
		IEQ,
		JMPF, 11,
		ICONST, 7,
		ICONST, 8,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 2 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 2", vm.stack.Size())
	}
}

func TestVM_GSTORE(t *testing.T) {
	// Size of the locals space should be 1 since value at 0 address is set for both 42 and 43
	program := []int{
		ICONST, 42,
		GSTORE, 0,
		ICONST, 43,
		GSTORE, 0,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 0 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 0", vm.stack.Size())
	}

	if len(vm.locals) != 1 {
		t.Fatalf("incorrect size of locals space. size of locals space is %d but is should be 1", len(vm.locals))
	}

	if vm.locals[0] != 43 {
		t.Fatalf("incorrect value at 0 address. value is %d but it should be 43", vm.locals[0])
	}
}

func TestVM_GLOAD(t *testing.T) {
	program := []int{
		ICONST, 42,
		GSTORE, 0,
		GLOAD, 0,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Size() != 1 {
		t.Fatalf("incorrect size of the stack. stack size is %d but it should be 1", vm.stack.Size())
	}

	if vm.stack.Peek() != 42 {
		t.Fatalf("incorrect value on stack. got %d but it should be 42", vm.stack.Peek())
	}
}

func TestVM_LOAD(t *testing.T) {
	program := []int{
		ICONST, 1,
		ICONST, 2,
		LOAD, 1,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if vm.stack.Peek() != 2 {
		t.Fatalf("top element on stack is %d but it should be 2", vm.stack.Peek())
	}

	if vm.stack.Size() != 3 {
		t.Fatalf("the size of the stack is %d but it should be 3", vm.stack.Size())
	}
}

func TestVM_STORE(t *testing.T) {
	program := []int{
		ICONST, 1,
		ICONST, 2,
		STORE, 1,
		HALT,
	}
	vm := NewVM(program)
	vm.Run()

	if len(vm.locals) != 1 {
		t.Fatalf("the size of the locals should be 1. got %d", len(vm.locals))
	}

	if vm.stack.Size() != 1 {
		t.Fatalf("the size of the stack should be 1. got %d", vm.stack.Size())
	}
}
