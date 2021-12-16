package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const day16input string = "4057231006FF2D2E1AD8025275E4EB45A9ED518E5F1AB4363C60084953FB09E008725772E8ECAC312F0C18025400D34F732333DCC8FCEDF7CFE504802B4B00426E1A129B86846441840193007E3041483E4008541F8490D4C01A89B0DE17280472FE937C8E6ECD2F0D63B0379AC72FF8CBC9CC01F4CCBE49777098D4169DE4BF2869DE6DACC015F005C401989D0423F0002111723AC289DED3E64401004B084F074BBECE829803D3A0D3AD51BD001D586B2BEAFFE0F1CC80267F005E54D254C272950F00119264DA7E9A3E9FE6BB2C564F5376A49625534C01B0004222B41D8A80008446A8990880010A83518A12B01A48C0639A0178060059801C404F990128AE007801002803AB1801A0030A280184026AA8014C01C9B005CE0011AB00304800694BE2612E00A45C97CC3C7C4020A600433253F696A7E74B54DE46F395EC5E2009C9FF91689D6F3005AC0119AF4698E4E2713B2609C7E92F57D2CB1CE0600063925CFE736DE04625CC6A2B71050055793B4679F08CA725CDCA1F4792CCB566494D8F4C69808010494499E469C289BA7B9E2720152EC0130004320FC1D8420008647E8230726FDFED6E6A401564EBA6002FD3417350D7C28400C8C8600A5003EB22413BED673AB8EC95ED0CE5D480285C00372755E11CCFB164920070B40118DB1AE5901C0199DCD8D616CFA89009BF600880021304E0EC52100623A4648AB33EB51BCC017C0040E490A490A532F86016CA064E2B4939CEABC99F9009632FDE3AE00660200D4398CD120401F8C70DE2DB004A9296C662750663EC89C1006AF34B9A00BCFDBB4BBFCB5FBFF98980273B5BD37FCC4DF00354100762EC258C6000854158750A2072001F9338AC05A1E800535230DDE318597E61567D88C013A00C2A63D5843D80A958FBBBF5F46F2947F952D7003E5E1AC4A854400404A069802B25618E008667B7BAFEF24A9DD024F72DBAAFCB312002A9336C20CE84"

func Test_DecodExamplePacket(t *testing.T) {
	packet := DecodePacket("D2FE28")

	assert.Equal(t, "110100101111111000101000", packet.binary)
	assert.Equal(t, 6, packet.version)
	assert.Equal(t, 4, packet.typeId)
	assert.Equal(t, 2021, packet.value)
	assert.Equal(t, len("110100101111111000101"), packet.binLenght)

	packet = DecodePacket("38006F45291200")

	assert.Equal(t, "00111000000000000110111101000101001010010001001000000000", packet.binary)
	assert.Equal(t, 1, packet.version)
	assert.Equal(t, 6, packet.typeId)
	assert.Equal(t, 0, packet.lenghtTypeId)
	assert.Len(t, packet.packets, 2)
	assert.Equal(t, len("0011100000000000011011110100010100101001000100100"), packet.binLenght)

	packet = DecodePacket("EE00D40C823060")

	assert.Equal(t, "11101110000000001101010000001100100000100011000001100000", packet.binary)
	assert.Equal(t, 7, packet.version)
	assert.Equal(t, 3, packet.typeId)
	assert.Equal(t, 1, packet.lenghtTypeId)
	assert.Len(t, packet.packets, 3)
}

func getTotalPacketVersion(p Packet) int {
	sub := 0
	if len(p.packets) > 0 {
		for _, s := range p.packets {
			sub += getTotalPacketVersion(s)
		}
	}
	return sub + p.version
}

func Test_DecodeWithVersionSum(t *testing.T) {
	packet := DecodePacket("8A004A801A8002F478")
	tpv := getTotalPacketVersion(packet)
	assert.Equal(t, 16, tpv)

	packet = DecodePacket("620080001611562C8802118E34")
	tpv = getTotalPacketVersion(packet)
	assert.Equal(t, 12, tpv)

	packet = DecodePacket("C0015000016115A2E0802F182340")
	tpv = getTotalPacketVersion(packet)
	assert.Equal(t, 23, tpv)

	packet = DecodePacket("A0016C880162017C3686B18A3D4780")
	tpv = getTotalPacketVersion(packet)
	assert.Equal(t, 31, tpv)
}

func Test_DecodePacketDay16P1(t *testing.T) {
	packet := DecodePacket(day16input)
	tpv := getTotalPacketVersion(packet)
	assert.Equal(t, 996, tpv)
}

func Test_DecodeTypePackets(t *testing.T) {
	// sum
	packet := DecodePacket("C200B40A82")
	assert.Equal(t, 3, packet.value)

	// product
	packet = DecodePacket("04005AC33890")
	assert.Equal(t, 54, packet.value)

	// minimum
	packet = DecodePacket("880086C3E88112")
	assert.Equal(t, 7, packet.value)

	// max
	packet = DecodePacket("CE00C43D881120")
	assert.Equal(t, 9, packet.value)

	// gt
	packet = DecodePacket("F600BC2D8F")
	assert.Equal(t, 0, packet.value)

	// lt
	packet = DecodePacket("D8005AC2A8F0")
	assert.Equal(t, 1, packet.value)

	// eq
	packet = DecodePacket("9C005AC2F8F0")
	assert.Equal(t, 0, packet.value)

	// 1 + 3 = 2 * 2
	packet = DecodePacket("9C0141080250320F1802104A08")
	assert.Equal(t, 1, packet.value)
}

func Test_DecodePacketDay16P2(t *testing.T) {
	packet := DecodePacket(day16input)

	assert.Equal(t, 96257984154, packet.value)
}
