package day16

import (
	"bufio"
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
)

type LiteralPacket struct {
	val int
}

type OperatorPacket struct {
	lengthTypeId int
	packets      []Packet
}

type TypeId int

const (
	TYPE_SUM TypeId = iota
	TYPE_PRODUCT
	TYPE_MIN
	TYPE_MAX
	TYPE_LITERAL
	TYPE_GT
	TYPE_LT
	TYPE_EQ
)

type Packet struct {
	version int
	typeId  TypeId
	LiteralPacket
	OperatorPacket
}

func newPacket(binary Binary) (packet *Packet, tail Binary, err error) {
	// first 3 bits = version
	head, tail := binary.splitAt(3)
	version, err := head.toInt()
	if err != nil {
		return nil, tail, err
	}

	// next 3 bits = type id
	head, tail = tail.splitAt(3)
	typeId, err := head.toInt()
	if err != nil {
		return nil, tail, err
	}

	packet = &Packet{
		version: version,
		typeId:  TypeId(typeId),
	}

	// typeId 4 is literal value;
	if packet.typeId == TYPE_LITERAL {
		tail, err = packet.parseLiteralValue(tail)
	} else {
		// every other typeId is an "operator"
		tail, err = packet.parseOperator(tail)
	}

	return packet, tail, err
}

func (p *Packet) parseLiteralValue(s Binary) (tail Binary, err error) {
	chunkSize := 5

	// get literal value from remaining bits
	var valueBin Binary
	head, tail := s.splitAt(chunkSize)

	for len(head) == 5 {
		// strip off the leading bit
		first, rest := head.splitAt(1)
		valueBin += rest

		// 0 means it is the last group
		if first == "0" {
			break
		}

		head, tail = tail.splitAt(5)
	}

	value, err := valueBin.toInt()

	if err != nil {
		return tail, err
	}

	p.val = value

	return tail, nil
}

func (p *Packet) parseOperator(s Binary) (tail Binary, err error) {
	head, tail := s.splitAt(1)

	if head == "0" {
		p.lengthTypeId = 0
		tail, err = p.parseOperatorBitCount(tail)
	} else {
		p.lengthTypeId = 1
		tail, err = p.parseOperatorPacketCount(tail)
	}

	return tail, err
}

func (p *Packet) parseOperatorBitCount(s Binary) (tail Binary, err error) {
	head, tail := s.splitAt(15)
	bitCount, err := head.toInt()
	if err != nil {
		return tail, err
	}

	subpackets, tail := tail.splitAt(bitCount)

	for subpackets != Binary(strings.Repeat("0", len(subpackets))) {
		subpacket, nextPackets, err := newPacket(subpackets)

		if err != nil {
			return tail, nil
		}

		p.packets = append(p.packets, *subpacket)

		subpackets = nextPackets
	}

	return tail, nil
}

func (p *Packet) parseOperatorPacketCount(binary Binary) (tail Binary, err error) {
	// 11 bits
	head, tail := binary.splitAt(11)

	packetCount, err := head.toInt()

	if err != nil {
		return binary, err
	}

	for packetCount > 0 {
		packetCount--

		subpacket, nextTail, err := newPacket(tail)

		if err != nil {
			return tail, nil
		}

		p.packets = append(p.packets, *subpacket)

		tail = nextTail
	}

	return tail, err
}

func (p *Packet) versionSum() (sum int) {
	sum += p.version

	// recursively check nested packets
	for _, subpacket := range p.packets {
		sum += subpacket.versionSum()
	}

	return
}

func(p *Packet) evaluateExpression() int {
	if p.typeId == TYPE_LITERAL {
		return p.val
	}

	values := []int{}

	// recursively check nested packets
	for _, subpacket := range p.packets {
		values = append(values, subpacket.evaluateExpression())
	}

	switch p.typeId {
	case TYPE_SUM:
		sum := 0
		for _, val := range values {
			sum += val
		}
		return sum
	case TYPE_PRODUCT:
		product := 1

		for _, val := range values {
			product *= val
		}

		return product
	case TYPE_MIN:
		min := math.MaxInt
		for _, val := range values {
			if val < min {
				min = val
			}
		}
		return min
	case TYPE_MAX:
		max := math.MinInt
		for _, val := range values {
			if val > max {
				max = val
			}
		}
		return max
	case TYPE_GT:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case TYPE_LT:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case TYPE_EQ:
		if values[0] == values[1] {
			return 1
		}
		return 0
	}

	return 0
}

type Binary string

func (binary Binary) splitAt(index int) (Binary, Binary) {
	return binary[:index], binary[index:]
}

func (binary Binary) toInt() (int, error) {
	val, err := strconv.ParseInt(string(binary), 2, 64)
	return int(val), err
}

func hexToBinaryString(s string) (Binary, error) {
	hexMapping := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
	val := make([]string, len(s))
	for i, v := range s {
		va, ok := hexMapping[v]
		if !ok {
			return "", errors.New("invalid hex")
		}
		val[i] = va
	}
	return Binary(strings.Join(val, "")), nil
}

func getScanner(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	return scanner, nil
}
