package com.botregistry.core

trait StorageItem[K] {
  def key: K
}

trait Storage[K, T <: StorageItem[K]] {
  def addOrUpdate(item: T): Option[Unit]
  def get(key: K): Option[T]
  def getLastKey: Option[K]
  def getAll: List[T]
  def delete(key: K): Option[Unit]
}