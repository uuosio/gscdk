package main

const cTableTemplate = `
type {{.StructName}}Table struct {
	database.MultiIndexInterface
}

func (mi *{{.StructName}}Table) Store(v *{{.StructName}}, payer chain.Name) {
	mi.MultiIndexInterface.Store(v, payer)
}

func (mi *{{.StructName}}Table) GetByKey(id uint64) (*database.Iterator, *{{.StructName}}) {
	it, data := mi.MultiIndexInterface.GetByKey(id)
	if !it.IsOk() {
		return it, nil
	}
	return it, data.(*{{.StructName}})
}

func (mi *{{.StructName}}Table) GetByIterator(it *database.Iterator) *{{.StructName}} {
	data := mi.MultiIndexInterface.GetByIterator(it)
	return data.(*{{.StructName}})
}

func (mi *{{.StructName}}Table) Update(it *database.Iterator, v *{{.StructName}}, payer chain.Name) {
	mi.MultiIndexInterface.Update(it, v, payer)
}
`

const cNewMultiIndexTemplate = `
func New{{.Name}}Table(code chain.Name, scope chain.Name) *{{.Name}}Table {
	table := chain.Name{N:uint64({{.TableName}})} //table name: {{.Name}}
	if table.N&uint64(0x0f) != 0 {
		// Limit table names to 12 characters so that the last character (4 bits) can be used to distinguish between the secondary indices.
		panic("NewMultiIndex:Invalid multi-index table name ")
	}

	mi := &database.MultiIndex{}
	mi.SetTable(code, scope, table)
	mi.Table = database.NewTableI64(code, scope, table, func(data []byte) database.TableValue {
		return mi.Unpack(data)
	})
	mi.IdxTableNameToIndex = {{.Name}}TableNameToIndex
	mi.IndexTypes = {{.Name}}SecondaryTypes
	mi.IDXTables = make([]database.SecondaryTable, len({{.Name}}SecondaryTypes))
	mi.Unpack = {{.Name}}Unpacker

{{- range $i, $val := .Indexes}}
	mi.IDXTables[{{$i}}] = database.New{{$val.TableType}}({{$i}}, code.N, scope.N, uint64({{$.FirstIdxTableName}})+{{$i}})
{{- end}}
	return &{{.Name}}Table{mi}
}

{{- range $i, $val := .Indexes}}
func (mi *{{$.Name}}Table) GetIdxTableBy{{$val.Name}}() *database.{{$val.TableType}} {
	return mi.GetIdxTableByIndex({{$i}}).(*database.{{$val.TableType}})
}
{{- end}}
`

const cDummyCode = `
//eliminate unused package errors
func dummy() {
	if false {
		v := 0;
		n := unsafe.Sizeof(v);
		chain.Printui(uint64(n));
		chain.Printui(database.IDX64);
	}
}`

const cMainCode = `
func main() {
	receiver, firstReceiver, action := chain.GetApplyArgs()
	contract := NewContract(receiver, firstReceiver, action)
	if contract == nil {
		return
	}
	data := chain.ReadActionData()
	
	//Fix data declared but not used error
	if false {
		println(len(data))
	}
`

const cSingletonCode = `
func (d *{{.Name}}) GetPrimary() uint64 {
	return uint64({{.TableName}})
}

type {{.Name}}Table struct {
	db *database.SingletonTable
}

func New{{.Name}}Table(code chain.Name, scope chain.Name) *{{.Name}}Table {
	chain.Check(code != chain.Name{0}, "bad code name")
	table := chain.Name{N:uint64({{.TableName}})}
	db := database.NewSingletonTable(code, scope, table, {{.Name}}Unpacker)
	return &{{.Name}}Table{db}
}

func (t *{{.Name}}Table) Set(data *{{.Name}}, payer chain.Name) {
	t.db.Set(data, payer)
}

func (t *{{.Name}}Table) Get() (*{{.Name}}) {
	data := t.db.Get()
	if data == nil {
		return nil
	}
	return data.(*{{.Name}})
}

func (t *{{.Name}}Table) Remove() {
	t.db.Remove()
}
`

const cImportCode = `package main
import (
	"github.com/uuosio/chain"
    "github.com/uuosio/chain/database"
    "unsafe"
)
`

