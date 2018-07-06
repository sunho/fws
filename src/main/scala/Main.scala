import java.io.PrintWriter

import com.twitter.finagle.param.Stats
import com.twitter.server.TwitterServer
import com.twitter.finagle.Http

import com.botregistry.core._
import com.twitter.util.Await
import com.botregistry.service._

object Main extends TwitterServer {
  def main(): Unit = {
    val config = Config.fromFile("Config.json")
    Bootstrap.setup(config)
    val api = new StandardService(config)
    api.startSaving()
    val server = Http.server
      .configured(Stats(statsReceiver))
      .serve(":8080", api.toService)

    onExit {
      server.close()
    }
    Await.ready(adminHttpServer)
  }
}
