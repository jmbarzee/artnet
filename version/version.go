package version

// Version 14 support for this package only
const Version ProtocolVersion = 0x1400

// ProtocolVersion is used to define the version of the packets
type ProtocolVersion uint16

// Bytes returns the ProtocolVersion high bit first.
func Bytes() []byte {
	return []byte{
		uint8(Version & 0xff),
		uint8(Version >> 8),
	}
}