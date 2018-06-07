// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: subStock.proto

#ifndef PROTOBUF_subStock_2eproto__INCLUDED
#define PROTOBUF_subStock_2eproto__INCLUDED

#include <string>

#include <google/protobuf/stubs/common.h>

#if GOOGLE_PROTOBUF_VERSION < 2006000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please update
#error your headers.
#endif
#if 2006001 < GOOGLE_PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers.  Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)

namespace comm {

// Internal implementation detail -- do not call these.
void  protobuf_AddDesc_subStock_2eproto();
void protobuf_AssignDesc_subStock_2eproto();
void protobuf_ShutdownFile_subStock_2eproto();

class SubStock;

// ===================================================================

class SubStock : public ::google::protobuf::Message {
 public:
  SubStock();
  virtual ~SubStock();

  SubStock(const SubStock& from);

  inline SubStock& operator=(const SubStock& from) {
    CopyFrom(from);
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const {
    return _unknown_fields_;
  }

  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields() {
    return &_unknown_fields_;
  }

  static const ::google::protobuf::Descriptor* descriptor();
  static const SubStock& default_instance();

  void Swap(SubStock* other);

  // implements Message ----------------------------------------------

  SubStock* New() const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const SubStock& from);
  void MergeFrom(const SubStock& from);
  void Clear();
  bool IsInitialized() const;

  int ByteSize() const;
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input);
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const;
  ::google::protobuf::uint8* SerializeWithCachedSizesToArray(::google::protobuf::uint8* output) const;
  int GetCachedSize() const { return _cached_size_; }
  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const;
  public:
  ::google::protobuf::Metadata GetMetadata() const;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // required bytes stock_code = 1;
  inline bool has_stock_code() const;
  inline void clear_stock_code();
  static const int kStockCodeFieldNumber = 1;
  inline const ::std::string& stock_code() const;
  inline void set_stock_code(const ::std::string& value);
  inline void set_stock_code(const char* value);
  inline void set_stock_code(const void* value, size_t size);
  inline ::std::string* mutable_stock_code();
  inline ::std::string* release_stock_code();
  inline void set_allocated_stock_code(::std::string* stock_code);

  // required bool sub_flag = 2;
  inline bool has_sub_flag() const;
  inline void clear_sub_flag();
  static const int kSubFlagFieldNumber = 2;
  inline bool sub_flag() const;
  inline void set_sub_flag(bool value);

  // @@protoc_insertion_point(class_scope:comm.SubStock)
 private:
  inline void set_has_stock_code();
  inline void clear_has_stock_code();
  inline void set_has_sub_flag();
  inline void clear_has_sub_flag();

  ::google::protobuf::UnknownFieldSet _unknown_fields_;

  ::google::protobuf::uint32 _has_bits_[1];
  mutable int _cached_size_;
  ::std::string* stock_code_;
  bool sub_flag_;
  friend void  protobuf_AddDesc_subStock_2eproto();
  friend void protobuf_AssignDesc_subStock_2eproto();
  friend void protobuf_ShutdownFile_subStock_2eproto();

  void InitAsDefaultInstance();
  static SubStock* default_instance_;
};
// ===================================================================


// ===================================================================

// SubStock

// required bytes stock_code = 1;
inline bool SubStock::has_stock_code() const {
  return (_has_bits_[0] & 0x00000001u) != 0;
}
inline void SubStock::set_has_stock_code() {
  _has_bits_[0] |= 0x00000001u;
}
inline void SubStock::clear_has_stock_code() {
  _has_bits_[0] &= ~0x00000001u;
}
inline void SubStock::clear_stock_code() {
  if (stock_code_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    stock_code_->clear();
  }
  clear_has_stock_code();
}
inline const ::std::string& SubStock::stock_code() const {
  // @@protoc_insertion_point(field_get:comm.SubStock.stock_code)
  return *stock_code_;
}
inline void SubStock::set_stock_code(const ::std::string& value) {
  set_has_stock_code();
  if (stock_code_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    stock_code_ = new ::std::string;
  }
  stock_code_->assign(value);
  // @@protoc_insertion_point(field_set:comm.SubStock.stock_code)
}
inline void SubStock::set_stock_code(const char* value) {
  set_has_stock_code();
  if (stock_code_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    stock_code_ = new ::std::string;
  }
  stock_code_->assign(value);
  // @@protoc_insertion_point(field_set_char:comm.SubStock.stock_code)
}
inline void SubStock::set_stock_code(const void* value, size_t size) {
  set_has_stock_code();
  if (stock_code_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    stock_code_ = new ::std::string;
  }
  stock_code_->assign(reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:comm.SubStock.stock_code)
}
inline ::std::string* SubStock::mutable_stock_code() {
  set_has_stock_code();
  if (stock_code_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    stock_code_ = new ::std::string;
  }
  // @@protoc_insertion_point(field_mutable:comm.SubStock.stock_code)
  return stock_code_;
}
inline ::std::string* SubStock::release_stock_code() {
  clear_has_stock_code();
  if (stock_code_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    return NULL;
  } else {
    ::std::string* temp = stock_code_;
    stock_code_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
    return temp;
  }
}
inline void SubStock::set_allocated_stock_code(::std::string* stock_code) {
  if (stock_code_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete stock_code_;
  }
  if (stock_code) {
    set_has_stock_code();
    stock_code_ = stock_code;
  } else {
    clear_has_stock_code();
    stock_code_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  // @@protoc_insertion_point(field_set_allocated:comm.SubStock.stock_code)
}

// required bool sub_flag = 2;
inline bool SubStock::has_sub_flag() const {
  return (_has_bits_[0] & 0x00000002u) != 0;
}
inline void SubStock::set_has_sub_flag() {
  _has_bits_[0] |= 0x00000002u;
}
inline void SubStock::clear_has_sub_flag() {
  _has_bits_[0] &= ~0x00000002u;
}
inline void SubStock::clear_sub_flag() {
  sub_flag_ = false;
  clear_has_sub_flag();
}
inline bool SubStock::sub_flag() const {
  // @@protoc_insertion_point(field_get:comm.SubStock.sub_flag)
  return sub_flag_;
}
inline void SubStock::set_sub_flag(bool value) {
  set_has_sub_flag();
  sub_flag_ = value;
  // @@protoc_insertion_point(field_set:comm.SubStock.sub_flag)
}


// @@protoc_insertion_point(namespace_scope)

}  // namespace comm

#ifndef SWIG
namespace google {
namespace protobuf {


}  // namespace google
}  // namespace protobuf
#endif  // SWIG

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_subStock_2eproto__INCLUDED