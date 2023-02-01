// Code generated by protoc-gen-go. DO NOT EDIT.
// source: install.proto

/*
Package install is a generated protocol buffer package.

It is generated from these files:
	install.proto

It has these top-level messages:
	InstallConfig
	ProxyConfig
	TLSSelfSigned
	TLSLetsEncrypt
	TLSCertificate
	CheckResult
	PerformCheckRequest
	PerformCheckResponse
	GetDefaultsRequest
	GetDefaultsResponse
	GetAgreementRequest
	GetAgreementResponse
	InstallRequest
	InstallResponse
*/
package install

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InstallConfig struct {
	InternalUrl              string         `protobuf:"bytes,32,opt,name=internalUrl" json:"internalUrl,omitempty"`
	DbConnectionType         string         `protobuf:"bytes,1,opt,name=dbConnectionType" json:"dbConnectionType,omitempty"`
	DbTCPHostname            string         `protobuf:"bytes,2,opt,name=dbTCPHostname" json:"dbTCPHostname,omitempty"`
	DbTCPPort                string         `protobuf:"bytes,3,opt,name=dbTCPPort" json:"dbTCPPort,omitempty"`
	DbTCPName                string         `protobuf:"bytes,4,opt,name=dbTCPName" json:"dbTCPName,omitempty"`
	DbTCPUser                string         `protobuf:"bytes,5,opt,name=dbTCPUser" json:"dbTCPUser,omitempty"`
	DbTCPPassword            string         `protobuf:"bytes,6,opt,name=dbTCPPassword" json:"dbTCPPassword,omitempty"`
	DbSocketFile             string         `protobuf:"bytes,7,opt,name=dbSocketFile" json:"dbSocketFile,omitempty"`
	DbSocketName             string         `protobuf:"bytes,8,opt,name=dbSocketName" json:"dbSocketName,omitempty"`
	DbSocketUser             string         `protobuf:"bytes,9,opt,name=dbSocketUser" json:"dbSocketUser,omitempty"`
	DbSocketPassword         string         `protobuf:"bytes,10,opt,name=dbSocketPassword" json:"dbSocketPassword,omitempty"`
	DbManualDSN              string         `protobuf:"bytes,11,opt,name=dbManualDSN" json:"dbManualDSN,omitempty"`
	DsName                   string         `protobuf:"bytes,12,opt,name=dsName" json:"dsName,omitempty"`
	DsPort                   string         `protobuf:"bytes,13,opt,name=dsPort" json:"dsPort,omitempty"`
	DsType                   string         `protobuf:"bytes,15,opt,name=dsType" json:"dsType,omitempty"`
	DsS3Custom               string         `protobuf:"bytes,16,opt,name=dsS3Custom" json:"dsS3Custom,omitempty"`
	DsS3CustomRegion         string         `protobuf:"bytes,17,opt,name=dsS3CustomRegion" json:"dsS3CustomRegion,omitempty"`
	DsS3ApiKey               string         `protobuf:"bytes,18,opt,name=dsS3ApiKey" json:"dsS3ApiKey,omitempty"`
	DsS3ApiSecret            string         `protobuf:"bytes,19,opt,name=dsS3ApiSecret" json:"dsS3ApiSecret,omitempty"`
	DsS3BucketDefault        string         `protobuf:"bytes,20,opt,name=dsS3BucketDefault" json:"dsS3BucketDefault,omitempty"`
	DsS3BucketPersonal       string         `protobuf:"bytes,21,opt,name=dsS3BucketPersonal" json:"dsS3BucketPersonal,omitempty"`
	DsS3BucketCells          string         `protobuf:"bytes,22,opt,name=dsS3BucketCells" json:"dsS3BucketCells,omitempty"`
	DsS3BucketBinaries       string         `protobuf:"bytes,23,opt,name=dsS3BucketBinaries" json:"dsS3BucketBinaries,omitempty"`
	DsS3BucketThumbs         string         `protobuf:"bytes,35,opt,name=dsS3BucketThumbs" json:"dsS3BucketThumbs,omitempty"`
	DsS3BucketVersions       string         `protobuf:"bytes,36,opt,name=dsS3BucketVersions" json:"dsS3BucketVersions,omitempty"`
	DsFolder                 string         `protobuf:"bytes,14,opt,name=dsFolder" json:"dsFolder,omitempty"`
	FrontendHosts            string         `protobuf:"bytes,24,opt,name=frontendHosts" json:"frontendHosts,omitempty"`
	FrontendLogin            string         `protobuf:"bytes,25,opt,name=frontendLogin" json:"frontendLogin,omitempty"`
	FrontendPassword         string         `protobuf:"bytes,26,opt,name=frontendPassword" json:"frontendPassword,omitempty"`
	FrontendRepeatPassword   string         `protobuf:"bytes,27,opt,name=frontendRepeatPassword" json:"frontendRepeatPassword,omitempty"`
	FrontendApplicationTitle string         `protobuf:"bytes,28,opt,name=frontendApplicationTitle" json:"frontendApplicationTitle,omitempty"`
	FrontendDefaultLanguage  string         `protobuf:"bytes,33,opt,name=frontendDefaultLanguage" json:"frontendDefaultLanguage,omitempty"`
	LicenseRequired          bool           `protobuf:"varint,29,opt,name=licenseRequired" json:"licenseRequired,omitempty"`
	LicenseString            string         `protobuf:"bytes,30,opt,name=licenseString" json:"licenseString,omitempty"`
	CheckResults             []*CheckResult `protobuf:"bytes,31,rep,name=CheckResults" json:"CheckResults,omitempty"`
	// Additional proxy config (optional)
	ProxyConfig *ProxyConfig `protobuf:"bytes,34,opt,name=ProxyConfig" json:"ProxyConfig,omitempty"`
}

