// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: request.proto

#define INTERNAL_SUPPRESS_PROTOBUF_FIELD_DEPRECATION
#include "request.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/stubs/port.h>
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

const ::google::protobuf::Descriptor* Request_descriptor_ = NULL;
const ::google::protobuf::internal::GeneratedMessageReflection*
  Request_reflection_ = NULL;

}  // namespace


void protobuf_AssignDesc_request_2eproto() GOOGLE_ATTRIBUTE_COLD;
void protobuf_AssignDesc_request_2eproto() {
  protobuf_AddDesc_request_2eproto();
  const ::google::protobuf::FileDescriptor* file =
    ::google::protobuf::DescriptorPool::generated_pool()->FindFileByName(
      "request.proto");
  GOOGLE_CHECK(file != NULL);
  Request_descriptor_ = file->message_type(0);
  static const int Request_offsets_[6] = {
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Request, qid_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Request, market_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Request, obj_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Request, starttime_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Request, endtime_),
    GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Request, dynatype_),
  };
  Request_reflection_ =
    ::google::protobuf::internal::GeneratedMessageReflection::NewGeneratedMessageReflection(
      Request_descriptor_,
      Request::internal_default_instance(),
      Request_offsets_,
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Request, _has_bits_),
      -1,
      -1,
      sizeof(Request),
      GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(Request, _internal_metadata_));
}

namespace {

GOOGLE_PROTOBUF_DECLARE_ONCE(protobuf_AssignDescriptors_once_);
void protobuf_AssignDescriptorsOnce() {
  ::google::protobuf::GoogleOnceInit(&protobuf_AssignDescriptors_once_,
                 &protobuf_AssignDesc_request_2eproto);
}

void protobuf_RegisterTypes(const ::std::string&) GOOGLE_ATTRIBUTE_COLD;
void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedMessage(
      Request_descriptor_, Request::internal_default_instance());
}

}  // namespace

void protobuf_ShutdownFile_request_2eproto() {
  Request_default_instance_.Shutdown();
  delete Request_reflection_;
}

void protobuf_InitDefaults_request_2eproto_impl() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  ::google::protobuf::internal::GetEmptyString();
  Request_default_instance_.DefaultConstruct();
  Request_default_instance_.get_mutable()->InitAsDefaultInstance();
}

GOOGLE_PROTOBUF_DECLARE_ONCE(protobuf_InitDefaults_request_2eproto_once_);
void protobuf_InitDefaults_request_2eproto() {
  ::google::protobuf::GoogleOnceInit(&protobuf_InitDefaults_request_2eproto_once_,
                 &protobuf_InitDefaults_request_2eproto_impl);
}
void protobuf_AddDesc_request_2eproto_impl() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  protobuf_InitDefaults_request_2eproto();
  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
    "\n\rrequest.proto\022\004comm\"i\n\007Request\022\013\n\003Qid\030"
    "\001 \002(\003\022\016\n\006Market\030\002 \002(\t\022\013\n\003Obj\030\003 \002(\t\022\021\n\tSt"
    "artTime\030\004 \002(\003\022\017\n\007EndTime\030\005 \002(\003\022\020\n\010DynaTy"
    "pe\030\006 \002(\003", 128);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "request.proto", &protobuf_RegisterTypes);
  ::google::protobuf::internal::OnShutdown(&protobuf_ShutdownFile_request_2eproto);
}

GOOGLE_PROTOBUF_DECLARE_ONCE(protobuf_AddDesc_request_2eproto_once_);
void protobuf_AddDesc_request_2eproto() {
  ::google::protobuf::GoogleOnceInit(&protobuf_AddDesc_request_2eproto_once_,
                 &protobuf_AddDesc_request_2eproto_impl);
}
// Force AddDescriptors() to be called at static initialization time.
struct StaticDescriptorInitializer_request_2eproto {
  StaticDescriptorInitializer_request_2eproto() {
    protobuf_AddDesc_request_2eproto();
  }
} static_descriptor_initializer_request_2eproto_;

namespace {

static void MergeFromFail(int line) GOOGLE_ATTRIBUTE_COLD GOOGLE_ATTRIBUTE_NORETURN;
static void MergeFromFail(int line) {
  ::google::protobuf::internal::MergeFromFail(__FILE__, line);
}

}  // namespace


// ===================================================================

#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int Request::kQidFieldNumber;
const int Request::kMarketFieldNumber;
const int Request::kObjFieldNumber;
const int Request::kStartTimeFieldNumber;
const int Request::kEndTimeFieldNumber;
const int Request::kDynaTypeFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

