// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token/tx.proto

package token

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgIssueToken defines an SDK message for issuing a new token.
type MsgIssueToken struct {
	Symbol        string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Scale         uint32 `protobuf:"varint,3,opt,name=scale,proto3" json:"scale,omitempty"`
	MinUnit       string `protobuf:"bytes,4,opt,name=min_unit,json=minUnit,proto3" json:"min_unit,omitempty" yaml:"min_unit"`
	InitialSupply uint64 `protobuf:"varint,5,opt,name=initial_supply,json=initialSupply,proto3" json:"initial_supply,omitempty" yaml:"initial_supply"`
	MaxSupply     uint64 `protobuf:"varint,6,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty" yaml:"max_supply"`
	Mintable      bool   `protobuf:"varint,7,opt,name=mintable,proto3" json:"mintable,omitempty"`
	Owner         string `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *MsgIssueToken) Reset()         { *m = MsgIssueToken{} }
func (m *MsgIssueToken) String() string { return proto.CompactTextString(m) }
func (*MsgIssueToken) ProtoMessage()    {}
func (*MsgIssueToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef78f47708126356, []int{0}
}
func (m *MsgIssueToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIssueToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIssueToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIssueToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIssueToken.Merge(m, src)
}
func (m *MsgIssueToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgIssueToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIssueToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIssueToken proto.InternalMessageInfo

// MsgMintToken defines an SDK message for transferring the token owner.
type MsgTransferTokenOwner struct {
	SrcOwner string `protobuf:"bytes,1,opt,name=src_owner,json=srcOwner,proto3" json:"src_owner,omitempty" yaml:"src_owner"`
	DstOwner string `protobuf:"bytes,2,opt,name=dst_owner,json=dstOwner,proto3" json:"dst_owner,omitempty" yaml:"dst_owner"`
	Symbol   string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (m *MsgTransferTokenOwner) Reset()         { *m = MsgTransferTokenOwner{} }
func (m *MsgTransferTokenOwner) String() string { return proto.CompactTextString(m) }
func (*MsgTransferTokenOwner) ProtoMessage()    {}
func (*MsgTransferTokenOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef78f47708126356, []int{1}
}
func (m *MsgTransferTokenOwner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferTokenOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferTokenOwner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferTokenOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferTokenOwner.Merge(m, src)
}
func (m *MsgTransferTokenOwner) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferTokenOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferTokenOwner.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferTokenOwner proto.InternalMessageInfo

// MsgEditToken defines an SDK message for editing a new token.
type MsgEditToken struct {
	Symbol    string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MaxSupply uint64 `protobuf:"varint,3,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty" yaml:"max_supply"`
	Mintable  Bool   `protobuf:"bytes,4,opt,name=mintable,proto3,casttype=Bool" json:"mintable,omitempty"`
	Owner     string `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *MsgEditToken) Reset()         { *m = MsgEditToken{} }
func (m *MsgEditToken) String() string { return proto.CompactTextString(m) }
func (*MsgEditToken) ProtoMessage()    {}
func (*MsgEditToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef78f47708126356, []int{2}
}
func (m *MsgEditToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEditToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEditToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEditToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEditToken.Merge(m, src)
}
func (m *MsgEditToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgEditToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEditToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEditToken proto.InternalMessageInfo

// MsgMintToken defines an SDK message for minting a new token.
type MsgMintToken struct {
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	To     string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Owner  string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *MsgMintToken) Reset()         { *m = MsgMintToken{} }
func (m *MsgMintToken) String() string { return proto.CompactTextString(m) }
func (*MsgMintToken) ProtoMessage()    {}
func (*MsgMintToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef78f47708126356, []int{3}
}
func (m *MsgMintToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintToken.Merge(m, src)
}
func (m *MsgMintToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintToken proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgIssueToken)(nil), "irismod.token.MsgIssueToken")
	proto.RegisterType((*MsgTransferTokenOwner)(nil), "irismod.token.MsgTransferTokenOwner")
	proto.RegisterType((*MsgEditToken)(nil), "irismod.token.MsgEditToken")
	proto.RegisterType((*MsgMintToken)(nil), "irismod.token.MsgMintToken")
}

func init() { proto.RegisterFile("token/tx.proto", fileDescriptor_ef78f47708126356) }

