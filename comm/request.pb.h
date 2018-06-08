// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: request.proto

#ifndef PROTOBUF_request_2eproto__INCLUDED
#define PROTOBUF_request_2eproto__INCLUDED

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 3001000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 3001000 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)

namespace comm {

// Internal implementation detail -- do not call these.
void protobuf_AddDesc_request_2eproto();
void protobuf_InitDefaults_request_2eproto();
void protobuf_AssignDesc_request_2eproto();
void protobuf_ShutdownFile_request_2eproto();

class Request;

// ===================================================================

class Request : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:comm.Request) */ {
 public:
  Request();
  virtual ~Request();

  Request(const Request& from);

  inline Request& operator=(const Request& from) {
    CopyFrom(from);
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const {
    return _internal_metadata_.unknown_fields();
  }

  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields() {
    return _internal_metadata_.mutable_unknown_fields();
  }

  static const ::google::protobuf::Descriptor* descriptor();
  static const Request& default_instance();

  static const Request* internal_default_instance();

  void Swap(Request* other);

  // implements Message ----------------------------------------------

  inline Request* New() const { return New(NULL); }

  Request* New(::google::protobuf::Arena* arena) const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const Request& from);
  void MergeFrom(const Request& from);
  void Clear();
  bool IsInitialized() const;

  size_t ByteSizeLong() const;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input);
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      bool deterministic, ::google::protobuf::uint8* output) const;
  ::google::protobuf::uint8* SerializeWithCachedSizesToArray(::google::protobuf::uint8* output) const {
    return InternalSerializeWithCachedSizesToArray(false, output);
  }
  int GetCachedSize() const { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  void InternalSwap(Request* other);
  void UnsafeMergeFrom(const Request& from);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return _internal_metadata_.arena();
  }
  inline void* MaybeArenaPtr() const {
    return _internal_metadata_.raw_arena_ptr();
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // required int64 Qid = 1;
  bool has_qid() const;
  void clear_qid();
  static const int kQidFieldNumber = 1;
  ::google::protobuf::int64 qid() const;
  void set_qid(::google::protobuf::int64 value);

  // required string Market = 2;
  bool has_market() const;
  void clear_market();
  static const int kMarketFieldNumber = 2;
  const ::std::string& market() const;
  void set_market(const ::std::string& value);
  void set_market(const char* value);
  void set_market(const char* value, size_t size);
  ::std::string* mutable_market();
  ::std::string* release_market();
  void set_allocated_market(::std::string* market);

  // required string Obj = 3;
  bool has_obj() const;
  void clear_obj();
  static const int kObjFieldNumber = 3;
  const ::std::string& obj() const;
  void set_obj(const ::std::string& value);
  void set_obj(const char* value);
  void set_obj(const char* value, size_t size);
  ::std::string* mutable_obj();
  ::std::string* release_obj();
  void set_allocated_obj(::std::string* obj);

  // required int64 StartTime = 4;
  bool has_starttime() const;
  void clear_starttime();
  static const int kStartTimeFieldNumber = 4;
  ::google::protobuf::int64 starttime() const;
  void set_starttime(::google::protobuf::int64 value);

  // required int64 EndTime = 5;
  bool has_endtime() const;
  void clear_endtime();
  static const int kEndTimeFieldNumber = 5;
  ::google::protobuf::int64 endtime() const;
  void set_endtime(::google::protobuf::int64 value);

  // required int64 DynaType = 6;
  bool has_dynatype() const;
  void clear_dynatype();
  static const int kDynaTypeFieldNumber = 6;
  ::google::protobuf::int64 dynatype() const;
  void set_dynatype(::google::protobuf::int64 value);

  // @@protoc_insertion_point(class_scope:comm.Request)
 private:
  inline void set_has_qid();
  inline void clear_has_qid();
  inline void set_has_market();
  inline void clear_has_market();
  inline void set_has_obj();
  inline void clear_has_obj();
  inline void set_has_starttime();
  inline void clear_has_starttime();
  inline void set_has_endtime();
  inline void clear_has_endtime();
  inline void set_has_dynatype();
  inline void clear_has_dynatype();

  // helper for ByteSizeLong()
  size_t RequiredFieldsByteSizeFallback() const;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::HasBits<1> _has_bits_;
  mutable int _cached_size_;
  ::google::protobuf::internal::ArenaStringPtr market_;
  ::google::protobuf::internal::ArenaStringPtr obj_;
  ::google::protobuf::int64 qid_;
  ::google::protobuf::int64 starttime_;
  ::google::protobuf::int64 endtime_;
  ::google::protobuf::int64 dynatype_;
  friend void  protobuf_InitDefaults_request_2eproto_impl();
  friend void  protobuf_AddDesc_request_2eproto_impl();
  friend void protobuf_AssignDesc_request_2eproto();
  friend void protobuf_ShutdownFile_request_2eproto();

  void InitAsDefaultInstance();
};
extern ::google::protobuf::internal::ExplicitlyConstructed<Request> Request_default_instance_;

