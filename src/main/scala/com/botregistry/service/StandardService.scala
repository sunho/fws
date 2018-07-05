package com.botregistry.service
import com.botregistry.core._
import io.finch.circe._
import io.circe.syntax._
import io.circe.generic.auto._

class StandardService(path: String) extends RepoService {
  override val config = Config.fromFile(s"$path/Config.json")
  override val userStore = FileStorage[Int, User](s"$path/User.json")
  override val repoStore = FileStorage[Int, Repo](s"$path/Repo.json")
  override val tokenStore = FileStorage[String, Token](s"$path/Token.json")
  def endpoints = getRepos :+: addRepos :+: authenticate
  def toService = endpoints.toService
}
