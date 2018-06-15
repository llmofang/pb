// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: code_table.proto

#ifndef PROTOBUF_code_5ftable_2eproto__INCLUDED
#define PROTOBUF_code_5ftable_2eproto__INCLUDED

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
void  protobuf_AddDesc_code_5ftable_2eproto();
void protobuf_AssignDesc_code_5ftable_2eproto();
void protobuf_ShutdownFile_code_5ftable_2eproto();

class TdfCode;
class TdfCodeTable;

// ===================================================================

class TdfCode : public ::google::protobuf::Message {
 public:
  TdfCode();
  virtual ~TdfCode();

  TdfCode(const TdfCode& from);

  inline TdfCode& operator=(const TdfCode& from) {
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
  static const TdfCode& default_instance();

  void Swap(TdfCode* other);

  // implements Message ----------------------------------------------

  TdfCode* New() const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const TdfCode& from);
  void MergeFrom(const TdfCode& from);
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

  // required bytes WindCode = 1;
  inline bool has_windcode() const;
  inline void clear_windcode();
  static const int kWindCodeFieldNumber = 1;
  inline const ::std::string& windcode() const;
  inline void set_windcode(const ::std::string& value);
  inline void set_windcode(const char* value);
  inline void set_windcode(const void* value, size_t size);
  inline ::std::string* mutable_windcode();
  inline ::std::string* release_windcode();
  inline void set_allocated_windcode(::std::string* windcode);

  // required bytes Market = 2;
  inline bool has_market() const;
  inline void clear_market();
  static const int kMarketFieldNumber = 2;
  inline const ::std::string& market() const;
  inline void set_market(const ::std::string& value);
  inline void set_market(const char* value);
  inline void set_market(const void* value, size_t size);
  inline ::std::string* mutable_market();
  inline ::std::string* release_market();
  inline void set_allocated_market(::std::string* market);

  // required bytes Code = 3;
  inline bool has_code() const;
  inline void clear_code();
  static const int kCodeFieldNumber = 3;
  inline const ::std::string& code() const;
  inline void set_code(const ::std::string& value);
  inline void set_code(const char* value);
  inline void set_code(const void* value, size_t size);
  inline ::std::string* mutable_code();
  inline ::std::string* release_code();
  inline void set_allocated_code(::std::string* code);

  // required bytes ENName = 4;
  inline bool has_enname() const;
  inline void clear_enname();
  static const int kENNameFieldNumber = 4;
  inline const ::std::string& enname() const;
  inline void set_enname(const ::std::string& value);
  inline void set_enname(const char* value);
  inline void set_enname(const void* value, size_t size);
  inline ::std::string* mutable_enname();
  inline ::std::string* release_enname();
  inline void set_allocated_enname(::std::string* enname);

  // required bytes CNName = 5;
  inline bool has_cnname() const;
  inline void clear_cnname();
  static const int kCNNameFieldNumber = 5;
  inline const ::std::string& cnname() const;
  inline void set_cnname(const ::std::string& value);
  inline void set_cnname(const char* value);
  inline void set_cnname(const void* value, size_t size);
  inline ::std::string* mutable_cnname();
  inline ::std::string* release_cnname();
  inline void set_allocated_cnname(::std::string* cnname);

  // required int32 Type = 6;
  inline bool has_type() const;
  inline void clear_type();
  static const int kTypeFieldNumber = 6;
  inline ::google::protobuf::int32 type() const;
  inline void set_type(::google::protobuf::int32 value);

  // @@protoc_insertion_point(class_scope:comm.TdfCode)
 private:
  inline void set_has_windcode();
  inline void clear_has_windcode();
  inline void set_has_market();
  inline void clear_has_market();
  inline void set_has_code();
  inline void clear_has_code();
  inline void set_has_enname();
  inline void clear_has_enname();
  inline void set_has_cnname();
  inline void clear_has_cnname();
  inline void set_has_type();
  inline void clear_has_type();