Request::Request()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  if (this != internal_default_instance()) protobuf_InitDefaults_request_2eproto();
  SharedCtor();
  // @@protoc_insertion_point(constructor:comm.Request)
}

void Request::InitAsDefaultInstance() {
}

Request::Request(const Request& from)
  : ::google::protobuf::Message(),
    _internal_metadata_(NULL) {
  SharedCtor();
  UnsafeMergeFrom(from);
  // @@protoc_insertion_point(copy_constructor:comm.Request)
}

void Request::SharedCtor() {
  _cached_size_ = 0;
  market_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  obj_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  ::memset(&qid_, 0, reinterpret_cast<char*>(&dynatype_) -
    reinterpret_cast<char*>(&qid_) + sizeof(dynatype_));
}

Request::~Request() {
  // @@protoc_insertion_point(destructor:comm.Request)
  SharedDtor();
}

void Request::SharedDtor() {
  market_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  obj_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}

void Request::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* Request::descriptor() {
  protobuf_AssignDescriptorsOnce();
  return Request_descriptor_;
}

const Request& Request::default_instance() {
  protobuf_InitDefaults_request_2eproto();
  return *internal_default_instance();
}

::google::protobuf::internal::ExplicitlyConstructed<Request> Request_default_instance_;

Request* Request::New(::google::protobuf::Arena* arena) const {
  Request* n = new Request;
  if (arena != NULL) {
    arena->Own(n);
  }
  return n;
}

void Request::Clear() {
// @@protoc_insertion_point(message_clear_start:comm.Request)
#if defined(__clang__)
#define ZR_HELPER_(f) \
  _Pragma("clang diagnostic push") \
  _Pragma("clang diagnostic ignored \"-Winvalid-offsetof\"") \
  __builtin_offsetof(Request, f) \
  _Pragma("clang diagnostic pop")
#else
#define ZR_HELPER_(f) reinterpret_cast<char*>(\
  &reinterpret_cast<Request*>(16)->f)
#endif

#define ZR_(first, last) do {\
  ::memset(&(first), 0,\
           ZR_HELPER_(last) - ZR_HELPER_(first) + sizeof(last));\
} while (0)

  if (_has_bits_[0 / 32] & 63u) {
    ZR_(qid_, dynatype_);
    if (has_market()) {
      market_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
    }
    if (has_obj()) {
      obj_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
    }
  }

#undef ZR_HELPER_
#undef ZR_

  _has_bits_.Clear();
  if (_internal_metadata_.have_unknown_fields()) {
    mutable_unknown_fields()->Clear();
  }
}

bool Request::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:comm.Request)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoff(127);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // required int64 Qid = 1;
      case 1: {
        if (tag == 8) {
          set_has_qid();
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int64, ::google::protobuf::internal::WireFormatLite::TYPE_INT64>(
                 input, &qid_)));
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(18)) goto parse_Market;
        break;
      }

      // required string Market = 2;
      case 2: {
        if (tag == 18) {
         parse_Market:
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_market()));
          ::google::protobuf::internal::WireFormat::VerifyUTF8StringNamedField(
            this->market().data(), this->market().length(),
            ::google::protobuf::internal::WireFormat::PARSE,
            "comm.Request.Market");
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(26)) goto parse_Obj;
        break;
      }

      // required string Obj = 3;
      case 3: {
        if (tag == 26) {
         parse_Obj:
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_obj()));
          ::google::protobuf::internal::WireFormat::VerifyUTF8StringNamedField(
            this->obj().data(), this->obj().length(),
            ::google::protobuf::internal::WireFormat::PARSE,
            "comm.Request.Obj");
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(32)) goto parse_StartTime;
        break;
      }

      // required int64 StartTime = 4;
      case 4: {
        if (tag == 32) {
         parse_StartTime:
          set_has_starttime();
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int64, ::google::protobuf::internal::WireFormatLite::TYPE_INT64>(
                 input, &starttime_)));
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(40)) goto parse_EndTime;
        break;
      }

      // required int64 EndTime = 5;
      case 5: {
        if (tag == 40) {
         parse_EndTime:
          set_has_endtime();
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int64, ::google::protobuf::internal::WireFormatLite::TYPE_INT64>(
                 input, &endtime_)));
        } else {
          goto handle_unusual;
        }
        if (input->ExpectTag(48)) goto parse_DynaType;
        break;
      }

      // required int64 DynaType = 6;
      case 6: {
        if (tag == 48) {
         parse_DynaType:
          set_has_dynatype();
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int64, ::google::protobuf::internal::WireFormatLite::TYPE_INT64>(
                 input, &dynatype_)));
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
  // @@protoc_insertion_point(parse_success:comm.Request)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:comm.Request)
  return false;
