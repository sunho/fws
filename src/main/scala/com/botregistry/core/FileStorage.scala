package com.botregistry.core

import java.io.PrintWriter
import io.circe.{Decoder, Encoder, Json}
import scala.io.Source
import scala.collection.concurrent.TrieMap
import io.circe.syntax._
import io.circe.parser._

class FileStorage[K <% Ordered[K], T <: StorageItem[K]] extends Storage[K, T] {
  val items = new TrieMap[K, T]
  val path = ""

  override def addOrUpdate(item: T): Option[Unit] = {
    items += item.key -> item; Some()
  }

  override def get(key: K): Option[T] = {
    items.get(key)
  }

  override def getLastKey: Option[K] = {
    items.keys reduceOption ((a, b) => if (a > b) a else b)
  }

  override def getAll: List[T] = {
    items.values.toList
  }

  override def delete(key: K): Option[Unit] = {
    items.get(key) match {
      case Some(_) => items -= key; Some()
      case None    => None
    }
  }

  def save()(implicit encoder: Encoder[T]) {
    new PrintWriter(path) {
      write(items.values.map(_.asJson).asJson.toString); close
    }
  }

}

object FileStorage {
  def apply[K <% Ordered[K], T <: StorageItem[K]](path: String)(
      implicit decoder: Decoder[T]): FileStorage[K, T] = {
    val path2 = path
    new FileStorage[K, T] {
      override val path = path2
      val raw = decode[List[Json]]({
        val src = Source.fromFile(s"$path")
        try src.mkString
        finally src.close
      }) match {
        case Right(decoded) => decoded
        case Left(_)        => throw new IllegalArgumentException
      }
      override val items = TrieMap[K, T](raw map (_.as[T]) map {
        case Right(decoded) => decoded
        case Left(_)        => throw new IllegalArgumentException
      } map { item: T =>
        (item.key, item)
      }: _*)
    }
  }
}
