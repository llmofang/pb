// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: subStock.proto

#define INTERNAL_SUPPRESS_PROTOBUF_FIELD_DEPRECATION
#include "subStock.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/stubs/once.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)

namespace comm {

namespace {

const ::google::protobuf::Descriptor* SubStock_descriptor_ = NULL;
const ::google::protobuf::internal::GeneratedMessageReflection*
  SubStock_reflection_ = NULL;

}  // namespace


void protobuf_AssignDesc_subStock_2eproto() {
  protobuf_AddDesc_subStock_2eproto();
  const ::google::protobuf::FileDescriptor* file =
    ::google::protobuf::DescriptorPool::generated_pool()->FindFileByName(
      "subStock.proto");
  GOOGLE_CHECK(file != NULL);
  SubStock_descriptor_ = file->message_type(0);
  static const int SubStock_offsets_[2] = {
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(SubStock, stock_code_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(SubStock, sub_flag_),
  };
  SubStock_reflection_ =
    new ::google::protobuf::internal::GeneratedMessageReflection(
      SubStock_descriptor_,
      SubStock::default_instance_,
      SubStock_offsets_,
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(SubStock, _has_bits_[0]),
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(SubStock, _unknown_fields_),
      -1,
      ::google::protobuf::DescriptorPool::generated_pool(),
      ::google::protobuf::MessageFactory::generated_factory(),
      sizeof(SubStock));
}

namespace {

GOOGLE_PROTOBUF_DECLARE_ONCE(protobuf_AssignDescriptors_once_);
inline void protobuf_AssignDescriptorsOnce() {
  ::google::protobuf::GoogleOnceInit(&protobuf_AssignDescriptors_once_,
                 &protobuf_AssignDesc_subStock_2eproto);
}

void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedMessage(
    SubStock_descriptor_, &SubStock::default_instance());
}

}  // namespace

void protobuf_ShutdownFile_subStock_2eproto() {
  delete SubStock::default_instance_;
  delete SubStock_reflection_;
}

void protobuf_AddDesc_subStock_2eproto() {
  static bool already_here = false;
  if (already_here) return;
  already_here = true;
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
    "\n\016subStock.proto\022\004comm\"0\n\010SubStock\022\022\n\nst"
    "ock_code\030\001 \002(\014\022\020\n\010sub_flag\030\002 \002(\010", 72);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "subStock.proto", &protobuf_RegisterTypes);
  SubStock::default_instance_ = new SubStock();
  SubStock::default_instance_->InitAsDefaultInstance();
  ::google::protobuf::internal::OnShutdown(&protobuf_ShutdownFile_subStock_2eproto);
}

// Force AddDescriptors() to be called at static initialization time.
struct StaticDescriptorInitializer_subStock_2eproto {
  StaticDescriptorInitializer_subStock_2eproto() {
    protobuf_AddDesc_subStock_2eproto();
  }
} static_descriptor_initializer_subStock_2eproto_;

// ===================================================================

#ifndef _MSC_VER
const int SubStock::kStockCodeFieldNumber;
const int SubStock::kSubFlagFieldNumber;
#endif  // !_MSC_VER

SubStock::SubStock()
  : ::google::protobuf::Message() {
  SharedCtor();
  // @@protoc_insertion_point(constructor:comm.SubStock)
}

void SubStock::InitAsDefaultInstance() {
}

SubStock::SubStock(const SubStock& from)
  : ::google::protobuf::Message() {
  SharedCtor();
  MergeFrom(from);
  // @@protoc_insertion_point(copy_constructor:comm.SubStock)
}

void SubStock::SharedCtor() {
  ::google::protobuf::internal::GetEmptyString();
  _cached_size_ = 0;
  stock_code_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  sub_flag_ = false;
  ::memset(_has_bits_, 0, sizeof(_has_bits_));
}

SubStock::~SubStock() {
  // @@protoc_insertion_point(destructor:comm.SubStock)
  SharedDtor();
}

void SubStock::SharedDtor() {
  if (stock_code_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete stock_code_;
  }
  if (this != default_instance_) {
  }
}

void SubStock::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* SubStock::descriptor() {
  protobuf_AssignDescriptorsOnce();
  return SubStock_descriptor_;
}

const SubStock& SubStock::default_instance() {
  if (default_instance_ == NULL) protobuf_AddDesc_subStock_2eproto();
  return *default_instance_;
}

SubStock* SubStock::default_instance_ = NULL;

SubStock* SubStock::New() const {
  return new SubStock;
}