func (m *InstallConfig) Reset()                    { *m = InstallConfig{} }
func (m *InstallConfig) String() string            { return proto.CompactTextString(m) }
func (*InstallConfig) ProtoMessage()               {}
func (*InstallConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InstallConfig) GetInternalUrl() string {
	if m != nil {
		return m.InternalUrl
	}
	return ""
}

func (m *InstallConfig) GetDbConnectionType() string {
	if m != nil {
		return m.DbConnectionType
	}
	return ""
}

func (m *InstallConfig) GetDbTCPHostname() string {
	if m != nil {
		return m.DbTCPHostname
	}
	return ""
}

func (m *InstallConfig) GetDbTCPPort() string {
	if m != nil {
		return m.DbTCPPort
	}
	return ""
}

func (m *InstallConfig) GetDbTCPName() string {
	if m != nil {
		return m.DbTCPName
	}
	return ""
}

func (m *InstallConfig) GetDbTCPUser() string {
	if m != nil {
		return m.DbTCPUser
	}
	return ""
}

func (m *InstallConfig) GetDbTCPPassword() string {
	if m != nil {
		return m.DbTCPPassword
	}
	return ""
}

func (m *InstallConfig) GetDbSocketFile() string {
	if m != nil {
		return m.DbSocketFile
	}
	return ""
}

func (m *InstallConfig) GetDbSocketName() string {
	if m != nil {
		return m.DbSocketName
	}
	return ""
}

func (m *InstallConfig) GetDbSocketUser() string {
	if m != nil {
		return m.DbSocketUser
	}
	return ""
}

func (m *InstallConfig) GetDbSocketPassword() string {
	if m != nil {
		return m.DbSocketPassword
	}
	return ""
}

func (m *InstallConfig) GetDbManualDSN() string {
	if m != nil {
		return m.DbManualDSN
	}
	return ""
}

func (m *InstallConfig) GetDsName() string {
	if m != nil {
		return m.DsName
	}
	return ""
}

func (m *InstallConfig) GetDsPort() string {
	if m != nil {
		return m.DsPort
	}
	return ""
}

func (m *InstallConfig) GetDsType() string {
	if m != nil {
		return m.DsType
	}
	return ""
}

func (m *InstallConfig) GetDsS3Custom() string {
	if m != nil {
		return m.DsS3Custom
	}
	return ""
}

func (m *InstallConfig) GetDsS3CustomRegion() string {
	if m != nil {
		return m.DsS3CustomRegion
	}
	return ""
}

func (m *InstallConfig) GetDsS3ApiKey() string {
	if m != nil {
		return m.DsS3ApiKey
	}
	return ""
}

func (m *InstallConfig) GetDsS3ApiSecret() string {
	if m != nil {
		return m.DsS3ApiSecret
	}
	return ""
}

func (m *InstallConfig) GetDsS3BucketDefault() string {
	if m != nil {
		return m.DsS3BucketDefault
	}
	return ""
}

func (m *InstallConfig) GetDsS3BucketPersonal() string {
	if m != nil {
		return m.DsS3BucketPersonal
	}
	return ""
}

func (m *InstallConfig) GetDsS3BucketCells() string {
	if m != nil {
		return m.DsS3BucketCells
	}
	return ""
}

func (m *InstallConfig) GetDsS3BucketBinaries() string {
	if m != nil {
		return m.DsS3BucketBinaries
	}
	return ""
}

func (m *InstallConfig) GetDsS3BucketThumbs() string {
	if m != nil {
		return m.DsS3BucketThumbs
	}
	return ""
}

