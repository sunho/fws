package com.botregistry.core

sealed trait BuildState

object BuildState {
  final case class None() extends BuildState
  final case class Building(step: Int, total: Int, name: String)
      extends BuildState
  final case class Completed(history: BuildHistory) extends BuildState

  def get(historyStore: Storage[Int, BuildHistory], repo: Repo): BuildState = {
    val res = BuildPool.building.find(_._1 == repo.id)
    if (res.isDefined) {
      val build = res.get._2
      Building(build.step, build.tasks.size, build.tasks(build.step).name)
    } else {
      historyStore.getAll
        .filter(_.repoId == repo.id)
        .reduceOption({ (a, b) =>
          if (a.time > b.time) a else b
        })
        .flatMap({ x =>
          Some(Completed(x))
        })
        .getOrElse(None())
    }
  }
}
