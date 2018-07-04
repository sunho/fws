package com.botregistry.core
import java.io.{IOException, PrintWriter}
import scala.reflect.runtime.universe.{TypeTag, typeOf}
import scala.io.Source
import io.circe.syntax._
import io.circe.parser.decode

class FileStore[K, T <: StorageItem[K]] extends Storage[K, T] {
  var items = new collection.mutable.HashMap[K, T]

  override def addOrUpdate(item: T): Option[Unit] = {
    items += item.key -> item; Some()
  }

  override def get(key: K): Option[T] = {
    items.get(key)
  }

  override def delete(key: K): Option[Unit] = {
    items.get(key) match {
      case Some(_) => items -= key; Some()
      case None => None
    }
  }

  def filename()(implicit tag: TypeTag[T]): String = s"${tag.tpe.toString}.json"

  def save(){
    new PrintWriter(filename){ write(items.asJson.toString); close }
  }

  def load(path: String): Unit = {
    val raw = {
      val src = Source.fromFile(filename)
      try src.mkString finally src.close
    }
    items = decode[collection.mutable.HashMap[K,T]](raw) match {
      case Left(error) => throw new IOException(error)
      case Right(decoded) => decoded
    }
  }
}
