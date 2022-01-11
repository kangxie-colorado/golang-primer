package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

var puzzleInput string = "60552F100693298A9EF0039D24B129BA56D67282E600A4B5857002439CE580E5E5AEF67803600D2E294B2FCE8AC489BAEF37FEACB31A678548034EA0086253B183F4F6BDDE864B13CBCFBC4C10066508E3F4B4B9965300470026E92DC2960691F7F3AB32CBE834C01A9B7A933E9D241003A520DF316647002E57C1331DFCE16A249802DA009CAD2117993CD2A253B33C8BA00277180390F60E45D30062354598AA4008641A8710FCC01492FB75004850EE5210ACEF68DE2A327B12500327D848028ED0046661A209986896041802DA0098002131621842300043E3C4168B12BCB6835C00B6033F480C493003C40080029F1400B70039808AC30024C009500208064C601674804E870025003AA400BED8024900066272D7A7F56A8FB0044B272B7C0E6F2392E3460094FAA5002512957B98717004A4779DAECC7E9188AB008B93B7B86CB5E47B2B48D7CAD3328FB76B40465243C8018F49CA561C979C182723D769642200412756271FC80460A00CC0401D8211A2270803D10A1645B947B3004A4BA55801494BC330A5BB6E28CCE60BE6012CB2A4A854A13CD34880572523898C7EDE1A9FA7EED53F1F38CD418080461B00440010A845152360803F0FA38C7798413005E4FB102D004E6492649CC017F004A448A44826AB9BFAB5E0AA8053306B0CE4D324BB2149ADDA2904028600021909E0AC7F0004221FC36826200FC3C8EB10940109DED1960CCE9A1008C731CB4FD0B8BD004872BC8C3A432BC8C3A4240231CF1C78028200F41485F100001098EB1F234900505224328612AF33A97367EA00CC4585F315073004E4C2B003530004363847889E200C45985F140C010A005565FD3F06C249F9E3BC8280804B234CA3C962E1F1C64ADED77D10C3002669A0C0109FB47D9EC58BC01391873141197DCBCEA401E2CE80D0052331E95F373798F4AF9B998802D3B64C9AB6617080"

func binStrToInt(bins string) int64 {
	var res int64 = 0
	for _, b := range bins {
		intval := int(b) - 48
		res = res*2 + int64(intval)
	}

	return res
}

func getBinaryString(input string) string {
	bins := ""

	for _, i := range input {
		intval := -1
		if int(i) >= 48 && int(i) <= 57 {
			intval = int(i) - 48
		}

		if int(i) >= 65 && int(i) <= 70 {
			intval = int(i) - 65 + 10
		}

		bins += fmt.Sprintf("%04b", intval)

	}

	return bins
}

const VerTypeLen int = 6
const LiteralUnitLen int = 5

func processLiteralPayload(payload string) (int64, int) {
	payloadLen := 0
	literal := ""
	readAll := false

	for i := 0; i+LiteralUnitLen <= len(payload) && !readAll; i += LiteralUnitLen {
		unit := payload[i : i+LiteralUnitLen]
		literal += unit[1:]
		if unit[0] == '0' {
			readAll = true
		}

		payloadLen += LiteralUnitLen
	}

	/*
		// no need to worry about the padding, it will be ignored by drive, 0xx something will not be matched to any version
			readMore := 0
			if (payloadLen+VerTypeLen)%4 != 0 {
				readMore = 4 - (payloadLen+VerTypeLen)%4
			}
			payloadLen += readMore

	*/

	return binStrToInt(literal), payloadLen
}

const LengthBitsType0 int = 15
const SubPacksBitsType1 int = 11

type decodeResult struct {
	vers         []int
	processedLen int
	literals     []int64
}