  ::google::protobuf::UnknownFieldSet _unknown_fields_;

  ::google::protobuf::uint32 _has_bits_[1];
  mutable int _cached_size_;
  ::std::string* windcode_;
  ::std::string* market_;
  ::std::string* code_;
  ::std::string* enname_;
  ::std::string* cnname_;
  ::google::protobuf::int32 type_;
  friend void  protobuf_AddDesc_code_5ftable_2eproto();
  friend void protobuf_AssignDesc_code_5ftable_2eproto();
  friend void protobuf_ShutdownFile_code_5ftable_2eproto();

  void InitAsDefaultInstance();
  static TdfCode* default_instance_;
};
// -------------------------------------------------------------------

class TdfCodeTable : public ::google::protobuf::Message {
 public:
  TdfCodeTable();
  virtual ~TdfCodeTable();

  TdfCodeTable(const TdfCodeTable& from);

  inline TdfCodeTable& operator=(const TdfCodeTable& from) {
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
  static const TdfCodeTable& default_instance();

  void Swap(TdfCodeTable* other);

  // implements Message ----------------------------------------------

  TdfCodeTable* New() const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const TdfCodeTable& from);
  void MergeFrom(const TdfCodeTable& from);
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

  // repeated .comm.TdfCode CodeTable = 1;
  inline int codetable_size() const;
  inline void clear_codetable();
  static const int kCodeTableFieldNumber = 1;
  inline const ::comm::TdfCode& codetable(int index) const;
  inline ::comm::TdfCode* mutable_codetable(int index);
  inline ::comm::TdfCode* add_codetable();
  inline const ::google::protobuf::RepeatedPtrField< ::comm::TdfCode >&
      codetable() const;
  inline ::google::protobuf::RepeatedPtrField< ::comm::TdfCode >*
      mutable_codetable();

  // @@protoc_insertion_point(class_scope:comm.TdfCodeTable)
 private:

  ::google::protobuf::UnknownFieldSet _unknown_fields_;

  ::google::protobuf::uint32 _has_bits_[1];
  mutable int _cached_size_;
  ::google::protobuf::RepeatedPtrField< ::comm::TdfCode > codetable_;
  friend void  protobuf_AddDesc_code_5ftable_2eproto();
  friend void protobuf_AssignDesc_code_5ftable_2eproto();
  friend void protobuf_ShutdownFile_code_5ftable_2eproto();

  void InitAsDefaultInstance();
  static TdfCodeTable* default_instance_;
};
// ===================================================================


// ===================================================================

// TdfCode

