package com.botregistry.service

import io.finch._
import io.finch.circe._
import io.finch.syntax._
import io.circe.generic.auto._
import com.botregistry.core._

trait HistoryService extends RepoService {
  protected def historyEndpoint = repoEndpoint :: repoPath :: "histories"

  val getRepoHistories: Endpoint[List[BuildHistory]] =
    get(historyEndpoint) { repo: Repo =>
      Ok(historyStore.getAll.filter(_.repoId == repo.id))
    }

  val getRepoHistory: Endpoint[String] =
    get(historyEndpoint :: path[Int]) { (repo: Repo, id: Int) =>
      historyStore.get(id) match {
        case Some(x) =>
          Ok(x.toString).withHeader("Content-Type", "text/plain;charset=utf-8")
        case None => NotFound(new IllegalArgumentException)
      }
    }

  val historyApi = getRepoHistories
}