const cExtensionTemplate = `
func (t *%[1]s) Pack() []byte {
	if !t.HasValue {
		return []byte{}
	}
	return t.%[2]s.Pack()
}

func (t *%[1]s) Unpack(data []byte) int {
	if len(data) == 0 {
		t.HasValue = false
		return 0
	} else {
		t.HasValue = true
	}

	dec := chain.NewDecoder(data)
	dec.Unpack(&t.%[2]s)
	return dec.Pos()
}

func (t *%[1]s) Size() int {
	return t.%[2]s.Size()
}
`

const cOptionalTemplate = `
func (t *%[1]s) Pack() []byte {
	if !t.IsValid {
		return []byte{0}
	}
	buf := make([]byte, 0, t.Size()+1)
	buf = append(buf, 1)
	buf = append(buf, t.%[2]s.Pack()...) //TODO: handle pack for different type
	return buf
}

func (t *%[1]s) Unpack(data []byte) int {
	chain.Check(len(data) >= 1, "invalid data size")
	valid := data[1]
	if valid == 0 {
		t.IsValid = false
	} else if valid == 1 {
		t.IsValid = true
	} else {
		chain.Check(false, "invalid optional value")
	}

	dec := chain.NewDecoder(data[1:])
	dec.Unpack(&t.%[2]s) //TODO: handle unpack for different type
	return dec.Pos() + 1
}

func (t *%[1]s) Size() int {
	return t.%[2]s.Size() + 1 //TODO: calculate size for different type
}
`

const cContractTemplate = `
package main

import (
	"github.com/uuosio/chain"
)

//contract %[1]s
type Contract struct {
	self, firstReceiver, action chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *Contract {
	return &Contract{receiver, firstReceiver, action}
}

//action sayhello
func (c *Contract) SayHello(name string) {
	chain.Println("Hello, ", name)
}

type MyOptional struct {
	chain.Optional
	value string
}

//example of optional abi type
//action testoptional
func (c *Contract) testoptional(opt *MyOptional) {
	if opt.IsValid {
		chain.Println(opt.value)
	}
}

type MyExtension struct {
	chain.BinaryExtension
	value string
}

//example of binary_extension abi type
//action testext
func (c *Contract) testext(ext *MyExtension) {
	if ext.HasValue {
		chain.Println(ext.value)
	}
}
`

const cContractCode = `
package main
import (
	"github.com/uuosio/chain"
)

//contract %[1]s
type Contract struct {
	self, firstReceiver, action chain.Name
}

func NewContract(receiver, firstReceiver, action chain.Name) *Contract {
	return &Contract{}
}

//action sayhello
func (c *Contract) SayHello(name string) {
	chain.Println("Hello, ", name)
}

type MyOptional struct {
	chain.Optional
	value string
}

//example of optional abi type
//action testoptional
func (c *Contract) testoptional(opt *MyOptional) {
	if opt.IsValid {
		chain.Println(value)
	}
}

type MyExtension struct {
	chain.BinaryExtension
	value string
}

//example of binary_extension abi type
//action testoptional
func (c *Contract) testoptional(ext *MyExtension) {
	if ext.HasValue {
		chain.Println(ext.value)
	}
}
`

const cUtils = `
package main

import (
	"github.com/uuosio/chain"
)

func check(b bool, msg string) {
	chain.Check(b, msg)
}
`

const cTables = `
package main

import (
	"github.com/uuosio/chain"
)

//table mytable
type MyData struct {
	primary uint64 		//primary : t.primary
	a1 uint64         	//IDX64 		: Bya1 : t.a1 : t.a1
	a2 chain.Uint128  	//IDX128 		: Bya2 : t.a2 : t.a2
	a3 chain.Uint256  	//IDX256 		: Bya3 : t.a3 : t.a3
	a4 float64        	//IDXFloat64 	: Bya4 : t.a4 : t.a4
	a5 chain.Float128 	//IDXFloat128 	: Bya5 : t.a5 : t.a5
}
`

const cStructs = `
package main

type MyStruct struct {
	a uint64
	b uint64
}
`

const cBuild = `
eosio-go build -o %[1]s.wasm .
`

