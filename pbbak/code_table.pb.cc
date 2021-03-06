// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: code_table.proto

#define INTERNAL_SUPPRESS_PROTOBUF_FIELD_DEPRECATION
#include "code_table.pb.h"

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

const ::google::protobuf::Descriptor* TdfCode_descriptor_ = NULL;
const ::google::protobuf::internal::GeneratedMessageReflection*
  TdfCode_reflection_ = NULL;
const ::google::protobuf::Descriptor* TdfCodeTable_descriptor_ = NULL;
const ::google::protobuf::internal::GeneratedMessageReflection*
  TdfCodeTable_reflection_ = NULL;

}  // namespace


void protobuf_AssignDesc_code_5ftable_2eproto() {
  protobuf_AddDesc_code_5ftable_2eproto();
  const ::google::protobuf::FileDescriptor* file =
    ::google::protobuf::DescriptorPool::generated_pool()->FindFileByName(
      "code_table.proto");
  GOOGLE_CHECK(file != NULL);
  TdfCode_descriptor_ = file->message_type(0);
  static const int TdfCode_offsets_[6] = {
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCode, windcode_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCode, market_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCode, code_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCode, enname_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCode, cnname_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCode, type_),
  };
  TdfCode_reflection_ =
    new ::google::protobuf::internal::GeneratedMessageReflection(
      TdfCode_descriptor_,
      TdfCode::default_instance_,
      TdfCode_offsets_,
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCode, _has_bits_[0]),
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCode, _unknown_fields_),
      -1,
      ::google::protobuf::DescriptorPool::generated_pool(),
      ::google::protobuf::MessageFactory::generated_factory(),
      sizeof(TdfCode));
  TdfCodeTable_descriptor_ = file->message_type(1);
  static const int TdfCodeTable_offsets_[1] = {
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCodeTable, codetable_),
  };
  TdfCodeTable_reflection_ =
    new ::google::protobuf::internal::GeneratedMessageReflection(
      TdfCodeTable_descriptor_,
      TdfCodeTable::default_instance_,
      TdfCodeTable_offsets_,
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCodeTable, _has_bits_[0]),
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(TdfCodeTable, _unknown_fields_),
      -1,
      ::google::protobuf::DescriptorPool::generated_pool(),
      ::google::protobuf::MessageFactory::generated_factory(),
      sizeof(TdfCodeTable));
}

namespace {

GOOGLE_PROTOBUF_DECLARE_ONCE(protobuf_AssignDescriptors_once_);
inline void protobuf_AssignDescriptorsOnce() {
  ::google::protobuf::GoogleOnceInit(&protobuf_AssignDescriptors_once_,
                 &protobuf_AssignDesc_code_5ftable_2eproto);
}

void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedMessage(
    TdfCode_descriptor_, &TdfCode::default_instance());
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedMessage(
    TdfCodeTable_descriptor_, &TdfCodeTable::default_instance());
}

}  // namespace

void protobuf_ShutdownFile_code_5ftable_2eproto() {
  delete TdfCode::default_instance_;
  delete TdfCode_reflection_;
  delete TdfCodeTable::default_instance_;
  delete TdfCodeTable_reflection_;
}

void protobuf_AddDesc_code_5ftable_2eproto() {
  static bool already_here = false;
  if (already_here) return;
  already_here = true;
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
    "\n\020code_table.proto\022\004comm\"g\n\007TdfCode\022\020\n\010W"
    "indCode\030\001 \002(\014\022\016\n\006Market\030\002 \002(\014\022\014\n\004Code\030\003 "
    "\002(\014\022\016\n\006ENName\030\004 \002(\014\022\016\n\006CNName\030\005 \002(\014\022\014\n\004T"
    "ype\030\006 \002(\005\"0\n\014TdfCodeTable\022 \n\tCodeTable\030\001"
    " \003(\0132\r.comm.TdfCode", 179);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "code_table.proto", &protobuf_RegisterTypes);
  TdfCode::default_instance_ = new TdfCode();
  TdfCodeTable::default_instance_ = new TdfCodeTable();
  TdfCode::default_instance_->InitAsDefaultInstance();
  TdfCodeTable::default_instance_->InitAsDefaultInstance();
  ::google::protobuf::internal::OnShutdown(&protobuf_ShutdownFile_code_5ftable_2eproto);
}

