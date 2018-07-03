package com.botregistry.core

case class Token(token: String, id: Int)

trait TokenStore {
  def addToken(token: Token): Option[Token]
  def getToken(token: String): Option[Token]
  def deleteToken(token: Token): Option[Unit]
}