// ===================================================================


// ===================================================================

#if !PROTOBUF_INLINE_NOT_IN_HEADERS
// Request

// required int64 Qid = 1;
inline bool Request::has_qid() const {
  return (_has_bits_[0] & 0x00000001u) != 0;
}
inline void Request::set_has_qid() {
  _has_bits_[0] |= 0x00000001u;
}
inline void Request::clear_has_qid() {
  _has_bits_[0] &= ~0x00000001u;
}
inline void Request::clear_qid() {
  qid_ = GOOGLE_LONGLONG(0);
  clear_has_qid();
}
inline ::google::protobuf::int64 Request::qid() const {
  // @@protoc_insertion_point(field_get:comm.Request.Qid)
  return qid_;
}
inline void Request::set_qid(::google::protobuf::int64 value) {
  set_has_qid();
  qid_ = value;
  // @@protoc_insertion_point(field_set:comm.Request.Qid)
}

// required string Market = 2;
inline bool Request::has_market() const {
  return (_has_bits_[0] & 0x00000002u) != 0;
}
inline void Request::set_has_market() {
  _has_bits_[0] |= 0x00000002u;
}
inline void Request::clear_has_market() {
  _has_bits_[0] &= ~0x00000002u;
}
inline void Request::clear_market() {
  market_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_market();
}
inline const ::std::string& Request::market() const {
  // @@protoc_insertion_point(field_get:comm.Request.Market)
  return market_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Request::set_market(const ::std::string& value) {
  set_has_market();
  market_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:comm.Request.Market)
}
inline void Request::set_market(const char* value) {
  set_has_market();
  market_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:comm.Request.Market)
}
inline void Request::set_market(const char* value, size_t size) {
  set_has_market();
  market_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:comm.Request.Market)
}
inline ::std::string* Request::mutable_market() {
  set_has_market();
  // @@protoc_insertion_point(field_mutable:comm.Request.Market)
  return market_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* Request::release_market() {
  // @@protoc_insertion_point(field_release:comm.Request.Market)
  clear_has_market();
  return market_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Request::set_allocated_market(::std::string* market) {
  if (market != NULL) {
    set_has_market();
  } else {
    clear_has_market();
  }
  market_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), market);
  // @@protoc_insertion_point(field_set_allocated:comm.Request.Market)
}

