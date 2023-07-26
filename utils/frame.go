package utils

import (
	"encoding/binary"
	"fmt"
	"sync"
)

var crcTable []uint16
var mux sync.Mutex

type Exception uint8

type RTUFrame struct {
	Address  uint8
	Function uint8
	Data     []byte
	CRC      uint16
}
type Framer interface {
	Bytes() []byte
	Copy() Framer
	GetUnitID() uint8
	GetData() []byte
	GetFunction() uint8
	SetException(exception *Exception)
	SetData(data []byte)
}

func NewRTUFrame(packet []byte) (*RTUFrame, error) {
	// Check the that the packet length.
	if len(packet) < 5 {
		return nil, fmt.Errorf("RTU Frame error: packet less than 5 bytes: %v", packet)
	}

	// Check the CRC.
	pLen := len(packet)
	crcExpect := binary.LittleEndian.Uint16(packet[pLen-2 : pLen])
	crcCalc := crcModbus(packet[0 : pLen-2])
	if crcCalc != crcExpect {
		return nil, fmt.Errorf("RTU Frame error: CRC (expected 0x%x, got 0x%x)", crcExpect, crcCalc)
	}

	frame := &RTUFrame{
		Address:  uint8(packet[0]),
		Function: uint8(packet[1]),
		Data:     packet[2 : pLen-2],
	}

	return frame, nil
}
func (frame *RTUFrame) Copy() Framer {
	copy := *frame
	return &copy
}

// Bytes returns the Modbus byte stream based on the RTUFrame fields
// Bytes returns the Modbus byte stream based on the RTUFrame fields
func (frame *RTUFrame) Bytes() []byte {
	bytes := make([]byte, 2)

	bytes[0] = frame.Address
	bytes[1] = frame.Function
	bytes = append(bytes, frame.Data...)

	// Calculate the CRC.
	pLen := len(bytes)
	crc := crcModbus(bytes[0:pLen])

	// Add the CRC.
	bytes = append(bytes, []byte{0, 0}...)
	binary.LittleEndian.PutUint16(bytes[pLen:pLen+2], crc)

	return bytes
}

// GetUnitID returns the Modbus RTU Slave ID.
func (frame *RTUFrame) GetUnitID() uint8 {
	return frame.Address
}

// GetFunction returns the Modbus function code.
func (frame *RTUFrame) GetFunction() uint8 {
	return frame.Function
}

// GetData returns the RTUFrame Data byte field.
func (frame *RTUFrame) GetData() []byte {
	return frame.Data
}

// SetData sets the RTUFrame Data byte field and updates the frame length
// accordingly.
func (frame *RTUFrame) SetData(data []byte) {
	frame.Data = data
}

// SetException sets the Modbus exception code in the frame.
func (frame *RTUFrame) SetException(exception *Exception) {
	frame.Function = frame.Function | 0x80
	frame.Data = []byte{byte(*exception)}
}

func crcInitTable() {
	crc16IBM := uint16(0xA001)
	crcTable = make([]uint16, 256)

	for i := uint16(0); i < 256; i++ {

		crc := uint16(0)
		c := uint16(i)

		for j := uint16(0); j < 8; j++ {
			if ((crc ^ c) & 0x0001) > 0 {
				crc = (crc >> 1) ^ crc16IBM
			} else {
				crc = crc >> 1
			}
			c = c >> 1
		}
		crcTable[i] = crc
	}
}
func crcModbus(data []byte) (crc uint16) {
	if crcTable == nil {
		// Thread safe initialization.
		mux.Lock()
		if crcTable == nil {
			crcInitTable()
		}
		mux.Unlock()
	}

	crc = 0xffff
	for _, v := range data {
		crc = (crc >> 8) ^ crcTable[(crc^uint16(v))&0x00FF]
	}

	return crc
}
