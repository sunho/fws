package com.botregistry.core

import com.botregistry.tasks.{BuildTask, BuildTaskFactory}
import com.botregistry.util.TimeUtil

import scala.concurrent.ExecutionContext.Implicits.global
import scala.util.{Failure, Success}

case class Build(repo: Repo, step: Int, tasks: List[BuildTask])

object Build {
  def apply(settings: BuildSettings, repo: Repo): Build = {
    val tasks = BuildTaskFactory(settings, repo)
    new Build(repo, 0, tasks)
  }

  def run(historyStore: Storage[Int, BuildHistory],
          build: Build,
          postBuild: (Repo, BuildHistory) => Unit): Unit = {
    if (BuildPool.building.exists(_._1 == build.repo.id)) {
      throw new IllegalArgumentException("the repo is already in build pool")
    }
    if (BuildHistory.getSubsequntBuilds(historyStore, build.repo).size >= 20) {
      throw new IllegalArgumentException(
        "you exceeded the number of builds allowed within an hour")
    }

    val fut = BuildPool.run(build)
    fut onComplete {
      case Success(t) => {
        val id = historyStore.getLastKey match {
          case Some(x) => x + 1
          case None    => 0
        }
        val time = TimeUtil.timestamp
        val history =
          BuildHistory(id, build.repo.id, time, t._1, t._2)
        if (historyStore.addOrUpdate(history).isEmpty) {
          println(s"history store failed ${build.repo} $history")
        } else {
          postBuild(build.repo, history)
        }
      }
      case Failure(e) => println(s"build failed $e ${build.repo}")
    }
  }
}
