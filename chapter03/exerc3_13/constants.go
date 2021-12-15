package main

import (
	"fmt"
	"math/big"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	var ZBs string = fmt.Sprint(EB) + "000"
	var YBs string = ZBs + "000"
	KiBBig := new(big.Int).SetInt64(1024)
	EiBBig := new(big.Int).SetInt64(EiB)
	ZiBBig := new(big.Int).Mul(EiBBig, KiBBig)
	YiBBig := new(big.Int).Mul(ZiBBig, KiBBig)

	fmt.Println("1KB =", KB, " bytes", "\t\t\t1KiB =", KiB, " bytes")
	fmt.Println("1MB =", MB, " bytes", "\t\t\t1MiB =", MiB, " bytes")
	fmt.Println("1GB =", GB, " bytes", "\t\t1GiB =", GiB, " bytes")
	fmt.Println("1TB =", TB, " bytes", "\t\t1TiB =", TiB, " bytes")
	fmt.Println("1PB =", PB, " bytes", "\t\t1PiB =", PiB, " bytes")
	fmt.Println("1EB =", EB, " bytes", "\t1EiB =", EiB, " bytes")
	fmt.Println("1ZB =", ZBs, " bytes", "\t1ZiB =", ZiBBig, "bytes")
	fmt.Println("1YB =", YBs, " bytes", "\t1YiB =", YiBBig, "bytes")
}
