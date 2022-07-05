package main

import "fmt"

func char_to_symbol(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return (c - 'a') + 6
	}

	if c >= '1' && c <= '5' {
		return (c - '1') + 1
	}
	return 0
}

func StringToName(str string) uint64 {
	length := len(str)
	value := uint64(0)

	for i := 0; i <= 12; i++ {
		c := uint64(0)
		if i < length && i <= 12 {
			c = uint64(char_to_symbol(str[i]))
		}
		if i < 12 {
			c &= 0x1f
			c <<= 64 - 5*(i+1)
		} else {
			c &= 0x0f
		}

		value |= c
	}

	return value
}

func NameToString(value uint64) string {
	charmap := []byte(".12345abcdefghijklmnopqrstuvwxyz")
	// 13 dots
	str := []byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'}

	tmp := value
	for i := 0; i <= 12; i++ {
		var c byte
		if i == 0 {
			c = charmap[tmp&0x0f]
		} else {
			c = charmap[tmp&0x1f]
		}
		str[12-i] = c
		if i == 0 {
			tmp >>= 4
		} else {
			tmp >>= 5
		}
	}

	i := len(str) - 1
	for ; i >= 0; i-- {
		if str[i] != '.' {
			break
		}
	}
	return string(str[:i+1])
}

var gPackerMap = map[string]string{
	"bool":   "enc.PackBool(%s)",
	"byte":   "enc.PackUint8(uint8(%s))",
	"int8":   "enc.PackUint8(uint8(%s))",
	"uint8":  "enc.PackUint8(%s)",
	"int16":  "enc.PackInt16(%s)",
	"uint16": "enc.PackUint16(%s)",
	"int32":  "enc.PackInt32(%s)",
	"uint32": "enc.PackUint32(%s)",
	"int64":  "enc.PackInt64(%s)",
	"uint64": "enc.PackUint64(%s)",

	"chain.Int128":    "enc.WriteBytes(%s[:])",
	"chain.Uint128":   "enc.WriteBytes(%s[:])",
	"chain.VarInt32":  "enc.PackVarInt32(int32(%s))",
	"chain.VarUint32": "enc.PackVarUint32(uint32(%s))",
	"float32":         "enc.PackFloat32(%s)",
	"float64":         "enc.PackFloat64(%s)",
	"float128":        "enc.WriteBytes(%s[:])",
	"bytes":           "enc.PackBytes(%s)",
	"string":          "enc.PackString(%s)",

	"Name":               "enc.PackUint64(%s.N)",
	"TimePoint":          "enc.Pack(&%s)",
	"TimePointSec":       "enc.Pack(&%s)",
	"BlockTimestampType": "enc.Pack(&%s)",
	"Checksum160":        "enc.Pack(&%s)",
	"Checksum256":        "enc.Pack(&%s)",
	"Checksum512":        "enc.Pack(&%s)",
	"PublicKeyType":      "enc.Pack(&%s)",
	"Signature":          "enc.Pack(&%s)",
	"Symbol":             "enc.Pack(&%s)",
	"SymbolCode":         "enc.Pack(&%s)",
	"Asset":              "enc.Pack(&%s)",
	"ExtendedAsset":      "enc.Pack(&%s)",

	"chain.Name":               "enc.PackUint64(%s.N)",
	"chain.TimePoint":          "enc.Pack(&%s)",
	"chain.TimePointSec":       "enc.Pack(&%s)",
	"chain.BlockTimestampType": "enc.Pack(&%s)",
	"chain.Checksum160":        "enc.Pack(&%s)",
	"chain.Checksum256":        "enc.Pack(&%s)",
	"chain.Checksum512":        "enc.Pack(&%s)",
	"chain.PublicKeyType":      "enc.Pack(&%s)",
	"chain.Signature":          "enc.Pack(&%s)",
	"chain.Symbol":             "enc.Pack(&%s)",
	"chain.SymbolCode":         "enc.Pack(&%s)",
	"chain.Asset":              "enc.Pack(&%s)",
	"chain.ExtendedAsset":      "enc.Pack(&%s)",
}

