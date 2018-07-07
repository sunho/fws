package com.botregistry.service

import io.finch._
import io.finch.circe._
import io.circe.syntax._
import io.finch.syntax._
import io.circe.generic.auto._
import com.botregistry.core._

trait StateService extends RepoService {
  val getRepoState: Endpoint[RepoState] =
    get(repoEndpoint :: repoPath :: "state") { repo: Repo =>
      Ok(
        RepoState(BuildState.get(historyStore, repo),
                  RunState
                    .get(config.kubeNamespace, repo.kubeName)
                    .getOrElse(RunState())))
    }

  val stateApi = getRepoState
}
