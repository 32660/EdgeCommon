// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_http_access_log.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// HTTP访问日志
type HTTPAccessLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId       string              `protobuf:"bytes,48,opt,name=requestId,proto3" json:"requestId,omitempty"`
	ServerId        int64               `protobuf:"varint,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	NodeId          int64               `protobuf:"varint,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	LocationId      int64               `protobuf:"varint,3,opt,name=locationId,proto3" json:"locationId,omitempty"`
	RewriteId       int64               `protobuf:"varint,4,opt,name=rewriteId,proto3" json:"rewriteId,omitempty"`
	OriginId        int64               `protobuf:"varint,5,opt,name=originId,proto3" json:"originId,omitempty"`
	RemoteAddr      string              `protobuf:"bytes,6,opt,name=remoteAddr,proto3" json:"remoteAddr,omitempty"`
	RawRemoteAddr   string              `protobuf:"bytes,7,opt,name=rawRemoteAddr,proto3" json:"rawRemoteAddr,omitempty"`
	RemotePort      int32               `protobuf:"varint,8,opt,name=remotePort,proto3" json:"remotePort,omitempty"`
	RemoteUser      string              `protobuf:"bytes,9,opt,name=remoteUser,proto3" json:"remoteUser,omitempty"`
	RequestURI      string              `protobuf:"bytes,10,opt,name=requestURI,proto3" json:"requestURI,omitempty"`
	RequestPath     string              `protobuf:"bytes,11,opt,name=requestPath,proto3" json:"requestPath,omitempty"`
	RequestLength   int64               `protobuf:"varint,12,opt,name=requestLength,proto3" json:"requestLength,omitempty"`
	RequestTime     float64             `protobuf:"fixed64,13,opt,name=requestTime,proto3" json:"requestTime,omitempty"`
	RequestMethod   string              `protobuf:"bytes,14,opt,name=requestMethod,proto3" json:"requestMethod,omitempty"`
	RequestFilename string              `protobuf:"bytes,15,opt,name=requestFilename,proto3" json:"requestFilename,omitempty"`
	RequestBody     []byte              `protobuf:"bytes,51,opt,name=requestBody,proto3" json:"requestBody,omitempty"`
	Scheme          string              `protobuf:"bytes,16,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Proto           string              `protobuf:"bytes,17,opt,name=proto,proto3" json:"proto,omitempty"`
	BytesSent       int64               `protobuf:"varint,18,opt,name=bytesSent,proto3" json:"bytesSent,omitempty"`
	BodyBytesSent   int64               `protobuf:"varint,19,opt,name=bodyBytesSent,proto3" json:"bodyBytesSent,omitempty"`
	Status          int32               `protobuf:"varint,20,opt,name=status,proto3" json:"status,omitempty"`
	StatusMessage   string              `protobuf:"bytes,21,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	SentHeader      map[string]*Strings `protobuf:"bytes,22,rep,name=sentHeader,proto3" json:"sentHeader,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TimeISO8601     string              `protobuf:"bytes,23,opt,name=timeISO8601,proto3" json:"timeISO8601,omitempty"`
	TimeLocal       string              `protobuf:"bytes,24,opt,name=timeLocal,proto3" json:"timeLocal,omitempty"`
	Msec            float64             `protobuf:"fixed64,25,opt,name=msec,proto3" json:"msec,omitempty"`
	Timestamp       int64               `protobuf:"varint,26,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Host            string              `protobuf:"bytes,27,opt,name=host,proto3" json:"host,omitempty"`
	Referer         string              `protobuf:"bytes,28,opt,name=referer,proto3" json:"referer,omitempty"`
	UserAgent       string              `protobuf:"bytes,29,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	Request         string              `protobuf:"bytes,30,opt,name=request,proto3" json:"request,omitempty"`
	ContentType     string              `protobuf:"bytes,31,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Cookie          map[string]string   `protobuf:"bytes,32,rep,name=cookie,proto3" json:"cookie,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Args            string              `protobuf:"bytes,34,opt,name=args,proto3" json:"args,omitempty"`
	QueryString     string              `protobuf:"bytes,35,opt,name=queryString,proto3" json:"queryString,omitempty"`
	Header          map[string]*Strings `protobuf:"bytes,36,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ServerName      string              `protobuf:"bytes,37,opt,name=serverName,proto3" json:"serverName,omitempty"`
	ServerPort      int32               `protobuf:"varint,38,opt,name=serverPort,proto3" json:"serverPort,omitempty"`
	ServerProtocol  string              `protobuf:"bytes,39,opt,name=serverProtocol,proto3" json:"serverProtocol,omitempty"`
	Hostname        string              `protobuf:"bytes,40,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// 源站相关
	OriginAddress string `protobuf:"bytes,41,opt,name=originAddress,proto3" json:"originAddress,omitempty"`
	OriginStatus  int32  `protobuf:"varint,52,opt,name=originStatus,proto3" json:"originStatus,omitempty"`
	// 错误信息
	Errors []string `protobuf:"bytes,42,rep,name=errors,proto3" json:"errors,omitempty"`
	// 扩展
	Attrs map[string]string `protobuf:"bytes,43,rep,name=attrs,proto3" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// WAF相关
	FirewallPolicyId    int64    `protobuf:"varint,44,opt,name=firewallPolicyId,proto3" json:"firewallPolicyId,omitempty"`
	FirewallRuleGroupId int64    `protobuf:"varint,45,opt,name=firewallRuleGroupId,proto3" json:"firewallRuleGroupId,omitempty"`
	FirewallRuleSetId   int64    `protobuf:"varint,46,opt,name=firewallRuleSetId,proto3" json:"firewallRuleSetId,omitempty"`
	FirewallRuleId      int64    `protobuf:"varint,47,opt,name=firewallRuleId,proto3" json:"firewallRuleId,omitempty"`
	FirewallActions     []string `protobuf:"bytes,49,rep,name=firewallActions,proto3" json:"firewallActions,omitempty"`
	Tags                []string `protobuf:"bytes,50,rep,name=tags,proto3" json:"tags,omitempty"`
	// 详情
	Node *Node `protobuf:"bytes,100,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *HTTPAccessLog) Reset() {
	*x = HTTPAccessLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_http_access_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPAccessLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPAccessLog) ProtoMessage() {}

