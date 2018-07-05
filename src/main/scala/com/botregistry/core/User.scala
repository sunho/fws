package com.botregistry.core

case class User(name: String, isAdmin: Boolean, repos: Set[Int])
    extends StorageItem[String] {
  val key = name
}

case class Token(token: String, name: String) extends StorageItem[String] {
  val key = token
}
