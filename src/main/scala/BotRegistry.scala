import com.twitter.finagle.Http
import com.twitter.util.Await
import io.finch._
import io.finch.syntax._
import io.circe.generic.auto._
import io.finch.circe._

object BotRegistry {
  def main(args: Array[String]): Unit = {
    val time: Endpoint[Unit] =
      get("time") {
        Ok()
      }
    Await.ready(Http.server.serve(":8080", time.toService))
  }
}