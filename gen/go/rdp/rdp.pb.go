// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.25.1
// source: rdp/rdp.proto

package rdpv1

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

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Content  []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rdp_rdp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rdp_rdp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_rdp_rdp_proto_rawDescGZIP(), []int{0}
}

func (x *FileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rdp_rdp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rdp_rdp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_rdp_rdp_proto_rawDescGZIP(), []int{1}
}

func (x *FileResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rdp_rdp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_rdp_rdp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_rdp_rdp_proto_rawDescGZIP(), []int{2}
}

func (x *FileChunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rdp_rdp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_rdp_rdp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_rdp_rdp_proto_rawDescGZIP(), []int{3}
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename  string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	CreatedAt string `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rdp_rdp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rdp_rdp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_rdp_rdp_proto_rawDescGZIP(), []int{4}
}

func (x *FileInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FileInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type FilesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*FileInfo `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *FilesList) Reset() {
	*x = FilesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rdp_rdp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesList) ProtoMessage() {}

func (x *FilesList) ProtoReflect() protoreflect.Message {
	mi := &file_rdp_rdp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesList.ProtoReflect.Descriptor instead.
func (*FilesList) Descriptor() ([]byte, []int) {
	return file_rdp_rdp_proto_rawDescGZIP(), []int{5}
}

func (x *FilesList) GetFiles() []*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_rdp_rdp_proto protoreflect.FileDescriptor

var file_rdp_rdp_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x64, 0x70, 0x2f, 0x72, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x22, 0x43, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x44, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x64, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x38, 0x0a,
	0x09, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x6d, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x32, 0xcd, 0x01, 0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x72,
	0x69, 0x75, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x72, 0x69,
	0x75, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x30, 0x01, 0x12, 0x37,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x16, 0x5a, 0x14, 0x6c, 0x65, 0x67, 0x6f, 0x39,
	0x36, 0x33, 0x2e, 0x72, 0x64, 0x70, 0x2e, 0x76, 0x31, 0x3b, 0x72, 0x64, 0x70, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rdp_rdp_proto_rawDescOnce sync.Once
	file_rdp_rdp_proto_rawDescData = file_rdp_rdp_proto_rawDesc
)

func file_rdp_rdp_proto_rawDescGZIP() []byte {
	file_rdp_rdp_proto_rawDescOnce.Do(func() {
		file_rdp_rdp_proto_rawDescData = protoimpl.X.CompressGZIP(file_rdp_rdp_proto_rawDescData)
	})
	return file_rdp_rdp_proto_rawDescData
}

var file_rdp_rdp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rdp_rdp_proto_goTypes = []interface{}{
	(*FileRequest)(nil),  // 0: imaginarium.FileRequest
	(*FileResponse)(nil), // 1: imaginarium.FileResponse
	(*FileChunk)(nil),    // 2: imaginarium.FileChunk
	(*Empty)(nil),        // 3: imaginarium.Empty
	(*FileInfo)(nil),     // 4: imaginarium.FileInfo
	(*FilesList)(nil),    // 5: imaginarium.FilesList
}
var file_rdp_rdp_proto_depIdxs = []int32{
	4, // 0: imaginarium.FilesList.files:type_name -> imaginarium.FileInfo
	0, // 1: imaginarium.Imaginarium.UploadFile:input_type -> imaginarium.FileRequest
	0, // 2: imaginarium.Imaginarium.DownloadFile:input_type -> imaginarium.FileRequest
	3, // 3: imaginarium.Imaginarium.ListFiles:input_type -> imaginarium.Empty
	1, // 4: imaginarium.Imaginarium.UploadFile:output_type -> imaginarium.FileResponse
	2, // 5: imaginarium.Imaginarium.DownloadFile:output_type -> imaginarium.FileChunk
	5, // 6: imaginarium.Imaginarium.ListFiles:output_type -> imaginarium.FilesList
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rdp_rdp_proto_init() }
func file_rdp_rdp_proto_init() {
	if File_rdp_rdp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rdp_rdp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
		file_rdp_rdp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponse); i {
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
		file_rdp_rdp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChunk); i {
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
		file_rdp_rdp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_rdp_rdp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_rdp_rdp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesList); i {
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
			RawDescriptor: file_rdp_rdp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rdp_rdp_proto_goTypes,
		DependencyIndexes: file_rdp_rdp_proto_depIdxs,
		MessageInfos:      file_rdp_rdp_proto_msgTypes,
	}.Build()
	File_rdp_rdp_proto = out.File
	file_rdp_rdp_proto_rawDesc = nil
	file_rdp_rdp_proto_goTypes = nil
	file_rdp_rdp_proto_depIdxs = nil
}