// Force AddDescriptors() to be called at static initialization time.
struct StaticDescriptorInitializer_code_5ftable_2eproto {
  StaticDescriptorInitializer_code_5ftable_2eproto() {
    protobuf_AddDesc_code_5ftable_2eproto();
  }
} static_descriptor_initializer_code_5ftable_2eproto_;

// ===================================================================

#ifndef _MSC_VER
const int TdfCode::kWindCodeFieldNumber;
const int TdfCode::kMarketFieldNumber;
const int TdfCode::kCodeFieldNumber;
const int TdfCode::kENNameFieldNumber;
const int TdfCode::kCNNameFieldNumber;
const int TdfCode::kTypeFieldNumber;
#endif  // !_MSC_VER

TdfCode::TdfCode()
  : ::google::protobuf::Message() {
  SharedCtor();
  // @@protoc_insertion_point(constructor:comm.TdfCode)
}

void TdfCode::InitAsDefaultInstance() {
}

TdfCode::TdfCode(const TdfCode& from)
  : ::google::protobuf::Message() {
  SharedCtor();
  MergeFrom(from);
  // @@protoc_insertion_point(copy_constructor:comm.TdfCode)
}

void TdfCode::SharedCtor() {
  ::google::protobuf::internal::GetEmptyString();
  _cached_size_ = 0;
  windcode_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  market_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  code_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  enname_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  cnname_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  type_ = 0;
  ::memset(_has_bits_, 0, sizeof(_has_bits_));
}

TdfCode::~TdfCode() {
  // @@protoc_insertion_point(destructor:comm.TdfCode)
  SharedDtor();
}

void TdfCode::SharedDtor() {
  if (windcode_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete windcode_;
  }
  if (market_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete market_;
  }
  if (code_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete code_;
  }
  if (enname_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete enname_;
  }
  if (cnname_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete cnname_;
  }
  if (this != default_instance_) {
  }
}

void TdfCode::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* TdfCode::descriptor() {
  protobuf_AssignDescriptorsOnce();
  return TdfCode_descriptor_;
}

const TdfCode& TdfCode::default_instance() {
  if (default_instance_ == NULL) protobuf_AddDesc_code_5ftable_2eproto();
  return *default_instance_;
}

TdfCode* TdfCode::default_instance_ = NULL;

TdfCode* TdfCode::New() const {
  return new TdfCode;
}

void TdfCode::Clear() {
  if (_has_bits_[0 / 32] & 63) {
    if (has_windcode()) {
      if (windcode_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
        windcode_->clear();
      }
    }
    if (has_market()) {
      if (market_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
        market_->clear();
      }
    }
    if (has_code()) {
      if (code_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
        code_->clear();
      }
    }
    if (has_enname()) {
      if (enname_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
        enname_->clear();
      }
    }
    if (has_cnname()) {
      if (cnname_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
        cnname_->clear();
      }
    }
    type_ = 0;
  }
  ::memset(_has_bits_, 0, sizeof(_has_bits_));
  mutable_unknown_fields()->Clear();
}