#undef DO_
}

void Request::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:comm.Request)
  // required int64 Qid = 1;
  if (has_qid()) {
    ::google::protobuf::internal::WireFormatLite::WriteInt64(1, this->qid(), output);
  }

  // required string Market = 2;
  if (has_market()) {
    ::google::protobuf::internal::WireFormat::VerifyUTF8StringNamedField(
      this->market().data(), this->market().length(),
      ::google::protobuf::internal::WireFormat::SERIALIZE,
      "comm.Request.Market");
    ::google::protobuf::internal::WireFormatLite::WriteStringMaybeAliased(
      2, this->market(), output);
  }

  // required string Obj = 3;
  if (has_obj()) {
    ::google::protobuf::internal::WireFormat::VerifyUTF8StringNamedField(
      this->obj().data(), this->obj().length(),
      ::google::protobuf::internal::WireFormat::SERIALIZE,
      "comm.Request.Obj");
    ::google::protobuf::internal::WireFormatLite::WriteStringMaybeAliased(
      3, this->obj(), output);
  }

  // required int64 StartTime = 4;
  if (has_starttime()) {
    ::google::protobuf::internal::WireFormatLite::WriteInt64(4, this->starttime(), output);
  }

  // required int64 EndTime = 5;
  if (has_endtime()) {
    ::google::protobuf::internal::WireFormatLite::WriteInt64(5, this->endtime(), output);
  }

  // required int64 DynaType = 6;
  if (has_dynatype()) {
    ::google::protobuf::internal::WireFormatLite::WriteInt64(6, this->dynatype(), output);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:comm.Request)
}

::google::protobuf::uint8* Request::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:comm.Request)
  // required int64 Qid = 1;
  if (has_qid()) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt64ToArray(1, this->qid(), target);
  }

  // required string Market = 2;
  if (has_market()) {
    ::google::protobuf::internal::WireFormat::VerifyUTF8StringNamedField(
      this->market().data(), this->market().length(),
      ::google::protobuf::internal::WireFormat::SERIALIZE,
      "comm.Request.Market");
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        2, this->market(), target);
  }

  // required string Obj = 3;
  if (has_obj()) {
    ::google::protobuf::internal::WireFormat::VerifyUTF8StringNamedField(
      this->obj().data(), this->obj().length(),
      ::google::protobuf::internal::WireFormat::SERIALIZE,
      "comm.Request.Obj");
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        3, this->obj(), target);
  }

  // required int64 StartTime = 4;
  if (has_starttime()) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt64ToArray(4, this->starttime(), target);
  }

  // required int64 EndTime = 5;
  if (has_endtime()) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt64ToArray(5, this->endtime(), target);
  }

  // required int64 DynaType = 6;
  if (has_dynatype()) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt64ToArray(6, this->dynatype(), target);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:comm.Request)
  return target;
}

size_t Request::RequiredFieldsByteSizeFallback() const {
// @@protoc_insertion_point(required_fields_byte_size_fallback_start:comm.Request)
  size_t total_size = 0;

  if (has_qid()) {
    // required int64 Qid = 1;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int64Size(
        this->qid());
  }

  if (has_market()) {
    // required string Market = 2;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::StringSize(
        this->market());
  }

  if (has_obj()) {
    // required string Obj = 3;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::StringSize(
        this->obj());
  }

  if (has_starttime()) {
    // required int64 StartTime = 4;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int64Size(
        this->starttime());
  }

  if (has_endtime()) {
    // required int64 EndTime = 5;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int64Size(
        this->endtime());
  }

  if (has_dynatype()) {
    // required int64 DynaType = 6;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int64Size(
        this->dynatype());
  }

  return total_size;
}
size_t Request::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:comm.Request)
  size_t total_size = 0;

  if (((_has_bits_[0] & 0x0000003f) ^ 0x0000003f) == 0) {  // All required fields are present.
    // required int64 Qid = 1;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int64Size(
        this->qid());

    // required string Market = 2;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::StringSize(
        this->market());

    // required string Obj = 3;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::StringSize(
        this->obj());

    // required int64 StartTime = 4;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int64Size(
        this->starttime());

    // required int64 EndTime = 5;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int64Size(
        this->endtime());

    // required int64 DynaType = 6;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int64Size(
        this->dynatype());

  } else {
    total_size += RequiredFieldsByteSizeFallback();
  }
  if (_internal_metadata_.have_unknown_fields()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        unknown_fields());
  }
  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = cached_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void Request::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:comm.Request)
  if (GOOGLE_PREDICT_FALSE(&from == this)) MergeFromFail(__LINE__);
  const Request* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const Request>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:comm.Request)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:comm.Request)
    UnsafeMergeFrom(*source);
  }
}

