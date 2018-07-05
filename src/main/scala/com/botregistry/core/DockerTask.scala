package com.botregistry.core

class DockerTask(settings: DockerSettings, dockerFile: String)
    extends BuildTask {
  def name = "Docker"
  def run(): (Boolean, String) = {
    (true, "Asdf")
  }
}
