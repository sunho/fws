import java.io.PrintWriter

import com.twitter.finagle.param.Stats
import com.twitter.server.TwitterServer
import com.twitter.finagle.Http
import java.nio.file.{Files, Paths}

import com.twitter.util.Await
import com.botregistry.service._
import skuber._
import skuber.json.format._
import skuber.apps.v1.Deployment
import LabelSelector.dsl._
import akka.actor.ActorSystem
import akka.stream.ActorMaterializer

import scala.util.{Failure, Success}

object Main extends TwitterServer {
  def create(path: String, content: String): Unit = {
    if (!Files.exists(Paths.get(path))) {
      new PrintWriter(path) {
        write(content); close
      }
    }
  }

  def main(): Unit = {
    implicit val system = ActorSystem()
    implicit val materializer = ActorMaterializer()
    implicit val dispatcher = system.dispatcher
    val k8s = k8sInit
    k8s.getInNamespace[Deployment]("bot-registry", "bot") onComplete {
      case Success(dep) => println(dep.status.get.replicas)
      case Failure(e)   => throw e
    }

    create("History.json", "[]")
    create("Repo.json", "[]")
    create("User.json", "[]")
    create("Token.json", "[]")
    val api = new StandardService(".")
    println(api.config.adminToken)
    val server = Http.server
      .configured(Stats(statsReceiver))
      .serve(":8080", api.toService)

    val t = new java.util.Timer()
    val task = new java.util.TimerTask {
      def run() = {
        println("saving")
        api.save()
      }
    }
    t.schedule(task, 10000L, 10000L)

    onExit {
      server.close()
    }
    Await.ready(adminHttpServer)
  }
}
