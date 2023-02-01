// Code generated by protoc-gen-go. DO NOT EDIT.
// source: share.proto

package rest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import service "github.com/pydio/cells/common/service/proto"
import idm "github.com/pydio/cells/common/proto/idm"
import tree "github.com/pydio/cells/common/proto/tree"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Known values for link permissions
type ShareLinkAccessType int32

const (
	ShareLinkAccessType_NoAccess ShareLinkAccessType = 0
	ShareLinkAccessType_Preview  ShareLinkAccessType = 1
	ShareLinkAccessType_Download ShareLinkAccessType = 2
	ShareLinkAccessType_Upload   ShareLinkAccessType = 3
)

var ShareLinkAccessType_name = map[int32]string{
	0: "NoAccess",
	1: "Preview",
	2: "Download",
	3: "Upload",
}
var ShareLinkAccessType_value = map[string]int32{
	"NoAccess": 0,
	"Preview":  1,
	"Download": 2,
	"Upload":   3,
}

func (x ShareLinkAccessType) String() string {
	return proto.EnumName(ShareLinkAccessType_name, int32(x))
}
func (ShareLinkAccessType) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

type ListSharedResourcesRequest_ListShareType int32

const (
	ListSharedResourcesRequest_ANY   ListSharedResourcesRequest_ListShareType = 0
	ListSharedResourcesRequest_LINKS ListSharedResourcesRequest_ListShareType = 1
	ListSharedResourcesRequest_CELLS ListSharedResourcesRequest_ListShareType = 2
)

var ListSharedResourcesRequest_ListShareType_name = map[int32]string{
	0: "ANY",
	1: "LINKS",
	2: "CELLS",
}
var ListSharedResourcesRequest_ListShareType_value = map[string]int32{
	"ANY":   0,
	"LINKS": 1,
	"CELLS": 2,
}

func (x ListSharedResourcesRequest_ListShareType) String() string {
	return proto.EnumName(ListSharedResourcesRequest_ListShareType_name, int32(x))
}
func (ListSharedResourcesRequest_ListShareType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor9, []int{12, 0}
}

// Group collected acls by subjects
type CellAcl struct {
	// Associated Role ID
	RoleId string `protobuf:"bytes,1,opt,name=RoleId" json:"RoleId,omitempty"`
	// List of Acl Actions and their effect
	Actions []*idm.ACLAction `protobuf:"bytes,2,rep,name=Actions" json:"Actions,omitempty"`
	// Flag for detecting if it's a user role or not
	IsUserRole bool `protobuf:"varint,3,opt,name=IsUserRole" json:"IsUserRole,omitempty"`
	// Associated User
	User *idm.User `protobuf:"bytes,4,opt,name=User" json:"User,omitempty"`
	// Associated Group
	Group *idm.User `protobuf:"bytes,5,opt,name=Group" json:"Group,omitempty"`
	// Associated Role
	Role *idm.Role `protobuf:"bytes,6,opt,name=Role" json:"Role,omitempty"`
}

func (m *CellAcl) Reset()                    { *m = CellAcl{} }
func (m *CellAcl) String() string            { return proto.CompactTextString(m) }
func (*CellAcl) ProtoMessage()               {}
func (*CellAcl) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *CellAcl) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *CellAcl) GetActions() []*idm.ACLAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *CellAcl) GetIsUserRole() bool {
	if m != nil {
		return m.IsUserRole
	}
	return false
}

func (m *CellAcl) GetUser() *idm.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CellAcl) GetGroup() *idm.User {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *CellAcl) GetRole() *idm.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

// Model for representing a Cell
type Cell struct {
	// Unique Id of the Cell
	Uuid string `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
	// Label of the Cell (max 500 chars)
	Label string `protobuf:"bytes,2,opt,name=Label" json:"Label,omitempty"`
	// Long description of the Cell (max 1000 chars)
	Description string `protobuf:"bytes,3,opt,name=Description" json:"Description,omitempty"`
	// Nodes attached as roots to this Cell
	RootNodes []*tree.Node `protobuf:"bytes,4,rep,name=RootNodes" json:"RootNodes,omitempty"`
	// Access control for this Cell
	ACLs map[string]*CellAcl `protobuf:"bytes,5,rep,name=ACLs" json:"ACLs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Associated access policies
	Policies []*service.ResourcePolicy `protobuf:"bytes,6,rep,name=Policies" json:"Policies,omitempty"`
	// Whether these policies are currently editable
	PoliciesContextEditable bool `protobuf:"varint,7,opt,name=PoliciesContextEditable" json:"PoliciesContextEditable,omitempty"`
}

func (m *Cell) Reset()                    { *m = Cell{} }
func (m *Cell) String() string            { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()               {}
func (*Cell) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *Cell) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Cell) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Cell) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Cell) GetRootNodes() []*tree.Node {
	if m != nil {
		return m.RootNodes
	}
	return nil
}

func (m *Cell) GetACLs() map[string]*CellAcl {
	if m != nil {
		return m.ACLs
	}
	return nil
}

func (m *Cell) GetPolicies() []*service.ResourcePolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *Cell) GetPoliciesContextEditable() bool {
	if m != nil {
		return m.PoliciesContextEditable
	}
	return false
}

