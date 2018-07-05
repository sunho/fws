import com.twitter.finagle.param.Stats
import com.twitter.server.TwitterServer
import com.twitter.finagle.{Http, Service}
import com.twitter.finagle.http.{Request, Response}
import com.twitter.util.Await
import com.botregistry.service._

object BotRegistry extends TwitterServer {
  def main(): Unit = {
    val api = new StandardService(".")
    println(api.config.adminToken)
    val server = Http.server
      .configured(Stats(statsReceiver))
      .serve(":8080", api.toService)

    onExit {
      server.close()
    }
    Await.ready(adminHttpServer)
  }
}