bool TdfCode::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:comm.TdfCode)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoff(127);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // required bytes WindCode = 1;
      case 1: {
        if (tag == 10) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadBytes(
                input, this->mutable_windcode()));
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(18)) goto parse_Market;
        break;
      }

      // required bytes Market = 2;
      case 2: {
        if (tag == 18) {
         parse_Market:
          DO_(::google::protobuf::internal::WireFormatLite::ReadBytes(
                input, this->mutable_market()));
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(26)) goto parse_Code;
        break;
      }

      // required bytes Code = 3;
      case 3: {
        if (tag == 26) {
         parse_Code:
          DO_(::google::protobuf::internal::WireFormatLite::ReadBytes(
                input, this->mutable_code()));
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(34)) goto parse_ENName;
        break;
      }

      // required bytes ENName = 4;
      case 4: {
        if (tag == 34) {
         parse_ENName:
          DO_(::google::protobuf::internal::WireFormatLite::ReadBytes(
                input, this->mutable_enname()));
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(42)) goto parse_CNName;
        break;
      }

      // required bytes CNName = 5;
      case 5: {
        if (tag == 42) {
         parse_CNName:
          DO_(::google::protobuf::internal::WireFormatLite::ReadBytes(
                input, this->mutable_cnname()));
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(48)) goto parse_Type;
        break;
      }

      // required int32 Type = 6;
      case 6: {
        if (tag == 48) {
         parse_Type:
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &type_)));
          set_has_type();
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
  // @@protoc_insertion_point(parse_success:comm.TdfCode)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:comm.TdfCode)
  return false;
#undef DO_
}

void TdfCode::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:comm.TdfCode)
  // required bytes WindCode = 1;
  if (has_windcode()) {
    ::google::protobuf::internal::WireFormatLite::WriteBytesMaybeAliased(
      1, this->windcode(), output);
  }

  // required bytes Market = 2;
  if (has_market()) {
    ::google::protobuf::internal::WireFormatLite::WriteBytesMaybeAliased(
      2, this->market(), output);
  }

  // required bytes Code = 3;
  if (has_code()) {
    ::google::protobuf::internal::WireFormatLite::WriteBytesMaybeAliased(
      3, this->code(), output);
  }

  // required bytes ENName = 4;
  if (has_enname()) {
    ::google::protobuf::internal::WireFormatLite::WriteBytesMaybeAliased(
      4, this->enname(), output);
  }

  // required bytes CNName = 5;
  if (has_cnname()) {
    ::google::protobuf::internal::WireFormatLite::WriteBytesMaybeAliased(
      5, this->cnname(), output);
  }

  // required int32 Type = 6;
  if (has_type()) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(6, this->type(), output);
  }

  if (!unknown_fields().empty()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:comm.TdfCode)
}

::google::protobuf::uint8* TdfCode::SerializeWithCachedSizesToArray(
    ::google::protobuf::uint8* target) const {
  // @@protoc_insertion_point(serialize_to_array_start:comm.TdfCode)
  // required bytes WindCode = 1;
  if (has_windcode()) {
    target =
      ::google::protobuf::internal::WireFormatLite::WriteBytesToArray(
        1, this->windcode(), target);
  }

  // required bytes Market = 2;
  if (has_market()) {
    target =
      ::google::protobuf::internal::WireFormatLite::WriteBytesToArray(
        2, this->market(), target);
  }

  // required bytes Code = 3;
  if (has_code()) {
    target =
      ::google::protobuf::internal::WireFormatLite::WriteBytesToArray(
        3, this->code(), target);
  }

  // required bytes ENName = 4;
  if (has_enname()) {
    target =
      ::google::protobuf::internal::WireFormatLite::WriteBytesToArray(
        4, this->enname(), target);
  }

  // required bytes CNName = 5;
  if (has_cnname()) {
    target =
      ::google::protobuf::internal::WireFormatLite::WriteBytesToArray(
        5, this->cnname(), target);
  }

  // required int32 Type = 6;
  if (has_type()) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(6, this->type(), target);
  }

  if (!unknown_fields().empty()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:comm.TdfCode)
  return target;
}

