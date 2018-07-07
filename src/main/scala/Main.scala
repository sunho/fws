import com.twitter.server.TwitterServer
import com.twitter.finagle.Http
import com.botregistry.core._
import com.botregistry.service._
import com.botregistry.util.Bootstrap
import com.twitter.util.Await

object Main extends TwitterServer {

  def main(): Unit = {
    val config = Config.fromFile("Config.json")
    Bootstrap.setup(config)

    val service = new StandardService(config)
    service.startSaving()

    val server = Http.server.serve(":8080", service.toService)

    onExit {
      server.close()
    }

    Await.ready(server)
  }
}
