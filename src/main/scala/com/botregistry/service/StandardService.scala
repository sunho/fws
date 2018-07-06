package com.botregistry.service

import com.botregistry.core._
import io.circe.{Encoder, Json}
import io.finch.circe._
import io.circe.generic.auto._

class StandardService(conf: Config)
    extends RepoService
    with UserService
    with UserRepoService
    with TokenService
    with WebhookService
    with BuildService
    with HistoryService {
  implicit val ee: Encoder[Exception] = Encoder.instance { e =>
    Json.obj(
      "message" -> Json.fromString(e.getMessage)
    )
  }

  override def postBuild(repo: Repo, history: BuildHistory): Unit = {
    println(repo, history)
  }
  override def buildSettings: BuildSettings = {
    BuildSettings(config.workspacePath,
                  config.dockerRegistry,
                  config.kubeNamespace)
  }
  override val config = conf
  override val userStore =
    FileStorage[String, User](s"${config.dataPath}/User.json")
  override val repoStore =
    FileStorage[Int, Repo](s"${config.dataPath}/Repo.json")
  override val tokenStore =
    FileStorage[String, Token](s"${config.dataPath}/Token.json")
  override val historyStore =
    FileStorage[Int, BuildHistory](s"${config.dataPath}/History.json")

  def api =
    repoApi :+: userApi :+: userRepoApi :+: tokenApi :+: webhookApi :+: buildApi :+: historyApi
  def toService = api.toService

  def save(): Unit = {
    historyStore.save()
    userStore.save()
    repoStore.save()
    tokenStore.save()
  }

  def startSaving(): Unit = {
    val t = new java.util.Timer()
    val task = new java.util.TimerTask {
      def run() = {
        println("saving")
        save()
      }
    }
    t.schedule(task, 10000L, 10000L)
  }
}
