package main

import "fmt"

func main() {
	virtualAddress := 0b01000001101000 // 14 位

	segMask := 0x3000 // 二进制 11000000000000

	segShift := 12

	offsetMask := 0xFFF // 二进制 111111111111

	segment := (virtualAddress & segMask) >> segShift

	fmt.Printf("%b \n", virtualAddress&segMask)
	fmt.Printf("%b \n", segment)

	offset := virtualAddress & offsetMask
	fmt.Printf("%b \n", offset)

	fmt.Printf("segment: %d \noffset: %d \n", segment, offset)
}