type ShareLinkTargetUser struct {
	Display       string `protobuf:"bytes,1,opt,name=Display" json:"Display,omitempty"`
	DownloadCount int32  `protobuf:"varint,2,opt,name=DownloadCount" json:"DownloadCount,omitempty"`
}

func (m *ShareLinkTargetUser) Reset()                    { *m = ShareLinkTargetUser{} }
func (m *ShareLinkTargetUser) String() string            { return proto.CompactTextString(m) }
func (*ShareLinkTargetUser) ProtoMessage()               {}
func (*ShareLinkTargetUser) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *ShareLinkTargetUser) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *ShareLinkTargetUser) GetDownloadCount() int32 {
	if m != nil {
		return m.DownloadCount
	}
	return 0
}

// Model for representing a public link
type ShareLink struct {
	// Internal identifier of the link
	Uuid string `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
	// Unique Hash for accessing the link
	LinkHash string `protobuf:"bytes,2,opt,name=LinkHash" json:"LinkHash,omitempty"`
	// Full URL for accessing the link
	LinkUrl string `protobuf:"bytes,3,opt,name=LinkUrl" json:"LinkUrl,omitempty"`
	// Label of the Link (max 500 chars)
	Label string `protobuf:"bytes,4,opt,name=Label" json:"Label,omitempty"`
	// Description of the Link (max 1000 chars)
	Description string `protobuf:"bytes,5,opt,name=Description" json:"Description,omitempty"`
	// Temporary user Uuid used to login automatically when accessing this link
	UserUuid string `protobuf:"bytes,6,opt,name=UserUuid" json:"UserUuid,omitempty"`
	// Temporary user Login used to login automatically when accessing this link
	UserLogin string `protobuf:"bytes,7,opt,name=UserLogin" json:"UserLogin,omitempty"`
	// Whether a password is required or not to access the link
	PasswordRequired bool `protobuf:"varint,8,opt,name=PasswordRequired" json:"PasswordRequired,omitempty"`
	// Timestamp of start date for enabling the share (not implemented yet)
	AccessStart int64 `protobuf:"varint,9,opt,name=AccessStart" json:"AccessStart,omitempty"`
	// Timestamp after which the share is disabled
	AccessEnd int64 `protobuf:"varint,10,opt,name=AccessEnd" json:"AccessEnd,omitempty"`
	// Maximum number of downloads until expiration
	MaxDownloads int64 `protobuf:"varint,11,opt,name=MaxDownloads" json:"MaxDownloads,omitempty"`
	// Current number of downloads
	CurrentDownloads int64 `protobuf:"varint,12,opt,name=CurrentDownloads" json:"CurrentDownloads,omitempty"`
	// Display Template for loading the public link
	ViewTemplateName string `protobuf:"bytes,13,opt,name=ViewTemplateName" json:"ViewTemplateName,omitempty"`
	// TargetUsers can be used to restrict access
	TargetUsers map[string]*ShareLinkTargetUser `protobuf:"bytes,14,rep,name=TargetUsers" json:"TargetUsers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// RestrictToTargetUsers enable users restriction
	RestrictToTargetUsers bool `protobuf:"varint,15,opt,name=RestrictToTargetUsers" json:"RestrictToTargetUsers,omitempty"`
	// Nodes in the tree that serve as root to this link
	RootNodes []*tree.Node `protobuf:"bytes,16,rep,name=RootNodes" json:"RootNodes,omitempty"`
	// Specific permissions for public links
	Permissions []ShareLinkAccessType `protobuf:"varint,17,rep,packed,name=Permissions,enum=rest.ShareLinkAccessType" json:"Permissions,omitempty"`
	// Security policies
	Policies []*service.ResourcePolicy `protobuf:"bytes,18,rep,name=Policies" json:"Policies,omitempty"`
	// Whether policies are currently editable or not
	PoliciesContextEditable bool `protobuf:"varint,19,opt,name=PoliciesContextEditable" json:"PoliciesContextEditable,omitempty"`
}

func (m *ShareLink) Reset()                    { *m = ShareLink{} }
func (m *ShareLink) String() string            { return proto.CompactTextString(m) }
func (*ShareLink) ProtoMessage()               {}
func (*ShareLink) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *ShareLink) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ShareLink) GetLinkHash() string {
	if m != nil {
		return m.LinkHash
	}
	return ""
}

func (m *ShareLink) GetLinkUrl() string {
	if m != nil {
		return m.LinkUrl
	}
	return ""
}

func (m *ShareLink) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ShareLink) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ShareLink) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *ShareLink) GetUserLogin() string {
	if m != nil {
		return m.UserLogin
	}
	return ""
}

func (m *ShareLink) GetPasswordRequired() bool {
	if m != nil {
		return m.PasswordRequired
	}
	return false
}

func (m *ShareLink) GetAccessStart() int64 {
	if m != nil {
		return m.AccessStart
	}
	return 0
}

func (m *ShareLink) GetAccessEnd() int64 {
	if m != nil {
		return m.AccessEnd
	}
	return 0
}

