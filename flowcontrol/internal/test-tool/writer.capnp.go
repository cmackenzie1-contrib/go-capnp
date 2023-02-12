// Code generated by capnpc-go. DO NOT EDIT.

package main

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
	fmt "fmt"
)

type Writer capnp.Client

// Writer_TypeID is the unique identifier for the type Writer.
const Writer_TypeID = 0xf82e58b4a78f136b

func (c Writer) Write(ctx context.Context, params func(Writer_write_Params) error) (Writer_write_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xf82e58b4a78f136b,
			MethodID:      0,
			InterfaceName: "writer.capnp:Writer",
			MethodName:    "write",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Writer_write_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Writer_write_Results_Future{Future: ans.Future()}, release

}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Writer) String() string {
	return fmt.Sprintf("%T(%v)", c, capnp.Client(c))
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Writer) AddRef() Writer {
	return Writer(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Writer) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Writer) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Writer) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Writer) DecodeFromPtr(p capnp.Ptr) Writer {
	return Writer(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Writer) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Writer) IsSame(other Writer) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Writer) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Writer) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
}

// A Writer_Server is a Writer with a local implementation.
type Writer_Server interface {
	Write(context.Context, Writer_write) error
}

// Writer_NewServer creates a new Server from an implementation of Writer_Server.
func Writer_NewServer(s Writer_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Writer_Methods(nil, s), s, c)
}

// Writer_ServerToClient creates a new Client from an implementation of Writer_Server.
// The caller is responsible for calling Release on the returned Client.
func Writer_ServerToClient(s Writer_Server) Writer {
	return Writer(capnp.NewClient(Writer_NewServer(s)))
}

// Writer_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Writer_Methods(methods []server.Method, s Writer_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf82e58b4a78f136b,
			MethodID:      0,
			InterfaceName: "writer.capnp:Writer",
			MethodName:    "write",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Write(ctx, Writer_write{call})
		},
	})

	return methods
}

// Writer_write holds the state for a server call to Writer.write.
// See server.Call for documentation.
type Writer_write struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Writer_write) Args() Writer_write_Params {
	return Writer_write_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Writer_write) AllocResults() (Writer_write_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Writer_write_Results(r), err
}

// Writer_List is a list of Writer.
type Writer_List = capnp.CapList[Writer]

// NewWriter creates a new list of Writer.
func NewWriter_List(s *capnp.Segment, sz int32) (Writer_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Writer](l), err
}

type Writer_write_Params capnp.Struct

// Writer_write_Params_TypeID is the unique identifier for the type Writer_write_Params.
const Writer_write_Params_TypeID = 0x80b8cd5f44e3c477

func NewWriter_write_Params(s *capnp.Segment) (Writer_write_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Writer_write_Params(st), err
}

func NewRootWriter_write_Params(s *capnp.Segment) (Writer_write_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Writer_write_Params(st), err
}

func ReadRootWriter_write_Params(msg *capnp.Message) (Writer_write_Params, error) {
	root, err := msg.Root()
	return Writer_write_Params(root.Struct()), err
}

func (s Writer_write_Params) String() string {
	str, _ := text.Marshal(0x80b8cd5f44e3c477, capnp.Struct(s))
	return str
}

func (s Writer_write_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Writer_write_Params) DecodeFromPtr(p capnp.Ptr) Writer_write_Params {
	return Writer_write_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Writer_write_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Writer_write_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Writer_write_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Writer_write_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Writer_write_Params) Data() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return []byte(p.Data()), err
}

func (s Writer_write_Params) HasData() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Writer_write_Params) SetData(v []byte) error {
	return capnp.Struct(s).SetData(0, v)
}

// Writer_write_Params_List is a list of Writer_write_Params.
type Writer_write_Params_List = capnp.StructList[Writer_write_Params]

// NewWriter_write_Params creates a new list of Writer_write_Params.
func NewWriter_write_Params_List(s *capnp.Segment, sz int32) (Writer_write_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Writer_write_Params](l), err
}

