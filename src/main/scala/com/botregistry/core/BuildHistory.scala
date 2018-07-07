package com.botregistry.core

case class BuildHistory(id: Int,
                        repoId: Int,
                        success: Boolean,
                        time: Int,
                        logs: List[(String, String)])
    extends StorageItem[Int] {
  val key = id
  override def toString: String = {
    s"Success: $success\n ${logs
      .map { x =>
        val (key, value) = x
        "#" * 20 + key + "#" * 20 + "\n" + value
      }
      .mkString("\n")}"
  }
}