var fileDescriptor_ef78f47708126356 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xae, 0xd3, 0xb4, 0x97, 0x5a, 0xb4, 0x80, 0x69, 0x4f, 0xe1, 0x86, 0xa4, 0xb2, 0x90, 0xe8,
	0x72, 0x8d, 0x10, 0x4c, 0x4c, 0x28, 0xd2, 0x0d, 0x0c, 0x15, 0x92, 0x39, 0x16, 0x96, 0xca, 0x6d,
	0x42, 0xb0, 0x2e, 0xb6, 0xab, 0xd8, 0x11, 0xed, 0xbf, 0x60, 0xe1, 0x57, 0xc0, 0x0f, 0xb9, 0xf1,
	0x46, 0xa6, 0x08, 0xda, 0x7f, 0xd0, 0x91, 0x09, 0xc5, 0xc9, 0xe5, 0x7a, 0x08, 0x09, 0x6e, 0xf3,
	0xf7, 0xbe, 0xf7, 0xe9, 0x7d, 0xfe, 0x9e, 0x1e, 0x1c, 0x68, 0x79, 0x11, 0x8b, 0x40, 0xaf, 0xa7,
	0xab, 0x4c, 0x6a, 0x89, 0xfa, 0x2c, 0x63, 0x8a, 0xcb, 0x68, 0x6a, 0xea, 0x27, 0xc3, 0x44, 0x26,
	0xd2, 0x30, 0x41, 0xf9, 0xaa, 0x9a, 0xf0, 0x57, 0x0b, 0xf6, 0x67, 0x2a, 0x79, 0xad, 0x54, 0x1e,
	0x9f, 0x97, 0x7d, 0xe8, 0x18, 0x76, 0xd5, 0x86, 0x2f, 0x64, 0xea, 0x82, 0x31, 0x98, 0xf4, 0x48,
	0x8d, 0x10, 0x82, 0xb6, 0xa0, 0x3c, 0x76, 0x2d, 0x53, 0x35, 0x6f, 0x34, 0x84, 0x1d, 0xb5, 0xa4,
	0x69, 0xec, 0xb6, 0xc7, 0x60, 0xd2, 0x27, 0x15, 0x40, 0x53, 0xe8, 0x70, 0x26, 0xe6, 0xb9, 0x60,
	0xda, 0xb5, 0xcb, 0xee, 0xf0, 0xd1, 0xbe, 0xf0, 0xef, 0x6f, 0x28, 0x4f, 0x5f, 0xe2, 0x6b, 0x06,
	0x93, 0x23, 0xce, 0xc4, 0x3b, 0xc1, 0x34, 0x7a, 0x05, 0x07, 0x4c, 0x30, 0xcd, 0x68, 0x3a, 0x57,
	0xf9, 0x6a, 0x95, 0x6e, 0xdc, 0xce, 0x18, 0x4c, 0xec, 0xf0, 0xf1, 0xbe, 0xf0, 0x47, 0x95, 0xea,
	0x36, 0x8f, 0x49, 0xbf, 0x2e, 0xbc, 0x35, 0x18, 0xbd, 0x80, 0x90, 0xd3, 0xf5, 0xb5, 0xba, 0x6b,
	0xd4, 0xa3, 0x7d, 0xe1, 0x3f, 0xac, 0x67, 0x36, 0x1c, 0x26, 0x3d, 0x4e, 0xd7, 0xb5, 0xea, 0xc4,
	0xf8, 0xd4, 0x74, 0x91, 0xc6, 0xee, 0xd1, 0x18, 0x4c, 0x1c, 0xd2, 0xe0, 0xf2, 0x67, 0xf2, 0x93,
	0x88, 0x33, 0xd7, 0x31, 0xdf, 0xad, 0x00, 0xfe, 0x02, 0xe0, 0x68, 0xa6, 0x92, 0xf3, 0x8c, 0x0a,
	0xf5, 0x21, 0xce, 0x4c, 0x60, 0x6f, 0x4a, 0x06, 0x3d, 0x83, 0x3d, 0x95, 0x2d, 0xe7, 0x95, 0xc6,
	0x04, 0x17, 0x0e, 0xf7, 0x85, 0xff, 0xa0, 0x32, 0xd0, 0x50, 0x98, 0x38, 0x2a, 0x5b, 0x36, 0x92,
	0x48, 0xe9, 0x5a, 0x62, 0xfd, 0x29, 0x69, 0x28, 0x4c, 0x9c, 0x48, 0xe9, 0x4a, 0x72, 0xb3, 0x9b,
	0xf6, 0xe1, 0x6e, 0xf0, 0x37, 0x00, 0xef, 0xcd, 0x54, 0x72, 0x16, 0x31, 0x7d, 0xf7, 0x25, 0xde,
	0x0e, 0xaf, 0xfd, 0x9f, 0xe1, 0x3d, 0x39, 0x08, 0xaf, 0x5a, 0xb2, 0xf3, 0xab, 0xf0, 0xed, 0x50,
	0xca, 0xf4, 0x6f, 0x31, 0x76, 0x0e, 0x63, 0x8c, 0x8c, 0xdb, 0x19, 0x13, 0xff, 0x70, 0x7b, 0x0c,
	0xbb, 0x94, 0xcb, 0x5c, 0x68, 0xe3, 0xd7, 0x26, 0x35, 0x42, 0x03, 0x68, 0x69, 0x59, 0x47, 0x60,
	0x69, 0x79, 0x33, 0xc5, 0x3e, 0x98, 0x12, 0x9e, 0x5d, 0xfe, 0xf4, 0x5a, 0x97, 0x5b, 0x0f, 0x5c,
	0x6d, 0x3d, 0xf0, 0x63, 0xeb, 0x81, 0xcf, 0x3b, 0xaf, 0x75, 0xb5, 0xf3, 0x5a, 0xdf, 0x77, 0x5e,
	0xeb, 0xfd, 0xd3, 0x84, 0xe9, 0x8f, 0xf9, 0x62, 0xba, 0x94, 0x3c, 0x28, 0x0f, 0x45, 0xc4, 0x3a,
	0xa8, 0x0f, 0xe6, 0x54, 0x45, 0x17, 0xa7, 0x89, 0x0c, 0xcc, 0xdd, 0x2c, 0xba, 0xe6, 0x50, 0x9e,
	0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x41, 0xfc, 0x68, 0x5f, 0x03, 0x00, 0x00,
}

