package com.botregistry.core
import io.circe.generic.auto._

case class User(id: Int, isAdmin: Boolean, repos: List[Int])
    extends StorageItem[Int] {
  val key = id
}

case class Token(token: String, id: Int) extends StorageItem[String] {
  val key = token
}
