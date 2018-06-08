// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: orderqueueData.proto

#ifndef PROTOBUF_orderqueueData_2eproto__INCLUDED
#define PROTOBUF_orderqueueData_2eproto__INCLUDED

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
void protobuf_AddDesc_orderqueueData_2eproto();
void protobuf_InitDefaults_orderqueueData_2eproto();
void protobuf_AssignDesc_orderqueueData_2eproto();
void protobuf_ShutdownFile_orderqueueData_2eproto();

class OrderQueueData;

// ===================================================================

class OrderQueueData : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:comm.OrderQueueData) */ {
 public:
  OrderQueueData();
  virtual ~OrderQueueData();

  OrderQueueData(const OrderQueueData& from);

  inline OrderQueueData& operator=(const OrderQueueData& from) {
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
  static const OrderQueueData& default_instance();

  static const OrderQueueData* internal_default_instance();

  void Swap(OrderQueueData* other);

  // implements Message ----------------------------------------------

  inline OrderQueueData* New() const { return New(NULL); }

  OrderQueueData* New(::google::protobuf::Arena* arena) const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const OrderQueueData& from);
  void MergeFrom(const OrderQueueData& from);
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
  void InternalSwap(OrderQueueData* other);
  void UnsafeMergeFrom(const OrderQueueData& from);
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

  // required string WindCode = 1;
  bool has_windcode() const;
  void clear_windcode();
  static const int kWindCodeFieldNumber = 1;
  const ::std::string& windcode() const;
  void set_windcode(const ::std::string& value);
  void set_windcode(const char* value);
  void set_windcode(const char* value, size_t size);
  ::std::string* mutable_windcode();
  ::std::string* release_windcode();
  void set_allocated_windcode(::std::string* windcode);

  // required string Code = 2;
  bool has_code() const;
  void clear_code();
  static const int kCodeFieldNumber = 2;
  const ::std::string& code() const;
  void set_code(const ::std::string& value);
  void set_code(const char* value);
  void set_code(const char* value, size_t size);
  ::std::string* mutable_code();
  ::std::string* release_code();
  void set_allocated_code(::std::string* code);

  // required int32 ActionDay = 3;
  bool has_actionday() const;
  void clear_actionday();
  static const int kActionDayFieldNumber = 3;
  ::google::protobuf::int32 actionday() const;
  void set_actionday(::google::protobuf::int32 value);

  // required int32 Time = 4;
  bool has_time() const;
  void clear_time();
  static const int kTimeFieldNumber = 4;
  ::google::protobuf::int32 time() const;
  void set_time(::google::protobuf::int32 value);

  // required int32 Side = 5;
  bool has_side() const;
  void clear_side();
  static const int kSideFieldNumber = 5;
  ::google::protobuf::int32 side() const;
  void set_side(::google::protobuf::int32 value);

  // required int32 Price = 6;
  bool has_price() const;
  void clear_price();
  static const int kPriceFieldNumber = 6;
  ::google::protobuf::int32 price() const;
  void set_price(::google::protobuf::int32 value);

  // required int32 Orders = 7;
  bool has_orders() const;
  void clear_orders();
  static const int kOrdersFieldNumber = 7;
  ::google::protobuf::int32 orders() const;
  void set_orders(::google::protobuf::int32 value);

  // required int32 ABItems = 8;
  bool has_abitems() const;
  void clear_abitems();
  static const int kABItemsFieldNumber = 8;
  ::google::protobuf::int32 abitems() const;
  void set_abitems(::google::protobuf::int32 value);

  // repeated int32 ABVolume = 9;
  int abvolume_size() const;
  void clear_abvolume();
  static const int kABVolumeFieldNumber = 9;
  ::google::protobuf::int32 abvolume(int index) const;
  void set_abvolume(int index, ::google::protobuf::int32 value);
  void add_abvolume(::google::protobuf::int32 value);
  const ::google::protobuf::RepeatedField< ::google::protobuf::int32 >&
      abvolume() const;
  ::google::protobuf::RepeatedField< ::google::protobuf::int32 >*
      mutable_abvolume();

