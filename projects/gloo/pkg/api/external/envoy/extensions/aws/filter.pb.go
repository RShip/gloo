// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/aws/filter.proto

package aws

import (
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// AWS Lambda contains the configuration necessary to perform transform regular
// http calls to AWS Lambda invocations.
type AWSLambdaPerRoute struct {
	// The name of the function
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The qualifier of the function (defaults to $LATEST if not specified)
	Qualifier string `protobuf:"bytes,2,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	// Invocation type - async or regular.
	Async bool `protobuf:"varint,3,opt,name=async,proto3" json:"async,omitempty"`
	// Optional default body if the body is empty. By default on default
	// body is used if the body empty, and an empty body will be sent upstream.
	EmptyBodyOverride    *types.StringValue `protobuf:"bytes,4,opt,name=empty_body_override,json=emptyBodyOverride,proto3" json:"empty_body_override,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AWSLambdaPerRoute) Reset()         { *m = AWSLambdaPerRoute{} }
func (m *AWSLambdaPerRoute) String() string { return proto.CompactTextString(m) }
func (*AWSLambdaPerRoute) ProtoMessage()    {}
func (*AWSLambdaPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec6280b82440c4d3, []int{0}
}
func (m *AWSLambdaPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSLambdaPerRoute.Unmarshal(m, b)
}
func (m *AWSLambdaPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSLambdaPerRoute.Marshal(b, m, deterministic)
}
func (m *AWSLambdaPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSLambdaPerRoute.Merge(m, src)
}
func (m *AWSLambdaPerRoute) XXX_Size() int {
	return xxx_messageInfo_AWSLambdaPerRoute.Size(m)
}
func (m *AWSLambdaPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSLambdaPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_AWSLambdaPerRoute proto.InternalMessageInfo

func (m *AWSLambdaPerRoute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AWSLambdaPerRoute) GetQualifier() string {
	if m != nil {
		return m.Qualifier
	}
	return ""
}

func (m *AWSLambdaPerRoute) GetAsync() bool {
	if m != nil {
		return m.Async
	}
	return false
}

func (m *AWSLambdaPerRoute) GetEmptyBodyOverride() *types.StringValue {
	if m != nil {
		return m.EmptyBodyOverride
	}
	return nil
}

type AWSLambdaProtocolExtension struct {
	// The host header for AWS this cluster
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The region for this cluster
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// The access_key for AWS this cluster
	AccessKey string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// The secret_key for AWS this cluster
	SecretKey string `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	// The session_token for AWS this cluster
	SessionToken         string   `protobuf:"bytes,5,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWSLambdaProtocolExtension) Reset()         { *m = AWSLambdaProtocolExtension{} }
func (m *AWSLambdaProtocolExtension) String() string { return proto.CompactTextString(m) }
func (*AWSLambdaProtocolExtension) ProtoMessage()    {}
func (*AWSLambdaProtocolExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec6280b82440c4d3, []int{1}
}
func (m *AWSLambdaProtocolExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSLambdaProtocolExtension.Unmarshal(m, b)
}
func (m *AWSLambdaProtocolExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSLambdaProtocolExtension.Marshal(b, m, deterministic)
}
func (m *AWSLambdaProtocolExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSLambdaProtocolExtension.Merge(m, src)
}
func (m *AWSLambdaProtocolExtension) XXX_Size() int {
	return xxx_messageInfo_AWSLambdaProtocolExtension.Size(m)
}
func (m *AWSLambdaProtocolExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSLambdaProtocolExtension.DiscardUnknown(m)
}

var xxx_messageInfo_AWSLambdaProtocolExtension proto.InternalMessageInfo

func (m *AWSLambdaProtocolExtension) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AWSLambdaProtocolExtension) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *AWSLambdaProtocolExtension) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *AWSLambdaProtocolExtension) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *AWSLambdaProtocolExtension) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

