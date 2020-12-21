package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Block) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ht":
			z.Height, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		case "bh":
			err = dc.ReadExactBytes((z.BlockHash)[:])
			if err != nil {
				err = msgp.WrapError(err, "BlockHash")
				return
			}
		case "bi":
			z.BlockInfo, err = dc.ReadBytes(z.BlockInfo)
			if err != nil {
				err = msgp.WrapError(err, "BlockInfo")
				return
			}
		case "tx":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "TxList")
				return
			}
			if cap(z.TxList) >= int(zb0002) {
				z.TxList = (z.TxList)[:zb0002]
			} else {
				z.TxList = make([]Tx, zb0002)
			}
			for za0002 := range z.TxList {
				err = z.TxList[za0002].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "TxList", za0002)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Block) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "ht"
	err = en.Append(0x84, 0xa2, 0x68, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Height)
	if err != nil {
		err = msgp.WrapError(err, "Height")
		return
	}
	// write "bh"
	err = en.Append(0xa2, 0x62, 0x68)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.BlockHash)[:])
	if err != nil {
		err = msgp.WrapError(err, "BlockHash")
		return
	}
	// write "bi"
	err = en.Append(0xa2, 0x62, 0x69)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.BlockInfo)
	if err != nil {
		err = msgp.WrapError(err, "BlockInfo")
		return
	}
	// write "tx"
	err = en.Append(0xa2, 0x74, 0x78)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.TxList)))
	if err != nil {
		err = msgp.WrapError(err, "TxList")
		return
	}
	for za0002 := range z.TxList {
		err = z.TxList[za0002].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "TxList", za0002)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Block) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "ht"
	o = append(o, 0x84, 0xa2, 0x68, 0x74)
	o = msgp.AppendInt64(o, z.Height)
	// string "bh"
	o = append(o, 0xa2, 0x62, 0x68)
	o = msgp.AppendBytes(o, (z.BlockHash)[:])
	// string "bi"
	o = append(o, 0xa2, 0x62, 0x69)
	o = msgp.AppendBytes(o, z.BlockInfo)
	// string "tx"
	o = append(o, 0xa2, 0x74, 0x78)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TxList)))
	for za0002 := range z.TxList {
		o, err = z.TxList[za0002].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "TxList", za0002)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Block) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ht":
			z.Height, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		case "bh":
			bts, err = msgp.ReadExactBytes(bts, (z.BlockHash)[:])
			if err != nil {
				err = msgp.WrapError(err, "BlockHash")
				return
			}
		case "bi":
			z.BlockInfo, bts, err = msgp.ReadBytesBytes(bts, z.BlockInfo)
			if err != nil {
				err = msgp.WrapError(err, "BlockInfo")
				return
			}
		case "tx":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxList")
				return
			}
			if cap(z.TxList) >= int(zb0002) {
				z.TxList = (z.TxList)[:zb0002]
			} else {
				z.TxList = make([]Tx, zb0002)
			}
			for za0002 := range z.TxList {
				bts, err = z.TxList[za0002].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "TxList", za0002)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Block) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 3 + msgp.ArrayHeaderSize + (32 * (msgp.ByteSize)) + 3 + msgp.BytesPrefixSize + len(z.BlockInfo) + 3 + msgp.ArrayHeaderSize
	for za0002 := range z.TxList {
		s += z.TxList[za0002].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *BlockIndex) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ht":
			z.Height, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		case "bh":
			z.BlockHash48, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "BlockHash48")
				return
			}
		case "tx":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "TxHash48List")
				return
			}
			if cap(z.TxHash48List) >= int(zb0002) {
				z.TxHash48List = (z.TxHash48List)[:zb0002]
			} else {
				z.TxHash48List = make([]uint64, zb0002)
			}
			for za0001 := range z.TxHash48List {
				z.TxHash48List[za0001], err = dc.ReadUint64()
				if err != nil {
					err = msgp.WrapError(err, "TxHash48List", za0001)
					return
				}
			}
		case "ai":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "AddrHashes")
				return
			}
			if cap(z.AddrHashes) >= int(zb0003) {
				z.AddrHashes = (z.AddrHashes)[:zb0003]
			} else {
				z.AddrHashes = make([]uint64, zb0003)
			}
			for za0002 := range z.AddrHashes {
				z.AddrHashes[za0002], err = dc.ReadUint64()
				if err != nil {
					err = msgp.WrapError(err, "AddrHashes", za0002)
					return
				}
			}
		case "ap":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "AddrPosLists")
				return
			}
			if cap(z.AddrPosLists) >= int(zb0004) {
				z.AddrPosLists = (z.AddrPosLists)[:zb0004]
			} else {
				z.AddrPosLists = make([][]uint32, zb0004)
			}
			for za0003 := range z.AddrPosLists {
				var zb0005 uint32
				zb0005, err = dc.ReadArrayHeader()
				if err != nil {
					err = msgp.WrapError(err, "AddrPosLists", za0003)
					return
				}
				if cap(z.AddrPosLists[za0003]) >= int(zb0005) {
					z.AddrPosLists[za0003] = (z.AddrPosLists[za0003])[:zb0005]
				} else {
					z.AddrPosLists[za0003] = make([]uint32, zb0005)
				}
				for za0004 := range z.AddrPosLists[za0003] {
					z.AddrPosLists[za0003][za0004], err = dc.ReadUint32()
					if err != nil {
						err = msgp.WrapError(err, "AddrPosLists", za0003, za0004)
						return
					}
				}
			}
		case "ti":
			var zb0006 uint32
			zb0006, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "TopicHashes")
				return
			}
			if cap(z.TopicHashes) >= int(zb0006) {
				z.TopicHashes = (z.TopicHashes)[:zb0006]
			} else {
				z.TopicHashes = make([]uint64, zb0006)
			}
			for za0005 := range z.TopicHashes {
				z.TopicHashes[za0005], err = dc.ReadUint64()
				if err != nil {
					err = msgp.WrapError(err, "TopicHashes", za0005)
					return
				}
			}
		case "tp":
			var zb0007 uint32
			zb0007, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "TopicPosLists")
				return
			}
			if cap(z.TopicPosLists) >= int(zb0007) {
				z.TopicPosLists = (z.TopicPosLists)[:zb0007]
			} else {
				z.TopicPosLists = make([][]uint32, zb0007)
			}
			for za0006 := range z.TopicPosLists {
				var zb0008 uint32
				zb0008, err = dc.ReadArrayHeader()
				if err != nil {
					err = msgp.WrapError(err, "TopicPosLists", za0006)
					return
				}
				if cap(z.TopicPosLists[za0006]) >= int(zb0008) {
					z.TopicPosLists[za0006] = (z.TopicPosLists[za0006])[:zb0008]
				} else {
					z.TopicPosLists[za0006] = make([]uint32, zb0008)
				}
				for za0007 := range z.TopicPosLists[za0006] {
					z.TopicPosLists[za0006][za0007], err = dc.ReadUint32()
					if err != nil {
						err = msgp.WrapError(err, "TopicPosLists", za0006, za0007)
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *BlockIndex) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "ht"
	err = en.Append(0x87, 0xa2, 0x68, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Height)
	if err != nil {
		err = msgp.WrapError(err, "Height")
		return
	}
	// write "bh"
	err = en.Append(0xa2, 0x62, 0x68)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.BlockHash48)
	if err != nil {
		err = msgp.WrapError(err, "BlockHash48")
		return
	}
	// write "tx"
	err = en.Append(0xa2, 0x74, 0x78)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.TxHash48List)))
	if err != nil {
		err = msgp.WrapError(err, "TxHash48List")
		return
	}
	for za0001 := range z.TxHash48List {
		err = en.WriteUint64(z.TxHash48List[za0001])
		if err != nil {
			err = msgp.WrapError(err, "TxHash48List", za0001)
			return
		}
	}
	// write "ai"
	err = en.Append(0xa2, 0x61, 0x69)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.AddrHashes)))
	if err != nil {
		err = msgp.WrapError(err, "AddrHashes")
		return
	}
	for za0002 := range z.AddrHashes {
		err = en.WriteUint64(z.AddrHashes[za0002])
		if err != nil {
			err = msgp.WrapError(err, "AddrHashes", za0002)
			return
		}
	}
	// write "ap"
	err = en.Append(0xa2, 0x61, 0x70)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.AddrPosLists)))
	if err != nil {
		err = msgp.WrapError(err, "AddrPosLists")
		return
	}
	for za0003 := range z.AddrPosLists {
		err = en.WriteArrayHeader(uint32(len(z.AddrPosLists[za0003])))
		if err != nil {
			err = msgp.WrapError(err, "AddrPosLists", za0003)
			return
		}
		for za0004 := range z.AddrPosLists[za0003] {
			err = en.WriteUint32(z.AddrPosLists[za0003][za0004])
			if err != nil {
				err = msgp.WrapError(err, "AddrPosLists", za0003, za0004)
				return
			}
		}
	}
	// write "ti"
	err = en.Append(0xa2, 0x74, 0x69)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.TopicHashes)))
	if err != nil {
		err = msgp.WrapError(err, "TopicHashes")
		return
	}
	for za0005 := range z.TopicHashes {
		err = en.WriteUint64(z.TopicHashes[za0005])
		if err != nil {
			err = msgp.WrapError(err, "TopicHashes", za0005)
			return
		}
	}
	// write "tp"
	err = en.Append(0xa2, 0x74, 0x70)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.TopicPosLists)))
	if err != nil {
		err = msgp.WrapError(err, "TopicPosLists")
		return
	}
	for za0006 := range z.TopicPosLists {
		err = en.WriteArrayHeader(uint32(len(z.TopicPosLists[za0006])))
		if err != nil {
			err = msgp.WrapError(err, "TopicPosLists", za0006)
			return
		}
		for za0007 := range z.TopicPosLists[za0006] {
			err = en.WriteUint32(z.TopicPosLists[za0006][za0007])
			if err != nil {
				err = msgp.WrapError(err, "TopicPosLists", za0006, za0007)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *BlockIndex) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "ht"
	o = append(o, 0x87, 0xa2, 0x68, 0x74)
	o = msgp.AppendUint32(o, z.Height)
	// string "bh"
	o = append(o, 0xa2, 0x62, 0x68)
	o = msgp.AppendUint64(o, z.BlockHash48)
	// string "tx"
	o = append(o, 0xa2, 0x74, 0x78)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TxHash48List)))
	for za0001 := range z.TxHash48List {
		o = msgp.AppendUint64(o, z.TxHash48List[za0001])
	}
	// string "ai"
	o = append(o, 0xa2, 0x61, 0x69)
	o = msgp.AppendArrayHeader(o, uint32(len(z.AddrHashes)))
	for za0002 := range z.AddrHashes {
		o = msgp.AppendUint64(o, z.AddrHashes[za0002])
	}
	// string "ap"
	o = append(o, 0xa2, 0x61, 0x70)
	o = msgp.AppendArrayHeader(o, uint32(len(z.AddrPosLists)))
	for za0003 := range z.AddrPosLists {
		o = msgp.AppendArrayHeader(o, uint32(len(z.AddrPosLists[za0003])))
		for za0004 := range z.AddrPosLists[za0003] {
			o = msgp.AppendUint32(o, z.AddrPosLists[za0003][za0004])
		}
	}
	// string "ti"
	o = append(o, 0xa2, 0x74, 0x69)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TopicHashes)))
	for za0005 := range z.TopicHashes {
		o = msgp.AppendUint64(o, z.TopicHashes[za0005])
	}
	// string "tp"
	o = append(o, 0xa2, 0x74, 0x70)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TopicPosLists)))
	for za0006 := range z.TopicPosLists {
		o = msgp.AppendArrayHeader(o, uint32(len(z.TopicPosLists[za0006])))
		for za0007 := range z.TopicPosLists[za0006] {
			o = msgp.AppendUint32(o, z.TopicPosLists[za0006][za0007])
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BlockIndex) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "ht":
			z.Height, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Height")
				return
			}
		case "bh":
			z.BlockHash48, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BlockHash48")
				return
			}
		case "tx":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxHash48List")
				return
			}
			if cap(z.TxHash48List) >= int(zb0002) {
				z.TxHash48List = (z.TxHash48List)[:zb0002]
			} else {
				z.TxHash48List = make([]uint64, zb0002)
			}
			for za0001 := range z.TxHash48List {
				z.TxHash48List[za0001], bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "TxHash48List", za0001)
					return
				}
			}
		case "ai":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AddrHashes")
				return
			}
			if cap(z.AddrHashes) >= int(zb0003) {
				z.AddrHashes = (z.AddrHashes)[:zb0003]
			} else {
				z.AddrHashes = make([]uint64, zb0003)
			}
			for za0002 := range z.AddrHashes {
				z.AddrHashes[za0002], bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "AddrHashes", za0002)
					return
				}
			}
		case "ap":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AddrPosLists")
				return
			}
			if cap(z.AddrPosLists) >= int(zb0004) {
				z.AddrPosLists = (z.AddrPosLists)[:zb0004]
			} else {
				z.AddrPosLists = make([][]uint32, zb0004)
			}
			for za0003 := range z.AddrPosLists {
				var zb0005 uint32
				zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "AddrPosLists", za0003)
					return
				}
				if cap(z.AddrPosLists[za0003]) >= int(zb0005) {
					z.AddrPosLists[za0003] = (z.AddrPosLists[za0003])[:zb0005]
				} else {
					z.AddrPosLists[za0003] = make([]uint32, zb0005)
				}
				for za0004 := range z.AddrPosLists[za0003] {
					z.AddrPosLists[za0003][za0004], bts, err = msgp.ReadUint32Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "AddrPosLists", za0003, za0004)
						return
					}
				}
			}
		case "ti":
			var zb0006 uint32
			zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TopicHashes")
				return
			}
			if cap(z.TopicHashes) >= int(zb0006) {
				z.TopicHashes = (z.TopicHashes)[:zb0006]
			} else {
				z.TopicHashes = make([]uint64, zb0006)
			}
			for za0005 := range z.TopicHashes {
				z.TopicHashes[za0005], bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "TopicHashes", za0005)
					return
				}
			}
		case "tp":
			var zb0007 uint32
			zb0007, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TopicPosLists")
				return
			}
			if cap(z.TopicPosLists) >= int(zb0007) {
				z.TopicPosLists = (z.TopicPosLists)[:zb0007]
			} else {
				z.TopicPosLists = make([][]uint32, zb0007)
			}
			for za0006 := range z.TopicPosLists {
				var zb0008 uint32
				zb0008, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "TopicPosLists", za0006)
					return
				}
				if cap(z.TopicPosLists[za0006]) >= int(zb0008) {
					z.TopicPosLists[za0006] = (z.TopicPosLists[za0006])[:zb0008]
				} else {
					z.TopicPosLists[za0006] = make([]uint32, zb0008)
				}
				for za0007 := range z.TopicPosLists[za0006] {
					z.TopicPosLists[za0006][za0007], bts, err = msgp.ReadUint32Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "TopicPosLists", za0006, za0007)
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BlockIndex) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint32Size + 3 + msgp.Uint64Size + 3 + msgp.ArrayHeaderSize + (len(z.TxHash48List) * (msgp.Uint64Size)) + 3 + msgp.ArrayHeaderSize + (len(z.AddrHashes) * (msgp.Uint64Size)) + 3 + msgp.ArrayHeaderSize
	for za0003 := range z.AddrPosLists {
		s += msgp.ArrayHeaderSize + (len(z.AddrPosLists[za0003]) * (msgp.Uint32Size))
	}
	s += 3 + msgp.ArrayHeaderSize + (len(z.TopicHashes) * (msgp.Uint64Size)) + 3 + msgp.ArrayHeaderSize
	for za0006 := range z.TopicPosLists {
		s += msgp.ArrayHeaderSize + (len(z.TopicPosLists[za0006]) * (msgp.Uint32Size))
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *IndexEntry) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "h":
			z.Hash48, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Hash48")
				return
			}
		case "l":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "PosList")
				return
			}
			if cap(z.PosList) >= int(zb0002) {
				z.PosList = (z.PosList)[:zb0002]
			} else {
				z.PosList = make([]uint32, zb0002)
			}
			for za0001 := range z.PosList {
				z.PosList[za0001], err = dc.ReadUint32()
				if err != nil {
					err = msgp.WrapError(err, "PosList", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *IndexEntry) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "h"
	err = en.Append(0x82, 0xa1, 0x68)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Hash48)
	if err != nil {
		err = msgp.WrapError(err, "Hash48")
		return
	}
	// write "l"
	err = en.Append(0xa1, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.PosList)))
	if err != nil {
		err = msgp.WrapError(err, "PosList")
		return
	}
	for za0001 := range z.PosList {
		err = en.WriteUint32(z.PosList[za0001])
		if err != nil {
			err = msgp.WrapError(err, "PosList", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *IndexEntry) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "h"
	o = append(o, 0x82, 0xa1, 0x68)
	o = msgp.AppendUint64(o, z.Hash48)
	// string "l"
	o = append(o, 0xa1, 0x6c)
	o = msgp.AppendArrayHeader(o, uint32(len(z.PosList)))
	for za0001 := range z.PosList {
		o = msgp.AppendUint32(o, z.PosList[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *IndexEntry) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "h":
			z.Hash48, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Hash48")
				return
			}
		case "l":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PosList")
				return
			}
			if cap(z.PosList) >= int(zb0002) {
				z.PosList = (z.PosList)[:zb0002]
			} else {
				z.PosList = make([]uint32, zb0002)
			}
			for za0001 := range z.PosList {
				z.PosList[za0001], bts, err = msgp.ReadUint32Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "PosList", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *IndexEntry) Msgsize() (s int) {
	s = 1 + 2 + msgp.Uint64Size + 2 + msgp.ArrayHeaderSize + (len(z.PosList) * (msgp.Uint32Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Log) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "a":
			err = dc.ReadExactBytes((z.Address)[:])
			if err != nil {
				err = msgp.WrapError(err, "Address")
				return
			}
		case "t":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Topics")
				return
			}
			if cap(z.Topics) >= int(zb0002) {
				z.Topics = (z.Topics)[:zb0002]
			} else {
				z.Topics = make([][32]byte, zb0002)
			}
			for za0002 := range z.Topics {
				err = dc.ReadExactBytes((z.Topics[za0002])[:])
				if err != nil {
					err = msgp.WrapError(err, "Topics", za0002)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Log) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "a"
	err = en.Append(0x82, 0xa1, 0x61)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.Address)[:])
	if err != nil {
		err = msgp.WrapError(err, "Address")
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Topics)))
	if err != nil {
		err = msgp.WrapError(err, "Topics")
		return
	}
	for za0002 := range z.Topics {
		err = en.WriteBytes((z.Topics[za0002])[:])
		if err != nil {
			err = msgp.WrapError(err, "Topics", za0002)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Log) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "a"
	o = append(o, 0x82, 0xa1, 0x61)
	o = msgp.AppendBytes(o, (z.Address)[:])
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Topics)))
	for za0002 := range z.Topics {
		o = msgp.AppendBytes(o, (z.Topics[za0002])[:])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Log) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "a":
			bts, err = msgp.ReadExactBytes(bts, (z.Address)[:])
			if err != nil {
				err = msgp.WrapError(err, "Address")
				return
			}
		case "t":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Topics")
				return
			}
			if cap(z.Topics) >= int(zb0002) {
				z.Topics = (z.Topics)[:zb0002]
			} else {
				z.Topics = make([][32]byte, zb0002)
			}
			for za0002 := range z.Topics {
				bts, err = msgp.ReadExactBytes(bts, (z.Topics[za0002])[:])
				if err != nil {
					err = msgp.WrapError(err, "Topics", za0002)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Log) Msgsize() (s int) {
	s = 1 + 2 + msgp.ArrayHeaderSize + (20 * (msgp.ByteSize)) + 2 + msgp.ArrayHeaderSize + (len(z.Topics) * (32 * (msgp.ByteSize)))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Tx) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "h":
			err = dc.ReadExactBytes((z.HashId)[:])
			if err != nil {
				err = msgp.WrapError(err, "HashId")
				return
			}
		case "c":
			z.Content, err = dc.ReadBytes(z.Content)
			if err != nil {
				err = msgp.WrapError(err, "Content")
				return
			}
		case "l":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "LogList")
				return
			}
			if cap(z.LogList) >= int(zb0002) {
				z.LogList = (z.LogList)[:zb0002]
			} else {
				z.LogList = make([]Log, zb0002)
			}
			for za0002 := range z.LogList {
				err = z.LogList[za0002].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "LogList", za0002)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Tx) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "h"
	err = en.Append(0x83, 0xa1, 0x68)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.HashId)[:])
	if err != nil {
		err = msgp.WrapError(err, "HashId")
		return
	}
	// write "c"
	err = en.Append(0xa1, 0x63)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Content)
	if err != nil {
		err = msgp.WrapError(err, "Content")
		return
	}
	// write "l"
	err = en.Append(0xa1, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.LogList)))
	if err != nil {
		err = msgp.WrapError(err, "LogList")
		return
	}
	for za0002 := range z.LogList {
		err = z.LogList[za0002].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "LogList", za0002)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Tx) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "h"
	o = append(o, 0x83, 0xa1, 0x68)
	o = msgp.AppendBytes(o, (z.HashId)[:])
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendBytes(o, z.Content)
	// string "l"
	o = append(o, 0xa1, 0x6c)
	o = msgp.AppendArrayHeader(o, uint32(len(z.LogList)))
	for za0002 := range z.LogList {
		o, err = z.LogList[za0002].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "LogList", za0002)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Tx) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "h":
			bts, err = msgp.ReadExactBytes(bts, (z.HashId)[:])
			if err != nil {
				err = msgp.WrapError(err, "HashId")
				return
			}
		case "c":
			z.Content, bts, err = msgp.ReadBytesBytes(bts, z.Content)
			if err != nil {
				err = msgp.WrapError(err, "Content")
				return
			}
		case "l":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LogList")
				return
			}
			if cap(z.LogList) >= int(zb0002) {
				z.LogList = (z.LogList)[:zb0002]
			} else {
				z.LogList = make([]Log, zb0002)
			}
			for za0002 := range z.LogList {
				bts, err = z.LogList[za0002].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "LogList", za0002)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tx) Msgsize() (s int) {
	s = 1 + 2 + msgp.ArrayHeaderSize + (32 * (msgp.ByteSize)) + 2 + msgp.BytesPrefixSize + len(z.Content) + 2 + msgp.ArrayHeaderSize
	for za0002 := range z.LogList {
		s += z.LogList[za0002].Msgsize()
	}
	return
}
