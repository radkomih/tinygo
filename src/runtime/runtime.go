package runtime

import (
	"unsafe"
)

//go:generate go run ../../tools/gen-critical-atomics -out ./atomics_critical.go

const Compiler = "tinygo"

// The compiler will fill this with calls to the initialization function of each
// package.
func initAll()

//go:linkname callMain main.main
func callMain()

func GOMAXPROCS(n int) int {
	// Note: setting GOMAXPROCS is ignored.
	return 1
}

func GOROOT() string {
	// TODO: don't hardcode but take the one at compile time.
	return "/usr/local/go"
}

// Copy size bytes from src to dst. The memory areas must not overlap.
// This function is implemented by the compiler as a call to a LLVM intrinsic
// like llvm.memcpy.p0i8.p0i8.i32(dst, src, size, false).
func memcpy(dst, src unsafe.Pointer, size uintptr) {
	for i := uintptr(0); i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(dst) + i)) = *(*uint8)(unsafe.Pointer(uintptr(src) + i))
	}
}

// Copy size bytes from src to dst. The memory areas may overlap and will do the
// correct thing.
// This function is implemented by the compiler as a call to a LLVM intrinsic
// like llvm.memmove.p0i8.p0i8.i32(dst, src, size, false).
func memmove(dst, src unsafe.Pointer, size uintptr) {
	if uintptr(dst) < uintptr(src) {
		// Copy forwards.
		memcpy(dst, src, size)
		return
	}
	// Copy backwards.
	for i := size; i != 0; {
		i--
		*(*uint8)(unsafe.Pointer(uintptr(dst) + i)) = *(*uint8)(unsafe.Pointer(uintptr(src) + i))
	}
}

// Set the given number of bytes to zero.
// This function is implemented by the compiler as a call to a LLVM intrinsic
// like llvm.memset.p0i8.i32(ptr, 0, size, false).
func memzero(ptr unsafe.Pointer, size uintptr) {
	for i := uintptr(0); i < size; i++ {
		*(*byte)(unsafe.Pointer(uintptr(ptr) + i)) = 0
	}
}

// This intrinsic returns the current stack pointer.
// It is normally used together with llvm.stackrestore but also works to get the
// current stack pointer in a platform-independent way.
//export llvm.stacksave
func stacksave() unsafe.Pointer

//export strlen
func strlen(ptr unsafe.Pointer) uintptr

//export malloc
func malloc(size uintptr) unsafe.Pointer

// Compare two same-size buffers for equality.
func memequal(x, y unsafe.Pointer, n uintptr) bool {
	for i := uintptr(0); i < n; i++ {
		cx := *(*uint8)(unsafe.Pointer(uintptr(x) + i))
		cy := *(*uint8)(unsafe.Pointer(uintptr(y) + i))
		if cx != cy {
			return false
		}
	}
	return true
}

func nanotime() int64 {
	return ticksToNanoseconds(ticks())
}

// Copied from the Go runtime source code.
//go:linkname os_sigpipe os.sigpipe
func os_sigpipe() {
	runtimePanic("too many writes on closed pipe")
}

// LockOSThread wires the calling goroutine to its current operating system thread.
// Stub for now
// Called by go1.18 standard library on windows, see https://github.com/golang/go/issues/49320
func LockOSThread() {
}

// UnlockOSThread undoes an earlier call to LockOSThread.
// Stub for now
func UnlockOSThread() {
}
