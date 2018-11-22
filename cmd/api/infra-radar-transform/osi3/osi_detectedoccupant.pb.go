// Code generated by protoc-gen-go. DO NOT EDIT.
// source: osi_detectedoccupant.proto

package osi3

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//
// \brief A vehicle occupant as detected and perceived by an interior sensor.
//
type DetectedOccupant struct {
	// Common information of one detected item.
	//
	Header *DetectedItemHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// A list of candidates for this occupant as estimated by the
	// sensor.
	//
	// \note OSI uses singular instead of plural for repeated field names.
	//
	Candidate            []*DetectedOccupant_CandidateOccupant `protobuf:"bytes,2,rep,name=candidate,proto3" json:"candidate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DetectedOccupant) Reset()         { *m = DetectedOccupant{} }
func (m *DetectedOccupant) String() string { return proto.CompactTextString(m) }
func (*DetectedOccupant) ProtoMessage()    {}
func (*DetectedOccupant) Descriptor() ([]byte, []int) {
	return fileDescriptor_f373320fc50bc7e9, []int{0}
}

func (m *DetectedOccupant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectedOccupant.Unmarshal(m, b)
}
func (m *DetectedOccupant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectedOccupant.Marshal(b, m, deterministic)
}
func (m *DetectedOccupant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectedOccupant.Merge(m, src)
}
func (m *DetectedOccupant) XXX_Size() int {
	return xxx_messageInfo_DetectedOccupant.Size(m)
}
func (m *DetectedOccupant) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectedOccupant.DiscardUnknown(m)
}

var xxx_messageInfo_DetectedOccupant proto.InternalMessageInfo

func (m *DetectedOccupant) GetHeader() *DetectedItemHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DetectedOccupant) GetCandidate() []*DetectedOccupant_CandidateOccupant {
	if m != nil {
		return m.Candidate
	}
	return nil
}

//
// \brief A candidate for a detected occupant as estimated by
// the sensor.
//
type DetectedOccupant_CandidateOccupant struct {
	// The estimated probability that this candidate is the true value.
	//
	// \note The sum of all \c #probability must be one. This probability is
	// given under the condition of
	// \c DetectedItemHeader::existence_probability.
	//
	// Range: [0,1]
	//
	Probability float64 `protobuf:"fixed64,1,opt,name=probability,proto3" json:"probability,omitempty"`
	// The detected vehicle occupant classification.
	//
	// \note IDs, which are referenced in this message, usually
	// reference to \c DetectedXXX::tracking_id IDs.
	//
	Classification       *Occupant_Classification `protobuf:"bytes,2,opt,name=classification,proto3" json:"classification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DetectedOccupant_CandidateOccupant) Reset()         { *m = DetectedOccupant_CandidateOccupant{} }
func (m *DetectedOccupant_CandidateOccupant) String() string { return proto.CompactTextString(m) }
func (*DetectedOccupant_CandidateOccupant) ProtoMessage()    {}
func (*DetectedOccupant_CandidateOccupant) Descriptor() ([]byte, []int) {
	return fileDescriptor_f373320fc50bc7e9, []int{0, 0}
}

func (m *DetectedOccupant_CandidateOccupant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectedOccupant_CandidateOccupant.Unmarshal(m, b)
}
func (m *DetectedOccupant_CandidateOccupant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectedOccupant_CandidateOccupant.Marshal(b, m, deterministic)
}
func (m *DetectedOccupant_CandidateOccupant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectedOccupant_CandidateOccupant.Merge(m, src)
}
func (m *DetectedOccupant_CandidateOccupant) XXX_Size() int {
	return xxx_messageInfo_DetectedOccupant_CandidateOccupant.Size(m)
}
func (m *DetectedOccupant_CandidateOccupant) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectedOccupant_CandidateOccupant.DiscardUnknown(m)
}

var xxx_messageInfo_DetectedOccupant_CandidateOccupant proto.InternalMessageInfo

func (m *DetectedOccupant_CandidateOccupant) GetProbability() float64 {
	if m != nil {
		return m.Probability
	}
	return 0
}

func (m *DetectedOccupant_CandidateOccupant) GetClassification() *Occupant_Classification {
	if m != nil {
		return m.Classification
	}
	return nil
}

func init() {
	proto.RegisterType((*DetectedOccupant)(nil), "osi3.DetectedOccupant")
	proto.RegisterType((*DetectedOccupant_CandidateOccupant)(nil), "osi3.DetectedOccupant.CandidateOccupant")
}

func init() { proto.RegisterFile("osi_detectedoccupant.proto", fileDescriptor_f373320fc50bc7e9) }

var fileDescriptor_f373320fc50bc7e9 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4b, 0x4a, 0x04, 0x31,
	0x10, 0x40, 0xe9, 0x28, 0x03, 0xa6, 0x41, 0x34, 0xab, 0x10, 0x10, 0x1a, 0x57, 0xbd, 0x0a, 0x32,
	0x73, 0x03, 0x7f, 0x8c, 0x2b, 0x21, 0x17, 0x90, 0x7c, 0x4a, 0x2c, 0x19, 0x3b, 0x4d, 0xa7, 0x5c,
	0x08, 0x9e, 0xc0, 0x53, 0xcb, 0xc4, 0x84, 0x31, 0xba, 0xad, 0xf7, 0x5e, 0xa5, 0x08, 0x57, 0x31,
	0xe1, 0x53, 0x00, 0x02, 0x4f, 0x10, 0xa2, 0xf7, 0xef, 0xb3, 0x9d, 0x48, 0xcf, 0x4b, 0xa4, 0x28,
	0x8e, 0x63, 0xc2, 0x8d, 0x12, 0x7b, 0xa3, 0x25, 0x4a, 0x36, 0x95, 0x7b, 0x05, 0x5f, 0xc8, 0xe5,
	0x17, 0xe3, 0x67, 0xb7, 0x05, 0x3c, 0x96, 0x48, 0x5c, 0xf1, 0xd5, 0x0b, 0xd8, 0x00, 0x8b, 0xec,
	0x86, 0x6e, 0xec, 0xd7, 0x52, 0xef, 0x37, 0xeb, 0xea, 0x3d, 0x10, 0xbc, 0x6d, 0x33, 0x37, 0xc5,
	0x13, 0xf7, 0xfc, 0xc4, 0xdb, 0x29, 0x60, 0xb0, 0x04, 0x92, 0x0d, 0x47, 0x63, 0xbf, 0x1e, 0xdb,
	0xa8, 0x2e, 0xd7, 0x37, 0xd5, 0xab, 0x13, 0x73, 0x48, 0xd5, 0x27, 0x3f, 0xff, 0xc7, 0xc5, 0xc0,
	0xfb, 0x79, 0x89, 0xce, 0x3a, 0xdc, 0x21, 0x7d, 0xe4, 0x9b, 0x3a, 0xf3, 0x7b, 0x24, 0xee, 0xf8,
	0xa9, 0xdf, 0xd9, 0x94, 0xf0, 0x19, 0xbd, 0x25, 0x8c, 0x93, 0x64, 0xf9, 0xf0, 0x8b, 0x9f, 0x1b,
	0x0e, 0x6f, 0x37, 0x92, 0xf9, 0x13, 0x5d, 0xb3, 0x6d, 0xe7, 0x56, 0xf9, 0x5f, 0x36, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x47, 0x4c, 0x70, 0x05, 0x69, 0x01, 0x00, 0x00,
}
