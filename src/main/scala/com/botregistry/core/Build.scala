package com.botregistry.core

case class Build(repo: Repo, step: Int, tasks: List[BuildTask])

object Build {
  def apply(settings: BuildSettings, repo: Repo): Build = {
    val tasks = BuildTaskFactory(settings, repo)
    new Build(repo, 0, tasks)
  }
}
