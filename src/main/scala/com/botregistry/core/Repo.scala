package com.botregistry.core

case class Repo(id: Int,
                name: String,
                repoURL: String,
                dockerFile: String,
                kubeName: String,
                configMaps: List[ConfigMap],
) extends StorageItem[Int] {
  val key = id
}