int TdfCode::ByteSize() const {
  int total_size = 0;

  if (_has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    // required bytes WindCode = 1;
    if (has_windcode()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::BytesSize(
          this->windcode());
    }

    // required bytes Market = 2;
    if (has_market()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::BytesSize(
          this->market());
    }

    // required bytes Code = 3;
    if (has_code()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::BytesSize(
          this->code());
    }

    // required bytes ENName = 4;
    if (has_enname()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::BytesSize(
          this->enname());
    }

    // required bytes CNName = 5;
    if (has_cnname()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::BytesSize(
          this->cnname());
    }

    // required int32 Type = 6;
    if (has_type()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::Int32Size(
          this->type());
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

void TdfCode::MergeFrom(const ::google::protobuf::Message& from) {
  GOOGLE_CHECK_NE(&from, this);
  const TdfCode* source =
    ::google::protobuf::internal::dynamic_cast_if_available<const TdfCode*>(
      &from);
  if (source == NULL) {
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
    MergeFrom(*source);
  }
}

void TdfCode::MergeFrom(const TdfCode& from) {
  GOOGLE_CHECK_NE(&from, this);
  if (from._has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    if (from.has_windcode()) {
      set_windcode(from.windcode());
    }
    if (from.has_market()) {
      set_market(from.market());
    }
    if (from.has_code()) {
      set_code(from.code());
    }
    if (from.has_enname()) {
      set_enname(from.enname());
    }
    if (from.has_cnname()) {
      set_cnname(from.cnname());
    }
    if (from.has_type()) {
      set_type(from.type());
    }
  }
  mutable_unknown_fields()->MergeFrom(from.unknown_fields());
}

void TdfCode::CopyFrom(const ::google::protobuf::Message& from) {
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void TdfCode::CopyFrom(const TdfCode& from) {
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool TdfCode::IsInitialized() const {
  if ((_has_bits_[0] & 0x0000003f) != 0x0000003f) return false;

  return true;
}

void TdfCode::Swap(TdfCode* other) {
  if (other != this) {
    std::swap(windcode_, other->windcode_);
    std::swap(market_, other->market_);
    std::swap(code_, other->code_);
    std::swap(enname_, other->enname_);
    std::swap(cnname_, other->cnname_);
    std::swap(type_, other->type_);
    std::swap(_has_bits_[0], other->_has_bits_[0]);
    _unknown_fields_.Swap(&other->_unknown_fields_);
    std::swap(_cached_size_, other->_cached_size_);
  }
}

::google::protobuf::Metadata TdfCode::GetMetadata() const {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::Metadata metadata;
  metadata.descriptor = TdfCode_descriptor_;
  metadata.reflection = TdfCode_reflection_;
  return metadata;
}


// ===================================================================

#ifndef _MSC_VER
const int TdfCodeTable::kCodeTableFieldNumber;
#endif  // !_MSC_VER

TdfCodeTable::TdfCodeTable()
  : ::google::protobuf::Message() {
  SharedCtor();
  // @@protoc_insertion_point(constructor:comm.TdfCodeTable)
}

void TdfCodeTable::InitAsDefaultInstance() {
}

TdfCodeTable::TdfCodeTable(const TdfCodeTable& from)
  : ::google::protobuf::Message() {
  SharedCtor();
  MergeFrom(from);
  // @@protoc_insertion_point(copy_constructor:comm.TdfCodeTable)
}

void TdfCodeTable::SharedCtor() {
  _cached_size_ = 0;
  ::memset(_has_bits_, 0, sizeof(_has_bits_));
}

TdfCodeTable::~TdfCodeTable() {
  // @@protoc_insertion_point(destructor:comm.TdfCodeTable)
  SharedDtor();
}

void TdfCodeTable::SharedDtor() {
  if (this != default_instance_) {
  }
}

void TdfCodeTable::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* TdfCodeTable::descriptor() {
  protobuf_AssignDescriptorsOnce();
  return TdfCodeTable_descriptor_;
}

const TdfCodeTable& TdfCodeTable::default_instance() {
  if (default_instance_ == NULL) protobuf_AddDesc_code_5ftable_2eproto();
  return *default_instance_;
}

TdfCodeTable* TdfCodeTable::default_instance_ = NULL;

TdfCodeTable* TdfCodeTable::New() const {
  return new TdfCodeTable;
}

void TdfCodeTable::Clear() {
  codetable_.Clear();
  ::memset(_has_bits_, 0, sizeof(_has_bits_));
  mutable_unknown_fields()->Clear();
}

bool TdfCodeTable::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:comm.TdfCodeTable)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoff(127);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // repeated .comm.TdfCode CodeTable = 1;
      case 1: {
        if (tag == 10) {
         parse_CodeTable:
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessageNoVirtual(
                input, add_codetable()));
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(10)) goto parse_CodeTable;
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
  // @@protoc_insertion_point(parse_success:comm.TdfCodeTable)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:comm.TdfCodeTable)
  return false;
#undef DO_
}

void TdfCodeTable::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:comm.TdfCodeTable)
  // repeated .comm.TdfCode CodeTable = 1;
  for (int i = 0; i < this->codetable_size(); i++) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      1, this->codetable(i), output);
  }

  if (!unknown_fields().empty()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:comm.TdfCodeTable)
}