var gUnpackMap = map[string]string{
	"bool":   "%s = dec.UnpackBool()",
	"byte":   "%s = dec.UnpackUint8()",
	"int8":   "%s = dec.UnpackInt8()",
	"uint8":  "%s = dec.UnpackUint8()",
	"int16":  "%s = dec.UnpackInt16()",
	"uint16": "%s = dec.UnpackUint16()",
	"int32":  "%s = dec.UnpackInt32()",
	"uint32": "%s = dec.UnpackUint32()",
	"int64":  "%s = dec.UnpackInt64()",
	"uint64": "%s = dec.UnpackUint64()",

	"bytes":   "%s = dec.UnpackBytes()",
	"string":  "%s = dec.UnpackString()",
	"float32": "%s = dec.UnpackFloat32()",
	"float64": "%s = dec.UnpackFloat64()",
	"[]byte":  "%s = dec.UnpackBytes()",

	"chain.Int128":    "dec.Unpack(&%s)",
	"chain.Uint128":   "dec.Unpack(&%s)",
	"chain.VarInt32":  "%s = dec.UnpackVarInt32()",
	"chain.VarUint32": "%s = dec.UnpackVarUint32()",
	"Float128":        "dec.Unpack(&%s)",

	"Name":               "%s.N = dec.UnpackUint64()",
	"TimePoint":          "dec.Unpack(&%s)",
	"TimePointSec":       "dec.Unpack(&%s)",
	"BlockTimestampType": "dec.Unpack(&%s)",
	"Checksum160":        "dec.Unpack(&%s)",
	"Checksum256":        "dec.Unpack(&%s)",
	"Checksum512":        "dec.Unpack(&%s)",
	"PublicKeyType":      "dec.Unpack(&%s)",
	"Signature":          "dec.Unpack(&%s)",
	"Symbol":             "dec.Unpack(&%s)",
	"SymbolCode":         "dec.Unpack(&%s)",
	"Asset":              "dec.Unpack(&%s)",
	"ExtendedAsset":      "dec.Unpack(&%s)",

	"chain.Name":               "%s.N = dec.UnpackUint64()",
	"chain.TimePoint":          "dec.Unpack(&%s)",
	"chain.TimePointSec":       "dec.Unpack(&%s)",
	"chain.BlockTimestampType": "dec.Unpack(&%s)",
	"chain.Checksum160":        "dec.Unpack(&%s)",
	"chain.Checksum256":        "dec.Unpack(&%s)",
	"chain.Checksum512":        "dec.Unpack(&%s)",
	"chain.PublicKeyType":      "dec.Unpack(&%s)",
	"chain.Signature":          "dec.Unpack(&%s)",
	"chain.Symbol":             "dec.Unpack(&%s)",
	"chain.SymbolCode":         "dec.Unpack(&%s)",
	"chain.Asset":              "dec.Unpack(&%s)",
	"chain.ExtendedAsset":      "dec.Unpack(&%s)",
}

func PackBasicType(member string, tp string) (string, bool) {
	if packer, ok := gPackerMap[tp]; ok {
		return fmt.Sprintf(packer, member), true
	}
	return "", false
}

func UnpackBasicType(member string, tp string) (string, bool) {
	if unpacker, ok := gUnpackMap[tp]; ok {
		return fmt.Sprintf(unpacker, member), true
	}
	return "", false
}

func abiTypes() []string {
	return []string{
		"bool",
		"int8",
		"uint8",
		"int16",
		"uint16",
		"int32",
		"uint32",
		"int64",
		"uint64",
		"int128",
		"uint128",
		"varint32",
		"varuint32",
		"float32",
		"float64",
		"float128",
		"time_point",
		"time_point_sec",
		"block_timestamp_type",
		"name",
		"bytes",
		"string",
		"checksum160",
		"checksum256",
		"checksum512",
		"public_key",
		"signature",
		"symbol",
		"symbol_code",
		"asset",
		"extended_asset",
	}
}

var gAbiTypeMap = map[string]string{
	"byte":    "uint8",
	"bool":    "bool",
	"int8":    "int8",
	"uint8":   "uint8",
	"int16":   "int16",
	"uint16":  "uint16",
	"int32":   "int32",
	"uint32":  "uint32",
	"int64":   "int64",
	"uint64":  "uint64",
	"string":  "string",
	"float32": "float32",
	"float64": "float64",

	"chain.VarInt32":           "varint32",
	"chain.VarUint32":          "varuint32",
	"chain.Int128":             "int128",
	"chain.Uint128":            "uint128",
	"chain.Float128":           "float128",
	"chain.Name":               "name",
	"chain.TimePoint":          "time_point",
	"chain.TimePointSec":       "time_point_sec",
	"chain.BlockTimestampType": "block_timestamp_type",
	"chain.Checksum160":        "checksum160",
	"chain.Checksum256":        "checksum256",
	"chain.Uint256":            "checksum256",
	"chain.Checksum512":        "checksum512",
	"chain.PublicKey":          "public_key",
	"chain.Signature":          "signature",
	"chain.Symbol":             "symbol",
	"chain.SymbolCode":         "symbol_code",
	"chain.Asset":              "asset",
	"chain.ExtendedAsset":      "extended_asset",

	"VarInt32":           "varint32",
	"VarUint32":          "varuint32",
	"Int128":             "int128",
	"Uint128":            "uint128",
	"Float128":           "float128",
	"Name":               "name",
	"TimePoint":          "time_point",
	"TimePointSec":       "time_point_sec",
	"BlockTimestampType": "block_timestamp_type",
	"Checksum160":        "checksum160",
	"Checksum256":        "checksum256",
	"Uint256":            "checksum256",
	"Checksum512":        "checksum512",
	"PublicKey":          "public_key",
	"Signature":          "signature",
	"Symbol":             "symbol",
	"SymbolCode":         "symbol_code",
	"Asset":              "asset",
	"ExtendedAsset":      "extended_asset",
}

