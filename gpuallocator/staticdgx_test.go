// Copyright (c) 2019, NVIDIA CORPORATION. All rights reserved.

package gpuallocator

import "testing"

func TestStaticDGX1PascalGPUAllocOne(t *testing.T) {
	devices := NewDGX1PascalNode().Devices()
	policy := NewStaticDGX1Policy(GPUTypePascal)
	allocator := newAllocatorFrom(devices, policy)

	tests := []AllocTest{
		{1, []int{0}},
		{1, []int{1}},
		{1, []int{2}},
		{1, []int{3}},
		{1, []int{4}},
		{1, []int{5}},
		{1, []int{6}},
		{1, []int{7}},
		{1, []int{}},
	}

	RunAllocTests(t, allocator, tests)
}

func TestStaticDGX1PascalGPUAllocTwo(t *testing.T) {
	devices := NewDGX1PascalNode().Devices()
	policy := NewStaticDGX1Policy(GPUTypePascal)
	allocator := newAllocatorFrom(devices, policy)

	tests := []AllocTest{
		{2, []int{0, 2}},
		{2, []int{1, 3}},
		{2, []int{4, 6}},
		{2, []int{5, 7}},
		{1, []int{}},
	}

	RunAllocTests(t, allocator, tests)
}

func TestStaticDGX1PascalGPUAllocFour(t *testing.T) {
	devices := NewDGX1PascalNode().Devices()
	policy := NewStaticDGX1Policy(GPUTypePascal)
	allocator := newAllocatorFrom(devices, policy)

	tests := []AllocTest{
		{4, []int{0, 1, 2, 3}},
		{4, []int{4, 5, 6, 7}},
		{1, []int{}},
	}

	RunAllocTests(t, allocator, tests)
}

func TestStaticDGX1PascalGPUAllocEight(t *testing.T) {
	devices := NewDGX1PascalNode().Devices()
	policy := NewStaticDGX1Policy(GPUTypePascal)
	allocator := newAllocatorFrom(devices, policy)

	tests := []AllocTest{
		{8, []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{1, []int{}},
	}

	RunAllocTests(t, allocator, tests)
}

func TestStaticDGX1PascalGPUAllocInvalid(t *testing.T) {
	devices := NewDGX1PascalNode().Devices()
	policy := NewStaticDGX1Policy(GPUTypePascal)
	allocator := newAllocatorFrom(devices, policy)

	tests := []AllocTest{
		{0, []int{}},
		{3, []int{}},
		{5, []int{}},
		{6, []int{}},
		{7, []int{}},
		{9, []int{}},
	}

	RunAllocTests(t, allocator, tests)
}

func TestStaticDGX1VoltaGPUAllocOne(t *testing.T) {
	devices := NewDGX1VoltaNode().Devices()
	policy := NewStaticDGX1Policy(GPUTypeVolta)
	allocator := newAllocatorFrom(devices, policy)

	tests := []AllocTest{
		{1, []int{0}},
		{1, []int{1}},
		{1, []int{2}},
		{1, []int{3}},
		{1, []int{4}},
		{1, []int{5}},
		{1, []int{6}},
		{1, []int{7}},
		{1, []int{}},
	}

	RunAllocTests(t, allocator, tests)
}

func TestStaticDGX1VoltaGPUAllocTwo(t *testing.T) {
	devices := NewDGX1VoltaNode().Devices()
	policy := NewStaticDGX1Policy(GPUTypeVolta)
	allocator := newAllocatorFrom(devices, policy)

	tests := []AllocTest{
		{2, []int{0, 3}},
		{2, []int{1, 2}},
		{2, []int{4, 7}},
		{2, []int{5, 6}},
		{1, []int{}},
	}

	RunAllocTests(t, allocator, tests)
}

func TestStaticDGX1VoltaGPUAllocFour(t *testing.T) {
	devices := NewDGX1VoltaNode().Devices()
	policy := NewStaticDGX1Policy(GPUTypeVolta)
	allocator := newAllocatorFrom(devices, policy)

	tests := []AllocTest{
		{4, []int{0, 1, 2, 3}},
		{4, []int{4, 5, 6, 7}},
		{1, []int{}},
	}

	RunAllocTests(t, allocator, tests)
}

func TestStaticDGX1VoltaGPUAllocEight(t *testing.T) {
	devices := NewDGX1VoltaNode().Devices()
	policy := NewStaticDGX1Policy(GPUTypeVolta)
	allocator := newAllocatorFrom(devices, policy)

	tests := []AllocTest{
		{8, []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{1, []int{}},
	}

	RunAllocTests(t, allocator, tests)
}

func TestStaticDGX1VoltaGPUAllocInvalid(t *testing.T) {
	devices := NewDGX1VoltaNode().Devices()
	policy := NewStaticDGX1Policy(GPUTypeVolta)
	allocator := newAllocatorFrom(devices, policy)

	tests := []AllocTest{
		{0, []int{}},
		{3, []int{}},
		{5, []int{}},
		{6, []int{}},
		{7, []int{}},
		{9, []int{}},
	}

	RunAllocTests(t, allocator, tests)
}
