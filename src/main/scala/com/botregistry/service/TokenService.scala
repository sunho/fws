package com.botregistry.service

import io.finch._
import io.finch.circe._
import io.finch.syntax._
import io.circe.generic.auto._
import com.botregistry.core._
import com.botregistry.util.TokenGenerator

trait TokenService extends UserService {
  protected def tokenEndpoint = admin :: "tokens"

  val getTokens: Endpoint[List[Token]] = get(tokenEndpoint) {
    Ok(tokenStore.getAll)
  }

  val recreateUserToken: Endpoint[String] = post(tokenEndpoint :: userPath) {
    u: User =>
      tokenStore.getAll
        .filter(_.name == u.name)
        .map(_.token)
        .foreach(tokenStore.delete)
      val tok = Token(TokenGenerator.generateSHAToken(u.name), u.name)
      tokenStore.addOrUpdate(tok) match {
        case Some(_) => Ok(tok.token)
        case None => InternalServerError(new ArithmeticException)
      }
  }

  val tokenApi = getTokens :+: recreateUserToken
}