::google::protobuf::uint8* TdfCodeTable::SerializeWithCachedSizesToArray(
    ::google::protobuf::uint8* target) const {
  // @@protoc_insertion_point(serialize_to_array_start:comm.TdfCodeTable)
  // repeated .comm.TdfCode CodeTable = 1;
  for (int i = 0; i < this->codetable_size(); i++) {
    target = ::google::protobuf::internal::WireFormatLite::
      WriteMessageNoVirtualToArray(
        1, this->codetable(i), target);
  }

  if (!unknown_fields().empty()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:comm.TdfCodeTable)
  return target;
}

int TdfCodeTable::ByteSize() const {
  int total_size = 0;

  // repeated .comm.TdfCode CodeTable = 1;
  total_size += 1 * this->codetable_size();
  for (int i = 0; i < this->codetable_size(); i++) {
    total_size +=
      ::google::protobuf::internal::WireFormatLite::MessageSizeNoVirtual(
        this->codetable(i));
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

void TdfCodeTable::MergeFrom(const ::google::protobuf::Message& from) {
  GOOGLE_CHECK_NE(&from, this);
  const TdfCodeTable* source =
    ::google::protobuf::internal::dynamic_cast_if_available<const TdfCodeTable*>(
      &from);
  if (source == NULL) {
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
    MergeFrom(*source);
  }
}

void TdfCodeTable::MergeFrom(const TdfCodeTable& from) {
  GOOGLE_CHECK_NE(&from, this);
  codetable_.MergeFrom(from.codetable_);
  mutable_unknown_fields()->MergeFrom(from.unknown_fields());
}

void TdfCodeTable::CopyFrom(const ::google::protobuf::Message& from) {
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void TdfCodeTable::CopyFrom(const TdfCodeTable& from) {
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool TdfCodeTable::IsInitialized() const {

  if (!::google::protobuf::internal::AllAreInitialized(this->codetable())) return false;
  return true;
}

void TdfCodeTable::Swap(TdfCodeTable* other) {
  if (other != this) {
    codetable_.Swap(&other->codetable_);
    std::swap(_has_bits_[0], other->_has_bits_[0]);
    _unknown_fields_.Swap(&other->_unknown_fields_);
    std::swap(_cached_size_, other->_cached_size_);
  }
}

::google::protobuf::Metadata TdfCodeTable::GetMetadata() const {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::Metadata metadata;
  metadata.descriptor = TdfCodeTable_descriptor_;
  metadata.reflection = TdfCodeTable_reflection_;
  return metadata;
}


// @@protoc_insertion_point(namespace_scope)

}  // namespace comm

// @@protoc_insertion_point(global_scope)
