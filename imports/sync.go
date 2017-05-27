// this file was generated by gomacro command: import _b "sync"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"sync"
)

// reflection: allow interpreted code to import "sync"
func init() {
	Packages["sync"] = Package{
	Binds: map[string]Value{
		"NewCond":	ValueOf(sync.NewCond),
	},Types: map[string]Type{
		"Cond":	TypeOf((*sync.Cond)(nil)).Elem(),
		"Locker":	TypeOf((*sync.Locker)(nil)).Elem(),
		"Mutex":	TypeOf((*sync.Mutex)(nil)).Elem(),
		"Once":	TypeOf((*sync.Once)(nil)).Elem(),
		"Pool":	TypeOf((*sync.Pool)(nil)).Elem(),
		"RWMutex":	TypeOf((*sync.RWMutex)(nil)).Elem(),
		"WaitGroup":	TypeOf((*sync.WaitGroup)(nil)).Elem(),
	},Proxies: map[string]Type{
		"Locker":	TypeOf((*Locker_sync)(nil)).Elem(),
	},
	}
}

// --------------- proxy for sync.Locker ---------------
type Locker_sync struct {
	Object	interface{}
	Lock_	func() 
	Unlock_	func() 
}
func (Proxy *Locker_sync) Lock()  {
	Proxy.Lock_()
}
func (Proxy *Locker_sync) Unlock()  {
	Proxy.Unlock_()
}
