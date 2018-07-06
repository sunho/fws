package com.botregistry.core

case class Repo(id: Int,
                name: String,
                repoURL: String,
                dockerFile: String,
                deployment: String
               )
    extends StorageItem[Int] {
  val key = id
}
