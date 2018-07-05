package com.botregistry.core

case class BuildTaskOptions(dockerSettings: DockerSettings, kubeSettings: KubeSettings)
object BuildTaskFactory {
  def apply(options: BuildTaskOptions, repo: Repo): List[BuildTask] = {
    new DockerTask(options.dockerSettings, repo.dockerFile) :: new KubeTask(options.kubeSettings, repo.deployment) :: Nil
  }
}