func (x *HTTPAccessLog) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_http_access_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPAccessLog.ProtoReflect.Descriptor instead.
func (*HTTPAccessLog) Descriptor() ([]byte, []int) {
	return file_models_model_http_access_log_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPAccessLog) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *HTTPAccessLog) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *HTTPAccessLog) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *HTTPAccessLog) GetLocationId() int64 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

func (x *HTTPAccessLog) GetRewriteId() int64 {
	if x != nil {
		return x.RewriteId
	}
	return 0
}

func (x *HTTPAccessLog) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

func (x *HTTPAccessLog) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *HTTPAccessLog) GetRawRemoteAddr() string {
	if x != nil {
		return x.RawRemoteAddr
	}
	return ""
}

func (x *HTTPAccessLog) GetRemotePort() int32 {
	if x != nil {
		return x.RemotePort
	}
	return 0
}

func (x *HTTPAccessLog) GetRemoteUser() string {
	if x != nil {
		return x.RemoteUser
	}
	return ""
}

func (x *HTTPAccessLog) GetRequestURI() string {
	if x != nil {
		return x.RequestURI
	}
	return ""
}

func (x *HTTPAccessLog) GetRequestPath() string {
	if x != nil {
		return x.RequestPath
	}
	return ""
}

func (x *HTTPAccessLog) GetRequestLength() int64 {
	if x != nil {
		return x.RequestLength
	}
	return 0
}

func (x *HTTPAccessLog) GetRequestTime() float64 {
	if x != nil {
		return x.RequestTime
	}
	return 0
}

