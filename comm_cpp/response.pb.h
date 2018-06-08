// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: response.proto

#ifndef PROTOBUF_response_2eproto__INCLUDED
#define PROTOBUF_response_2eproto__INCLUDED

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
void protobuf_AddDesc_response_2eproto();
void protobuf_InitDefaults_response_2eproto();
void protobuf_AssignDesc_response_2eproto();
void protobuf_ShutdownFile_response_2eproto();

class Response;

// ===================================================================

class Response : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:comm.Response) */ {
 public:
  Response();
  virtual ~Response();

  Response(const Response& from);

  inline Response& operator=(const Response& from) {
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
  static const Response& default_instance();

  static const Response* internal_default_instance();

  void Swap(Response* other);

  // implements Message ----------------------------------------------

  inline Response* New() const { return New(NULL); }

  Response* New(::google::protobuf::Arena* arena) const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const Response& from);
  void MergeFrom(const Response& from);
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
  void InternalSwap(Response* other);
  void UnsafeMergeFrom(const Response& from);
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

  // required int64 errorno = 2;
  bool has_errorno() const;
  void clear_errorno();
  static const int kErrornoFieldNumber = 2;
  ::google::protobuf::int64 errorno() const;
  void set_errorno(::google::protobuf::int64 value);

  // repeated bytes data = 3;
  int data_size() const;
  void clear_data();
  static const int kDataFieldNumber = 3;
  const ::std::string& data(int index) const;
  ::std::string* mutable_data(int index);
  void set_data(int index, const ::std::string& value);
  void set_data(int index, const char* value);
  void set_data(int index, const void* value, size_t size);
  ::std::string* add_data();
  void add_data(const ::std::string& value);
  void add_data(const char* value);
  void add_data(const void* value, size_t size);
  const ::google::protobuf::RepeatedPtrField< ::std::string>& data() const;
  ::google::protobuf::RepeatedPtrField< ::std::string>* mutable_data();

  // @@protoc_insertion_point(class_scope:comm.Response)
 private:
  inline void set_has_qid();
  inline void clear_has_qid();
  inline void set_has_errorno();
  inline void clear_has_errorno();

  // helper for ByteSizeLong()
  size_t RequiredFieldsByteSizeFallback() const;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::HasBits<1> _has_bits_;
  mutable int _cached_size_;
  ::google::protobuf::RepeatedPtrField< ::std::string> data_;
  ::google::protobuf::int64 qid_;
  ::google::protobuf::int64 errorno_;
  friend void  protobuf_InitDefaults_response_2eproto_impl();
  friend void  protobuf_AddDesc_response_2eproto_impl();
  friend void protobuf_AssignDesc_response_2eproto();
  friend void protobuf_ShutdownFile_response_2eproto();

  void InitAsDefaultInstance();
};
extern ::google::protobuf::internal::ExplicitlyConstructed<Response> Response_default_instance_;

// ===================================================================


// ===================================================================

#if !PROTOBUF_INLINE_NOT_IN_HEADERS
// Response

// required int64 Qid = 1;
inline bool Response::has_qid() const {
  return (_has_bits_[0] & 0x00000001u) != 0;
}
inline void Response::set_has_qid() {
  _has_bits_[0] |= 0x00000001u;
}
inline void Response::clear_has_qid() {
  _has_bits_[0] &= ~0x00000001u;
}
inline void Response::clear_qid() {
  qid_ = GOOGLE_LONGLONG(0);
  clear_has_qid();
}
inline ::google::protobuf::int64 Response::qid() const {
  // @@protoc_insertion_point(field_get:comm.Response.Qid)
  return qid_;
}
inline void Response::set_qid(::google::protobuf::int64 value) {
  set_has_qid();
  qid_ = value;
  // @@protoc_insertion_point(field_set:comm.Response.Qid)
}

// required int64 errorno = 2;
inline bool Response::has_errorno() const {
  return (_has_bits_[0] & 0x00000002u) != 0;
}
inline void Response::set_has_errorno() {
  _has_bits_[0] |= 0x00000002u;
}
inline void Response::clear_has_errorno() {
  _has_bits_[0] &= ~0x00000002u;
}
inline void Response::clear_errorno() {
  errorno_ = GOOGLE_LONGLONG(0);
  clear_has_errorno();
}
inline ::google::protobuf::int64 Response::errorno() const {
  // @@protoc_insertion_point(field_get:comm.Response.errorno)
  return errorno_;
}
inline void Response::set_errorno(::google::protobuf::int64 value) {
  set_has_errorno();
  errorno_ = value;
  // @@protoc_insertion_point(field_set:comm.Response.errorno)
}

// repeated bytes data = 3;
inline int Response::data_size() const {
  return data_.size();
}
inline void Response::clear_data() {
  data_.Clear();
}
inline const ::std::string& Response::data(int index) const {
  // @@protoc_insertion_point(field_get:comm.Response.data)
  return data_.Get(index);
}
inline ::std::string* Response::mutable_data(int index) {
  // @@protoc_insertion_point(field_mutable:comm.Response.data)
  return data_.Mutable(index);
}
inline void Response::set_data(int index, const ::std::string& value) {
  // @@protoc_insertion_point(field_set:comm.Response.data)
  data_.Mutable(index)->assign(value);
}
inline void Response::set_data(int index, const char* value) {
  data_.Mutable(index)->assign(value);
  // @@protoc_insertion_point(field_set_char:comm.Response.data)
}
inline void Response::set_data(int index, const void* value, size_t size) {
  data_.Mutable(index)->assign(
    reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:comm.Response.data)
}
inline ::std::string* Response::add_data() {
  // @@protoc_insertion_point(field_add_mutable:comm.Response.data)
  return data_.Add();
}
inline void Response::add_data(const ::std::string& value) {
  data_.Add()->assign(value);
  // @@protoc_insertion_point(field_add:comm.Response.data)
}
inline void Response::add_data(const char* value) {
  data_.Add()->assign(value);
  // @@protoc_insertion_point(field_add_char:comm.Response.data)
}
inline void Response::add_data(const void* value, size_t size) {
  data_.Add()->assign(reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_add_pointer:comm.Response.data)
}
inline const ::google::protobuf::RepeatedPtrField< ::std::string>&
Response::data() const {
  // @@protoc_insertion_point(field_list:comm.Response.data)
  return data_;
}
inline ::google::protobuf::RepeatedPtrField< ::std::string>*
Response::mutable_data() {
  // @@protoc_insertion_point(field_mutable_list:comm.Response.data)
  return &data_;
}

inline const Response* Response::internal_default_instance() {
  return &Response_default_instance_.get();
}
#endif  // !PROTOBUF_INLINE_NOT_IN_HEADERS

// @@protoc_insertion_point(namespace_scope)

}  // namespace comm

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_response_2eproto__INCLUDED
