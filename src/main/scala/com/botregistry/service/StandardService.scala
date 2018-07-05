package com.botregistry.service

import com.botregistry.core._
import io.circe.{Encoder, Json}
import io.finch.circe._
import io.circe.generic.auto._

class StandardService(path: String)
    extends RepoService
    with UserService
    with UserRepoService
    with TokenService
    with WebhookService {

  implicit val ee: Encoder[Exception] = Encoder.instance { e =>
    Json.obj(
      "message" -> Json.fromString(e.getMessage)
    )
  }

  override val config = Config.fromFile(s"$path/Config.json")
  override val userStore = FileStorage[String, User](s"$path/User.json")
  override val repoStore = FileStorage[Int, Repo](s"$path/Repo.json")
  override val tokenStore = FileStorage[String, Token](s"$path/Token.json")

  def api = repoApi :+: userApi :+: userRepoApi :+: tokenApi :+: webhookApi
  def toService = api.toService

  def save() {
    userStore.save()
    repoStore.save()
    tokenStore.save()
  }
}
