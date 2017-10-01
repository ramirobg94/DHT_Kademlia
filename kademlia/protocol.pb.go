// Code generated by protoc-gen-go.
// source: protocol.proto
// DO NOT EDIT!

/*
Package kademlia is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	ProtocolPackage
*/
package kademlia

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ProtocolPackage_MessageSent int32

const (
	ProtocolPackage_PING      ProtocolPackage_MessageSent = 1
	ProtocolPackage_STORE     ProtocolPackage_MessageSent = 2
	ProtocolPackage_FINDNODE  ProtocolPackage_MessageSent = 3
	ProtocolPackage_FINDVALUE ProtocolPackage_MessageSent = 4
)

var ProtocolPackage_MessageSent_name = map[int32]string{
	1: "PING",
	2: "STORE",
	3: "FINDNODE",
	4: "FINDVALUE",
}
var ProtocolPackage_MessageSent_value = map[string]int32{
	"PING":      1,
	"STORE":     2,
	"FINDNODE":  3,
	"FINDVALUE": 4,
}

func (x ProtocolPackage_MessageSent) Enum() *ProtocolPackage_MessageSent {
	p := new(ProtocolPackage_MessageSent)
	*p = x
	return p
}
func (x ProtocolPackage_MessageSent) String() string {
	return proto.EnumName(ProtocolPackage_MessageSent_name, int32(x))
}
func (x *ProtocolPackage_MessageSent) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProtocolPackage_MessageSent_value, data, "ProtocolPackage_MessageSent")
	if err != nil {
		return err
	}
	*x = ProtocolPackage_MessageSent(value)
	return nil
}

type ProtocolPackage struct {
	ClientID         []byte                         `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	Address          *string                        `protobuf:"bytes,2,req,name=address" json:"address,omitempty"`
	MessageSent      *ProtocolPackage_MessageSent   `protobuf:"varint,3,req,name=messageSent,enum=kademlia.ProtocolPackage_MessageSent" json:"messageSent,omitempty"`
	FindID           []byte                         `protobuf:"bytes,4,opt,name=findID" json:"findID,omitempty"`
	FindValue        []byte                         `protobuf:"bytes,5,opt,name=findValue" json:"findValue,omitempty"`
	ContactsKNearest []*ProtocolPackage_ContactInfo `protobuf:"bytes,6,rep,name=contactsKNearest" json:"contactsKNearest,omitempty"`
	StoredeID        *string                        `protobuf:"bytes,7,opt,name=storedeID" json:"storedeID,omitempty"`
	File             *string                        `protobuf:"bytes,8,opt,name=file" json:"file,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *ProtocolPackage) Reset()         { *m = ProtocolPackage{} }
func (m *ProtocolPackage) String() string { return proto.CompactTextString(m) }
func (*ProtocolPackage) ProtoMessage()    {}

func (m *ProtocolPackage) GetClientID() []byte {
	if m != nil {
		return m.ClientID
	}
	return nil
}

func (m *ProtocolPackage) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *ProtocolPackage) GetMessageSent() ProtocolPackage_MessageSent {
	if m != nil && m.MessageSent != nil {
		return *m.MessageSent
	}
	return ProtocolPackage_PING
}

func (m *ProtocolPackage) GetFindID() []byte {
	if m != nil {
		return m.FindID
	}
	return nil
}

func (m *ProtocolPackage) GetFindValue() []byte {
	if m != nil {
		return m.FindValue
	}
	return nil
}

func (m *ProtocolPackage) GetContactsKNearest() []*ProtocolPackage_ContactInfo {
	if m != nil {
		return m.ContactsKNearest
	}
	return nil
}

func (m *ProtocolPackage) GetStoredeID() string {
	if m != nil && m.StoredeID != nil {
		return *m.StoredeID
	}
	return ""
}

func (m *ProtocolPackage) GetFile() string {
	if m != nil && m.File != nil {
		return *m.File
	}
	return ""
}

type ProtocolPackage_ContactInfo struct {
	ContactID        []byte  `protobuf:"bytes,1,opt,name=contactID" json:"contactID,omitempty"`
	Address          *string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Distance         []byte  `protobuf:"bytes,3,opt,name=distance" json:"distance,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtocolPackage_ContactInfo) Reset()         { *m = ProtocolPackage_ContactInfo{} }
func (m *ProtocolPackage_ContactInfo) String() string { return proto.CompactTextString(m) }
func (*ProtocolPackage_ContactInfo) ProtoMessage()    {}

func (m *ProtocolPackage_ContactInfo) GetContactID() []byte {
	if m != nil {
		return m.ContactID
	}
	return nil
}

func (m *ProtocolPackage_ContactInfo) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *ProtocolPackage_ContactInfo) GetDistance() []byte {
	if m != nil {
		return m.Distance
	}
	return nil
}

func init() {
	proto.RegisterEnum("kademlia.ProtocolPackage_MessageSent", ProtocolPackage_MessageSent_name, ProtocolPackage_MessageSent_value)
}
