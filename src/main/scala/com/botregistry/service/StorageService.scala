package com.botregistry.service

import com.botregistry.core._

trait StorageService {
  val userStore: Storage[String, User]
  val repoStore: Storage[Int, Repo]
  val tokenStore: Storage[String, Token]
}
