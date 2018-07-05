package com.botregistry.core

class KubeTask(settings: KubeSettings, deployment: String) extends BuildTask{
  def name = "Kubernetes"
  def run(): (Boolean, String) = {
    (true, "sadfasdf")
  }
}
