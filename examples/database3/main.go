package main

import (
	"github.com/uuosio/chain"
	"github.com/uuosio/chain/database"
)

//table mydata
type MyData struct {
	primary uint64        //primary:t.primary
	a1      uint64        //IDX64: bya1 : t.a1 : t.a1=%v
	a2      chain.Uint128 //IDX128: bya2 : t.a2 : t.a2=%v
}

//contract test
type MyContract struct {
	self          chain.Name
	firstReceiver chain.Name
	action        chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *MyContract {
	return &MyContract{receiver, firstReceiver, action}
}

//action test
func (t *MyContract) Test() {
	payer := t.self

	mi := NewMyDataDB(t.self, t.self)
	primary := uint64(1000)

	if it, data := mi.Get(primary); it.IsOk() {
		n := data.a2.Uint64()
		chain.Println("++++data.a2:", n)
		n += 1
		data.a2.SetUint64(n)
		mi.Update(it, data, payer)
	} else {
		data := &MyData{}
		data.primary = primary
		data.a1 = uint64(1)
		data.a2.SetUint64(1)
		mi.Store(data, payer)
	}

	secondary := uint64(0)
	{
		var it database.SecondaryIterator
		idxDB := mi.GetIdxDBbya1()
		it, secondary = idxDB.FindByPrimary(primary)
		chain.Check(it.IsOk(), "Invalid secondary iterator")
		secondary += 1
		chain.Println(idxDB.GetIndex(), it.I, it.Primary, secondary)
		if it.IsOk() {
			mi.IdxUpdate(it, secondary, payer)
		}
	}
	chain.Println("++++++++secondary:", secondary)
	{
		idxDB := mi.GetIdxDB("bya1")
		it := idxDB.Find(secondary)
		chain.Check(it.IsOk(), "Invalid secondary iterator")
		secondary += 1
		chain.Println(idxDB.GetIndex(), it.I, it.Primary, secondary)
		if it.IsOk() {
			mi.IdxUpdate(it, secondary, payer)
		}
	}
}