void SubStock::Clear() {
  if (_has_bits_[0 / 32] & 3) {
    if (has_stock_code()) {
      if (stock_code_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
        stock_code_->clear();
      }
    }
    sub_flag_ = false;
  }
  ::memset(_has_bits_, 0, sizeof(_has_bits_));
  mutable_unknown_fields()->Clear();
}

bool SubStock::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:comm.SubStock)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoff(127);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // required bytes stock_code = 1;
      case 1: {
        if (tag == 10) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadBytes(
                input, this->mutable_stock_code()));
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(16)) goto parse_sub_flag;
        break;
      }

      // required bool sub_flag = 2;
      case 2: {
        if (tag == 16) {
         parse_sub_flag:
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   bool, ::google::protobuf::internal::WireFormatLite::TYPE_BOOL>(
                 input, &sub_flag_)));
          set_has_sub_flag();
        } else {
          goto handle_unusual;
        }
        if (input->ExpectAtEnd()) goto success;
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0 ||
            ::google::protobuf::internal::WireFormatLite::GetTagWireType(tag) ==
            ::google::protobuf::internal::WireFormatLite::WIRETYPE_END_GROUP) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:comm.SubStock)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:comm.SubStock)
  return false;
#undef DO_
}

void SubStock::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:comm.SubStock)
  // required bytes stock_code = 1;
  if (has_stock_code()) {
    ::google::protobuf::internal::WireFormatLite::WriteBytesMaybeAliased(
      1, this->stock_code(), output);
  }

  // required bool sub_flag = 2;
  if (has_sub_flag()) {
    ::google::protobuf::internal::WireFormatLite::WriteBool(2, this->sub_flag(), output);
  }

  if (!unknown_fields().empty()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:comm.SubStock)
}

::google::protobuf::uint8* SubStock::SerializeWithCachedSizesToArray(
    ::google::protobuf::uint8* target) const {
  // @@protoc_insertion_point(serialize_to_array_start:comm.SubStock)
  // required bytes stock_code = 1;
  if (has_stock_code()) {
    target =
      ::google::protobuf::internal::WireFormatLite::WriteBytesToArray(
        1, this->stock_code(), target);
  }

  // required bool sub_flag = 2;
  if (has_sub_flag()) {
    target = ::google::protobuf::internal::WireFormatLite::WriteBoolToArray(2, this->sub_flag(), target);
  }

  if (!unknown_fields().empty()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:comm.SubStock)
  return target;
}

int SubStock::ByteSize() const {
  int total_size = 0;

  if (_has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    // required bytes stock_code = 1;
    if (has_stock_code()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::BytesSize(
          this->stock_code());
    }

    // required bool sub_flag = 2;
    if (has_sub_flag()) {
      total_size += 1 + 1;
    }

  }
  if (!unknown_fields().empty()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        unknown_fields());
  }
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = total_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void SubStock::MergeFrom(const ::google::protobuf::Message& from) {
  GOOGLE_CHECK_NE(&from, this);
  const SubStock* source =
    ::google::protobuf::internal::dynamic_cast_if_available<const SubStock*>(
      &from);
  if (source == NULL) {
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
    MergeFrom(*source);
  }
}

void SubStock::MergeFrom(const SubStock& from) {
  GOOGLE_CHECK_NE(&from, this);
  if (from._has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    if (from.has_stock_code()) {
      set_stock_code(from.stock_code());
    }
    if (from.has_sub_flag()) {
      set_sub_flag(from.sub_flag());
    }
  }
  mutable_unknown_fields()->MergeFrom(from.unknown_fields());
}

void SubStock::CopyFrom(const ::google::protobuf::Message& from) {
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void SubStock::CopyFrom(const SubStock& from) {
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool SubStock::IsInitialized() const {
  if ((_has_bits_[0] & 0x00000003) != 0x00000003) return false;

  return true;
}

void SubStock::Swap(SubStock* other) {
  if (other != this) {
    std::swap(stock_code_, other->stock_code_);
    std::swap(sub_flag_, other->sub_flag_);
    std::swap(_has_bits_[0], other->_has_bits_[0]);
    _unknown_fields_.Swap(&other->_unknown_fields_);
    std::swap(_cached_size_, other->_cached_size_);
  }
}

::google::protobuf::Metadata SubStock::GetMetadata() const {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::Metadata metadata;
  metadata.descriptor = SubStock_descriptor_;
  metadata.reflection = SubStock_reflection_;
  return metadata;
}


// @@protoc_insertion_point(namespace_scope)

}  // namespace comm

// @@protoc_insertion_point(global_scope)
