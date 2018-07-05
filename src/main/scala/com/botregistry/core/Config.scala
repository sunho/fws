package com.botregistry.core
import io.circe.generic.auto._
import io.circe.syntax._
import io.circe.parser._
import scala.io.Source

case class Config(adminToken: String)

object Config {
  def fromFile(path: String): Config = {
    decode[Config]({
      val src = Source.fromFile(path)
      try src.mkString
      finally src.close
    }) match {
      case Right(decoded) => decoded
      case Left(_)        => throw new IllegalArgumentException
    }
  }
}