func (m *MsgIssueToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIssueToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIssueToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x42
	}
	if m.Mintable {
		i--
		if m.Mintable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.MaxSupply != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x30
	}
	if m.InitialSupply != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.InitialSupply))
		i--
		dAtA[i] = 0x28
	}
	if len(m.MinUnit) > 0 {
		i -= len(m.MinUnit)
		copy(dAtA[i:], m.MinUnit)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MinUnit)))
		i--
		dAtA[i] = 0x22
	}
	if m.Scale != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Scale))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferTokenOwner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferTokenOwner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferTokenOwner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DstOwner) > 0 {
		i -= len(m.DstOwner)
		copy(dAtA[i:], m.DstOwner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DstOwner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SrcOwner) > 0 {
		i -= len(m.SrcOwner)
		copy(dAtA[i:], m.SrcOwner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SrcOwner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgEditToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEditToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEditToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Mintable) > 0 {
		i -= len(m.Mintable)
		copy(dAtA[i:], m.Mintable)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Mintable)))
		i--
		dAtA[i] = 0x22
	}
	if m.MaxSupply != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Amount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgIssueToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Scale != 0 {
		n += 1 + sovTx(uint64(m.Scale))
	}
	l = len(m.MinUnit)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.InitialSupply != 0 {
		n += 1 + sovTx(uint64(m.InitialSupply))
	}
	if m.MaxSupply != 0 {
		n += 1 + sovTx(uint64(m.MaxSupply))
	}
	if m.Mintable {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgTransferTokenOwner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SrcOwner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DstOwner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgEditToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MaxSupply != 0 {
		n += 1 + sovTx(uint64(m.MaxSupply))
	}
	l = len(m.Mintable)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgMintToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTx(uint64(m.Amount))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgIssueToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgIssueToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIssueToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scale", wireType)
			}
			m.Scale = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Scale |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUnit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinUnit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialSupply", wireType)
			}
			m.InitialSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mintable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Mintable = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgTransferTokenOwner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgTransferTokenOwner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferTokenOwner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DstOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DstOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgEditToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgEditToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEditToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mintable", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mintable = Bool(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMintToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMintToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
