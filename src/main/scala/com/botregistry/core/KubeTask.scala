package com.botregistry.core

import skuber._
import skuber.apps.v1.Deployment
import akka.actor.ActorSystem
import akka.stream.ActorMaterializer
import scala.concurrent.duration.Duration
import scala.concurrent.Await
import scala.util.{Failure, Success}

class KubeTask(namespace: String, appName: String, image: String)
    extends BuildTask {
  implicit val system = ActorSystem()
  implicit val materializer = ActorMaterializer()
  implicit val dispatcher = system.dispatcher

  def name: String = "Deploy to Kubernetes"

  //skjdfdsfkd;klfdsjfks;fka
  def run(): (Boolean, String) = {
    val k8s = k8sInit
    val fut = k8s.getInNamespace[Deployment](appName, namespace)
    val fut2 = fut.map { x =>
      val container =
        (for (y <- x.getPodSpec;
              z <- y.containers.headOption)
          yield z) match {
          case Some(x) => x
          case None    => throw new IllegalStateException
        }
      val updated = x.updateContainer(container.copy(image = image))
      val fut3 = k8s update updated
      Await.ready(fut3, Duration.Inf).value.get
    }
    Await.ready(fut2, Duration.Inf).value.get match {
      case Success(_) => (true, "Success")
      case Failure(e) => (false, e.getMessage)
    }
  }
}

object KubeTask {
  def apply(namespace: String, appName: String, image: String): KubeTask = {
    new KubeTask(namespace, appName, image)
  }
}