  // @@protoc_insertion_point(class_scope:comm.OrderQueueData)
 private:
  inline void set_has_windcode();
  inline void clear_has_windcode();
  inline void set_has_code();
  inline void clear_has_code();
  inline void set_has_actionday();
  inline void clear_has_actionday();
  inline void set_has_time();
  inline void clear_has_time();
  inline void set_has_side();
  inline void clear_has_side();
  inline void set_has_price();
  inline void clear_has_price();
  inline void set_has_orders();
  inline void clear_has_orders();
  inline void set_has_abitems();
  inline void clear_has_abitems();

  // helper for ByteSizeLong()
  size_t RequiredFieldsByteSizeFallback() const;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::HasBits<1> _has_bits_;
  mutable int _cached_size_;
  ::google::protobuf::RepeatedField< ::google::protobuf::int32 > abvolume_;
  ::google::protobuf::internal::ArenaStringPtr windcode_;
  ::google::protobuf::internal::ArenaStringPtr code_;
  ::google::protobuf::int32 actionday_;
  ::google::protobuf::int32 time_;
  ::google::protobuf::int32 side_;
  ::google::protobuf::int32 price_;
  ::google::protobuf::int32 orders_;
  ::google::protobuf::int32 abitems_;
  friend void  protobuf_InitDefaults_orderqueueData_2eproto_impl();
  friend void  protobuf_AddDesc_orderqueueData_2eproto_impl();
  friend void protobuf_AssignDesc_orderqueueData_2eproto();
  friend void protobuf_ShutdownFile_orderqueueData_2eproto();

  void InitAsDefaultInstance();
};
extern ::google::protobuf::internal::ExplicitlyConstructed<OrderQueueData> OrderQueueData_default_instance_;

// ===================================================================


// ===================================================================

#if !PROTOBUF_INLINE_NOT_IN_HEADERS
// OrderQueueData

// required string WindCode = 1;
inline bool OrderQueueData::has_windcode() const {
  return (_has_bits_[0] & 0x00000001u) != 0;
}
inline void OrderQueueData::set_has_windcode() {
  _has_bits_[0] |= 0x00000001u;
}
inline void OrderQueueData::clear_has_windcode() {
  _has_bits_[0] &= ~0x00000001u;
}
inline void OrderQueueData::clear_windcode() {
  windcode_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_windcode();
}
inline const ::std::string& OrderQueueData::windcode() const {
  // @@protoc_insertion_point(field_get:comm.OrderQueueData.WindCode)
  return windcode_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void OrderQueueData::set_windcode(const ::std::string& value) {
  set_has_windcode();
  windcode_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:comm.OrderQueueData.WindCode)
}
inline void OrderQueueData::set_windcode(const char* value) {
  set_has_windcode();
  windcode_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:comm.OrderQueueData.WindCode)
}
inline void OrderQueueData::set_windcode(const char* value, size_t size) {
  set_has_windcode();
  windcode_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:comm.OrderQueueData.WindCode)
}
inline ::std::string* OrderQueueData::mutable_windcode() {
  set_has_windcode();
  // @@protoc_insertion_point(field_mutable:comm.OrderQueueData.WindCode)
  return windcode_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* OrderQueueData::release_windcode() {
  // @@protoc_insertion_point(field_release:comm.OrderQueueData.WindCode)
  clear_has_windcode();
  return windcode_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void OrderQueueData::set_allocated_windcode(::std::string* windcode) {
  if (windcode != NULL) {
    set_has_windcode();
  } else {
    clear_has_windcode();
  }
  windcode_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), windcode);
  // @@protoc_insertion_point(field_set_allocated:comm.OrderQueueData.WindCode)
}

