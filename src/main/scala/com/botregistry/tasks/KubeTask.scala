package com.botregistry.tasks

import com.botregistry.util.KubeUtil

class KubeTask(namespace: String, appName: String, image: String)
    extends BuildTask {

  def name: String = "Deploy to Kubernetes"

  def run(): (Boolean, String) = {
    KubeUtil.updateImage(namespace, appName, image) match {
      case Left(err) => (false, err)
      case Right(_)  => (true, "Success")
    }
  }
}

object KubeTask {
  def apply(namespace: String, appName: String, image: String): KubeTask = {
    new KubeTask(namespace, appName, image)
  }
}
