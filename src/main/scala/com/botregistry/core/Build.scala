package com.botregistry.core

import com.botregistry.util.TimeUtil
import scala.util.{Failure, Success}
import scala.concurrent.ExecutionContext.Implicits.global

case class Build(repo: Repo, step: Int, tasks: List[BuildTask])

object Build {
  def apply(settings: BuildSettings, repo: Repo): Build = {
    val tasks = BuildTaskFactory(settings, repo)
    new Build(repo, 0, tasks)
  }

  private def makeHistory(id: Int,
                          time: Int,
                          repo: Repo,
                          res: (Boolean, List[(String, String)])) = {
    BuildHistory(id, repo.id, res._1, time, res._2)
  }

  def getSubsequntBuilds(historyStore: Storage[Int, BuildHistory],
                         repo: Repo): List[BuildHistory] = {
    val time = TimeUtil.timestamp
    historyStore.getAll.filter { x =>
      x.repoId == repo.id && time - x.time <= 60 * 60
    }
  }

  def buildRepo(historyStore: Storage[Int, BuildHistory],
                postBuild: (Repo, BuildHistory) => Unit,
                buildSettings: BuildSettings,
                repo: Repo): Unit = {
    if (BuildPool.building.exists(_._1 == repo.id)) {
      throw new IllegalArgumentException("the repo is already in build pool")
    }
    if (getSubsequntBuilds(historyStore, repo).size >= 20) {
      throw new IllegalArgumentException(
        "you exceeded the number of builds allowed within an hour")
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
}