// required string Obj = 3;
inline bool Request::has_obj() const {
  return (_has_bits_[0] & 0x00000004u) != 0;
}
inline void Request::set_has_obj() {
  _has_bits_[0] |= 0x00000004u;
}
inline void Request::clear_has_obj() {
  _has_bits_[0] &= ~0x00000004u;
}
inline void Request::clear_obj() {
  obj_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_obj();
}
inline const ::std::string& Request::obj() const {
  // @@protoc_insertion_point(field_get:comm.Request.Obj)
  return obj_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Request::set_obj(const ::std::string& value) {
  set_has_obj();
  obj_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:comm.Request.Obj)
}
inline void Request::set_obj(const char* value) {
  set_has_obj();
  obj_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:comm.Request.Obj)
}
inline void Request::set_obj(const char* value, size_t size) {
  set_has_obj();
  obj_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:comm.Request.Obj)
}
inline ::std::string* Request::mutable_obj() {
  set_has_obj();
  // @@protoc_insertion_point(field_mutable:comm.Request.Obj)
  return obj_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* Request::release_obj() {
  // @@protoc_insertion_point(field_release:comm.Request.Obj)
  clear_has_obj();
  return obj_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Request::set_allocated_obj(::std::string* obj) {
  if (obj != NULL) {
    set_has_obj();
  } else {
    clear_has_obj();
  }
  obj_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), obj);
  // @@protoc_insertion_point(field_set_allocated:comm.Request.Obj)
}

// required int64 StartTime = 4;
inline bool Request::has_starttime() const {
  return (_has_bits_[0] & 0x00000008u) != 0;
}
inline void Request::set_has_starttime() {
  _has_bits_[0] |= 0x00000008u;
}
inline void Request::clear_has_starttime() {
  _has_bits_[0] &= ~0x00000008u;
}
inline void Request::clear_starttime() {
  starttime_ = GOOGLE_LONGLONG(0);
  clear_has_starttime();
}
inline ::google::protobuf::int64 Request::starttime() const {
  // @@protoc_insertion_point(field_get:comm.Request.StartTime)
  return starttime_;
}
inline void Request::set_starttime(::google::protobuf::int64 value) {
  set_has_starttime();
  starttime_ = value;
  // @@protoc_insertion_point(field_set:comm.Request.StartTime)
}

// required int64 EndTime = 5;
inline bool Request::has_endtime() const {
  return (_has_bits_[0] & 0x00000010u) != 0;
}
inline void Request::set_has_endtime() {
  _has_bits_[0] |= 0x00000010u;
}
inline void Request::clear_has_endtime() {
  _has_bits_[0] &= ~0x00000010u;
}
inline void Request::clear_endtime() {
  endtime_ = GOOGLE_LONGLONG(0);
  clear_has_endtime();
}
inline ::google::protobuf::int64 Request::endtime() const {
  // @@protoc_insertion_point(field_get:comm.Request.EndTime)
  return endtime_;
}
inline void Request::set_endtime(::google::protobuf::int64 value) {
  set_has_endtime();
  endtime_ = value;
  // @@protoc_insertion_point(field_set:comm.Request.EndTime)
}

// required int64 DynaType = 6;
inline bool Request::has_dynatype() const {
  return (_has_bits_[0] & 0x00000020u) != 0;
}
inline void Request::set_has_dynatype() {
  _has_bits_[0] |= 0x00000020u;
}
inline void Request::clear_has_dynatype() {
  _has_bits_[0] &= ~0x00000020u;
}
inline void Request::clear_dynatype() {
  dynatype_ = GOOGLE_LONGLONG(0);
  clear_has_dynatype();
}
inline ::google::protobuf::int64 Request::dynatype() const {
  // @@protoc_insertion_point(field_get:comm.Request.DynaType)
  return dynatype_;
}
inline void Request::set_dynatype(::google::protobuf::int64 value) {
  set_has_dynatype();
  dynatype_ = value;
  // @@protoc_insertion_point(field_set:comm.Request.DynaType)
}

inline const Request* Request::internal_default_instance() {
  return &Request_default_instance_.get();
}
#endif  // !PROTOBUF_INLINE_NOT_IN_HEADERS

// @@protoc_insertion_point(namespace_scope)

}  // namespace comm

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_request_2eproto__INCLUDED