const cTestScript = `
import os
import sys
try:
	from pyeoskit import eosapi, wallet
except:
	print('pyeoskit not found, please install it with "pip install pyeoskit"')
	sys.exit(-1)
from pyeoskit.exceptions import ChainException

# modify your test account here
test_account1 = 'helloworld11'
# modify your test account private key here
wallet.import_key('test', '5JRYimgLBrRLCBAcjHUWCYRv3asNedTYYzVgmiU4q2ZVxMBiJXL')
# modify test node here
eosapi.set_node('https://testnode.uuos.network:8443')

with open('%[1]s.wasm', 'rb') as f:
    code = f.read()
with open('%[1]s.abi', 'rb') as f:
    abi = f.read()

try:
    eosapi.deploy_contract(test_account1, code, abi, vm_type=0)
except ChainException as e:
    if not e.json['error']['details'][0]['message'] == 'contract is already running this version of code':
        raise e

r = eosapi.push_action(test_account1, 'sayhello', {'name': 'alice'})
print(r['processed']['action_traces'][0]['console'])
`

const cReadMe = "# Building\n\n```bash\neosio-go build -o %[1]s.wasm .\n```\n\n# Testing\n```\npython3 test.py\n```"

const cSecondaryValueTemplate = `
var (
	{{.StructInfo.StructName}}SecondaryTypes = []int{
	{{- range $i, $val := .SecondaryIndexes}}
		database.{{$val.Type}},
	{{- end}}
	}
)

func {{.StructInfo.StructName}}TableNameToIndex(indexName string) int {
	switch indexName {
	{{- range $i, $val := .SecondaryIndexes}}
		case "{{$val.Name}}":
			return {{$i}}
	{{- end}}
	default:
		panic("unknow indexName")
	}
}

func {{.StructInfo.StructName}}Unpacker(buf []byte) database.MultiIndexValue {
	v := &{{.StructInfo.StructName}}{}
	v.Unpack(buf)
	return v
}

func (t *{{.StructInfo.StructName}}) GetSecondaryValue(index int) interface{} {
	switch index {
	{{- range $i, $val := .SecondaryIndexes}}
		case {{$i}}:
			return {{$val.Getter}}
	{{- end}}
		default:
			panic("index out of bound")
	}
}

func (t *{{.StructInfo.StructName}}) SetSecondaryValue(index int, v interface{}) {
	switch index {
		{{- range $i, $val := .SecondaryIndexes}}
	case {{$i}}:
		{{$val.GetSetter}}
{{- end}}
	default:
		panic("unknown index")
	}
}
`

const cSerializerTemplate = `
func (t *{{.StructName}}) Pack() []byte {
    enc := chain.NewEncoder(t.Size())
	{{- range $i, $member := .Members}}
	{{$member.PackMember}}
	{{- end}}
    return enc.GetBytes()
}

func (t *{{.StructName}}) Unpack(data []byte) int {
    dec := chain.NewDecoder(data)
	{{- range $i, $member := .Members}}
	{{$member.UnpackMember}}
	{{- end}}
    return dec.Pos()
}

func (t *{{.StructName}}) Size() int {
    size := 0
	{{- range $i, $member := .Members}}
	{{$member.GetSize}}
	{{- end}}
    return size
}
`

const cVariantTemplate = `
func New{{.StructName}}(value interface{}) *{{.StructName}} {
	ret := &{{.StructName}}{}
	switch value.(type) {
		{{- range $i, $member := .Members}}
		case *{{$member.Type}}:
				ret.value = value
		{{- end}}
		default:
			chain.Check(false, "unknown variant type")	
	}
	return ret
}

func (t *{{.StructName}}) Pack() []byte {
    enc := chain.NewEncoder(t.Size())
	{{- range $i, $member := .Members}}
	if _, ok := t.value.(*{{$member.Type}}); ok {
		enc.PackUint8(uint8({{$i}}))
		{{$member.PackVariantMember}}
		return enc.GetBytes()
	}
	{{- end}}
    return enc.GetBytes()
}

func (t *{{.StructName}}) Unpack(data []byte) int {
    dec := chain.NewDecoder(data)
	tp := dec.UnpackUint8()
	{{- range $i, $member := .Members}}
	if tp == uint8({{$i}}) {
		type A struct {
			value {{$member.Type}}
		}
		a := &A{}
		dec.Unpack(&a.value)
		t.value = &a.value
	}
	{{- end}}
    return dec.Pos()
}

func (t *{{.StructName}}) Size() int {
    size := 1
	{{- range $i, $member := .Members}}
	if _, ok := t.value.({{$member.Type}}); ok {
		{{$member.GetVariantSize}}
		return size
	}
	{{- end}}
    return size
}
`