func (m *InstallConfig) GetDsS3BucketVersions() string {
	if m != nil {
		return m.DsS3BucketVersions
	}
	return ""
}

func (m *InstallConfig) GetDsFolder() string {
	if m != nil {
		return m.DsFolder
	}
	return ""
}

func (m *InstallConfig) GetFrontendHosts() string {
	if m != nil {
		return m.FrontendHosts
	}
	return ""
}

func (m *InstallConfig) GetFrontendLogin() string {
	if m != nil {
		return m.FrontendLogin
	}
	return ""
}

func (m *InstallConfig) GetFrontendPassword() string {
	if m != nil {
		return m.FrontendPassword
	}
	return ""
}

func (m *InstallConfig) GetFrontendRepeatPassword() string {
	if m != nil {
		return m.FrontendRepeatPassword
	}
	return ""
}

func (m *InstallConfig) GetFrontendApplicationTitle() string {
	if m != nil {
		return m.FrontendApplicationTitle
	}
	return ""
}

func (m *InstallConfig) GetFrontendDefaultLanguage() string {
	if m != nil {
		return m.FrontendDefaultLanguage
	}
	return ""
}

func (m *InstallConfig) GetLicenseRequired() bool {
	if m != nil {
		return m.LicenseRequired
	}
	return false
}

func (m *InstallConfig) GetLicenseString() string {
	if m != nil {
		return m.LicenseString
	}
	return ""
}

func (m *InstallConfig) GetCheckResults() []*CheckResult {
	if m != nil {
		return m.CheckResults
	}
	return nil
}

func (m *InstallConfig) GetProxyConfig() *ProxyConfig {
	if m != nil {
		return m.ProxyConfig
	}
	return nil
}

// ProxyConfig gives necessary URL and TLS configurations to start proxy
type ProxyConfig struct {
	BindURL      string   `protobuf:"bytes,1,opt,name=BindURL" json:"BindURL,omitempty"`
	ExternalURL  string   `protobuf:"bytes,2,opt,name=ExternalURL" json:"ExternalURL,omitempty"`
	RedirectURLs []string `protobuf:"bytes,3,rep,name=RedirectURLs" json:"RedirectURLs,omitempty"`
	// Types that are valid to be assigned to TLSConfig:
	//	*ProxyConfig_SelfSigned
	//	*ProxyConfig_LetsEncrypt
	//	*ProxyConfig_Certificate
	TLSConfig isProxyConfig_TLSConfig `protobuf_oneof:"TLSConfig"`
}

func (m *ProxyConfig) Reset()                    { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string            { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()               {}
func (*ProxyConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isProxyConfig_TLSConfig interface{ isProxyConfig_TLSConfig() }

type ProxyConfig_SelfSigned struct {
	SelfSigned *TLSSelfSigned `protobuf:"bytes,4,opt,name=SelfSigned,oneof"`
}
type ProxyConfig_LetsEncrypt struct {
	LetsEncrypt *TLSLetsEncrypt `protobuf:"bytes,5,opt,name=LetsEncrypt,oneof"`
}
type ProxyConfig_Certificate struct {
	Certificate *TLSCertificate `protobuf:"bytes,6,opt,name=Certificate,oneof"`
}

func (*ProxyConfig_SelfSigned) isProxyConfig_TLSConfig()  {}
func (*ProxyConfig_LetsEncrypt) isProxyConfig_TLSConfig() {}
func (*ProxyConfig_Certificate) isProxyConfig_TLSConfig() {}

func (m *ProxyConfig) GetTLSConfig() isProxyConfig_TLSConfig {
	if m != nil {
		return m.TLSConfig
	}
	return nil
}

func (m *ProxyConfig) GetBindURL() string {
	if m != nil {
		return m.BindURL
	}
	return ""
}

func (m *ProxyConfig) GetExternalURL() string {
	if m != nil {
		return m.ExternalURL
	}
	return ""
}

func (m *ProxyConfig) GetRedirectURLs() []string {
	if m != nil {
		return m.RedirectURLs
	}
	return nil
}

func (m *ProxyConfig) GetSelfSigned() *TLSSelfSigned {
	if x, ok := m.GetTLSConfig().(*ProxyConfig_SelfSigned); ok {
		return x.SelfSigned
	}
	return nil
}

func (m *ProxyConfig) GetLetsEncrypt() *TLSLetsEncrypt {
	if x, ok := m.GetTLSConfig().(*ProxyConfig_LetsEncrypt); ok {
		return x.LetsEncrypt
	}
	return nil
}

func (m *ProxyConfig) GetCertificate() *TLSCertificate {
	if x, ok := m.GetTLSConfig().(*ProxyConfig_Certificate); ok {
		return x.Certificate
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ProxyConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ProxyConfig_OneofMarshaler, _ProxyConfig_OneofUnmarshaler, _ProxyConfig_OneofSizer, []interface{}{
		(*ProxyConfig_SelfSigned)(nil),
		(*ProxyConfig_LetsEncrypt)(nil),
		(*ProxyConfig_Certificate)(nil),
	}
}

func _ProxyConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ProxyConfig)
	// TLSConfig
	switch x := m.TLSConfig.(type) {
	case *ProxyConfig_SelfSigned:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SelfSigned); err != nil {
			return err
		}
	case *ProxyConfig_LetsEncrypt:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LetsEncrypt); err != nil {
			return err
		}
	case *ProxyConfig_Certificate:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Certificate); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ProxyConfig.TLSConfig has unexpected type %T", x)
	}
	return nil
}

