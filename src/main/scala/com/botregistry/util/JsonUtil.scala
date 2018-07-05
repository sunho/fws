package com.botregistry.util

import io.circe.{ACursor, Decoder}

object JsonUtil {
  def parse[T](cursor: ACursor, fields: String*)(
      implicit decoder: Decoder[T]) = {
    var cur = cursor
    for (s <- fields) {
      cur = cur.downField(s)
    }
    cur.as[T] match {
      case Left(error) => throw error
      case Right(x)    => x
    }
  }
}
