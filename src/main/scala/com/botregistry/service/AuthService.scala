package com.botregistry.service
import io.finch._
import com.botregistry.core.{Repo, Storage, Token, User}

trait AuthService extends StorageService with ConfigService {
  def getUser(token: String): User = {
    if (token == config.adminToken)
      return User(0, true, List[Int]())
    val tok = tokenStore.get(token) match {
      case Some(x) => x
      case None    => throw new IllegalArgumentException
    }
    userStore.get(tok.id) match {
      case Some(x) => x
      case None    => throw new NoSuchElementException
    }
  }

  val authenticate: Endpoint[User] = header("Authorization")
    .mapOutput { t =>
      Ok(getUser(t))
    }
    .handle {
      case e: Exception => Unauthorized(e)
    }
}