func _ProxyConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ProxyConfig)
	switch tag {
	case 4: // TLSConfig.SelfSigned
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TLSSelfSigned)
		err := b.DecodeMessage(msg)
		m.TLSConfig = &ProxyConfig_SelfSigned{msg}
		return true, err
	case 5: // TLSConfig.LetsEncrypt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TLSLetsEncrypt)
		err := b.DecodeMessage(msg)
		m.TLSConfig = &ProxyConfig_LetsEncrypt{msg}
		return true, err
	case 6: // TLSConfig.Certificate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TLSCertificate)
		err := b.DecodeMessage(msg)
		m.TLSConfig = &ProxyConfig_Certificate{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ProxyConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ProxyConfig)
	// TLSConfig
	switch x := m.TLSConfig.(type) {
	case *ProxyConfig_SelfSigned:
		s := proto.Size(x.SelfSigned)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProxyConfig_LetsEncrypt:
		s := proto.Size(x.LetsEncrypt)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProxyConfig_Certificate:
		s := proto.Size(x.Certificate)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// TLSSelfSigned generates a selfsigned certificate
type TLSSelfSigned struct {
	Hostnames []string `protobuf:"bytes,1,rep,name=Hostnames" json:"Hostnames,omitempty"`
}

func (m *TLSSelfSigned) Reset()                    { *m = TLSSelfSigned{} }
func (m *TLSSelfSigned) String() string            { return proto.CompactTextString(m) }
func (*TLSSelfSigned) ProtoMessage()               {}
func (*TLSSelfSigned) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TLSSelfSigned) GetHostnames() []string {
	if m != nil {
		return m.Hostnames
	}
	return nil
}

// TLSLetsEncrypt set up proxy to automatically get a valid certificate from let's encrypt servers
type TLSLetsEncrypt struct {
	Email      string `protobuf:"bytes,1,opt,name=Email" json:"Email,omitempty"`
	AcceptEULA bool   `protobuf:"varint,2,opt,name=AcceptEULA" json:"AcceptEULA,omitempty"`
	StagingCA  bool   `protobuf:"varint,3,opt,name=StagingCA" json:"StagingCA,omitempty"`
}

func (m *TLSLetsEncrypt) Reset()                    { *m = TLSLetsEncrypt{} }
func (m *TLSLetsEncrypt) String() string            { return proto.CompactTextString(m) }
func (*TLSLetsEncrypt) ProtoMessage()               {}
func (*TLSLetsEncrypt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TLSLetsEncrypt) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *TLSLetsEncrypt) GetAcceptEULA() bool {
	if m != nil {
		return m.AcceptEULA
	}
	return false
}

func (m *TLSLetsEncrypt) GetStagingCA() bool {
	if m != nil {
		return m.StagingCA
	}
	return false
}

// TLSCertificate is a TLSConfig where user passes
type TLSCertificate struct {
	CertFile string `protobuf:"bytes,1,opt,name=CertFile" json:"CertFile,omitempty"`
	KeyFile  string `protobuf:"bytes,2,opt,name=KeyFile" json:"KeyFile,omitempty"`
}

func (m *TLSCertificate) Reset()                    { *m = TLSCertificate{} }
func (m *TLSCertificate) String() string            { return proto.CompactTextString(m) }
func (*TLSCertificate) ProtoMessage()               {}
func (*TLSCertificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TLSCertificate) GetCertFile() string {
	if m != nil {
		return m.CertFile
	}
	return ""
}

func (m *TLSCertificate) GetKeyFile() string {
	if m != nil {
		return m.KeyFile
	}
	return ""
}

type CheckResult struct {
	Name       string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Success    bool   `protobuf:"varint,2,opt,name=Success" json:"Success,omitempty"`
	JsonResult string `protobuf:"bytes,3,opt,name=JsonResult" json:"JsonResult,omitempty"`
}

