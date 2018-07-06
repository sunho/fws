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
    with WebhookService
    with BuildService {
  implicit val ee: Encoder[Exception] = Encoder.instance { e =>
    Json.obj(
      "message" -> Json.fromString(e.getMessage)
    )
  }

  override def postBuild(repo: Repo, history: BuildHistory): Unit = {
    println(repo, history)
  }
  override def buildSettings: BuildSettings = {
    BuildSettings(config.basePath, config.dockerRegistry, config.kubeNamespace)
  }
  override val config = Config.fromFile(s"$path/Config.json")
  override val userStore = FileStorage[String, User](s"$path/User.json")
  override val repoStore = FileStorage[Int, Repo](s"$path/Repo.json")
  override val tokenStore = FileStorage[String, Token](s"$path/Token.json")
  override val historyStore =
    FileStorage[Int, BuildHistory](s"$path/History.json")

  def api =
    repoApi :+: userApi :+: userRepoApi :+: tokenApi :+: webhookApi :+: buildApi
  def toService = api.toService

  def save() {
    historyStore.save()
    userStore.save()
    repoStore.save()
    tokenStore.save()
  }
}
