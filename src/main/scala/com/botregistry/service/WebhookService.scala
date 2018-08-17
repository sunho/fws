package com.botregistry.service

import com.botregistry.core._
import io.finch._
import io.finch.circe._
import io.finch.syntax._
import io.circe._
import com.botregistry.util.JsonUtil

trait WebhookService extends BuildService {
  val webhook: Endpoint[String] =
    post(
      "webhook" :: path[String] :: header("X-GitHub-Event") :: jsonBody[Json]) {
      (token: String, event: String, body: Json) =>
        val user = getUser(token)
        event match {
          case "push" => {
            val cursor = body.hcursor
            val url = JsonUtil.parse[String](cursor, "repository", "html_url")
            val ref = JsonUtil.parse[String](cursor, "ref")
            if (!ref.contains("master")){
              throw new IllegalArgumentException("invalid branch")
            }
            val repo =
              user.repos.flatMap(repoStore.get).find(_.repoURL == url) match {
                case Some(x) => x
                case None    => throw new IllegalArgumentException("invalid repo")
              }
            Build.run(historyStore, Build(buildSettings, repo), postBuild)
            Ok("build requested")
          }
          case "ping" => Ok("pong!")
          case _ =>
            throw new IllegalArgumentException(
              "only ping and push events are supported")
        }
    }.handle {
      case e: Exception => BadRequest(e)
    }

  val webhookApi = webhook
}