void Request::MergeFrom(const Request& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:comm.Request)
  if (GOOGLE_PREDICT_TRUE(&from != this)) {
    UnsafeMergeFrom(from);
  } else {
    MergeFromFail(__LINE__);
  }
}

void Request::UnsafeMergeFrom(const Request& from) {
  GOOGLE_DCHECK(&from != this);
  if (from._has_bits_[0 / 32] & (0xffu << (0 % 32))) {
    if (from.has_qid()) {
      set_qid(from.qid());
    }
    if (from.has_market()) {
      set_has_market();
      market_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.market_);
    }
    if (from.has_obj()) {
      set_has_obj();
      obj_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.obj_);
    }
    if (from.has_starttime()) {
      set_starttime(from.starttime());
    }
    if (from.has_endtime()) {
      set_endtime(from.endtime());
    }
    if (from.has_dynatype()) {
      set_dynatype(from.dynatype());
    }
  }
  if (from._internal_metadata_.have_unknown_fields()) {
    ::google::protobuf::UnknownFieldSet::MergeToInternalMetdata(
      from.unknown_fields(), &_internal_metadata_);
  }
}

void Request::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:comm.Request)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void Request::CopyFrom(const Request& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:comm.Request)
  if (&from == this) return;
  Clear();
  UnsafeMergeFrom(from);
}

bool Request::IsInitialized() const {
  if ((_has_bits_[0] & 0x0000003f) != 0x0000003f) return false;

  return true;
}

void Request::Swap(Request* other) {
  if (other == this) return;
  InternalSwap(other);
}
void Request::InternalSwap(Request* other) {
  std::swap(qid_, other->qid_);
  market_.Swap(&other->market_);
  obj_.Swap(&other->obj_);
  std::swap(starttime_, other->starttime_);
  std::swap(endtime_, other->endtime_);
  std::swap(dynatype_, other->dynatype_);
  std::swap(_has_bits_[0], other->_has_bits_[0]);
  _internal_metadata_.Swap(&other->_internal_metadata_);
  std::swap(_cached_size_, other->_cached_size_);
}

::google::protobuf::Metadata Request::GetMetadata() const {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::Metadata metadata;
  metadata.descriptor = Request_descriptor_;
  metadata.reflection = Request_reflection_;
  return metadata;
}

#if PROTOBUF_INLINE_NOT_IN_HEADERS
// Request

// required int64 Qid = 1;
bool Request::has_qid() const {
  return (_has_bits_[0] & 0x00000001u) != 0;
}
void Request::set_has_qid() {
  _has_bits_[0] |= 0x00000001u;
}
void Request::clear_has_qid() {
  _has_bits_[0] &= ~0x00000001u;
}
void Request::clear_qid() {
  qid_ = GOOGLE_LONGLONG(0);
  clear_has_qid();
}
::google::protobuf::int64 Request::qid() const {
  // @@protoc_insertion_point(field_get:comm.Request.Qid)
  return qid_;
}
void Request::set_qid(::google::protobuf::int64 value) {
  set_has_qid();
  qid_ = value;
  // @@protoc_insertion_point(field_set:comm.Request.Qid)
}

// required string Market = 2;
bool Request::has_market() const {
  return (_has_bits_[0] & 0x00000002u) != 0;
}
void Request::set_has_market() {
  _has_bits_[0] |= 0x00000002u;
}
void Request::clear_has_market() {
  _has_bits_[0] &= ~0x00000002u;
}
void Request::clear_market() {
  market_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_market();
}
const ::std::string& Request::market() const {
  // @@protoc_insertion_point(field_get:comm.Request.Market)
  return market_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
void Request::set_market(const ::std::string& value) {
  set_has_market();
  market_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:comm.Request.Market)
}
void Request::set_market(const char* value) {
  set_has_market();
  market_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:comm.Request.Market)
}
void Request::set_market(const char* value, size_t size) {
  set_has_market();
  market_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:comm.Request.Market)
}
::std::string* Request::mutable_market() {
  set_has_market();
  // @@protoc_insertion_point(field_mutable:comm.Request.Market)
  return market_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
::std::string* Request::release_market() {
  // @@protoc_insertion_point(field_release:comm.Request.Market)
  clear_has_market();
  return market_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
void Request::set_allocated_market(::std::string* market) {
  if (market != NULL) {
    set_has_market();
  } else {
    clear_has_market();
  }
  market_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), market);
  // @@protoc_insertion_point(field_set_allocated:comm.Request.Market)
}

