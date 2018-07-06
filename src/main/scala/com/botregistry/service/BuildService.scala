package com.botregistry.service

import io.finch._
import io.finch.circe._
import io.finch.syntax._
import io.circe.generic.auto._
import com.botregistry.core._
import com.botregistry.util.{TimeUtil, TokenGenerator}
import scala.concurrent.ExecutionContext.Implicits.global

import scala.util.{Failure, Success}

trait BuildService extends RepoService {
  def buildSettings: BuildSettings
  def postBuild(repo: Repo, history: BuildHistory): Unit

  protected def makeHistory(id: Int,
                            time: Int,
                            repo: Repo,
                            res: (Boolean, Map[String, String])) = {
    BuildHistory(id, repo.id, res._1, time, res._2)
  }

  protected def build(repo: Repo): Unit = {
    if (BuildPool.building.exists(_._1 == repo.id)) {
      throw new IllegalArgumentException("the repo is already in build pool")
    }
    val fut = BuildPool.run(Build(buildSettings, repo))
    fut onComplete {
      case Success(t) => {
        val id = historyStore.getLastKey match {
          case Some(x) => x + 1
          case None    => 0
        }
        val time = TimeUtil.timestamp
        val history = makeHistory(id, time, repo, t)
        if (historyStore.addOrUpdate(history).isEmpty) {
          println(s"history store failed $repo $history")
          return
        }
        postBuild(repo, history)
      }
      case Failure(e) => println(s"build failed $e $repo")
    }
  }

  val buildRepo: Endpoint[Unit] =
    get(authenticate :: repoEndpoint :: repoPath) { (u: User, repo: Repo) =>
      if (u.isAdmin || u.repos.contains(repo.id)) {
        build(repo)
        Ok()
      } else {
        Unauthorized(new IllegalAccessException)
      }
    }.handle {
      case e: Exception => BadRequest(e)
    }

  val buildApi = buildRepo
}
