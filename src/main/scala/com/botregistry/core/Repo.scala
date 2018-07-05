package com.botregistry.core
import io.circe.generic.auto._

case class Repo(id: Int,
                repoURL: String,
                dockerFile: String,
                deployment: String)
    extends StorageItem[Int] {
  val key = id
}
