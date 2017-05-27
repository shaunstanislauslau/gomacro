// this file was generated by gomacro command: import _b "hash"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"hash"
)

// reflection: allow interpreted code to import "hash"
func init() {
	Packages["hash"] = Package{
	Types: map[string]Type{
		"Hash":	TypeOf((*hash.Hash)(nil)).Elem(),
		"Hash32":	TypeOf((*hash.Hash32)(nil)).Elem(),
		"Hash64":	TypeOf((*hash.Hash64)(nil)).Elem(),
	},Proxies: map[string]Type{
		"Hash":	TypeOf((*Hash_hash)(nil)).Elem(),
		"Hash32":	TypeOf((*Hash32_hash)(nil)).Elem(),
		"Hash64":	TypeOf((*Hash64_hash)(nil)).Elem(),
	},
	}
}

// --------------- proxy for hash.Hash ---------------
type Hash_hash struct {
	Object	interface{}
	BlockSize_	func() int
	Reset_	func() 
	Size_	func() int
	Sum_	func(b []byte) []byte
	Write_	func(p []byte) (n int, err error)
}
func (Proxy *Hash_hash) BlockSize() int {
	return Proxy.BlockSize_()
}
func (Proxy *Hash_hash) Reset()  {
	Proxy.Reset_()
}
func (Proxy *Hash_hash) Size() int {
	return Proxy.Size_()
}
func (Proxy *Hash_hash) Sum(b []byte) []byte {
	return Proxy.Sum_(b)
}
func (Proxy *Hash_hash) Write(p []byte) (n int, err error) {
	return Proxy.Write_(p)
}

// --------------- proxy for hash.Hash32 ---------------
type Hash32_hash struct {
	Object	interface{}
	BlockSize_	func() int
	Reset_	func() 
	Size_	func() int
	Sum_	func(b []byte) []byte
	Sum32_	func() uint32
	Write_	func(p []byte) (n int, err error)
}
func (Proxy *Hash32_hash) BlockSize() int {
	return Proxy.BlockSize_()
}
func (Proxy *Hash32_hash) Reset()  {
	Proxy.Reset_()
}
func (Proxy *Hash32_hash) Size() int {
	return Proxy.Size_()
}
func (Proxy *Hash32_hash) Sum(b []byte) []byte {
	return Proxy.Sum_(b)
}
func (Proxy *Hash32_hash) Sum32() uint32 {
	return Proxy.Sum32_()
}
func (Proxy *Hash32_hash) Write(p []byte) (n int, err error) {
	return Proxy.Write_(p)
}

// --------------- proxy for hash.Hash64 ---------------
type Hash64_hash struct {
	Object	interface{}
	BlockSize_	func() int
	Reset_	func() 
	Size_	func() int
	Sum_	func(b []byte) []byte
	Sum64_	func() uint64
	Write_	func(p []byte) (n int, err error)
}
func (Proxy *Hash64_hash) BlockSize() int {
	return Proxy.BlockSize_()
}
func (Proxy *Hash64_hash) Reset()  {
	Proxy.Reset_()
}
func (Proxy *Hash64_hash) Size() int {
	return Proxy.Size_()
}
func (Proxy *Hash64_hash) Sum(b []byte) []byte {
	return Proxy.Sum_(b)
}
func (Proxy *Hash64_hash) Sum64() uint64 {
	return Proxy.Sum64_()
}
func (Proxy *Hash64_hash) Write(p []byte) (n int, err error) {
	return Proxy.Write_(p)
}
