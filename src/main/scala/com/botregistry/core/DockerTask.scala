package com.botregistry.core

object DockerTask {
  def build(path: String, image: String): CommandTask = {
    CommandTask(s"docker build -t $image $path")
  }

  def push(image: String): CommandTask = {
    CommandTask(s"docker push $image")
  }

  def loginCmd(config: Config): String = {
    s"docker login ${config.dockerHost} -u ${config.dockerUsername} -p ${config.dockerPassword}"
  }
}
