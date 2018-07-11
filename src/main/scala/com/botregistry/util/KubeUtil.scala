package com.botregistry.util

import skuber._
import skuber.json.format._
import skuber.apps.v1.Deployment

import akka.actor.ActorSystem
import akka.stream.ActorMaterializer

import scala.concurrent.Await
import scala.concurrent.duration.Duration
import scala.util.{Failure, Success}

object KubeUtil {
  implicit val system = ActorSystem()
  implicit val materializer = ActorMaterializer()
  implicit val dispatcher = system.dispatcher
  val k8s = k8sInit
  def updateImage(namespace: String,
                  deployment: String,
                  image: String): Either[String, Unit] = {
    val fut = k8s.getInNamespace[Deployment](deployment, namespace)
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
      case Success(_) => Right()
      case Failure(e) => Left(e.getMessage)
    }
  }

  def restartDeployment(namespace: String,
                  deployment: String): Either[String, Unit] = {
    val fut = k8s.getInNamespace[Deployment](deployment, namespace)
    val fut2 = fut.map { x =>
      val container =
        (for (y <- x.getPodSpec;
              z <- y.containers.headOption)
          yield z) match {
          case Some(x) => x
          case None => throw new IllegalStateException
        }
      val newEnv = container.env.filter(_.name != "TIME") :+ EnvVar("TIME",TimeUtil.timestamp.toString)
      val updated = x.updateContainer(container.copy(env = newEnv))
      val fut3 = k8s update updated
      Await.ready(fut3, Duration.Inf).value.get
    }
    Await.ready(fut2, Duration.Inf).value.get match {
      case Success(_) => Right()
      case Failure(e) => Left(e.getMessage)
    }
  }

  def getConfigMap(namespace: String,
                   configMap: String,
                   key: String): Either[String, String] = {
    val fut = k8s.getInNamespace[ConfigMap](configMap, namespace)
    val fut2 = fut.map { x =>
      x.data.get(key) match {
        case Some(x) => x
        case None    => throw new IllegalStateException
      }
    }
    Await.ready(fut2, Duration.Inf).value.get match {
      case Success(x) => Right(x)
      case Failure(e) => Left(e.getMessage)
    }
  }

  def setConfigMap(namespace: String,
                   configMap: String,
                   key: String,
                   content: String): Either[String, Unit] = {
    val fut = k8s.getInNamespace[ConfigMap](configMap, namespace)
    val fut2 = fut.map { x =>
      val data = x.data + (key -> content)
      val updated = x.copy(data = data)
      val fut3 = k8s update updated
      Await.ready(fut3, Duration.Inf).value.get
    }
    Await.ready(fut2, Duration.Inf).value.get match {
      case Success(_) => Right()
      case Failure(e) => Left(e.getMessage)
    }
  }

  def getDeploymentStatus(
      namespace: String,
      deployment: String): Either[String, Deployment.Status] = {
    val fut = k8s.getInNamespace[Deployment](deployment, namespace)
    val fut2 = fut.map { x =>
      x.status.get
    }
    Await.ready(fut2, Duration.Inf).value.get match {
      case Success(x) => Right(x)
      case Failure(e) => Left(e.getMessage)
    }
  }
}