func (m *ShareLink) GetMaxDownloads() int64 {
	if m != nil {
		return m.MaxDownloads
	}
	return 0
}

func (m *ShareLink) GetCurrentDownloads() int64 {
	if m != nil {
		return m.CurrentDownloads
	}
	return 0
}

func (m *ShareLink) GetViewTemplateName() string {
	if m != nil {
		return m.ViewTemplateName
	}
	return ""
}

func (m *ShareLink) GetTargetUsers() map[string]*ShareLinkTargetUser {
	if m != nil {
		return m.TargetUsers
	}
	return nil
}

func (m *ShareLink) GetRestrictToTargetUsers() bool {
	if m != nil {
		return m.RestrictToTargetUsers
	}
	return false
}

func (m *ShareLink) GetRootNodes() []*tree.Node {
	if m != nil {
		return m.RootNodes
	}
	return nil
}

func (m *ShareLink) GetPermissions() []ShareLinkAccessType {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *ShareLink) GetPolicies() []*service.ResourcePolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *ShareLink) GetPoliciesContextEditable() bool {
	if m != nil {
		return m.PoliciesContextEditable
	}
	return false
}

// Request for creating a Cell
type PutCellRequest struct {
	// Content of the Cell (Room is legacy name)
	Room *Cell `protobuf:"bytes,1,opt,name=Room" json:"Room,omitempty"`
	// Whether to create a dedicated folder for this cell at creation
	CreateEmptyRoot bool `protobuf:"varint,2,opt,name=CreateEmptyRoot" json:"CreateEmptyRoot,omitempty"`
}

func (m *PutCellRequest) Reset()                    { *m = PutCellRequest{} }
func (m *PutCellRequest) String() string            { return proto.CompactTextString(m) }
func (*PutCellRequest) ProtoMessage()               {}
func (*PutCellRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *PutCellRequest) GetRoom() *Cell {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *PutCellRequest) GetCreateEmptyRoot() bool {
	if m != nil {
		return m.CreateEmptyRoot
	}
	return false
}

// Load a Cell request
type GetCellRequest struct {
	// Cell Uuid
	Uuid string `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
}

func (m *GetCellRequest) Reset()                    { *m = GetCellRequest{} }
func (m *GetCellRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCellRequest) ProtoMessage()               {}
func (*GetCellRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *GetCellRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// Request for deleting a Cell
type DeleteCellRequest struct {
	// Cell Uuid
	Uuid string `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
}

