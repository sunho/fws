package com.botregistry.core

import com.botregistry.util.TimeUtil

object BuildTaskFactory {
  private def makeTag(settings: BuildSettings,
                      repo: Repo,
                      time: Int): String = {
    s"${settings.dockerRegistry}/${repo.name}:$time"
  }
  private def makePath(settings: BuildSettings, repo: Repo): String = {
    s"${settings.basePath}/${repo.name}"
  }

  def apply(settings: BuildSettings, repo: Repo): List[BuildTask] = {
    val tag = makeTag(settings, repo, TimeUtil.timestamp)
    val path = makePath(settings, repo)
    CleanTask(path) ::
      CheckoutTask(path, repo) ::
      CleanTask(path + "/Dockerfile") ::
      WriteTask(path + "/Dockerfile", repo.dockerFile) ::
      DockerTask.build(path, tag) ::
      DockerTask.push(tag) ::
      Nil

  }
}