func GoType2PrimitiveABIType(goType string) (string, bool) {
	abiType, ok := gAbiTypeMap[goType]
	if ok {
		return abiType, true
	} else {
		return "", false
	}
}

var gGetSizeMap = map[string]string{
	"byte":    "size += 1",
	"bool":    "size += 1",
	"int8":    "size += 1",
	"uint8":   "size += 1",
	"int16":   "size += 2",
	"uint16":  "size += 2",
	"int":     "size += 4",
	"int32":   "size += 4",
	"uint32":  "size += 4",
	"int64":   "size += 8",
	"uint64":  "size += 8",
	"float32": "size += 4",
	"float64": "size += 8",

	"chain.Int128":  "size += 16",
	"chain.Uint128": "size += 16",
	"chain.Uint256": "size += 32",
	"chain.Name":    "size += 8",
	"chain.Symbol":  "size += 8",

	"Int128":  "size += 16",
	"Uint128": "size += 16",
	"Uint256": "size += 32",
	"Name":    "size += 8",
	"Symbol":  "size += 8",
}

func calcNotArrayMemberSize(name string, goType string) string {
	var code string

	if getSizeCode, ok := gGetSizeMap[goType]; ok {
		code = getSizeCode
	} else {
		if goType == "string" {
			code = fmt.Sprintf("size += chain.PackedVarUint32Length(uint32(len(t.%s))) + len(t.%s)", name, name)
		} else {
			// "chain.Signature":
			// "chain.PublicKey":
			code = fmt.Sprintf("size += t.%[1]s.Size()", name)
		}
	}

	return code + " //" + name
}

var gGetArraySizeMap = map[string]string{
	"byte":          "size += len(%s)",
	"bool":          "size += len(%s)",
	"int8":          "size += len(%s)",
	"uint8":         "size += len(%s)",
	"int16":         "size += len(%s)*2",
	"uint16":        "size += len(%s)*2",
	"int":           "size += len(%s)*4",
	"int32":         "size += len(%s)*4",
	"uint32":        "size += len(%s)*4",
	"int64":         "size += len(%s)*8",
	"uint64":        "size += len(%s)*8",
	"chain.Uint128": "size += len(%s)*16",
	"chain.Uint256": "size += len(%s)*32",
	"float32":       "size += len(%s)*4",
	"float64":       "size += len(%s)*8",
	"chain.Name":    "size += len(%s)*8",
}

func calcArrayMemberSize(name string, goType string) string {
	if getArraySizeCode, ok := gGetArraySizeMap[goType]; ok {
		return fmt.Sprintf(getArraySizeCode, "t."+name)
	}
	switch goType {
	case "[]byte":
		return fmt.Sprintf(`for i := range t.%[1]s {
	size += chain.PackedVarUint32Length(uint32(len(t.%[1]s[i]))) + len(t.%[1]s[i])
}`, name)
	case "string":
		return fmt.Sprintf(`for i := range t.%[1]s {
	 size += chain.PackedVarUint32Length(uint32(len(t.%[1]s[i]))) + len(t.%[1]s[i])
}`, name)
	default:
		return fmt.Sprintf(`
    for i := range t.%[1]s {
        size += t.%[1]s[i].Size()
    }`, name)
	}
}

func GetIndexType(index string) string {
	switch index {
	case "IDX64":
		return "uint64"
	case "IDX128":
		return "chain.Uint128"
	case "IDX256":
		return "chain.Uint256"
	case "IDXFloat64":
		return "float64"
	case "IDXFloat128":
		return "chain.Float128"
	default:
		panic(fmt.Sprintf("unknown secondary index type: %s", index))
	}
}

func indexTypeToSecondaryType(indexType string) string {
	switch indexType {
	case "IDX64":
		return "uint64"
	case "IDX128":
		return "chain.Uint128"
	case "IDX256":
		return "chain.Uint256"
	case "IDXFloat64":
		return "float64"
	case "IDXFloat128":
		return "chain.Float128"
	default:
		panic(fmt.Sprintf("unknown secondary index type: %s", indexType))
	}
	return ""
}

func indexTypeToSecondaryTableName(indexType string) string {
	switch indexType {
	case "IDX64":
		return "IdxTable64"
	case "IDX128":
		return "IdxTable128"
	case "IDX256":
		return "IdxTable256"
	case "IDXFloat64":
		return "IdxTableFloat64"
	case "IDXFloat128":
		return "IdxTableFloat128"
	default:
		panic(fmt.Sprintf("unknown secondary index type: %s", indexType))
	}
	return ""
}
