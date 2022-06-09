// Copyright 2014 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Request and reply to the "checkin server" devices poll every few hours.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: checkin.proto

package checkin_proto

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

// A concrete name/value pair sent to the device's Gservices database.
type GservicesSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value []byte `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
}

func (x *GservicesSetting) Reset() {
	*x = GservicesSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GservicesSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GservicesSetting) ProtoMessage() {}

func (x *GservicesSetting) ProtoReflect() protoreflect.Message {
	mi := &file_checkin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GservicesSetting.ProtoReflect.Descriptor instead.
func (*GservicesSetting) Descriptor() ([]byte, []int) {
	return file_checkin_proto_rawDescGZIP(), []int{0}
}

func (x *GservicesSetting) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *GservicesSetting) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Devices send this every few hours to tell us how they're doing.
type AndroidCheckinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IMEI (used by GSM phones) is sent and stored as 15 decimal
	// digits; the 15th is a check digit.
	Imei *string `protobuf:"bytes,1,opt,name=imei" json:"imei,omitempty"` // IMEI, reported but not logged.
	// MEID (used by CDMA phones) is sent and stored as 14 hexadecimal
	// digits (no check digit).
	Meid *string `protobuf:"bytes,10,opt,name=meid" json:"meid,omitempty"` // MEID, reported but not logged.
	// MAC address (used by non-phone devices).  12 hexadecimal digits;
	// no separators (eg "0016E6513AC2", not "00:16:E6:51:3A:C2").
	MacAddr []string `protobuf:"bytes,9,rep,name=mac_addr,json=macAddr" json:"mac_addr,omitempty"` // MAC address, reported but not logged.
	// An array parallel to mac_addr, describing the type of interface.
	// Currently accepted values: "wifi", "ethernet", "bluetooth".  If
	// not present, "wifi" is assumed.
	MacAddrType []string `protobuf:"bytes,19,rep,name=mac_addr_type,json=macAddrType" json:"mac_addr_type,omitempty"`
	// Serial number (a manufacturer-defined unique hardware
	// identifier).  Alphanumeric, case-insensitive.
	SerialNumber *string `protobuf:"bytes,16,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	// Older CDMA networks use an ESN (8 hex digits) instead of an MEID.
	Esn       *string              `protobuf:"bytes,17,opt,name=esn" json:"esn,omitempty"`                              // ESN, reported but not logged
	Id        *int64               `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`                                // Android device ID, not logged
	LoggingId *int64               `protobuf:"varint,7,opt,name=logging_id,json=loggingId" json:"logging_id,omitempty"` // Pseudonymous logging ID for Sawmill
	Digest    *string              `protobuf:"bytes,3,opt,name=digest" json:"digest,omitempty"`                         // Digest of device provisioning, not logged.
	Locale    *string              `protobuf:"bytes,6,opt,name=locale" json:"locale,omitempty"`                         // Current locale in standard (xx_XX) format
	Checkin   *AndroidCheckinProto `protobuf:"bytes,4,req,name=checkin" json:"checkin,omitempty"`
	// DEPRECATED, see AndroidCheckinProto.requested_group
	DesiredBuild *string `protobuf:"bytes,5,opt,name=desired_build,json=desiredBuild" json:"desired_build,omitempty"`
	// Blob of data from the Market app to be passed to Market API server
	MarketCheckin *string `protobuf:"bytes,8,opt,name=market_checkin,json=marketCheckin" json:"market_checkin,omitempty"`
	// SID cookies of any google accounts stored on the phone.  Not logged.
	AccountCookie []string `protobuf:"bytes,11,rep,name=account_cookie,json=accountCookie" json:"account_cookie,omitempty"`
	// Time zone.  Not currently logged.
	TimeZone *string `protobuf:"bytes,12,opt,name=time_zone,json=timeZone" json:"time_zone,omitempty"`
	// Security token used to validate the checkin request.
	// Required for android IDs issued to Froyo+ devices, not for legacy IDs.
	SecurityToken *uint64 `protobuf:"fixed64,13,opt,name=security_token,json=securityToken" json:"security_token,omitempty"`
	// Version of checkin protocol.
	//
	// There are currently two versions:
	//
	// - version field missing: android IDs are assigned based on
	//   hardware identifiers.  unsecured in the sense that you can
	//   "unregister" someone's phone by sending a registration request
	//   with their IMEI/MEID/MAC.
	//
	// - version=2: android IDs are assigned randomly.  The device is
	//   sent a security token that must be included in all future
	//   checkins for that android id.
	//
	// - version=3: same as version 2, but the 'fragment' field is
	//   provided, and the device understands incremental updates to the
	//   gservices table (ie, only returning the keys whose values have
	//   changed.)
	//
	// (version=1 was skipped to avoid confusion with the "missing"
	// version field that is effectively version 1.)
	Version *int32 `protobuf:"varint,14,opt,name=version" json:"version,omitempty"`
	// OTA certs accepted by device (base-64 SHA-1 of cert files).  Not
	// logged.
	OtaCert []string `protobuf:"bytes,15,rep,name=ota_cert,json=otaCert" json:"ota_cert,omitempty"`
	// A single CheckinTask on the device may lead to multiple checkin
	// requests if there is too much log data to upload in a single
	// request.  For version 3 and up, this field will be filled in with
	// the number of the request, starting with 0.
	Fragment *int32 `protobuf:"varint,20,opt,name=fragment" json:"fragment,omitempty"`
	// For devices supporting multiple users, the name of the current
	// profile (they all check in independently, just as if they were
	// multiple physical devices).  This may not be set, even if the
	// device is using multiuser.  (checkin.user_number should be set to
	// the ordinal of the user.)
	UserName *string `protobuf:"bytes,21,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	// For devices supporting multiple user profiles, the serial number
	// for the user checking in.  Not logged.  May not be set, even if
	// the device supportes multiuser.  checkin.user_number is the
	// ordinal of the user (0, 1, 2, ...), which may be reused if users
	// are deleted and re-created.  user_serial_number is never reused
	// (unless the device is wiped).
	UserSerialNumber *int32 `protobuf:"varint,22,opt,name=user_serial_number,json=userSerialNumber" json:"user_serial_number,omitempty"`
}

func (x *AndroidCheckinRequest) Reset() {
	*x = AndroidCheckinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndroidCheckinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidCheckinRequest) ProtoMessage() {}

func (x *AndroidCheckinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidCheckinRequest.ProtoReflect.Descriptor instead.
func (*AndroidCheckinRequest) Descriptor() ([]byte, []int) {
	return file_checkin_proto_rawDescGZIP(), []int{1}
}

func (x *AndroidCheckinRequest) GetImei() string {
	if x != nil && x.Imei != nil {
		return *x.Imei
	}
	return ""
}

func (x *AndroidCheckinRequest) GetMeid() string {
	if x != nil && x.Meid != nil {
		return *x.Meid
	}
	return ""
}

func (x *AndroidCheckinRequest) GetMacAddr() []string {
	if x != nil {
		return x.MacAddr
	}
	return nil
}

func (x *AndroidCheckinRequest) GetMacAddrType() []string {
	if x != nil {
		return x.MacAddrType
	}
	return nil
}

func (x *AndroidCheckinRequest) GetSerialNumber() string {
	if x != nil && x.SerialNumber != nil {
		return *x.SerialNumber
	}
	return ""
}

func (x *AndroidCheckinRequest) GetEsn() string {
	if x != nil && x.Esn != nil {
		return *x.Esn
	}
	return ""
}

func (x *AndroidCheckinRequest) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *AndroidCheckinRequest) GetLoggingId() int64 {
	if x != nil && x.LoggingId != nil {
		return *x.LoggingId
	}
	return 0
}

func (x *AndroidCheckinRequest) GetDigest() string {
	if x != nil && x.Digest != nil {
		return *x.Digest
	}
	return ""
}

func (x *AndroidCheckinRequest) GetLocale() string {
	if x != nil && x.Locale != nil {
		return *x.Locale
	}
	return ""
}

func (x *AndroidCheckinRequest) GetCheckin() *AndroidCheckinProto {
	if x != nil {
		return x.Checkin
	}
	return nil
}

func (x *AndroidCheckinRequest) GetDesiredBuild() string {
	if x != nil && x.DesiredBuild != nil {
		return *x.DesiredBuild
	}
	return ""
}

func (x *AndroidCheckinRequest) GetMarketCheckin() string {
	if x != nil && x.MarketCheckin != nil {
		return *x.MarketCheckin
	}
	return ""
}

func (x *AndroidCheckinRequest) GetAccountCookie() []string {
	if x != nil {
		return x.AccountCookie
	}
	return nil
}

func (x *AndroidCheckinRequest) GetTimeZone() string {
	if x != nil && x.TimeZone != nil {
		return *x.TimeZone
	}
	return ""
}

func (x *AndroidCheckinRequest) GetSecurityToken() uint64 {
	if x != nil && x.SecurityToken != nil {
		return *x.SecurityToken
	}
	return 0
}

func (x *AndroidCheckinRequest) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *AndroidCheckinRequest) GetOtaCert() []string {
	if x != nil {
		return x.OtaCert
	}
	return nil
}

func (x *AndroidCheckinRequest) GetFragment() int32 {
	if x != nil && x.Fragment != nil {
		return *x.Fragment
	}
	return 0
}

func (x *AndroidCheckinRequest) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *AndroidCheckinRequest) GetUserSerialNumber() int32 {
	if x != nil && x.UserSerialNumber != nil {
		return *x.UserSerialNumber
	}
	return 0
}

// The response to the device.
type AndroidCheckinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatsOk  *bool  `protobuf:"varint,1,req,name=stats_ok,json=statsOk" json:"stats_ok,omitempty"`    // Whether statistics were recorded properly.
	TimeMsec *int64 `protobuf:"varint,3,opt,name=time_msec,json=timeMsec" json:"time_msec,omitempty"` // Time of day from server (Java epoch).
	// Provisioning is sent if the request included an obsolete digest.
	//
	// For version <= 2, 'digest' contains the digest that should be
	// sent back to the server on the next checkin, and 'setting'
	// contains the entire gservices table (which replaces the entire
	// current table on the device).
	//
	// for version >= 3, 'digest' will be absent.  If 'settings_diff'
	// is false, then 'setting' contains the entire table, as in version
	// 2.  If 'settings_diff' is true, then 'delete_setting' contains
	// the keys to delete, and 'setting' contains only keys to be added
	// or for which the value has changed.  All other keys in the
	// current table should be left untouched.  If 'settings_diff' is
	// absent, don't touch the existing gservices table.
	//
	Digest        *string             `protobuf:"bytes,4,opt,name=digest" json:"digest,omitempty"`
	SettingsDiff  *bool               `protobuf:"varint,9,opt,name=settings_diff,json=settingsDiff" json:"settings_diff,omitempty"`
	DeleteSetting []string            `protobuf:"bytes,10,rep,name=delete_setting,json=deleteSetting" json:"delete_setting,omitempty"`
	Setting       []*GservicesSetting `protobuf:"bytes,5,rep,name=setting" json:"setting,omitempty"`
	MarketOk      *bool               `protobuf:"varint,6,opt,name=market_ok,json=marketOk" json:"market_ok,omitempty"`                 // If Market got the market_checkin data OK.
	AndroidId     *uint64             `protobuf:"fixed64,7,opt,name=android_id,json=androidId" json:"android_id,omitempty"`             // From the request, or newly assigned
	SecurityToken *uint64             `protobuf:"fixed64,8,opt,name=security_token,json=securityToken" json:"security_token,omitempty"` // The associated security token
	VersionInfo   *string             `protobuf:"bytes,11,opt,name=version_info,json=versionInfo" json:"version_info,omitempty"`        // NEXT TAG: 12
}

func (x *AndroidCheckinResponse) Reset() {
	*x = AndroidCheckinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndroidCheckinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidCheckinResponse) ProtoMessage() {}

func (x *AndroidCheckinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidCheckinResponse.ProtoReflect.Descriptor instead.
func (*AndroidCheckinResponse) Descriptor() ([]byte, []int) {
	return file_checkin_proto_rawDescGZIP(), []int{2}
}

func (x *AndroidCheckinResponse) GetStatsOk() bool {
	if x != nil && x.StatsOk != nil {
		return *x.StatsOk
	}
	return false
}

func (x *AndroidCheckinResponse) GetTimeMsec() int64 {
	if x != nil && x.TimeMsec != nil {
		return *x.TimeMsec
	}
	return 0
}

func (x *AndroidCheckinResponse) GetDigest() string {
	if x != nil && x.Digest != nil {
		return *x.Digest
	}
	return ""
}

func (x *AndroidCheckinResponse) GetSettingsDiff() bool {
	if x != nil && x.SettingsDiff != nil {
		return *x.SettingsDiff
	}
	return false
}

func (x *AndroidCheckinResponse) GetDeleteSetting() []string {
	if x != nil {
		return x.DeleteSetting
	}
	return nil
}

func (x *AndroidCheckinResponse) GetSetting() []*GservicesSetting {
	if x != nil {
		return x.Setting
	}
	return nil
}

func (x *AndroidCheckinResponse) GetMarketOk() bool {
	if x != nil && x.MarketOk != nil {
		return *x.MarketOk
	}
	return false
}

func (x *AndroidCheckinResponse) GetAndroidId() uint64 {
	if x != nil && x.AndroidId != nil {
		return *x.AndroidId
	}
	return 0
}

func (x *AndroidCheckinResponse) GetSecurityToken() uint64 {
	if x != nil && x.SecurityToken != nil {
		return *x.SecurityToken
	}
	return 0
}

func (x *AndroidCheckinResponse) GetVersionInfo() string {
	if x != nil && x.VersionInfo != nil {
		return *x.VersionInfo
	}
	return ""
}

var File_checkin_proto protoreflect.FileDescriptor

var file_checkin_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xa5, 0x05, 0x0a, 0x15, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6d, 0x65, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x65,
	0x69, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x65, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x73, 0x6e,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x73, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
	0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x74, 0x61,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x74, 0x61,
	0x43, 0x65, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xf5, 0x02, 0x0a, 0x16,
	0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x73, 0x4f,
	0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x65, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x65, 0x63, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x44, 0x69, 0x66, 0x66, 0x12, 0x25, 0x0a, 0x0e, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x39, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4f, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x06, 0x52, 0x09,
	0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x06, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x13, 0x48, 0x03, 0x5a, 0x0f, 0x2e, 0x3b, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_checkin_proto_rawDescOnce sync.Once
	file_checkin_proto_rawDescData = file_checkin_proto_rawDesc
)

func file_checkin_proto_rawDescGZIP() []byte {
	file_checkin_proto_rawDescOnce.Do(func() {
		file_checkin_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkin_proto_rawDescData)
	})
	return file_checkin_proto_rawDescData
}

var file_checkin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_checkin_proto_goTypes = []interface{}{
	(*GservicesSetting)(nil),       // 0: checkin_proto.GservicesSetting
	(*AndroidCheckinRequest)(nil),  // 1: checkin_proto.AndroidCheckinRequest
	(*AndroidCheckinResponse)(nil), // 2: checkin_proto.AndroidCheckinResponse
	(*AndroidCheckinProto)(nil),    // 3: checkin_proto.AndroidCheckinProto
}
var file_checkin_proto_depIdxs = []int32{
	3, // 0: checkin_proto.AndroidCheckinRequest.checkin:type_name -> checkin_proto.AndroidCheckinProto
	0, // 1: checkin_proto.AndroidCheckinResponse.setting:type_name -> checkin_proto.GservicesSetting
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_checkin_proto_init() }
func file_checkin_proto_init() {
	if File_checkin_proto != nil {
		return
	}
	file_android_checkin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_checkin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GservicesSetting); i {
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
		file_checkin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AndroidCheckinRequest); i {
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
		file_checkin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AndroidCheckinResponse); i {
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
			RawDescriptor: file_checkin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_checkin_proto_goTypes,
		DependencyIndexes: file_checkin_proto_depIdxs,
		MessageInfos:      file_checkin_proto_msgTypes,
	}.Build()
	File_checkin_proto = out.File
	file_checkin_proto_rawDesc = nil
	file_checkin_proto_goTypes = nil
	file_checkin_proto_depIdxs = nil
}
