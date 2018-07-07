package com.botregistry.core

import com.botregistry.util.KubeUtil

case class RunState(available: Int = 0, unavailable: Int = 0)

object RunState {
  def get(namespace: String, deployment: String): Option[RunState] = {
    KubeUtil.getDeploymentStatus(namespace, deployment) match {
      case Left(e) => println(e); None
      case Right(x) =>
        Some(RunState(x.availableReplicas, x.unavailableReplicas))
    }
  }
}