func (m *CheckResult) Reset()                    { *m = CheckResult{} }
func (m *CheckResult) String() string            { return proto.CompactTextString(m) }
func (*CheckResult) ProtoMessage()               {}
func (*CheckResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CheckResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CheckResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CheckResult) GetJsonResult() string {
	if m != nil {
		return m.JsonResult
	}
	return ""
}

type PerformCheckRequest struct {
	Name   string         `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Config *InstallConfig `protobuf:"bytes,2,opt,name=Config" json:"Config,omitempty"`
}

func (m *PerformCheckRequest) Reset()                    { *m = PerformCheckRequest{} }
func (m *PerformCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*PerformCheckRequest) ProtoMessage()               {}
func (*PerformCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PerformCheckRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PerformCheckRequest) GetConfig() *InstallConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type PerformCheckResponse struct {
	Result *CheckResult `protobuf:"bytes,1,opt,name=Result" json:"Result,omitempty"`
}

func (m *PerformCheckResponse) Reset()                    { *m = PerformCheckResponse{} }
func (m *PerformCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*PerformCheckResponse) ProtoMessage()               {}
func (*PerformCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PerformCheckResponse) GetResult() *CheckResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type GetDefaultsRequest struct {
}

func (m *GetDefaultsRequest) Reset()                    { *m = GetDefaultsRequest{} }
func (m *GetDefaultsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDefaultsRequest) ProtoMessage()               {}
func (*GetDefaultsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type GetDefaultsResponse struct {
	Config *InstallConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
}

func (m *GetDefaultsResponse) Reset()                    { *m = GetDefaultsResponse{} }
func (m *GetDefaultsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDefaultsResponse) ProtoMessage()               {}
func (*GetDefaultsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetDefaultsResponse) GetConfig() *InstallConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type GetAgreementRequest struct {
}

func (m *GetAgreementRequest) Reset()                    { *m = GetAgreementRequest{} }
func (m *GetAgreementRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAgreementRequest) ProtoMessage()               {}
func (*GetAgreementRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type GetAgreementResponse struct {
	Text string `protobuf:"bytes,1,opt,name=Text" json:"Text,omitempty"`
}

func (m *GetAgreementResponse) Reset()                    { *m = GetAgreementResponse{} }
func (m *GetAgreementResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAgreementResponse) ProtoMessage()               {}
func (*GetAgreementResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetAgreementResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type InstallRequest struct {
	Config *InstallConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
}

func (m *InstallRequest) Reset()                    { *m = InstallRequest{} }
func (m *InstallRequest) String() string            { return proto.CompactTextString(m) }
func (*InstallRequest) ProtoMessage()               {}
func (*InstallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *InstallRequest) GetConfig() *InstallConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type InstallResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *InstallResponse) Reset()                    { *m = InstallResponse{} }
func (m *InstallResponse) String() string            { return proto.CompactTextString(m) }
func (*InstallResponse) ProtoMessage()               {}
func (*InstallResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *InstallResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*InstallConfig)(nil), "install.InstallConfig")
	proto.RegisterType((*ProxyConfig)(nil), "install.ProxyConfig")
	proto.RegisterType((*TLSSelfSigned)(nil), "install.TLSSelfSigned")
	proto.RegisterType((*TLSLetsEncrypt)(nil), "install.TLSLetsEncrypt")
	proto.RegisterType((*TLSCertificate)(nil), "install.TLSCertificate")
	proto.RegisterType((*CheckResult)(nil), "install.CheckResult")
	proto.RegisterType((*PerformCheckRequest)(nil), "install.PerformCheckRequest")
	proto.RegisterType((*PerformCheckResponse)(nil), "install.PerformCheckResponse")
	proto.RegisterType((*GetDefaultsRequest)(nil), "install.GetDefaultsRequest")
	proto.RegisterType((*GetDefaultsResponse)(nil), "install.GetDefaultsResponse")
	proto.RegisterType((*GetAgreementRequest)(nil), "install.GetAgreementRequest")
	proto.RegisterType((*GetAgreementResponse)(nil), "install.GetAgreementResponse")
	proto.RegisterType((*InstallRequest)(nil), "install.InstallRequest")
	proto.RegisterType((*InstallResponse)(nil), "install.InstallResponse")
}

func init() { proto.RegisterFile("install.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1052 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x5f, 0x53, 0xdb, 0x46,
	0x10, 0x8f, 0x81, 0x18, 0x58, 0xf1, 0x27, 0x39, 0x1c, 0xb8, 0x12, 0x92, 0xba, 0x6a, 0x1e, 0x98,
	0x34, 0xe5, 0x81, 0xcc, 0x64, 0x98, 0xb6, 0x0f, 0x35, 0x06, 0x4a, 0x8b, 0xcb, 0x78, 0x64, 0xe8,
	0x4c, 0xa7, 0x4f, 0xb2, 0xb4, 0x76, 0x6e, 0x22, 0x9f, 0x1c, 0xdd, 0x79, 0x8a, 0xbf, 0x54, 0xfb,
	0x8d, 0xfa, 0x59, 0x3a, 0xb7, 0x3a, 0x49, 0x27, 0x03, 0x9d, 0xe9, 0x9b, 0xf6, 0xf7, 0xfb, 0xed,
	0x1f, 0xed, 0x9e, 0xf6, 0x04, 0x9b, 0x42, 0x2a, 0x1d, 0x26, 0xc9, 0xd1, 0x34, 0x4b, 0x75, 0xca,
	0x56, 0xad, 0xe9, 0xff, 0xed, 0xc1, 0xe6, 0xcf, 0xf9, 0x73, 0x37, 0x95, 0x23, 0x31, 0x66, 0x6d,
	0xf0, 0x84, 0xd4, 0x98, 0xc9, 0x30, 0xb9, 0xcd, 0x12, 0xde, 0x6e, 0x37, 0x0e, 0xd7, 0x03, 0x17,
	0x62, 0x6f, 0xe1, 0x59, 0x3c, 0xec, 0xa6, 0x52, 0x62, 0xa4, 0x45, 0x2a, 0x6f, 0xe6, 0x53, 0xe4,
	0x0d, 0x92, 0xdd, 0xc3, 0xd9, 0x1b, 0xd8, 0x8c, 0x87, 0x37, 0xdd, 0xfe, 0x65, 0xaa, 0xb4, 0x0c,
	0x27, 0xc8, 0x97, 0x48, 0x58, 0x07, 0xd9, 0x01, 0xac, 0x13, 0xd0, 0x4f, 0x33, 0xcd, 0x97, 0x49,
	0x51, 0x01, 0x25, 0x7b, 0x6d, 0xfc, 0x57, 0x1c, 0xf6, 0xda, 0xf5, 0xbd, 0x55, 0x98, 0xf1, 0xa7,
	0x0e, 0x6b, 0x80, 0x32, 0x7f, 0x3f, 0x54, 0xea, 0xcf, 0x34, 0x8b, 0x79, 0xd3, 0xc9, 0x5f, 0x80,
	0xcc, 0x87, 0x8d, 0x78, 0x38, 0x48, 0xa3, 0x4f, 0xa8, 0x2f, 0x44, 0x82, 0x7c, 0x95, 0x44, 0x35,
	0xcc, 0xd5, 0x50, 0x21, 0x6b, 0x75, 0x0d, 0xd5, 0xe2, 0x68, 0xa8, 0x9c, 0xf5, 0xba, 0x86, 0x2a,
	0xa2, 0xee, 0xe5, 0x76, 0x59, 0x14, 0x14, 0xdd, 0xab, 0xe3, 0x66, 0x16, 0xf1, 0xf0, 0xd7, 0x50,
	0xce, 0xc2, 0xe4, 0x6c, 0x70, 0xcd, 0xbd, 0x7c, 0x16, 0x0e, 0xc4, 0x76, 0xa1, 0x19, 0x2b, 0xaa,
	0x67, 0x83, 0x48, 0x6b, 0xe5, 0x38, 0xb5, 0x73, 0xb3, 0xc0, 0xa9, 0x97, 0x84, 0xd3, 0xc4, 0xb6,
	0x0b, 0x9c, 0xe6, 0xf4, 0x1a, 0x20, 0x56, 0x83, 0xf7, 0xdd, 0x99, 0xd2, 0xe9, 0x84, 0x3f, 0x23,
	0xce, 0x41, 0xa8, 0xea, 0xd2, 0x0a, 0x70, 0x2c, 0x52, 0xc9, 0x9f, 0xdb, 0xaa, 0x17, 0xf0, 0x22,
	0x56, 0x67, 0x2a, 0xae, 0x70, 0xce, 0x59, 0x15, 0x2b, 0x47, 0x68, 0x26, 0xb9, 0x35, 0xc0, 0x28,
	0x43, 0xcd, 0x77, 0xec, 0x4c, 0x5c, 0x90, 0xbd, 0x83, 0xe7, 0x06, 0x38, 0x9d, 0x99, 0x8e, 0x9c,
	0xe1, 0x28, 0x9c, 0x25, 0x9a, 0xb7, 0x48, 0x79, 0x9f, 0x60, 0x47, 0xc0, 0x2a, 0xb0, 0x8f, 0x99,
	0x4a, 0x65, 0x98, 0xf0, 0x17, 0x24, 0x7f, 0x80, 0x61, 0x87, 0xb0, 0x5d, 0xa1, 0x5d, 0x4c, 0x12,
	0xc5, 0x77, 0x49, 0xbc, 0x08, 0xd7, 0x23, 0x9f, 0x0a, 0x19, 0x66, 0x02, 0x15, 0xdf, 0x5b, 0x8c,
	0x5c, 0x30, 0x45, 0xa7, 0x72, 0xf4, 0xe6, 0xe3, 0x6c, 0x32, 0x54, 0xfc, 0xeb, 0xaa, 0x53, 0x2e,
	0x5e, 0x8f, 0xfd, 0x1b, 0x66, 0x4a, 0xa4, 0x52, 0xf1, 0x37, 0x8b, 0xb1, 0x0b, 0x86, 0xed, 0xc3,
	0x5a, 0xac, 0x2e, 0xd2, 0x24, 0xc6, 0x8c, 0x6f, 0x91, 0xaa, 0xb4, 0x4d, 0x57, 0x47, 0x59, 0x2a,
	0x35, 0xca, 0xd8, 0x7c, 0x57, 0x8a, 0xf3, 0xbc, 0xab, 0x35, 0xd0, 0x55, 0xf5, 0xd2, 0xb1, 0x90,
	0xfc, 0x8b, 0xba, 0x8a, 0x40, 0xf3, 0x0e, 0x05, 0x50, 0x9e, 0xd1, 0xfd, 0xfc, 0x1d, 0x16, 0x71,
	0xf6, 0x01, 0x76, 0x0b, 0x2c, 0xc0, 0x29, 0x86, 0xd5, 0xa9, 0x7e, 0x49, 0x1e, 0x8f, 0xb0, 0xec,
	0x3b, 0xe0, 0x05, 0xd3, 0x99, 0x4e, 0x13, 0x11, 0x85, 0xb4, 0x34, 0x84, 0x4e, 0x90, 0x1f, 0x90,
	0xe7, 0xa3, 0x3c, 0x3b, 0x81, 0xbd, 0x82, 0xb3, 0x07, 0xa0, 0x17, 0xca, 0xf1, 0x2c, 0x1c, 0x23,
	0xff, 0x8a, 0x5c, 0x1f, 0xa3, 0xcd, 0xdc, 0x13, 0x11, 0xa1, 0x54, 0x18, 0xe0, 0xe7, 0x99, 0xc8,
	0x30, 0xe6, 0xaf, 0xda, 0x8d, 0xc3, 0xb5, 0x60, 0x11, 0x36, 0x9d, 0xb2, 0xd0, 0x40, 0x67, 0x42,
	0x8e, 0xf9, 0xeb, 0xbc, 0x53, 0x35, 0x90, 0x9d, 0xc0, 0x46, 0xf7, 0x23, 0x46, 0x9f, 0x02, 0x54,
	0xb3, 0x44, 0x2b, 0xfe, 0x65, 0x7b, 0xf9, 0xd0, 0x3b, 0x6e, 0x1d, 0x15, 0xeb, 0xd6, 0x21, 0x83,
	0x9a, 0x92, 0x7d, 0x00, 0xaf, 0x9f, 0xa5, 0x77, 0xf3, 0x7c, 0xed, 0x72, 0xbf, 0xdd, 0xa8, 0x39,
	0x3a, 0x5c, 0xe0, 0x0a, 0xfd, 0xbf, 0x96, 0x6a, 0x8e, 0x8c, 0xc3, 0xea, 0xa9, 0x90, 0xf1, 0x6d,
	0xd0, 0xb3, 0x4b, 0xb8, 0x30, 0xcd, 0xf6, 0x38, 0xbf, 0xb3, 0x6b, 0x3b, 0xe8, 0xd9, 0xcd, 0xeb,
	0x42, 0x66, 0x5f, 0x05, 0x18, 0x8b, 0x0c, 0x23, 0x7d, 0x1b, 0xf4, 0x14, 0x5f, 0x6e, 0x2f, 0x9b,
	0x7d, 0xe5, 0x62, 0xec, 0x04, 0x60, 0x80, 0xc9, 0x68, 0x20, 0xc6, 0x12, 0x63, 0x5a, 0xbf, 0xde,
	0xf1, 0x6e, 0x59, 0xe6, 0x4d, 0x6f, 0x50, 0xb1, 0x97, 0x4f, 0x02, 0x47, 0xcb, 0xbe, 0x07, 0xaf,
	0x87, 0x5a, 0x9d, 0xcb, 0x28, 0x9b, 0x4f, 0x35, 0xed, 0x66, 0xef, 0x78, 0xcf, 0x75, 0x75, 0xe8,
	0xcb, 0x27, 0x81, 0xab, 0x36, 0xce, 0x5d, 0xcc, 0xb4, 0x18, 0x99, 0xc1, 0x23, 0xad, 0xed, 0x05,
	0x67, 0x87, 0x36, 0xce, 0x8e, 0x79, 0xea, 0xc1, 0xba, 0x11, 0xe4, 0x0d, 0xfb, 0x16, 0x36, 0x6b,
	0x55, 0x9a, 0x1b, 0xa3, 0xb8, 0x79, 0x14, 0x6f, 0xd0, 0x2b, 0x57, 0x80, 0x1f, 0xc3, 0x56, 0xbd,
	0x32, 0xd6, 0x82, 0xa7, 0xe7, 0x93, 0x50, 0x24, 0xb6, 0xbf, 0xb9, 0x61, 0xb6, 0x5c, 0x27, 0x8a,
	0x70, 0xaa, 0xcf, 0x6f, 0x7b, 0x1d, 0x6a, 0xee, 0x5a, 0xe0, 0x20, 0x26, 0xcb, 0x40, 0x87, 0x63,
	0x21, 0xc7, 0xdd, 0x0e, 0xdd, 0x69, 0x6b, 0x41, 0x05, 0xf8, 0x17, 0x94, 0xc5, 0xa9, 0xd9, 0x7c,
	0xdb, 0xc6, 0xa4, 0xfb, 0x27, 0x4f, 0x54, 0xda, 0x66, 0xc6, 0x57, 0x38, 0x27, 0x2a, 0x9f, 0x62,
	0x61, 0xfa, 0x7f, 0x80, 0xe7, 0x9c, 0x2a, 0xc6, 0x60, 0x85, 0x2e, 0x83, 0x3c, 0x00, 0x3d, 0x1b,
	0xe7, 0xc1, 0x2c, 0x8a, 0x50, 0x29, 0x5b, 0x65, 0x61, 0x9a, 0x57, 0xf8, 0x45, 0xa5, 0x32, 0xf7,
	0xb5, 0xf7, 0xae, 0x83, 0xf8, 0xbf, 0xc3, 0x4e, 0x1f, 0xb3, 0x51, 0x9a, 0x4d, 0x6c, 0x8e, 0xcf,
	0x33, 0x54, 0x0f, 0x27, 0x39, 0x82, 0xa6, 0x3d, 0xc8, 0x4b, 0x0b, 0x27, 0xa4, 0xf6, 0x77, 0x11,
	0x58, 0x95, 0x7f, 0x06, 0xad, 0x7a, 0x68, 0x35, 0x4d, 0xa5, 0x42, 0xf6, 0x0e, 0x9a, 0xb6, 0x9c,
	0xc6, 0xc2, 0x07, 0xe1, 0x7e, 0x49, 0x56, 0xe3, 0xb7, 0x80, 0xfd, 0x54, 0xde, 0x01, 0xca, 0xd6,
	0xe7, 0x9f, 0xc3, 0x4e, 0x0d, 0xb5, 0xa1, 0x8f, 0xa0, 0x19, 0xe5, 0x25, 0x36, 0xfe, 0xbb, 0xc4,
	0x5c, 0xe5, 0xbf, 0xa0, 0x30, 0x9d, 0x71, 0x86, 0x38, 0x41, 0xa9, 0x8b, 0xe8, 0x6f, 0xa1, 0x55,
	0x87, 0x6d, 0x78, 0x06, 0x2b, 0x37, 0x78, 0xa7, 0x8b, 0xae, 0x98, 0x67, 0xff, 0x47, 0xd8, 0xb2,
	0xb1, 0x8b, 0xde, 0xfd, 0xdf, 0x22, 0xbe, 0x81, 0xed, 0x32, 0x82, 0x4d, 0xc4, 0x61, 0x55, 0xd9,
	0x79, 0x36, 0xf2, 0x79, 0x5a, 0xf3, 0xf8, 0x9f, 0x06, 0xac, 0x5a, 0x35, 0xbb, 0x04, 0xcf, 0x69,
	0x02, 0x7b, 0x59, 0xe6, 0xb9, 0xdf, 0xb0, 0xfd, 0x83, 0x87, 0x49, 0x9b, 0xef, 0x87, 0x2a, 0xe8,
	0xde, 0x62, 0xb5, 0x45, 0x04, 0x7e, 0x9f, 0xb0, 0xde, 0x57, 0xb0, 0xe1, 0x0e, 0x9a, 0x55, 0xb9,
	0x1e, 0x38, 0x5a, 0xfb, 0xaf, 0x1e, 0x61, 0xf3, 0x60, 0xc3, 0x26, 0xfd, 0xbd, 0xbe, 0xff, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xcb, 0x41, 0x39, 0xf7, 0xce, 0x0a, 0x00, 0x00,
}