func (x *HTTPAccessLog) GetRequestMethod() string {
	if x != nil {
		return x.RequestMethod
	}
	return ""
}

func (x *HTTPAccessLog) GetRequestFilename() string {
	if x != nil {
		return x.RequestFilename
	}
	return ""
}

func (x *HTTPAccessLog) GetRequestBody() []byte {
	if x != nil {
		return x.RequestBody
	}
	return nil
}

func (x *HTTPAccessLog) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *HTTPAccessLog) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *HTTPAccessLog) GetBytesSent() int64 {
	if x != nil {
		return x.BytesSent
	}
	return 0
}

func (x *HTTPAccessLog) GetBodyBytesSent() int64 {
	if x != nil {
		return x.BodyBytesSent
	}
	return 0
}

func (x *HTTPAccessLog) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HTTPAccessLog) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *HTTPAccessLog) GetSentHeader() map[string]*Strings {
	if x != nil {
		return x.SentHeader
	}
	return nil
}

func (x *HTTPAccessLog) GetTimeISO8601() string {
	if x != nil {
		return x.TimeISO8601
	}
	return ""
}

func (x *HTTPAccessLog) GetTimeLocal() string {
	if x != nil {
		return x.TimeLocal
	}
	return ""
}

func (x *HTTPAccessLog) GetMsec() float64 {
	if x != nil {
		return x.Msec
	}
	return 0
}

func (x *HTTPAccessLog) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *HTTPAccessLog) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *HTTPAccessLog) GetReferer() string {
	if x != nil {
		return x.Referer
	}
	return ""
}

func (x *HTTPAccessLog) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *HTTPAccessLog) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *HTTPAccessLog) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *HTTPAccessLog) GetCookie() map[string]string {
	if x != nil {
		return x.Cookie
	}
	return nil
}

func (x *HTTPAccessLog) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *HTTPAccessLog) GetQueryString() string {
	if x != nil {
		return x.QueryString
	}
	return ""
}

func (x *HTTPAccessLog) GetHeader() map[string]*Strings {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HTTPAccessLog) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *HTTPAccessLog) GetServerPort() int32 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

func (x *HTTPAccessLog) GetServerProtocol() string {
	if x != nil {
		return x.ServerProtocol
	}
	return ""
}

func (x *HTTPAccessLog) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *HTTPAccessLog) GetOriginAddress() string {
	if x != nil {
		return x.OriginAddress
	}
	return ""
}

func (x *HTTPAccessLog) GetOriginStatus() int32 {
	if x != nil {
		return x.OriginStatus
	}
	return 0
}

func (x *HTTPAccessLog) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *HTTPAccessLog) GetAttrs() map[string]string {
	if x != nil {
		return x.Attrs
	}
	return nil
}

func (x *HTTPAccessLog) GetFirewallPolicyId() int64 {
	if x != nil {
		return x.FirewallPolicyId
	}
	return 0
}

func (x *HTTPAccessLog) GetFirewallRuleGroupId() int64 {
	if x != nil {
		return x.FirewallRuleGroupId
	}
	return 0
}

func (x *HTTPAccessLog) GetFirewallRuleSetId() int64 {
	if x != nil {
		return x.FirewallRuleSetId
	}
	return 0
}

func (x *HTTPAccessLog) GetFirewallRuleId() int64 {
	if x != nil {
		return x.FirewallRuleId
	}
	return 0
}

func (x *HTTPAccessLog) GetFirewallActions() []string {
	if x != nil {
		return x.FirewallActions
	}
	return nil
}

func (x *HTTPAccessLog) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *HTTPAccessLog) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

type Strings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Strings) Reset() {
	*x = Strings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_http_access_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strings) ProtoMessage() {}

func (x *Strings) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_http_access_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strings.ProtoReflect.Descriptor instead.
func (*Strings) Descriptor() ([]byte, []int) {
	return file_models_model_http_access_log_proto_rawDescGZIP(), []int{1}
}

