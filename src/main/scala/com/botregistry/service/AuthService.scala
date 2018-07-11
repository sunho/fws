package com.botregistry.service

import io.finch._
import com.botregistry.core._
import shapeless.HNil

trait AuthService extends StorageService with ConfigService {
  def getUser(token: String): User = {
    if (token == config.adminToken)
      return User("", true, Set[Int]())
    val tok = tokenStore.get(token) match {
      case Some(x) => x
      case None    => throw new IllegalArgumentException("invalid token")
    }
    userStore.get(tok.name) match {
      case Some(x) => x
      case None    => throw new IllegalStateException("internal error")
    }
  }

  protected def authenticate: Endpoint[User] =
    header("Authorization")
      .mapOutput { t =>
        Ok(getUser(t))
      }
      .handle {
        case e: Exception => Unauthorized(e)
      }

  protected def admin: Endpoint[HNil] =
    authenticate
      .mapOutput { u =>
        if (u.isAdmin)
          Ok(null)
        else
          throw new IllegalAccessException
      }
      .handle {
        case e: Exception => Forbidden(e)
      }
}
