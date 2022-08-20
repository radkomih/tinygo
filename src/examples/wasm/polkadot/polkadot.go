package main

//go:wasm-module env
//export ext_allocator_malloc_version_1
func extAllocatorMallocVersion1(size uint32) uint32

//go:wasm-module env
//export ext_allocator_free_version_1
func extAllocatorFreeVersion1(ptr uint32)

type VersionData struct {
	SpecName []byte
}

func newVersionData() *VersionData {
	return &VersionData{SpecName: []byte{}}
}

//export Core_version
func CoreVersion(dataPtr uint32, dataLen uint32) uint64 {
	extAllocatorMallocVersion1(0)
	extAllocatorFreeVersion1(0)

	vd := newVersionData()
	vd.SpecName = []byte("gosemble")

	return 1713
}

func main() {}