type AWSLambdaConfig struct {
	// Use AWS default credentials chain to get credentials.
	// This will search environment variables, ECS metadata and instance metadata
	// to get the credentials. credentials will be rotated automatically.
	//
	// If credentials are provided on the cluster (using the
	// AWSLambdaProtocolExtension), it will override these credentials. This
	// defaults to false, but may change in the future to true.
	UseDefaultCredentials *types.BoolValue `protobuf:"bytes,1,opt,name=use_default_credentials,json=useDefaultCredentials,proto3" json:"use_default_credentials,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *AWSLambdaConfig) Reset()         { *m = AWSLambdaConfig{} }
func (m *AWSLambdaConfig) String() string { return proto.CompactTextString(m) }
func (*AWSLambdaConfig) ProtoMessage()    {}
func (*AWSLambdaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec6280b82440c4d3, []int{2}
}
func (m *AWSLambdaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSLambdaConfig.Unmarshal(m, b)
}
func (m *AWSLambdaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSLambdaConfig.Marshal(b, m, deterministic)
}
func (m *AWSLambdaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSLambdaConfig.Merge(m, src)
}
func (m *AWSLambdaConfig) XXX_Size() int {
	return xxx_messageInfo_AWSLambdaConfig.Size(m)
}
func (m *AWSLambdaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSLambdaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AWSLambdaConfig proto.InternalMessageInfo

func (m *AWSLambdaConfig) GetUseDefaultCredentials() *types.BoolValue {
	if m != nil {
		return m.UseDefaultCredentials
	}
	return nil
}

func init() {
	proto.RegisterType((*AWSLambdaPerRoute)(nil), "envoy.config.filter.http.aws_lambda.v2.AWSLambdaPerRoute")
	proto.RegisterType((*AWSLambdaProtocolExtension)(nil), "envoy.config.filter.http.aws_lambda.v2.AWSLambdaProtocolExtension")
	proto.RegisterType((*AWSLambdaConfig)(nil), "envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/aws/filter.proto", fileDescriptor_ec6280b82440c4d3)
}

var fileDescriptor_ec6280b82440c4d3 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0x95, 0xbd, 0xdd, 0x5e, 0xea, 0xf1, 0x47, 0x0b, 0xa0, 0x45, 0x65, 0x40, 0x55, 0x24,
	0xd4, 0x1b, 0x6c, 0xa9, 0xf0, 0x05, 0x96, 0xc1, 0xd5, 0x26, 0x31, 0x65, 0x13, 0x48, 0xdc, 0x44,
	0x4e, 0x72, 0x92, 0x9a, 0xba, 0x3e, 0xc6, 0x76, 0xda, 0xe5, 0x83, 0xf0, 0x45, 0xb8, 0xe7, 0x8b,
	0x71, 0x85, 0x62, 0xb7, 0x14, 0x69, 0x17, 0xec, 0x2e, 0xe7, 0x79, 0xce, 0x63, 0xfd, 0xce, 0xc9,
	0x21, 0xd7, 0x8d, 0x70, 0xf3, 0xb6, 0xa0, 0x25, 0x2e, 0x99, 0x45, 0x89, 0x6f, 0x04, 0xb2, 0x46,
	0x22, 0x32, 0x6d, 0xf0, 0x2b, 0x94, 0xce, 0x86, 0x8a, 0x6b, 0xc1, 0xe0, 0xc6, 0x81, 0x51, 0x5c,
	0x32, 0x50, 0x2b, 0xec, 0x7c, 0xa9, 0xac, 0x40, 0x65, 0x19, 0x5f, 0x5b, 0x56, 0x0b, 0xe9, 0xc0,
	0x50, 0x6d, 0xd0, 0x61, 0xfc, 0xda, 0xb7, 0xd0, 0x12, 0x55, 0x2d, 0x1a, 0xba, 0xb1, 0xe6, 0xce,
	0x69, 0xca, 0xd7, 0x36, 0x97, 0x7c, 0x59, 0x54, 0x9c, 0xae, 0x66, 0xa3, 0x17, 0x0d, 0x62, 0x23,
	0x81, 0xf9, 0x54, 0xd1, 0xd6, 0x6c, 0x6d, 0xb8, 0xd6, 0x60, 0x6c, 0x78, 0x67, 0x74, 0xbc, 0xe2,
	0x52, 0x54, 0xdc, 0x01, 0xdb, 0x7e, 0x04, 0x63, 0xf2, 0x23, 0x22, 0x47, 0xa7, 0x9f, 0xaf, 0x2e,
	0xfc, 0x4b, 0x97, 0x60, 0x32, 0x6c, 0x1d, 0xc4, 0xcf, 0xc8, 0x40, 0xf1, 0x25, 0x24, 0xd1, 0x38,
	0x9a, 0x0e, 0xd3, 0xff, 0x7f, 0xa5, 0x03, 0xb3, 0x37, 0x8e, 0x32, 0x2f, 0xc6, 0x27, 0x64, 0xf8,
	0xad, 0xe5, 0x52, 0xd4, 0x02, 0x4c, 0xb2, 0xd7, 0x77, 0x64, 0x3b, 0x21, 0x7e, 0x42, 0xf6, 0xb9,
	0xed, 0x54, 0x99, 0xfc, 0x37, 0x8e, 0xa6, 0xf7, 0xb2, 0x50, 0xc4, 0x17, 0xe4, 0x31, 0x2c, 0xb5,
	0xeb, 0xf2, 0x02, 0xab, 0x2e, 0xc7, 0x15, 0x18, 0x23, 0x2a, 0x48, 0x06, 0xe3, 0x68, 0x7a, 0x38,
	0x3b, 0xa1, 0x81, 0x9e, 0x6e, 0xe9, 0xe9, 0x95, 0x33, 0x42, 0x35, 0x9f, 0xb8, 0x6c, 0x21, 0x3b,
	0xf2, 0xc1, 0x14, 0xab, 0xee, 0xe3, 0x26, 0x36, 0xf9, 0x19, 0x91, 0xd1, 0x0e, 0xba, 0x0f, 0x95,
	0x28, 0x3f, 0x6c, 0xd7, 0xd8, 0xd3, 0xcf, 0xd1, 0xba, 0x5b, 0xf4, 0xbd, 0x18, 0xbf, 0x24, 0x07,
	0x06, 0x1a, 0x81, 0x2a, 0xa0, 0xef, 0xec, 0x8d, 0x1c, 0x3f, 0x27, 0x84, 0x97, 0x25, 0x58, 0x9b,
	0x2f, 0xa0, 0xf3, 0x53, 0x0c, 0xb3, 0x61, 0x50, 0xce, 0xa1, 0xeb, 0x6d, 0x0b, 0xa5, 0x01, 0xe7,
	0xed, 0x41, 0xb0, 0x83, 0xd2, 0xdb, 0xaf, 0xc8, 0x03, 0x0b, 0xb6, 0xc7, 0xc8, 0x1d, 0x2e, 0x40,
	0x25, 0xfb, 0xbe, 0xe3, 0xfe, 0x46, 0xbc, 0xee, 0xb5, 0x09, 0x90, 0x47, 0x7f, 0xf0, 0xcf, 0xfc,
	0xaf, 0x8d, 0x33, 0x72, 0xdc, 0x5a, 0xc8, 0x2b, 0xa8, 0x79, 0x2b, 0x5d, 0x5e, 0x1a, 0xa8, 0x40,
	0x39, 0xc1, 0xa5, 0xf5, 0x63, 0x1c, 0xce, 0x46, 0xb7, 0x96, 0x94, 0x22, 0xca, 0xb0, 0xa2, 0xa7,
	0xad, 0x85, 0xf7, 0x21, 0x79, 0xb6, 0x0b, 0xa6, 0xdf, 0x23, 0xf2, 0x4e, 0x20, 0xf5, 0x27, 0xa4,
	0x0d, 0xde, 0x74, 0xf4, 0x6e, 0xd7, 0x94, 0x3e, 0x3c, 0x5d, 0xdb, 0xbf, 0x96, 0x7b, 0x19, 0x7d,
	0x39, 0xbf, 0xdb, 0x75, 0xeb, 0x45, 0xf3, 0xef, 0x0b, 0x2f, 0x0e, 0xfc, 0x08, 0x6f, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xe3, 0x74, 0x3e, 0x14, 0x33, 0x03, 0x00, 0x00,
}
