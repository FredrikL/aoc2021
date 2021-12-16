package aoc

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type Packet struct {
	version      int
	typeId       int
	lenghtTypeId int
	value        int
	binary       string
	binLenght    int
	packets      []Packet
}

func toBinary(input []byte) string {
	arr := make([]string, len(input))

	for i, b := range input {
		arr[i] = fmt.Sprintf("%08b", b)
	}

	return strings.Join(arr, "")
}

func DecodePacket(data string) Packet {
	b, _ := hex.DecodeString(data)
	bin := toBinary(b)

	return DecodePacketBin(bin, 0)
}

func DecodePacketBin(bin string, pos int) Packet {
	version, typeId := DecodePacketHeader(bin, pos)
	packet := Packet{
		version: version,
		typeId:  typeId,
		binary:  bin,
	}

	if typeId == 4 {
		endPos, value := DecodeLiteralValuePacket(bin, pos+6)
		packet.value = value
		packet.binLenght = endPos - pos
	} else if typeId == 0 {
		lenghtTypeId, endPos, packets := DecodeOperatorPacket(bin, pos+6)
		packet.lenghtTypeId = lenghtTypeId
		packet.packets = packets
		packet.binLenght = endPos - pos
		value := 0
		for _, p := range packets {
			value += p.value
		}
		packet.value = value
	} else if typeId == 1 {
		lenghtTypeId, endPos, packets := DecodeOperatorPacket(bin, pos+6)
		packet.lenghtTypeId = lenghtTypeId
		packet.packets = packets
		packet.binLenght = endPos - pos
		value := 1
		for _, p := range packets {
			value *= p.value
		}
		packet.value = value
	} else if typeId == 2 {
		lenghtTypeId, endPos, packets := DecodeOperatorPacket(bin, pos+6)
		packet.lenghtTypeId = lenghtTypeId
		packet.packets = packets
		packet.binLenght = endPos - pos
		value := packets[0].value
		for _, p := range packets {
			if p.value < value {
				value = p.value
			}
		}
		packet.value = value
	} else if typeId == 3 {
		lenghtTypeId, endPos, packets := DecodeOperatorPacket(bin, pos+6)
		packet.lenghtTypeId = lenghtTypeId
		packet.packets = packets
		packet.binLenght = endPos - pos
		value := packets[0].value
		for _, p := range packets {
			if p.value > value {
				value = p.value
			}
		}
		packet.value = value
	} else if typeId == 5 {
		lenghtTypeId, endPos, packets := DecodeOperatorPacket(bin, pos+6)
		packet.lenghtTypeId = lenghtTypeId
		packet.packets = packets
		packet.binLenght = endPos - pos
		if packets[0].value > packets[1].value {
			packet.value = 1
		} else {
			packet.value = 0
		}
	} else if typeId == 6 {
		lenghtTypeId, endPos, packets := DecodeOperatorPacket(bin, pos+6)
		packet.lenghtTypeId = lenghtTypeId
		packet.packets = packets
		packet.binLenght = endPos - pos
		if packets[0].value < packets[1].value {
			packet.value = 1
		} else {
			packet.value = 0
		}
	} else if typeId == 7 {
		lenghtTypeId, endPos, packets := DecodeOperatorPacket(bin, pos+6)
		packet.lenghtTypeId = lenghtTypeId
		packet.packets = packets
		packet.binLenght = endPos - pos
		if packets[0].value == packets[1].value {
			packet.value = 1
		} else {
			packet.value = 0
		}
	} else {
		lenghtTypeId, endPos, packets := DecodeOperatorPacket(bin, pos+6)
		packet.lenghtTypeId = lenghtTypeId
		packet.packets = packets
		packet.binLenght = endPos - pos
	}

	return packet
}

func DecodeOperatorPacket(bin string, pos int) (int, int, []Packet) {
	packets := []Packet{}
	lenghtTypeId := bin[pos : pos+1]
	l, _ := strconv.ParseInt(lenghtTypeId, 2, 64)

	if l == 0 {
		subPacketSizeStr := bin[pos+1 : pos+16]
		s, _ := strconv.ParseInt(subPacketSizeStr, 2, 64)
		subPacketSize := int(s)
		read := 0
		pos = pos + 16
		for {
			p := DecodePacketBin(bin, pos)
			read += p.binLenght
			pos += p.binLenght
			packets = append(packets, p)
			if read >= subPacketSize {
				break
			}
		}
	}

	if l == 1 {
		subPacketCountStr := bin[pos+1 : pos+12]
		s, _ := strconv.ParseInt(subPacketCountStr, 2, 64)
		subPacketCount := int(s)
		count := 0
		pos = pos + 12
		for {
			p := DecodePacketBin(bin, pos)
			count += 1
			pos += p.binLenght
			packets = append(packets, p)
			if count >= subPacketCount {
				break
			}
		}
	}

	return int(l), pos, packets
}

func DecodeLiteralValuePacket(bin string, pos int) (int, int) {
	endPos, groups := readGroups(bin, pos, 5)

	value_str := strings.Join(groups, "")
	value, _ := strconv.ParseInt(value_str, 2, 64)
	return endPos, int(value)
}

func DecodePacketHeader(bin string, start int) (version, typeId int) {
	bin_version := bin[start : start+3]
	v, _ := strconv.ParseInt(bin_version, 2, 64)
	bin_type := bin[start+3 : start+6]
	t, _ := strconv.ParseInt(bin_type, 2, 64)

	version = int(v)
	typeId = int(t)
	return
}

func readGroups(str string, start, size int) (int, []string) {
	result := []string{}

	pos := start
	for {
		slice := str[pos : (pos)+size]
		result = append(result, slice[1:])
		pos = pos + size
		if slice[0] == '0' {
			break
		}
	}

	return pos, result
}
