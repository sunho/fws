package com.botregistry.core

case class BuildHistory(id: Int,
                        repoId: Int,
                        success: Boolean,
                        time: Int,
                        logs: Map[String, String])
    extends StorageItem[Int] {
  val key = id
}
