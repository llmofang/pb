// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: marketDataSet.proto

#ifndef PROTOBUF_marketDataSet_2eproto__INCLUDED
#define PROTOBUF_marketDataSet_2eproto__INCLUDED

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
#include "marketData.pb.h"
#include "transactionData.pb.h"
#include "orderqueueData.pb.h"
#include "indexData.pb.h"
#include "futureData.pb.h"
#include "orderData.pb.h"
// @@protoc_insertion_point(includes)

namespace comm {

// Internal implementation detail -- do not call these.
void protobuf_AddDesc_marketDataSet_2eproto();
void protobuf_InitDefaults_marketDataSet_2eproto();
void protobuf_AssignDesc_marketDataSet_2eproto();
void protobuf_ShutdownFile_marketDataSet_2eproto();

class MarketDataSet;

// ===================================================================

class MarketDataSet : public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:comm.MarketDataSet) */ {
 public:
  MarketDataSet();
  virtual ~MarketDataSet();

  MarketDataSet(const MarketDataSet& from);

  inline MarketDataSet& operator=(const MarketDataSet& from) {
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
  static const MarketDataSet& default_instance();

  static const MarketDataSet* internal_default_instance();

  void Swap(MarketDataSet* other);

  // implements Message ----------------------------------------------

  inline MarketDataSet* New() const { return New(NULL); }

  MarketDataSet* New(::google::protobuf::Arena* arena) const;
  void CopyFrom(const ::google::protobuf::Message& from);
  void MergeFrom(const ::google::protobuf::Message& from);
  void CopyFrom(const MarketDataSet& from);
  void MergeFrom(const MarketDataSet& from);
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
  void InternalSwap(MarketDataSet* other);
  void UnsafeMergeFrom(const MarketDataSet& from);
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

  // required int32 DataType = 1;
  bool has_datatype() const;
  void clear_datatype();
  static const int kDataTypeFieldNumber = 1;
  ::google::protobuf::int32 datatype() const;
  void set_datatype(::google::protobuf::int32 value);

  // optional .comm.MarketData Tick = 2;
  bool has_tick() const;
  void clear_tick();
  static const int kTickFieldNumber = 2;
  const ::comm::MarketData& tick() const;
  ::comm::MarketData* mutable_tick();
  ::comm::MarketData* release_tick();
  void set_allocated_tick(::comm::MarketData* tick);

  // optional .comm.TransactionData Transaction = 3;
  bool has_transaction() const;
  void clear_transaction();
  static const int kTransactionFieldNumber = 3;
  const ::comm::TransactionData& transaction() const;
  ::comm::TransactionData* mutable_transaction();
  ::comm::TransactionData* release_transaction();
  void set_allocated_transaction(::comm::TransactionData* transaction);

  // optional .comm.OrderQueueData OrderQueue = 4;
  bool has_orderqueue() const;
  void clear_orderqueue();
  static const int kOrderQueueFieldNumber = 4;
  const ::comm::OrderQueueData& orderqueue() const;
  ::comm::OrderQueueData* mutable_orderqueue();
  ::comm::OrderQueueData* release_orderqueue();
  void set_allocated_orderqueue(::comm::OrderQueueData* orderqueue);

  // optional .comm.IndexData Index = 5;
  bool has_index() const;
  void clear_index();
  static const int kIndexFieldNumber = 5;
  const ::comm::IndexData& index() const;
  ::comm::IndexData* mutable_index();
  ::comm::IndexData* release_index();
  void set_allocated_index(::comm::IndexData* index);

  // optional .comm.FutureData future = 6;
  bool has_future() const;
  void clear_future();
  static const int kFutureFieldNumber = 6;
  const ::comm::FutureData& future() const;
  ::comm::FutureData* mutable_future();
  ::comm::FutureData* release_future();
  void set_allocated_future(::comm::FutureData* future);

  // optional .comm.OrderData order = 7;
  bool has_order() const;
  void clear_order();
  static const int kOrderFieldNumber = 7;
  const ::comm::OrderData& order() const;
  ::comm::OrderData* mutable_order();
  ::comm::OrderData* release_order();
  void set_allocated_order(::comm::OrderData* order);

  // @@protoc_insertion_point(class_scope:comm.MarketDataSet)
 private:
  inline void set_has_datatype();
  inline void clear_has_datatype();
  inline void set_has_tick();
  inline void clear_has_tick();
  inline void set_has_transaction();
  inline void clear_has_transaction();
  inline void set_has_orderqueue();
  inline void clear_has_orderqueue();
  inline void set_has_index();
  inline void clear_has_index();
  inline void set_has_future();
  inline void clear_has_future();
  inline void set_has_order();
  inline void clear_has_order();

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::HasBits<1> _has_bits_;
  mutable int _cached_size_;
  ::comm::MarketData* tick_;
  ::comm::TransactionData* transaction_;
  ::comm::OrderQueueData* orderqueue_;
  ::comm::IndexData* index_;
  ::comm::FutureData* future_;
  ::comm::OrderData* order_;
  ::google::protobuf::int32 datatype_;
  friend void  protobuf_InitDefaults_marketDataSet_2eproto_impl();
  friend void  protobuf_AddDesc_marketDataSet_2eproto_impl();
  friend void protobuf_AssignDesc_marketDataSet_2eproto();
  friend void protobuf_ShutdownFile_marketDataSet_2eproto();

  void InitAsDefaultInstance();
};
extern ::google::protobuf::internal::ExplicitlyConstructed<MarketDataSet> MarketDataSet_default_instance_;

// ===================================================================


// ===================================================================

#if !PROTOBUF_INLINE_NOT_IN_HEADERS
// MarketDataSet

// required int32 DataType = 1;
inline bool MarketDataSet::has_datatype() const {
  return (_has_bits_[0] & 0x00000001u) != 0;
}
inline void MarketDataSet::set_has_datatype() {
  _has_bits_[0] |= 0x00000001u;
}
inline void MarketDataSet::clear_has_datatype() {
  _has_bits_[0] &= ~0x00000001u;
}
inline void MarketDataSet::clear_datatype() {
  datatype_ = 0;
  clear_has_datatype();
}
inline ::google::protobuf::int32 MarketDataSet::datatype() const {
  // @@protoc_insertion_point(field_get:comm.MarketDataSet.DataType)
  return datatype_;
}
inline void MarketDataSet::set_datatype(::google::protobuf::int32 value) {
  set_has_datatype();
  datatype_ = value;
  // @@protoc_insertion_point(field_set:comm.MarketDataSet.DataType)
}

// optional .comm.MarketData Tick = 2;
inline bool MarketDataSet::has_tick() const {
  return (_has_bits_[0] & 0x00000002u) != 0;
}
inline void MarketDataSet::set_has_tick() {
  _has_bits_[0] |= 0x00000002u;
}
inline void MarketDataSet::clear_has_tick() {
  _has_bits_[0] &= ~0x00000002u;
}
inline void MarketDataSet::clear_tick() {
  if (tick_ != NULL) tick_->::comm::MarketData::Clear();
  clear_has_tick();
}
inline const ::comm::MarketData& MarketDataSet::tick() const {
  // @@protoc_insertion_point(field_get:comm.MarketDataSet.Tick)
  return tick_ != NULL ? *tick_
                         : *::comm::MarketData::internal_default_instance();
}
inline ::comm::MarketData* MarketDataSet::mutable_tick() {
  set_has_tick();
  if (tick_ == NULL) {
    tick_ = new ::comm::MarketData;
  }
  // @@protoc_insertion_point(field_mutable:comm.MarketDataSet.Tick)
  return tick_;
}
inline ::comm::MarketData* MarketDataSet::release_tick() {
  // @@protoc_insertion_point(field_release:comm.MarketDataSet.Tick)
  clear_has_tick();
  ::comm::MarketData* temp = tick_;
  tick_ = NULL;
  return temp;
}
inline void MarketDataSet::set_allocated_tick(::comm::MarketData* tick) {
  delete tick_;
  tick_ = tick;
  if (tick) {
    set_has_tick();
  } else {
    clear_has_tick();
  }
  // @@protoc_insertion_point(field_set_allocated:comm.MarketDataSet.Tick)
}

// optional .comm.TransactionData Transaction = 3;
inline bool MarketDataSet::has_transaction() const {
  return (_has_bits_[0] & 0x00000004u) != 0;
}
inline void MarketDataSet::set_has_transaction() {
  _has_bits_[0] |= 0x00000004u;
}
inline void MarketDataSet::clear_has_transaction() {
  _has_bits_[0] &= ~0x00000004u;
}
inline void MarketDataSet::clear_transaction() {
  if (transaction_ != NULL) transaction_->::comm::TransactionData::Clear();
  clear_has_transaction();
}
inline const ::comm::TransactionData& MarketDataSet::transaction() const {
  // @@protoc_insertion_point(field_get:comm.MarketDataSet.Transaction)
  return transaction_ != NULL ? *transaction_
                         : *::comm::TransactionData::internal_default_instance();
}
inline ::comm::TransactionData* MarketDataSet::mutable_transaction() {
  set_has_transaction();
  if (transaction_ == NULL) {
    transaction_ = new ::comm::TransactionData;
  }
  // @@protoc_insertion_point(field_mutable:comm.MarketDataSet.Transaction)
  return transaction_;
}
inline ::comm::TransactionData* MarketDataSet::release_transaction() {
  // @@protoc_insertion_point(field_release:comm.MarketDataSet.Transaction)
  clear_has_transaction();
  ::comm::TransactionData* temp = transaction_;
  transaction_ = NULL;
  return temp;
}
inline void MarketDataSet::set_allocated_transaction(::comm::TransactionData* transaction) {
  delete transaction_;
  transaction_ = transaction;
  if (transaction) {
    set_has_transaction();
  } else {
    clear_has_transaction();
  }
  // @@protoc_insertion_point(field_set_allocated:comm.MarketDataSet.Transaction)
}

// optional .comm.OrderQueueData OrderQueue = 4;
inline bool MarketDataSet::has_orderqueue() const {
  return (_has_bits_[0] & 0x00000008u) != 0;
}
inline void MarketDataSet::set_has_orderqueue() {
  _has_bits_[0] |= 0x00000008u;
}
inline void MarketDataSet::clear_has_orderqueue() {
  _has_bits_[0] &= ~0x00000008u;
}
inline void MarketDataSet::clear_orderqueue() {
  if (orderqueue_ != NULL) orderqueue_->::comm::OrderQueueData::Clear();
  clear_has_orderqueue();
}
inline const ::comm::OrderQueueData& MarketDataSet::orderqueue() const {
  // @@protoc_insertion_point(field_get:comm.MarketDataSet.OrderQueue)
  return orderqueue_ != NULL ? *orderqueue_
                         : *::comm::OrderQueueData::internal_default_instance();
}
inline ::comm::OrderQueueData* MarketDataSet::mutable_orderqueue() {
  set_has_orderqueue();
  if (orderqueue_ == NULL) {
    orderqueue_ = new ::comm::OrderQueueData;
  }
  // @@protoc_insertion_point(field_mutable:comm.MarketDataSet.OrderQueue)
  return orderqueue_;
}
inline ::comm::OrderQueueData* MarketDataSet::release_orderqueue() {
  // @@protoc_insertion_point(field_release:comm.MarketDataSet.OrderQueue)
  clear_has_orderqueue();
  ::comm::OrderQueueData* temp = orderqueue_;
  orderqueue_ = NULL;
  return temp;
}
inline void MarketDataSet::set_allocated_orderqueue(::comm::OrderQueueData* orderqueue) {
  delete orderqueue_;
  orderqueue_ = orderqueue;
  if (orderqueue) {
    set_has_orderqueue();
  } else {
    clear_has_orderqueue();
  }
  // @@protoc_insertion_point(field_set_allocated:comm.MarketDataSet.OrderQueue)
}

// optional .comm.IndexData Index = 5;
inline bool MarketDataSet::has_index() const {
  return (_has_bits_[0] & 0x00000010u) != 0;
}
inline void MarketDataSet::set_has_index() {
  _has_bits_[0] |= 0x00000010u;
}
inline void MarketDataSet::clear_has_index() {
  _has_bits_[0] &= ~0x00000010u;
}
inline void MarketDataSet::clear_index() {
  if (index_ != NULL) index_->::comm::IndexData::Clear();
  clear_has_index();
}
inline const ::comm::IndexData& MarketDataSet::index() const {
  // @@protoc_insertion_point(field_get:comm.MarketDataSet.Index)
  return index_ != NULL ? *index_
                         : *::comm::IndexData::internal_default_instance();
}
inline ::comm::IndexData* MarketDataSet::mutable_index() {
  set_has_index();
  if (index_ == NULL) {
    index_ = new ::comm::IndexData;
  }
  // @@protoc_insertion_point(field_mutable:comm.MarketDataSet.Index)
  return index_;
}
inline ::comm::IndexData* MarketDataSet::release_index() {
  // @@protoc_insertion_point(field_release:comm.MarketDataSet.Index)
  clear_has_index();
  ::comm::IndexData* temp = index_;
  index_ = NULL;
  return temp;
}
inline void MarketDataSet::set_allocated_index(::comm::IndexData* index) {
  delete index_;
  index_ = index;
  if (index) {
    set_has_index();
  } else {
    clear_has_index();
  }
  // @@protoc_insertion_point(field_set_allocated:comm.MarketDataSet.Index)
}

// optional .comm.FutureData future = 6;
inline bool MarketDataSet::has_future() const {
  return (_has_bits_[0] & 0x00000020u) != 0;
}
inline void MarketDataSet::set_has_future() {
  _has_bits_[0] |= 0x00000020u;
}
inline void MarketDataSet::clear_has_future() {
  _has_bits_[0] &= ~0x00000020u;
}
inline void MarketDataSet::clear_future() {
  if (future_ != NULL) future_->::comm::FutureData::Clear();
  clear_has_future();
}
inline const ::comm::FutureData& MarketDataSet::future() const {
  // @@protoc_insertion_point(field_get:comm.MarketDataSet.future)
  return future_ != NULL ? *future_
                         : *::comm::FutureData::internal_default_instance();
}
inline ::comm::FutureData* MarketDataSet::mutable_future() {
  set_has_future();
  if (future_ == NULL) {
    future_ = new ::comm::FutureData;
  }
  // @@protoc_insertion_point(field_mutable:comm.MarketDataSet.future)
  return future_;
}
inline ::comm::FutureData* MarketDataSet::release_future() {
  // @@protoc_insertion_point(field_release:comm.MarketDataSet.future)
  clear_has_future();
  ::comm::FutureData* temp = future_;
  future_ = NULL;
  return temp;
}
inline void MarketDataSet::set_allocated_future(::comm::FutureData* future) {
  delete future_;
  future_ = future;
  if (future) {
    set_has_future();
  } else {
    clear_has_future();
  }
  // @@protoc_insertion_point(field_set_allocated:comm.MarketDataSet.future)
}

// optional .comm.OrderData order = 7;
inline bool MarketDataSet::has_order() const {
  return (_has_bits_[0] & 0x00000040u) != 0;
}
inline void MarketDataSet::set_has_order() {
  _has_bits_[0] |= 0x00000040u;
}
inline void MarketDataSet::clear_has_order() {
  _has_bits_[0] &= ~0x00000040u;
}
inline void MarketDataSet::clear_order() {
  if (order_ != NULL) order_->::comm::OrderData::Clear();
  clear_has_order();
}
inline const ::comm::OrderData& MarketDataSet::order() const {
  // @@protoc_insertion_point(field_get:comm.MarketDataSet.order)
  return order_ != NULL ? *order_
                         : *::comm::OrderData::internal_default_instance();
}
inline ::comm::OrderData* MarketDataSet::mutable_order() {
  set_has_order();
  if (order_ == NULL) {
    order_ = new ::comm::OrderData;
  }
  // @@protoc_insertion_point(field_mutable:comm.MarketDataSet.order)
  return order_;
}
inline ::comm::OrderData* MarketDataSet::release_order() {
  // @@protoc_insertion_point(field_release:comm.MarketDataSet.order)
  clear_has_order();
  ::comm::OrderData* temp = order_;
  order_ = NULL;
  return temp;
}
inline void MarketDataSet::set_allocated_order(::comm::OrderData* order) {
  delete order_;
  order_ = order;
  if (order) {
    set_has_order();
  } else {
    clear_has_order();
  }
  // @@protoc_insertion_point(field_set_allocated:comm.MarketDataSet.order)
}

inline const MarketDataSet* MarketDataSet::internal_default_instance() {
  return &MarketDataSet_default_instance_.get();
}
#endif  // !PROTOBUF_INLINE_NOT_IN_HEADERS

// @@protoc_insertion_point(namespace_scope)

}  // namespace comm

// @@protoc_insertion_point(global_scope)

#endif  // PROTOBUF_marketDataSet_2eproto__INCLUDED