// required bytes WindCode = 1;
inline bool TdfCode::has_windcode() const {
  return (_has_bits_[0] & 0x00000001u) != 0;
}
inline void TdfCode::set_has_windcode() {
  _has_bits_[0] |= 0x00000001u;
}
inline void TdfCode::clear_has_windcode() {
  _has_bits_[0] &= ~0x00000001u;
}
inline void TdfCode::clear_windcode() {
  if (windcode_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    windcode_->clear();
  }
  clear_has_windcode();
}
inline const ::std::string& TdfCode::windcode() const {
  // @@protoc_insertion_point(field_get:comm.TdfCode.WindCode)
  return *windcode_;
}
inline void TdfCode::set_windcode(const ::std::string& value) {
  set_has_windcode();
  if (windcode_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    windcode_ = new ::std::string;
  }
  windcode_->assign(value);
  // @@protoc_insertion_point(field_set:comm.TdfCode.WindCode)
}
inline void TdfCode::set_windcode(const char* value) {
  set_has_windcode();
  if (windcode_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    windcode_ = new ::std::string;
  }
  windcode_->assign(value);
  // @@protoc_insertion_point(field_set_char:comm.TdfCode.WindCode)
}
inline void TdfCode::set_windcode(const void* value, size_t size) {
  set_has_windcode();
  if (windcode_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    windcode_ = new ::std::string;
  }
  windcode_->assign(reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:comm.TdfCode.WindCode)
}
inline ::std::string* TdfCode::mutable_windcode() {
  set_has_windcode();
  if (windcode_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    windcode_ = new ::std::string;
  }
  // @@protoc_insertion_point(field_mutable:comm.TdfCode.WindCode)
  return windcode_;
}
inline ::std::string* TdfCode::release_windcode() {
  clear_has_windcode();
  if (windcode_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    return NULL;
  } else {
    ::std::string* temp = windcode_;
    windcode_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
    return temp;
  }
}
inline void TdfCode::set_allocated_windcode(::std::string* windcode) {
  if (windcode_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete windcode_;
  }
  if (windcode) {
    set_has_windcode();
    windcode_ = windcode;
  } else {
    clear_has_windcode();
    windcode_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  // @@protoc_insertion_point(field_set_allocated:comm.TdfCode.WindCode)
}

// required bytes Market = 2;
inline bool TdfCode::has_market() const {
  return (_has_bits_[0] & 0x00000002u) != 0;
}
inline void TdfCode::set_has_market() {
  _has_bits_[0] |= 0x00000002u;
}
inline void TdfCode::clear_has_market() {
  _has_bits_[0] &= ~0x00000002u;
}
inline void TdfCode::clear_market() {
  if (market_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    market_->clear();
  }
  clear_has_market();
}
inline const ::std::string& TdfCode::market() const {
  // @@protoc_insertion_point(field_get:comm.TdfCode.Market)
  return *market_;
}
inline void TdfCode::set_market(const ::std::string& value) {
  set_has_market();
  if (market_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    market_ = new ::std::string;
  }
  market_->assign(value);
  // @@protoc_insertion_point(field_set:comm.TdfCode.Market)
}
inline void TdfCode::set_market(const char* value) {
  set_has_market();
  if (market_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    market_ = new ::std::string;
  }
  market_->assign(value);
  // @@protoc_insertion_point(field_set_char:comm.TdfCode.Market)
}
inline void TdfCode::set_market(const void* value, size_t size) {
  set_has_market();
  if (market_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    market_ = new ::std::string;
  }
  market_->assign(reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:comm.TdfCode.Market)
}
inline ::std::string* TdfCode::mutable_market() {
  set_has_market();
  if (market_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    market_ = new ::std::string;
  }
  // @@protoc_insertion_point(field_mutable:comm.TdfCode.Market)
  return market_;
}
inline ::std::string* TdfCode::release_market() {
  clear_has_market();
  if (market_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    return NULL;
  } else {
    ::std::string* temp = market_;
    market_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
    return temp;
  }
}
inline void TdfCode::set_allocated_market(::std::string* market) {
  if (market_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete market_;
  }
  if (market) {
    set_has_market();
    market_ = market;
  } else {
    clear_has_market();
    market_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  // @@protoc_insertion_point(field_set_allocated:comm.TdfCode.Market)
}

// required bytes Code = 3;
inline bool TdfCode::has_code() const {
  return (_has_bits_[0] & 0x00000004u) != 0;
}
inline void TdfCode::set_has_code() {
  _has_bits_[0] |= 0x00000004u;
}
inline void TdfCode::clear_has_code() {
  _has_bits_[0] &= ~0x00000004u;
}
inline void TdfCode::clear_code() {
  if (code_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    code_->clear();
  }
  clear_has_code();
}
inline const ::std::string& TdfCode::code() const {
  // @@protoc_insertion_point(field_get:comm.TdfCode.Code)
  return *code_;
}
inline void TdfCode::set_code(const ::std::string& value) {
  set_has_code();
  if (code_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    code_ = new ::std::string;
  }
  code_->assign(value);
  // @@protoc_insertion_point(field_set:comm.TdfCode.Code)
}
inline void TdfCode::set_code(const char* value) {
  set_has_code();
  if (code_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    code_ = new ::std::string;
  }
  code_->assign(value);
  // @@protoc_insertion_point(field_set_char:comm.TdfCode.Code)
}
inline void TdfCode::set_code(const void* value, size_t size) {
  set_has_code();
  if (code_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    code_ = new ::std::string;
  }
  code_->assign(reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:comm.TdfCode.Code)
}
inline ::std::string* TdfCode::mutable_code() {
  set_has_code();
  if (code_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    code_ = new ::std::string;
  }
  // @@protoc_insertion_point(field_mutable:comm.TdfCode.Code)
  return code_;
}
inline ::std::string* TdfCode::release_code() {
  clear_has_code();
  if (code_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    return NULL;
  } else {
    ::std::string* temp = code_;
    code_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
    return temp;
  }
}
inline void TdfCode::set_allocated_code(::std::string* code) {
  if (code_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete code_;
  }
  if (code) {
    set_has_code();
    code_ = code;
  } else {
    clear_has_code();
    code_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  // @@protoc_insertion_point(field_set_allocated:comm.TdfCode.Code)
}

// required bytes ENName = 4;
inline bool TdfCode::has_enname() const {
  return (_has_bits_[0] & 0x00000008u) != 0;
}
inline void TdfCode::set_has_enname() {
  _has_bits_[0] |= 0x00000008u;
}
inline void TdfCode::clear_has_enname() {
  _has_bits_[0] &= ~0x00000008u;
}
inline void TdfCode::clear_enname() {
  if (enname_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    enname_->clear();
  }
  clear_has_enname();
}
inline const ::std::string& TdfCode::enname() const {
  // @@protoc_insertion_point(field_get:comm.TdfCode.ENName)
  return *enname_;
}
inline void TdfCode::set_enname(const ::std::string& value) {
  set_has_enname();
  if (enname_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    enname_ = new ::std::string;
  }
  enname_->assign(value);
  // @@protoc_insertion_point(field_set:comm.TdfCode.ENName)
}
inline void TdfCode::set_enname(const char* value) {
  set_has_enname();
  if (enname_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    enname_ = new ::std::string;
  }
  enname_->assign(value);
  // @@protoc_insertion_point(field_set_char:comm.TdfCode.ENName)
}
inline void TdfCode::set_enname(const void* value, size_t size) {
  set_has_enname();
  if (enname_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    enname_ = new ::std::string;
  }
  enname_->assign(reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:comm.TdfCode.ENName)
}
inline ::std::string* TdfCode::mutable_enname() {
  set_has_enname();
  if (enname_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    enname_ = new ::std::string;
  }
  // @@protoc_insertion_point(field_mutable:comm.TdfCode.ENName)
  return enname_;
}
inline ::std::string* TdfCode::release_enname() {
  clear_has_enname();
  if (enname_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    return NULL;
  } else {
    ::std::string* temp = enname_;
    enname_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
    return temp;
  }
}
inline void TdfCode::set_allocated_enname(::std::string* enname) {
  if (enname_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete enname_;
  }
  if (enname) {
    set_has_enname();
    enname_ = enname;
  } else {
    clear_has_enname();
    enname_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  // @@protoc_insertion_point(field_set_allocated:comm.TdfCode.ENName)
}

// required bytes CNName = 5;
inline bool TdfCode::has_cnname() const {
  return (_has_bits_[0] & 0x00000010u) != 0;
}
inline void TdfCode::set_has_cnname() {
  _has_bits_[0] |= 0x00000010u;
}
inline void TdfCode::clear_has_cnname() {
  _has_bits_[0] &= ~0x00000010u;
}
inline void TdfCode::clear_cnname() {
  if (cnname_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    cnname_->clear();
  }
  clear_has_cnname();
}
inline const ::std::string& TdfCode::cnname() const {
  // @@protoc_insertion_point(field_get:comm.TdfCode.CNName)
  return *cnname_;
}
inline void TdfCode::set_cnname(const ::std::string& value) {
  set_has_cnname();
  if (cnname_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    cnname_ = new ::std::string;
  }
  cnname_->assign(value);
  // @@protoc_insertion_point(field_set:comm.TdfCode.CNName)
}
inline void TdfCode::set_cnname(const char* value) {
  set_has_cnname();
  if (cnname_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    cnname_ = new ::std::string;
  }
  cnname_->assign(value);
  // @@protoc_insertion_point(field_set_char:comm.TdfCode.CNName)
}
inline void TdfCode::set_cnname(const void* value, size_t size) {
  set_has_cnname();
  if (cnname_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    cnname_ = new ::std::string;
  }
  cnname_->assign(reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:comm.TdfCode.CNName)
}
inline ::std::string* TdfCode::mutable_cnname() {
  set_has_cnname();
  if (cnname_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    cnname_ = new ::std::string;
  }
  // @@protoc_insertion_point(field_mutable:comm.TdfCode.CNName)
  return cnname_;
}
inline ::std::string* TdfCode::release_cnname() {
  clear_has_cnname();
  if (cnname_ == &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    return NULL;
  } else {
    ::std::string* temp = cnname_;
    cnname_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
    return temp;
  }
}
inline void TdfCode::set_allocated_cnname(::std::string* cnname) {
  if (cnname_ != &::google::protobuf::internal::GetEmptyStringAlreadyInited()) {
    delete cnname_;
  }
  if (cnname) {
    set_has_cnname();
    cnname_ = cnname;
  } else {
    clear_has_cnname();
    cnname_ = const_cast< ::std::string*>(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  }
  // @@protoc_insertion_point(field_set_allocated:comm.TdfCode.CNName)
}

// required int32 Type = 6;
inline bool TdfCode::has_type() const {
  return (_has_bits_[0] & 0x00000020u) != 0;
}
inline void TdfCode::set_has_type() {
  _has_bits_[0] |= 0x00000020u;
}
inline void TdfCode::clear_has_type() {
  _has_bits_[0] &= ~0x00000020u;
}
inline void TdfCode::clear_type() {
  type_ = 0;
  clear_has_type();
}
inline ::google::protobuf::int32 TdfCode::type() const {
  // @@protoc_insertion_point(field_get:comm.TdfCode.Type)
  return type_;
}
inline void TdfCode::set_type(::google::protobuf::int32 value) {
  set_has_type();
  type_ = value;
  // @@protoc_insertion_point(field_set:comm.TdfCode.Type)
}

// -------------------------------------------------------------------

// TdfCodeTable

// repeated .comm.TdfCode CodeTable = 1;
inline int TdfCodeTable::codetable_size() const {
  return codetable_.size();
}
inline void TdfCodeTable::clear_codetable() {
  codetable_.Clear();
}
inline const ::comm::TdfCode& TdfCodeTable::codetable(int index) const {
  // @@protoc_insertion_point(field_get:comm.TdfCodeTable.CodeTable)
  return codetable_.Get(index);
}
inline ::comm::TdfCode* TdfCodeTable::mutable_codetable(int index) {
  // @@protoc_insertion_point(field_mutable:comm.TdfCodeTable.CodeTable)
  return codetable_.Mutable(index);
}
inline ::comm::TdfCode* TdfCodeTable::add_codetable() {
  // @@protoc_insertion_point(field_add:comm.TdfCodeTable.CodeTable)
  return codetable_.Add();
}
inline const ::google::protobuf::RepeatedPtrField< ::comm::TdfCode >&
TdfCodeTable::codetable() const {
  // @@protoc_insertion_point(field_list:comm.TdfCodeTable.CodeTable)
  return codetable_;
}
inline ::google::protobuf::RepeatedPtrField< ::comm::TdfCode >*
TdfCodeTable::mutable_codetable() {
  // @@protoc_insertion_point(field_mutable_list:comm.TdfCodeTable.CodeTable)
  return &codetable_;
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

#endif  // PROTOBUF_code_5ftable_2eproto__INCLUDED