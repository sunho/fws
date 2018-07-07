package com.botregistry.service

import io.finch._
import io.finch.syntax._
import com.botregistry.core._
import com.botregistry.util.TimeUtil

import scala.util.{Success, Failure}

trait BuildService extends RepoService {
  def buildSettings: BuildSettings
  def postBuild(repo: Repo, history: BuildHistory): Unit

  val buildRepo: Endpoint[Unit] =
    post(authenticate :: repoEndpoint :: repoPath :: "build") {
      (u: User, repo: Repo) =>
        if (u.isAdmin || u.repos.contains(repo.id)) {
          Build.run(historyStore, Build(buildSettings, repo), postBuild)
          Ok()
        } else {
          Unauthorized(new IllegalAccessException)
        }
    }.handle {
      case e: Exception => BadRequest(e)
    }

  val buildApi = buildRepo
}
