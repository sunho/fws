package com.botregistry.util

import java.util.TimeZone
import java.util.stream.Collectors

import akka.NotUsed
import skuber._
import skuber.json.format._
import skuber.apps.v1.Deployment
import akka.actor.ActorSystem
import akka.stream.ActorMaterializer
import akka.stream.scaladsl.{Flow, Framing, Keep, Sink}
import akka.util.ByteString
import org.joda.time.{DateTime, DateTimeZone}

import scala.collection.JavaConverters._
import scala.concurrent.{Await, Future}
import scala.concurrent.duration._
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
      Await.ready(fut3, 20 seconds).value.get
    }
    Await.ready(fut2, 20 seconds).value.get match {
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
      Await.ready(fut3, 20 seconds).value.get
    }
    Await.ready(fut2, 20 seconds).value.get match {
      case Success(_) => Right()
      case Failure(e) => Left(e.getMessage)
    }
  }

  def UTC2KST(s: String): String = {
    val list = s.split(" ")
    val tz = DateTimeZone.forOffsetHours(9)

    try {
      val time = new DateTime(list(0), tz)
      time.toString + " " + list.drop(1).mkString(" ")
    } catch {
      case e: Exception => s
    }
  }

  def getDeploymentLog(namespace: String,
                       kubeName: String): Either[String, String]= {
    val logFlow = Flow[ByteString]
      .via(Framing.delimiter(
        ByteString("\n"),
        maximumFrameLength = 256,
        allowTruncation = true))
      .map(_.utf8String)

    val fut = k8s.listInNamespace[PodList](namespace)
    val fut2 = fut.map { x =>
      val pods = x.toList.filter {
        y => y.metadata.labels.exists(_ == "app" -> kubeName)
      }
      pods.map { p =>
        val fut3  = if (p.status.get.phase.get == Pod.Phase.Running) {
          k8s.getPodLogSource(p.name, Pod.LogQueryParams(timestamps = Some(true), tailLines = Some(100)))
        } else {
          k8s.getPodLogSource(p.name, Pod.LogQueryParams(timestamps = Some(true), tailLines = Some(100), previous = Some(true)))
        }
        val src = Await.ready(fut3, 5 seconds).value.get match {
          case Success(x) => x
          case Failure(e) => throw e
        }

        val time = p.metadata.creationTimestamp.get.toEpochSecond
        var buf = new StringBuilder(s"Logs from ${p.name} $time\n")
        val fut4 = src.via(logFlow).runForeach(buf ++= UTC2KST(_) + "\n")
        Await.result(fut4, 10 seconds)
        (time, buf.toString)
      }
    }
    Await.ready(fut2, 20 seconds).value.get match {
      case Success(x) => Right(x.sortBy(_._1).map(_._2).mkString("\n"))
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
    Await.ready(fut2, 20 seconds).value.get match {
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
      Await.ready(fut3, 20 seconds).value.get
    }
    Await.ready(fut2, 20 seconds).value.get match {
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
    Await.ready(fut2, 20 seconds).value.get match {
      case Success(x) => Right(x)
      case Failure(e) => Left(e.getMessage)
    }
  }
}