func processOperatorPayload(payload string, level int) decodeResult {
	lenType := payload[0]
	payloadLen := 0
	vers := []int{}
	result := decodeResult{}
	literals := []int64{}

	switch lenType {
	case '0':
		subpacketsLen := int(binStrToInt(payload[1 : 1+LengthBitsType0]))
		subpacketsStart := 1 + LengthBitsType0
		subpacketsEnd := subpacketsStart + subpacketsLen

		result = packetDecoder(payload[subpacketsStart:subpacketsEnd], false, level+1)
		literals = append(literals, result.literals...)
		vers = append(vers, result.vers...)

		payloadLen += 1 + LengthBitsType0 + subpacketsLen
	case '1':
		numOfPackets := binStrToInt(payload[1 : 1+SubPacksBitsType1])
		// If the length type ID is 1, then the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet.
		// shall I assume they must be literal packates: notice the "immediately contained"
		// nah.. I don't think so, it is just immediately contained beng three, sub-sub-packets can be any number..
		payloadLen += 1 + SubPacksBitsType1

		subpacksLen := 0
		for numOfPackets != 0 {
			// lets take a guess that this is indeed literal packets only
			payloadStart := 1 + SubPacksBitsType1 + subpacksLen

			result = packetDecoder(payload[payloadStart:], true, level+1)
			vers = append(vers, result.vers...)
			literals = append(literals, result.literals...)

			subpacksLen += result.processedLen

			numOfPackets--
		}

		payloadLen += subpacksLen
	default:
		panic("This cannot be right!")
	}

	return decodeResult{vers, payloadLen, literals}
}

// this func will trim the processed bits
// other func won't
func packetDecoder(packet string, onlyDecodeOne bool, level int) decodeResult {
	vers := []int{}
	processedLen := 0
	literals := []int64{}

	// first three bits are version
FORLOOP:
	for len(packet) > 6 {
		packVer := int(binStrToInt(packet[:3]))
		packType := int(binStrToInt(packet[3:6]))
		packet = packet[6:]

		fmt.Println("level:", level, "ver:", packVer, "type:", packType)

		switch packType {
		case 4:
			// literal packet
			litVal, payloadLen := processLiteralPayload(packet)
			fmt.Printf("Got a literal: %v\n", litVal)
			packet = packet[payloadLen:]
			processedLen += payloadLen + 6

			literals = append(literals, litVal)

		default:
			if len(packet) < 2 {
				break FORLOOP
			}

			// operator packet
			result := processOperatorPayload(packet, level)
			vers = append(vers, result.vers...)
			packet = packet[result.processedLen:]

			processedLen += result.processedLen + 6

			fmt.Printf("At level:%v, Func %v, result.literals: %v\n", level, GetFunctionName(typeOpMap[packType].ops), result.literals)

			val := typeOpMap[packType].ops(result.literals)
			literals = append(literals, val)

		}

		vers = append(vers, packVer)
		fmt.Printf("At level:%v, Literals rn are: %v\n", level, literals)

		if onlyDecodeOne {
			break
		}

	}

	fmt.Printf("At level:%v, Literal to return: %v\n", level, literals)
	return decodeResult{vers, processedLen, literals}

}

type opFunc func([]int64) int64
type opsAndNumOfLits struct {
	ops opFunc
	num int
}

var typeOpMap map[int]opsAndNumOfLits = map[int]opsAndNumOfLits{
	0: {sumLits, -1},
	1: {prodLits, -1},
	2: {minLits, -1},
	3: {maxLits, -1},
	5: {gtLits, 2},
	6: {ltLits, 2},
	7: {equalLits, 2},
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func getAllVersions(input string) int {
	bins := getBinaryString(input)
	fmt.Println(bins)

	result := packetDecoder(bins, false, 0)

	//fmt.Println(result.vers, result.processedLen, result.literals)

	sumVer := 0
	for _, v := range result.vers {
		sumVer += v
	}

	//fmt.Println(sumVer)

	return sumVer
}

func calculatePacket(packet string) []int64 {
	bins := getBinaryString(packet)
	fmt.Println(bins)

	result := packetDecoder(bins, false, 0)

	return result.literals
}

func sumLits(lits []int64) int64 {
	var res int64 = 0
	for _, l := range lits {
		res += l
	}

	return res
}

func prodLits(lits []int64) int64 {
	var res int64 = 1
	for _, l := range lits {
		res *= l
	}

	return res
}

func minLits(lits []int64) int64 {
	var res int64 = math.MaxInt64
	for _, l := range lits {
		if l < res {
			res = l
		}
	}

	return res
}

func maxLits(lits []int64) int64 {
	var res int64 = math.MinInt64
	for _, l := range lits {
		if l > res {
			res = l
		}
	}

	return res
}

func gtLits(lits []int64) int64 {
	if lits[0] > lits[1] {
		return 1
	}

	return 0
}

func ltLits(lits []int64) int64 {
	if lits[0] < lits[1] {
		return 1
	}

	return 0
}

func equalLits(lits []int64) int64 {
	if lits[0] == lits[1] {
		return 1
	}

	return 0
}