// required string Obj = 3;
bool Request::has_obj() const {
  return (_has_bits_[0] & 0x00000004u) != 0;
}
void Request::set_has_obj() {
  _has_bits_[0] |= 0x00000004u;
}
void Request::clear_has_obj() {
  _has_bits_[0] &= ~0x00000004u;
}
void Request::clear_obj() {
  obj_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_obj();
}
const ::std::string& Request::obj() const {
  // @@protoc_insertion_point(field_get:comm.Request.Obj)
  return obj_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
void Request::set_obj(const ::std::string& value) {
  set_has_obj();
  obj_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:comm.Request.Obj)
}
void Request::set_obj(const char* value) {
  set_has_obj();
  obj_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:comm.Request.Obj)
}
void Request::set_obj(const char* value, size_t size) {
  set_has_obj();
  obj_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:comm.Request.Obj)
}
::std::string* Request::mutable_obj() {
  set_has_obj();
  // @@protoc_insertion_point(field_mutable:comm.Request.Obj)
  return obj_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
::std::string* Request::release_obj() {
  // @@protoc_insertion_point(field_release:comm.Request.Obj)
  clear_has_obj();
  return obj_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
void Request::set_allocated_obj(::std::string* obj) {
  if (obj != NULL) {
    set_has_obj();
  } else {
    clear_has_obj();
  }
  obj_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), obj);
  // @@protoc_insertion_point(field_set_allocated:comm.Request.Obj)
}

// required int64 StartTime = 4;
bool Request::has_starttime() const {
  return (_has_bits_[0] & 0x00000008u) != 0;
}
void Request::set_has_starttime() {
  _has_bits_[0] |= 0x00000008u;
}
void Request::clear_has_starttime() {
  _has_bits_[0] &= ~0x00000008u;
}
void Request::clear_starttime() {
  starttime_ = GOOGLE_LONGLONG(0);
  clear_has_starttime();
}
::google::protobuf::int64 Request::starttime() const {
  // @@protoc_insertion_point(field_get:comm.Request.StartTime)
  return starttime_;
}
void Request::set_starttime(::google::protobuf::int64 value) {
  set_has_starttime();
  starttime_ = value;
  // @@protoc_insertion_point(field_set:comm.Request.StartTime)
}

// required int64 EndTime = 5;
bool Request::has_endtime() const {
  return (_has_bits_[0] & 0x00000010u) != 0;
}
void Request::set_has_endtime() {
  _has_bits_[0] |= 0x00000010u;
}
void Request::clear_has_endtime() {
  _has_bits_[0] &= ~0x00000010u;
}
void Request::clear_endtime() {
  endtime_ = GOOGLE_LONGLONG(0);
  clear_has_endtime();
}
::google::protobuf::int64 Request::endtime() const {
  // @@protoc_insertion_point(field_get:comm.Request.EndTime)
  return endtime_;
}
void Request::set_endtime(::google::protobuf::int64 value) {
  set_has_endtime();
  endtime_ = value;
  // @@protoc_insertion_point(field_set:comm.Request.EndTime)
}

// required int64 DynaType = 6;
bool Request::has_dynatype() const {
  return (_has_bits_[0] & 0x00000020u) != 0;
}
void Request::set_has_dynatype() {
  _has_bits_[0] |= 0x00000020u;
}
void Request::clear_has_dynatype() {
  _has_bits_[0] &= ~0x00000020u;
}
void Request::clear_dynatype() {
  dynatype_ = GOOGLE_LONGLONG(0);
  clear_has_dynatype();
}
::google::protobuf::int64 Request::dynatype() const {
  // @@protoc_insertion_point(field_get:comm.Request.DynaType)
  return dynatype_;
}
void Request::set_dynatype(::google::protobuf::int64 value) {
  set_has_dynatype();
  dynatype_ = value;
  // @@protoc_insertion_point(field_set:comm.Request.DynaType)
}

inline const Request* Request::internal_default_instance() {
  return &Request_default_instance_.get();
}
#endif  // PROTOBUF_INLINE_NOT_IN_HEADERS

// @@protoc_insertion_point(namespace_scope)

}  // namespace comm

// @@protoc_insertion_point(global_scope)
