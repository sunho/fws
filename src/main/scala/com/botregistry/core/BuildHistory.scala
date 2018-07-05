package com.botregistry.core

import com.github.nscala_time.time.Imports._

case class BuildHistory(id: Int,
                        repoId: Int,
                        time: DateTime,
                        logs: Map[String, String])
    extends StorageItem[Int] {
  val key = id
}