func (x *Strings) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_models_model_http_access_log_proto protoreflect.FileDescriptor

var file_models_model_http_access_log_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x81, 0x10, 0x0a, 0x0d, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x6f, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x30, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x24, 0x0a, 0x0d, 0x72, 0x61, 0x77, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x61, 0x77, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x55, 0x52, 0x49, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x55, 0x52, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x33,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x41, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x53, 0x4f, 0x38,
	0x36, 0x30, 0x31, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x49,
	0x53, 0x4f, 0x38, 0x36, 0x30, 0x31, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x73, 0x65, 0x63, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x6d, 0x73, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x72, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35,
	0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x20, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63,
	0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x35, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x24, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62,
	0x2e, 0x48, 0x54, 0x54, 0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x18, 0x26, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x34, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x2a, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x61, 0x74, 0x74, 0x72,
	0x73, 0x18, 0x2b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54,
	0x50, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x10,
	0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64,
	0x18, 0x2c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x2d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x69,
	0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x18,
	0x2e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x2f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x0f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x31, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x32, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1c,
	0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x4a, 0x0a, 0x0f,
	0x53, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x46, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x41,
	0x74, 0x74, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x21, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_http_access_log_proto_rawDescOnce sync.Once
	file_models_model_http_access_log_proto_rawDescData = file_models_model_http_access_log_proto_rawDesc
)

func file_models_model_http_access_log_proto_rawDescGZIP() []byte {
	file_models_model_http_access_log_proto_rawDescOnce.Do(func() {
		file_models_model_http_access_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_http_access_log_proto_rawDescData)
	})
	return file_models_model_http_access_log_proto_rawDescData
}

var file_models_model_http_access_log_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_models_model_http_access_log_proto_goTypes = []interface{}{
	(*HTTPAccessLog)(nil), // 0: pb.HTTPAccessLog
	(*Strings)(nil),       // 1: pb.Strings
	nil,                   // 2: pb.HTTPAccessLog.SentHeaderEntry
	nil,                   // 3: pb.HTTPAccessLog.CookieEntry
	nil,                   // 4: pb.HTTPAccessLog.HeaderEntry
	nil,                   // 5: pb.HTTPAccessLog.AttrsEntry
	(*Node)(nil),          // 6: pb.Node
}
var file_models_model_http_access_log_proto_depIdxs = []int32{
	2, // 0: pb.HTTPAccessLog.sentHeader:type_name -> pb.HTTPAccessLog.SentHeaderEntry
	3, // 1: pb.HTTPAccessLog.cookie:type_name -> pb.HTTPAccessLog.CookieEntry
	4, // 2: pb.HTTPAccessLog.header:type_name -> pb.HTTPAccessLog.HeaderEntry
	5, // 3: pb.HTTPAccessLog.attrs:type_name -> pb.HTTPAccessLog.AttrsEntry
	6, // 4: pb.HTTPAccessLog.node:type_name -> pb.Node
	1, // 5: pb.HTTPAccessLog.SentHeaderEntry.value:type_name -> pb.Strings
	1, // 6: pb.HTTPAccessLog.HeaderEntry.value:type_name -> pb.Strings
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_models_model_http_access_log_proto_init() }
func file_models_model_http_access_log_proto_init() {
	if File_models_model_http_access_log_proto != nil {
		return
	}
	file_models_model_node_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_http_access_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPAccessLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_models_model_http_access_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_model_http_access_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_http_access_log_proto_goTypes,
		DependencyIndexes: file_models_model_http_access_log_proto_depIdxs,
		MessageInfos:      file_models_model_http_access_log_proto_msgTypes,
	}.Build()
	File_models_model_http_access_log_proto = out.File
	file_models_model_http_access_log_proto_rawDesc = nil
	file_models_model_http_access_log_proto_goTypes = nil
	file_models_model_http_access_log_proto_depIdxs = nil
}
