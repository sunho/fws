package com.botregistry.service

import io.finch._
import io.finch.circe._
import io.finch.syntax._
import io.circe.generic.auto._
import com.botregistry.core._

trait UserService extends StorageService with AuthService {
  protected def userEndpoint = admin :: "users"
  protected def oneUserEndpoint = userEndpoint :: userPath
  protected def userPath: Endpoint[User] =
    path[String]
      .mapOutput { n =>
        userStore.get(n) match {
          case Some(x) => Ok(x)
          case None    => throw new ArithmeticException
        }
      }
      .handle {
        case e: Exception => NotFound(e)
      }

  val getUsers: Endpoint[List[User]] = get(userEndpoint) {
    Ok(userStore.getAll)
  }

  val putUser: Endpoint[Unit] = put(userEndpoint :: jsonBody[User]) { u: User =>
    userStore.addOrUpdate(u) match {
      case Some(_) => Ok()
      case None    => InternalServerError(new ArithmeticException())
    }
  }

  val deleteUser: Endpoint[Unit] = delete(oneUserEndpoint) { u: User =>
    userStore.delete(u.name) match {
      case Some(_) => Ok()
      case None    => InternalServerError(new ArithmeticException)
    }
  }

  val userApi = getUsers :+: putUser :+: deleteUser
}