// Writer_write_Params_Future is a wrapper for a Writer_write_Params promised by a client call.
type Writer_write_Params_Future struct{ *capnp.Future }

func (f Writer_write_Params_Future) Struct() (Writer_write_Params, error) {
	p, err := f.Future.Ptr()
	return Writer_write_Params(p.Struct()), err
}

type Writer_write_Results capnp.Struct

// Writer_write_Results_TypeID is the unique identifier for the type Writer_write_Results.
const Writer_write_Results_TypeID = 0xd939de8c6024e7f8

func NewWriter_write_Results(s *capnp.Segment) (Writer_write_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Writer_write_Results(st), err
}

func NewRootWriter_write_Results(s *capnp.Segment) (Writer_write_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Writer_write_Results(st), err
}

func ReadRootWriter_write_Results(msg *capnp.Message) (Writer_write_Results, error) {
	root, err := msg.Root()
	return Writer_write_Results(root.Struct()), err
}

func (s Writer_write_Results) String() string {
	str, _ := text.Marshal(0xd939de8c6024e7f8, capnp.Struct(s))
	return str
}

func (s Writer_write_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Writer_write_Results) DecodeFromPtr(p capnp.Ptr) Writer_write_Results {
	return Writer_write_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Writer_write_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Writer_write_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Writer_write_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Writer_write_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Writer_write_Results_List is a list of Writer_write_Results.
type Writer_write_Results_List = capnp.StructList[Writer_write_Results]

// NewWriter_write_Results creates a new list of Writer_write_Results.
func NewWriter_write_Results_List(s *capnp.Segment, sz int32) (Writer_write_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Writer_write_Results](l), err
}

// Writer_write_Results_Future is a wrapper for a Writer_write_Results promised by a client call.
type Writer_write_Results_Future struct{ *capnp.Future }

func (f Writer_write_Results_Future) Struct() (Writer_write_Results, error) {
	p, err := f.Future.Ptr()
	return Writer_write_Results(p.Struct()), err
}

const schema_aca73f831c7ebfdd = "x\xda\x12\xa8u`1\xe4\xcdgb`\x0a\x94ae" +
	"\xfb_~\xe4\xb1K\xfc\xd9\x1d\x0d\x0c\x82\"\x8c\x0c\x0c" +
	"\xac\x8c\xec\x0c\x0c\xc6\xb2\x8c\\\x8c\x0c\x8c\xc2\xaa\x8c\xf6" +
	"\x0c\x8c\xff\x7f<WI\xe8\xb9gy\x13\xa2\x80\x05$" +
	"\xef\xca(\xc4\xc8\xc0\xf2?[\xb8\x7f\xf9\x96\x08\xbd\x1f" +
	"\x0c\x82\xbc\xcc\xff\xef\xee\xaf\x93i\xb6_\xbe\x86\x81\x81" +
	"QX\x97q\x91\xb0)\xc8$aCFw\xe1HF" +
	"v\x06\x9d\xff\xe5E\x99%\xa9Ez\xc9\xcc\x89\x05y" +
	"\x05V\xe1\x10\x1eXP% \xb1(1\xb7\x98\x81!" +
	"\x90\x85\x99\x85\x81\x81\x85\x91\x81A\x90W\x8b\x81!\x90" +
	"\x83\x991P\x84\x89\x91?%\xb1$\x91\x91\x97\x81\x89" +
	"\x91\x97\x81\x11\x9f9A\xa9\xc5\xa59%\x8c\xc5p5" +
	"\x8c05\xec%\xa9E\x01\x8c\x8c\x81,\xcc\xac\x0c\x0c" +
	"p/3\xc2\xbc&(h\xc4\xc0$\xc8\xca.\x0f\xd6" +
	"\xe8\xc0\x18\xc0\xc8\x08\x08\x00\x00\xff\xff\x1d\x88R\x0a"

func init() {
	schemas.Register(schema_aca73f831c7ebfdd,
		0x80b8cd5f44e3c477,
		0xd939de8c6024e7f8,
		0xf82e58b4a78f136b)
}