// required string Code = 2;
inline bool OrderQueueData::has_code() const {
  return (_has_bits_[0] & 0x00000002u) != 0;
}
inline void OrderQueueData::set_has_code() {
  _has_bits_[0] |= 0x00000002u;
}
inline void OrderQueueData::clear_has_code() {
  _has_bits_[0] &= ~0x00000002u;
}
inline void OrderQueueData::clear_code() {
  code_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  clear_has_code();
}
inline const ::std::string& OrderQueueData::code() const {
  // @@protoc_insertion_point(field_get:comm.OrderQueueData.Code)
  return code_.GetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void OrderQueueData::set_code(const ::std::string& value) {
  set_has_code();
  code_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:comm.OrderQueueData.Code)
}
inline void OrderQueueData::set_code(const char* value) {
  set_has_code();
  code_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:comm.OrderQueueData.Code)
}
inline void OrderQueueData::set_code(const char* value, size_t size) {
  set_has_code();
  code_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:comm.OrderQueueData.Code)
}
inline ::std::string* OrderQueueData::mutable_code() {
  set_has_code();
  // @@protoc_insertion_point(field_mutable:comm.OrderQueueData.Code)
  return code_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* OrderQueueData::release_code() {
  // @@protoc_insertion_point(field_release:comm.OrderQueueData.Code)
  clear_has_code();
  return code_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void OrderQueueData::set_allocated_code(::std::string* code) {
  if (code != NULL) {
    set_has_code();
  } else {
    clear_has_code();
  }
  code_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), code);
  // @@protoc_insertion_point(field_set_allocated:comm.OrderQueueData.Code)
}

// required int32 ActionDay = 3;
inline bool OrderQueueData::has_actionday() const {
  return (_has_bits_[0] & 0x00000004u) != 0;
}
inline void OrderQueueData::set_has_actionday() {
  _has_bits_[0] |= 0x00000004u;
}
inline void OrderQueueData::clear_has_actionday() {
  _has_bits_[0] &= ~0x00000004u;
}
inline void OrderQueueData::clear_actionday() {
  actionday_ = 0;
  clear_has_actionday();
}
inline ::google::protobuf::int32 OrderQueueData::actionday() const {
  // @@protoc_insertion_point(field_get:comm.OrderQueueData.ActionDay)
  return actionday_;
}
inline void OrderQueueData::set_actionday(::google::protobuf::int32 value) {
  set_has_actionday();
  actionday_ = value;
  // @@protoc_insertion_point(field_set:comm.OrderQueueData.ActionDay)
}

// required int32 Time = 4;
inline bool OrderQueueData::has_time() const {
  return (_has_bits_[0] & 0x00000008u) != 0;
}
inline void OrderQueueData::set_has_time() {
  _has_bits_[0] |= 0x00000008u;
}
inline void OrderQueueData::clear_has_time() {
  _has_bits_[0] &= ~0x00000008u;
}
inline void OrderQueueData::clear_time() {
  time_ = 0;
  clear_has_time();
}
inline ::google::protobuf::int32 OrderQueueData::time() const {
  // @@protoc_insertion_point(field_get:comm.OrderQueueData.Time)
  return time_;
}
inline void OrderQueueData::set_time(::google::protobuf::int32 value) {
  set_has_time();
  time_ = value;
  // @@protoc_insertion_point(field_set:comm.OrderQueueData.Time)
}

// required int32 Side = 5;
inline bool OrderQueueData::has_side() const {
  return (_has_bits_[0] & 0x00000010u) != 0;
}
inline void OrderQueueData::set_has_side() {
  _has_bits_[0] |= 0x00000010u;
}
inline void OrderQueueData::clear_has_side() {
  _has_bits_[0] &= ~0x00000010u;
}
inline void OrderQueueData::clear_side() {
  side_ = 0;
  clear_has_side();
}
inline ::google::protobuf::int32 OrderQueueData::side() const {
  // @@protoc_insertion_point(field_get:comm.OrderQueueData.Side)
  return side_;
}
inline void OrderQueueData::set_side(::google::protobuf::int32 value) {
  set_has_side();
  side_ = value;
  // @@protoc_insertion_point(field_set:comm.OrderQueueData.Side)
}

