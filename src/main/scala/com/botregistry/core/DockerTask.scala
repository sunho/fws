package com.botregistry.core

object DockerTask {
  def build(path: String, tag: String): CommandTask = {
    CommandTask(s"docker build -t $tag $path")
  }

  def push(tag: String): CommandTask = {
    CommandTask(s"docker push $tag")
  }
}
