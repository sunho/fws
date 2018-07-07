package com.botregistry.core

import com.botregistry.util.TimeUtil

case class BuildHistory(id: Int,
                        repoId: Int,
                        time: Int,
                        success: Boolean,
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

object BuildHistory {
  def getSubsequntBuilds(historyStore: Storage[Int, BuildHistory],
                         repo: Repo): List[BuildHistory] = {
    val time = TimeUtil.timestamp
    historyStore.getAll.filter { x =>
      x.repoId == repo.id && time - x.time <= 60 * 60
    }
  }
}
