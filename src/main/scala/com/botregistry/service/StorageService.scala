package com.botregistry.service

import com.botregistry.core._

trait StorageService {
  val userStore: Storage[Int, User]
  val repoStore: Storage[Int, Repo]
  val tokenStore: Storage[String, Token]
}