func (m *DeleteCellRequest) Reset()                    { *m = DeleteCellRequest{} }
func (m *DeleteCellRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteCellRequest) ProtoMessage()               {}
func (*DeleteCellRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

func (m *DeleteCellRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type DeleteCellResponse struct {
	// Delete result
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *DeleteCellResponse) Reset()                    { *m = DeleteCellResponse{} }
func (m *DeleteCellResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteCellResponse) ProtoMessage()               {}
func (*DeleteCellResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }

func (m *DeleteCellResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetShareLinkRequest struct {
	// Link Uuid
	Uuid string `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
}

func (m *GetShareLinkRequest) Reset()                    { *m = GetShareLinkRequest{} }
func (m *GetShareLinkRequest) String() string            { return proto.CompactTextString(m) }
func (*GetShareLinkRequest) ProtoMessage()               {}
func (*GetShareLinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{8} }

func (m *GetShareLinkRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// Request for create/update a link
type PutShareLinkRequest struct {
	// Content of the link to create
	ShareLink *ShareLink `protobuf:"bytes,1,opt,name=ShareLink" json:"ShareLink,omitempty"`
	// Whether it has Password enabled
	PasswordEnabled bool `protobuf:"varint,2,opt,name=PasswordEnabled" json:"PasswordEnabled,omitempty"`
	// Set if switching from no password to password
	CreatePassword string `protobuf:"bytes,3,opt,name=CreatePassword" json:"CreatePassword,omitempty"`
	// Set if updating an existing password
	UpdatePassword string `protobuf:"bytes,4,opt,name=UpdatePassword" json:"UpdatePassword,omitempty"`
	// Change the ShareLink Hash with a custom value
	UpdateCustomHash string `protobuf:"bytes,5,opt,name=UpdateCustomHash" json:"UpdateCustomHash,omitempty"`
}

func (m *PutShareLinkRequest) Reset()                    { *m = PutShareLinkRequest{} }
func (m *PutShareLinkRequest) String() string            { return proto.CompactTextString(m) }
func (*PutShareLinkRequest) ProtoMessage()               {}
func (*PutShareLinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{9} }

func (m *PutShareLinkRequest) GetShareLink() *ShareLink {
	if m != nil {
		return m.ShareLink
	}
	return nil
}

func (m *PutShareLinkRequest) GetPasswordEnabled() bool {
	if m != nil {
		return m.PasswordEnabled
	}
	return false
}

func (m *PutShareLinkRequest) GetCreatePassword() string {
	if m != nil {
		return m.CreatePassword
	}
	return ""
}

func (m *PutShareLinkRequest) GetUpdatePassword() string {
	if m != nil {
		return m.UpdatePassword
	}
	return ""
}

func (m *PutShareLinkRequest) GetUpdateCustomHash() string {
	if m != nil {
		return m.UpdateCustomHash
	}
	return ""
}

// Request for deleting a link
type DeleteShareLinkRequest struct {
	// Id of Link to delete
	Uuid string `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
}

func (m *DeleteShareLinkRequest) Reset()                    { *m = DeleteShareLinkRequest{} }
func (m *DeleteShareLinkRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteShareLinkRequest) ProtoMessage()               {}
func (*DeleteShareLinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{10} }

func (m *DeleteShareLinkRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// Response for deleting a share link
type DeleteShareLinkResponse struct {
	// If delete sucess or failed
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *DeleteShareLinkResponse) Reset()                    { *m = DeleteShareLinkResponse{} }
func (m *DeleteShareLinkResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteShareLinkResponse) ProtoMessage()               {}
func (*DeleteShareLinkResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{11} }

func (m *DeleteShareLinkResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListSharedResourcesRequest struct {
	// Filter output to a given type
	ShareType ListSharedResourcesRequest_ListShareType `protobuf:"varint,1,opt,name=ShareType,enum=rest.ListSharedResourcesRequest_ListShareType" json:"ShareType,omitempty"`
	// Will restrict the list to the shares readable by a specific subject.
	// In user-context, current user is used by default. In admin-context, this can
	// be any resource policy subject
	Subject string `protobuf:"bytes,3,opt,name=Subject" json:"Subject,omitempty"`
	// If true, will also check filter the output to shares actually owned by subject
	OwnedBySubject bool `protobuf:"varint,4,opt,name=OwnedBySubject" json:"OwnedBySubject,omitempty"`
	// Start listing at a given offset
	Offset int32 `protobuf:"varint,5,opt,name=Offset" json:"Offset,omitempty"`
	// Limit number of results
	Limit int32 `protobuf:"varint,6,opt,name=Limit" json:"Limit,omitempty"`
}

func (m *ListSharedResourcesRequest) Reset()                    { *m = ListSharedResourcesRequest{} }
func (m *ListSharedResourcesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSharedResourcesRequest) ProtoMessage()               {}
func (*ListSharedResourcesRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{12} }

func (m *ListSharedResourcesRequest) GetShareType() ListSharedResourcesRequest_ListShareType {
	if m != nil {
		return m.ShareType
	}
	return ListSharedResourcesRequest_ANY
}

func (m *ListSharedResourcesRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ListSharedResourcesRequest) GetOwnedBySubject() bool {
	if m != nil {
		return m.OwnedBySubject
	}
	return false
}

func (m *ListSharedResourcesRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListSharedResourcesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListSharedResourcesResponse struct {
	// Actual results
	Resources []*ListSharedResourcesResponse_SharedResource `protobuf:"bytes,1,rep,name=Resources" json:"Resources,omitempty"`
	// Cursor informations
	Offset int32 `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,3,opt,name=Limit" json:"Limit,omitempty"`
	Total  int32 `protobuf:"varint,4,opt,name=Total" json:"Total,omitempty"`
}

func (m *ListSharedResourcesResponse) Reset()                    { *m = ListSharedResourcesResponse{} }
func (m *ListSharedResourcesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSharedResourcesResponse) ProtoMessage()               {}
func (*ListSharedResourcesResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{13} }

func (m *ListSharedResourcesResponse) GetResources() []*ListSharedResourcesResponse_SharedResource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *ListSharedResourcesResponse) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListSharedResourcesResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListSharedResourcesResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

// Container for ShareLink or Cell
type ListSharedResourcesResponse_SharedResource struct {
	Node  *tree.Node `protobuf:"bytes,1,opt,name=Node" json:"Node,omitempty"`
	Link  *ShareLink `protobuf:"bytes,2,opt,name=Link" json:"Link,omitempty"`
	Cells []*Cell    `protobuf:"bytes,3,rep,name=Cells" json:"Cells,omitempty"`
}

func (m *ListSharedResourcesResponse_SharedResource) Reset() {
	*m = ListSharedResourcesResponse_SharedResource{}
}
func (m *ListSharedResourcesResponse_SharedResource) String() string {
	return proto.CompactTextString(m)
}
func (*ListSharedResourcesResponse_SharedResource) ProtoMessage() {}
func (*ListSharedResourcesResponse_SharedResource) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{13, 0}
}

func (m *ListSharedResourcesResponse_SharedResource) GetNode() *tree.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *ListSharedResourcesResponse_SharedResource) GetLink() *ShareLink {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *ListSharedResourcesResponse_SharedResource) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type UpdateSharePoliciesRequest struct {
	// Cell or Link UUID
	Uuid string `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
	// List of policies to update
	Policies []*service.ResourcePolicy `protobuf:"bytes,2,rep,name=Policies" json:"Policies,omitempty"`
}

func (m *UpdateSharePoliciesRequest) Reset()                    { *m = UpdateSharePoliciesRequest{} }
func (m *UpdateSharePoliciesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateSharePoliciesRequest) ProtoMessage()               {}
func (*UpdateSharePoliciesRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{14} }

func (m *UpdateSharePoliciesRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UpdateSharePoliciesRequest) GetPolicies() []*service.ResourcePolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type UpdateSharePoliciesResponse struct {
	Success                 bool                      `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Policies                []*service.ResourcePolicy `protobuf:"bytes,2,rep,name=Policies" json:"Policies,omitempty"`
	PoliciesContextEditable bool                      `protobuf:"varint,3,opt,name=PoliciesContextEditable" json:"PoliciesContextEditable,omitempty"`
}

func (m *UpdateSharePoliciesResponse) Reset()                    { *m = UpdateSharePoliciesResponse{} }
func (m *UpdateSharePoliciesResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateSharePoliciesResponse) ProtoMessage()               {}
func (*UpdateSharePoliciesResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{15} }

func (m *UpdateSharePoliciesResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UpdateSharePoliciesResponse) GetPolicies() []*service.ResourcePolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *UpdateSharePoliciesResponse) GetPoliciesContextEditable() bool {
	if m != nil {
		return m.PoliciesContextEditable
	}
	return false
}

func init() {
	proto.RegisterType((*CellAcl)(nil), "rest.CellAcl")
	proto.RegisterType((*Cell)(nil), "rest.Cell")
	proto.RegisterType((*ShareLinkTargetUser)(nil), "rest.ShareLinkTargetUser")
	proto.RegisterType((*ShareLink)(nil), "rest.ShareLink")
	proto.RegisterType((*PutCellRequest)(nil), "rest.PutCellRequest")
	proto.RegisterType((*GetCellRequest)(nil), "rest.GetCellRequest")
	proto.RegisterType((*DeleteCellRequest)(nil), "rest.DeleteCellRequest")
	proto.RegisterType((*DeleteCellResponse)(nil), "rest.DeleteCellResponse")
	proto.RegisterType((*GetShareLinkRequest)(nil), "rest.GetShareLinkRequest")
	proto.RegisterType((*PutShareLinkRequest)(nil), "rest.PutShareLinkRequest")
	proto.RegisterType((*DeleteShareLinkRequest)(nil), "rest.DeleteShareLinkRequest")
	proto.RegisterType((*DeleteShareLinkResponse)(nil), "rest.DeleteShareLinkResponse")
	proto.RegisterType((*ListSharedResourcesRequest)(nil), "rest.ListSharedResourcesRequest")
	proto.RegisterType((*ListSharedResourcesResponse)(nil), "rest.ListSharedResourcesResponse")
	proto.RegisterType((*ListSharedResourcesResponse_SharedResource)(nil), "rest.ListSharedResourcesResponse.SharedResource")
	proto.RegisterType((*UpdateSharePoliciesRequest)(nil), "rest.UpdateSharePoliciesRequest")
	proto.RegisterType((*UpdateSharePoliciesResponse)(nil), "rest.UpdateSharePoliciesResponse")
	proto.RegisterEnum("rest.ShareLinkAccessType", ShareLinkAccessType_name, ShareLinkAccessType_value)
	proto.RegisterEnum("rest.ListSharedResourcesRequest_ListShareType", ListSharedResourcesRequest_ListShareType_name, ListSharedResourcesRequest_ListShareType_value)
}

func init() { proto.RegisterFile("share.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 1250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xdb, 0x46,
	0x13, 0x8d, 0x7e, 0xa8, 0x9f, 0x51, 0x2c, 0x2b, 0xeb, 0x7c, 0x09, 0x3f, 0xa5, 0x49, 0x04, 0x26,
	0x68, 0x95, 0xa0, 0xa1, 0x5a, 0xbb, 0x28, 0x82, 0xf6, 0x4a, 0x91, 0xdd, 0x34, 0xad, 0xea, 0x18,
	0x6b, 0xab, 0x40, 0x7a, 0x47, 0x93, 0x1b, 0x87, 0x0d, 0xc9, 0x55, 0xb9, 0x4b, 0x2b, 0x7a, 0x85,
	0x3e, 0x48, 0x2f, 0xfa, 0x28, 0x7d, 0x89, 0x02, 0xbd, 0x2a, 0x8a, 0x02, 0xbd, 0xe9, 0x03, 0x14,
	0x3b, 0x4b, 0x8a, 0x94, 0xac, 0x28, 0x46, 0xd1, 0x0b, 0x1b, 0x3b, 0x67, 0x66, 0x76, 0x77, 0xce,
	0xce, 0x19, 0x11, 0x5a, 0xe2, 0x95, 0x13, 0x33, 0x7b, 0x1a, 0x73, 0xc9, 0x49, 0x35, 0x66, 0x42,
	0x76, 0x1f, 0x9f, 0xf9, 0xf2, 0x55, 0x72, 0x6a, 0xbb, 0x3c, 0x1c, 0x4c, 0xe7, 0x9e, 0xcf, 0x07,
	0x2e, 0x0b, 0x02, 0x31, 0x70, 0x79, 0x18, 0xf2, 0x68, 0x20, 0x58, 0x7c, 0xee, 0xbb, 0x6c, 0x80,
	0x29, 0x29, 0xa8, 0xf3, 0xbb, 0x1f, 0x6f, 0xce, 0xd4, 0x19, 0xbe, 0x17, 0xaa, 0xbf, 0x34, 0x65,
	0xef, 0x32, 0x29, 0x32, 0x66, 0x0c, 0xff, 0xa5, 0x49, 0x9f, 0x16, 0x92, 0xc2, 0x99, 0x2f, 0x5f,
	0xf3, 0xd9, 0xe0, 0x8c, 0x3f, 0x42, 0xe7, 0xa3, 0x73, 0x27, 0xf0, 0x3d, 0x47, 0xf2, 0x58, 0x0c,
	0x16, 0x4b, 0x9d, 0x67, 0xfd, 0x52, 0x82, 0xfa, 0x88, 0x05, 0xc1, 0xd0, 0x0d, 0xc8, 0x0d, 0xa8,
	0x51, 0x1e, 0xb0, 0x67, 0x9e, 0x59, 0xea, 0x95, 0xfa, 0x4d, 0x9a, 0x5a, 0xa4, 0x0f, 0xf5, 0xa1,
	0x2b, 0x7d, 0x1e, 0x09, 0xb3, 0xdc, 0xab, 0xf4, 0x5b, 0xbb, 0x6d, 0x5b, 0xdd, 0x76, 0x38, 0x1a,
	0x6b, 0x98, 0x66, 0x6e, 0x72, 0x07, 0xe0, 0x99, 0x98, 0x08, 0x16, 0xab, 0x4c, 0xb3, 0xd2, 0x2b,
	0xf5, 0x1b, 0xb4, 0x80, 0x90, 0xdb, 0x50, 0x55, 0x6b, 0xb3, 0xda, 0x2b, 0xf5, 0x5b, 0xbb, 0x4d,
	0xdc, 0x06, 0x9d, 0x08, 0x93, 0xbb, 0x60, 0x3c, 0x8d, 0x79, 0x32, 0x35, 0x8d, 0x55, 0xbf, 0xc6,
	0x55, 0x3e, 0xee, 0x5c, 0x2b, 0xf8, 0x15, 0x40, 0x11, 0xb6, 0xfe, 0x2a, 0x43, 0x55, 0x15, 0x43,
	0x08, 0x54, 0x27, 0x89, 0x9f, 0xd5, 0x81, 0x6b, 0x72, 0x1b, 0x8c, 0xb1, 0x73, 0xca, 0x02, 0xb3,
	0xac, 0xc0, 0x27, 0xf5, 0xdf, 0x7e, 0xbd, 0x5b, 0x79, 0xf3, 0x77, 0x85, 0x6a, 0x94, 0x3c, 0x80,
	0xd6, 0x3e, 0x13, 0x6e, 0xec, 0x4f, 0x55, 0x29, 0x78, 0xf7, 0x2c, 0xe8, 0xf7, 0x3a, 0x2d, 0xfa,
	0x48, 0x1f, 0x9a, 0x94, 0x73, 0x79, 0xc8, 0x3d, 0x26, 0xcc, 0x2a, 0x32, 0x02, 0x36, 0xbe, 0x85,
	0x82, 0x68, 0xee, 0x24, 0x7d, 0xa8, 0x0e, 0x47, 0x63, 0x61, 0x1a, 0x18, 0x74, 0xdd, 0x56, 0xcd,
	0x64, 0xab, 0x1b, 0x2a, 0xf2, 0xc4, 0x41, 0x24, 0xe3, 0x39, 0xc5, 0x08, 0xb2, 0x07, 0x8d, 0x23,
	0x1e, 0xf8, 0xae, 0xcf, 0x84, 0x59, 0xc3, 0xe8, 0x9b, 0x76, 0xda, 0x56, 0x36, 0x65, 0x82, 0x27,
	0xb1, 0xcb, 0x30, 0x60, 0x4e, 0x17, 0x81, 0xe4, 0x31, 0xdc, 0xcc, 0xd6, 0x23, 0x1e, 0x49, 0xf6,
	0x46, 0x1e, 0x78, 0xbe, 0x74, 0x4e, 0x03, 0x66, 0xd6, 0x91, 0xfb, 0xb7, 0xb9, 0xbb, 0x5f, 0x40,
	0x73, 0x71, 0x03, 0xd2, 0x81, 0xca, 0x6b, 0x36, 0x4f, 0xc9, 0x52, 0x4b, 0x72, 0x0f, 0x8c, 0x73,
	0x27, 0x48, 0x18, 0x72, 0xd5, 0xda, 0xdd, 0xca, 0x2f, 0x3e, 0x74, 0x03, 0xaa, 0x7d, 0x9f, 0x95,
	0x1f, 0x97, 0xac, 0x09, 0xec, 0x1c, 0x2b, 0xb5, 0x8c, 0xfd, 0xe8, 0xf5, 0x89, 0x13, 0x9f, 0x31,
	0x89, 0x0f, 0x69, 0x42, 0x7d, 0xdf, 0x17, 0xd3, 0xc0, 0xc9, 0x76, 0xcd, 0x4c, 0x72, 0x1f, 0xb6,
	0xf6, 0xf9, 0x2c, 0x0a, 0xb8, 0xe3, 0x8d, 0x78, 0x12, 0x49, 0x3c, 0xc1, 0xa0, 0xcb, 0xa0, 0xf5,
	0x67, 0x0d, 0x9a, 0x8b, 0x7d, 0xd7, 0xbe, 0x66, 0x17, 0x1a, 0xca, 0xf7, 0xa5, 0x23, 0x5e, 0xe9,
	0x07, 0xa5, 0x0b, 0x5b, 0x9d, 0xae, 0xd6, 0x93, 0x38, 0xd0, 0xcf, 0x48, 0x33, 0x33, 0xef, 0x81,
	0xea, 0x65, 0x7a, 0xc0, 0xd8, 0xd0, 0x03, 0x5d, 0x68, 0xa8, 0x4a, 0xf1, 0x5e, 0x35, 0x7d, 0x7e,
	0x66, 0x93, 0xf7, 0xa0, 0xa9, 0xd6, 0x63, 0x7e, 0xe6, 0x47, 0xf8, 0x10, 0x4d, 0x9a, 0x03, 0xe4,
	0x21, 0x74, 0x8e, 0x1c, 0x21, 0x66, 0x3c, 0xf6, 0x28, 0xfb, 0x21, 0xf1, 0x63, 0xe6, 0x99, 0x0d,
	0x7c, 0xad, 0x0b, 0x38, 0xe9, 0x41, 0x6b, 0xe8, 0xba, 0x4c, 0x88, 0x63, 0xe9, 0xc4, 0xd2, 0x6c,
	0xf6, 0x4a, 0xfd, 0x0a, 0x2d, 0x42, 0xea, 0x2c, 0x6d, 0x1e, 0x44, 0x9e, 0x09, 0xe8, 0xcf, 0x01,
	0x62, 0xc1, 0xd5, 0x6f, 0x9c, 0x37, 0x19, 0xb7, 0xc2, 0x6c, 0x61, 0xc0, 0x12, 0xa6, 0xee, 0x33,
	0x4a, 0xe2, 0x98, 0x45, 0x32, 0x8f, 0xbb, 0x8a, 0x71, 0x17, 0x70, 0x15, 0xfb, 0xad, 0xcf, 0x66,
	0x27, 0x2c, 0x9c, 0x06, 0x8e, 0x64, 0x87, 0x4e, 0xc8, 0xcc, 0x2d, 0x2c, 0xf0, 0x02, 0x4e, 0x9e,
	0x40, 0x2b, 0xef, 0x08, 0x61, 0xb6, 0xb1, 0xa9, 0x7b, 0xba, 0x93, 0x16, 0x6f, 0x6b, 0x17, 0x42,
	0xb4, 0x1c, 0x8a, 0x49, 0xe4, 0x13, 0xf8, 0x1f, 0x65, 0x42, 0xc6, 0xbe, 0x2b, 0x4f, 0x78, 0x71,
	0xb7, 0x6d, 0x24, 0x6c, 0xbd, 0x73, 0x59, 0x9f, 0x9d, 0x4d, 0xfa, 0xfc, 0x1c, 0x5a, 0x47, 0x2c,
	0x0e, 0x7d, 0x21, 0x70, 0xba, 0x5d, 0xeb, 0x55, 0xfa, 0xed, 0xdd, 0xff, 0xaf, 0xdc, 0x51, 0xd3,
	0x79, 0x32, 0x9f, 0x32, 0x5a, 0x8c, 0x5e, 0x92, 0x2c, 0xf9, 0x0f, 0x24, 0xbb, 0xb3, 0x59, 0xb2,
	0x2f, 0xa0, 0xb3, 0x4a, 0xd6, 0x1a, 0xe5, 0x0e, 0x96, 0x95, 0xbb, 0x5a, 0x4b, 0xbe, 0x43, 0x51,
	0xc5, 0xdf, 0x41, 0xfb, 0x28, 0x91, 0x4a, 0xde, 0xaa, 0xf3, 0x98, 0x90, 0xe4, 0x8e, 0x1a, 0xb4,
	0x3c, 0xc4, 0x9d, 0x15, 0x7b, 0x0b, 0xfd, 0x53, 0xc4, 0x49, 0x1f, 0xb6, 0x47, 0x31, 0x73, 0x24,
	0x3b, 0x08, 0xa7, 0x72, 0xae, 0x08, 0xc5, 0x03, 0x1b, 0x74, 0x15, 0xb6, 0xee, 0x43, 0xfb, 0x29,
	0x5b, 0xda, 0x7b, 0x8d, 0x9c, 0xad, 0x0f, 0xe0, 0xda, 0x3e, 0x0b, 0x98, 0x64, 0xef, 0x0a, 0xb4,
	0x81, 0x14, 0x03, 0xc5, 0x94, 0x47, 0x82, 0x29, 0xc5, 0x1f, 0x27, 0xf8, 0x4c, 0x18, 0xdc, 0xa0,
	0x99, 0x69, 0x3d, 0x80, 0x9d, 0xa7, 0x4c, 0x2e, 0xea, 0xdf, 0xb4, 0xf5, 0x1f, 0x25, 0xd8, 0x39,
	0x4a, 0x2e, 0xc6, 0x3e, 0x2a, 0xcc, 0xa2, 0x94, 0x90, 0xed, 0x15, 0x5a, 0x69, 0x61, 0x5a, 0xf5,
	0x61, 0x3b, 0xd3, 0xf1, 0x41, 0xa4, 0x5e, 0xce, 0xcb, 0xa8, 0x59, 0x81, 0xc9, 0xfb, 0xd0, 0xd6,
	0x6c, 0x65, 0x8e, 0x74, 0x5c, 0xad, 0xa0, 0x2a, 0x6e, 0x32, 0xf5, 0x8a, 0x71, 0x55, 0x1d, 0xb7,
	0x8c, 0x2a, 0x75, 0x6a, 0x64, 0x94, 0x08, 0xc9, 0x43, 0x9c, 0x8d, 0x86, 0x56, 0xe7, 0x2a, 0x6e,
	0x7d, 0x08, 0x37, 0x34, 0x8f, 0x97, 0xa2, 0x66, 0x0f, 0x6e, 0x5e, 0x88, 0x7e, 0x27, 0xf5, 0x3f,
	0x96, 0xa1, 0x3b, 0xf6, 0x85, 0x26, 0xd4, 0xcb, 0x14, 0x21, 0xb2, 0x73, 0xc6, 0x29, 0xad, 0x4a,
	0x58, 0x98, 0xda, 0xde, 0xb5, 0x35, 0xad, 0x6f, 0x4f, 0xca, 0x5d, 0x28, 0xc7, 0x7c, 0x03, 0x7d,
	0x8d, 0xd3, 0xef, 0x99, 0x2b, 0xb3, 0x99, 0x9f, 0x9a, 0x8a, 0xbd, 0xe7, 0xb3, 0x88, 0x79, 0x4f,
	0xe6, 0x59, 0x40, 0x15, 0xef, 0xb9, 0x82, 0xaa, 0xaf, 0x9f, 0xe7, 0x2f, 0x5f, 0x0a, 0x26, 0x91,
	0x33, 0x83, 0xa6, 0x16, 0xb9, 0x0e, 0xc6, 0xd8, 0x0f, 0x7d, 0x89, 0x63, 0xde, 0xa0, 0xda, 0xb0,
	0x6c, 0xd8, 0x5a, 0xba, 0x0b, 0xa9, 0x43, 0x65, 0x78, 0xf8, 0xa2, 0x73, 0x85, 0x34, 0xc1, 0x18,
	0x3f, 0x3b, 0xfc, 0xfa, 0xb8, 0x53, 0x52, 0xcb, 0xd1, 0xc1, 0x78, 0x7c, 0xdc, 0x29, 0x5b, 0x3f,
	0x97, 0xe1, 0xd6, 0xda, 0xba, 0x52, 0x1a, 0x0f, 0xa1, 0xb9, 0x00, 0xcd, 0x12, 0x4e, 0x93, 0x8f,
	0x36, 0xb0, 0xa1, 0xb3, 0xec, 0x65, 0x9c, 0xe6, 0x5b, 0x14, 0xaa, 0x29, 0xaf, 0xaf, 0xa6, 0x52,
	0xa8, 0x46, 0xa1, 0x27, 0x5c, 0x3a, 0xfa, 0x77, 0xd1, 0xa0, 0xda, 0xe8, 0xce, 0xa0, 0xbd, 0x7c,
	0x80, 0x1a, 0x0b, 0x6a, 0x70, 0x2e, 0xc6, 0x42, 0x3e, 0x54, 0x11, 0x27, 0xf7, 0xa0, 0x8a, 0x2a,
	0x29, 0xaf, 0x57, 0x09, 0x3a, 0x49, 0x0f, 0x0c, 0x25, 0x5e, 0x61, 0x56, 0xd2, 0xd1, 0x9c, 0x0f,
	0x17, 0xed, 0xb0, 0x18, 0x74, 0x75, 0xc3, 0x62, 0x6a, 0x36, 0x10, 0x37, 0x34, 0xe8, 0xd2, 0x2c,
	0x2e, 0x5f, 0x72, 0x16, 0x5b, 0x3f, 0x95, 0xe0, 0xd6, 0xda, 0x73, 0xde, 0xd5, 0xda, 0xff, 0xea,
	0xb8, 0x4d, 0xa3, 0xbf, 0xb2, 0x71, 0xf4, 0x3f, 0xfc, 0xaa, 0xf0, 0x95, 0x95, 0xff, 0x1a, 0x91,
	0xab, 0xd0, 0x38, 0xe4, 0xda, 0xee, 0x5c, 0x21, 0x2d, 0xa8, 0x1f, 0xc5, 0xec, 0xdc, 0x67, 0xb3,
	0x4e, 0x49, 0xb9, 0xb2, 0x5f, 0xed, 0x4e, 0x99, 0x00, 0xd4, 0x26, 0x53, 0x5c, 0x57, 0x4e, 0x6b,
	0xf8, 0xdd, 0xbf, 0xf7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xa2, 0x0f, 0xc2, 0xe6, 0x0c,
	0x00, 0x00,
}