// required int32 Price = 6;
inline bool OrderQueueData::has_price() const {
  return (_has_bits_[0] & 0x00000020u) != 0;
}
inline void OrderQueueData::set_has_price() {
  _has_bits_[0] |= 0x00000020u;
}
inline void OrderQueueData::clear_has_price() {
  _has_bits_[0] &= ~0x00000020u;
}
inline void OrderQueueData::clear_price() {
  price_ = 0;
  clear_has_price();
}
inline ::google::protobuf::int32 OrderQueueData::price() const {
  // @@protoc_insertion_point(field_get:comm.OrderQueueData.Price)
  return price_;
}
inline void OrderQueueData::set_price(::google::protobuf::int32 value) {
  set_has_price();
  price_ = value;
  // @@protoc_insertion_point(field_set:comm.OrderQueueData.Price)
}

// required int32 Orders = 7;
inline bool OrderQueueData::has_orders() const {
  return (_has_bits_[0] & 0x00000040u) != 0;
}
inline void OrderQueueData::set_has_orders() {
  _has_bits_[0] |= 0x00000040u;
}
inline void OrderQueueData::clear_has_orders() {
  _has_bits_[0] &= ~0x00000040u;
}
inline void OrderQueueData::clear_orders() {
  orders_ = 0;
  clear_has_orders();
}
inline ::google::protobuf::int32 OrderQueueData::orders() const {
  // @@protoc_insertion_point(field_get:comm.OrderQueueData.Orders)
  return orders_;
}
inline void OrderQueueData::set_orders(::google::protobuf::int32 value) {
  set_has_orders();
  orders_ = value;
  // @@protoc_insertion_point(field_set:comm.OrderQueueData.Orders)
}

// required int32 ABItems = 8;
inline bool OrderQueueData::has_abitems() const {
  return (_has_bits_[0] & 0x00000080u) != 0;
}
inline void OrderQueueData::set_has_abitems() {
  _has_bits_[0] |= 0x00000080u;
}
inline void OrderQueueData::clear_has_abitems() {
  _has_bits_[0] &= ~0x00000080u;
}
inline void OrderQueueData::clear_abitems() {
  abitems_ = 0;
  clear_has_abitems();
}
inline ::google::protobuf::int32 OrderQueueData::abitems() const {
  // @@protoc_insertion_point(field_get:comm.OrderQueueData.ABItems)
  return abitems_;
}
inline void OrderQueueData::set_abitems(::google::protobuf::int32 value) {
  set_has_abitems();
  abitems_ = value;
  // @@protoc_insertion_point(field_set:comm.OrderQueueData.ABItems)
}

// repeated int32 ABVolume = 9;
inline int OrderQueueData::abvolume_size() const {
  return abvolume_.size();
}
inline void OrderQueueData::clear_abvolume() {
  abvolume_.Clear();
}
inline ::google::protobuf::int32 OrderQueueData::abvolume(int index) const {
  // @@protoc_insertion_point(field_get:comm.OrderQueueData.ABVolume)
  return abvolume_.Get(index);
}
inline void OrderQueueData::set_abvolume(int index, ::google::protobuf::int32 value) {
  abvolume_.Set(index, value);
  // @@protoc_insertion_point(field_set:comm.OrderQueueData.ABVolume)
}
inline void OrderQueueData::add_abvolume(::google::protobuf::int32 value) {
  abvolume_.Add(value);
  // @@protoc_insertion_point(field_add:comm.OrderQueueData.ABVolume)
}
inline const ::google::protobuf::RepeatedField< ::google::protobuf::int32 >&
OrderQueueData::abvolume() const {
  // @@protoc_insertion_point(field_list:comm.OrderQueueData.ABVolume)
  return abvolume_;
}
inline ::google::protobuf::RepeatedField< ::google::protobuf::int32 >*
OrderQueueData::mutable_abvolume() {
  // @@protoc_insertion_point(field_mutable_list:comm.OrderQueueData.ABVolume)
  return &abvolume_;
}

inline const OrderQueueData* OrderQueueData::internal_default_instance() {
  return &OrderQueueData_default_instance_.get();
}
#endif  // !PROTOBUF_INLINE_NOT_IN_HEADERS

// @@protoc_insertion_point(namespace_scope)

}  // namespace comm

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_orderqueueData_2eproto__INCLUDED
