package com.botregistry.service

import io.finch._
import io.finch.circe._
import io.finch.syntax._
import io.circe.generic.auto._
import com.botregistry.core._

trait UserRepoService extends RepoService with UserService {
  protected def userRepoEndpoint = oneUserEndpoint :: "repos"

  val addUserRepo: Endpoint[Unit] = post(userRepoEndpoint :: repoPath) {
    (u: User, repo: Repo) =>
      userStore.addOrUpdate(u.copy(repos = u.repos + repo.id)) match {
        case Some(_) => Ok()
        case None    => InternalServerError(new IllegalStateException)
      }
  }

  val getUserRepos: Endpoint[List[Repo]] = get(userRepoEndpoint) { u: User =>
    Ok(u.repos.flatMap(repoStore.get).toList)
  }

  val deleteUserRepo: Endpoint[Unit] = delete(userRepoEndpoint :: repoPath) {
    (u: User, repo: Repo) =>
      userStore.addOrUpdate(u.copy(repos = u.repos - repo.id)) match {
        case Some(_) => Ok()
        case None    => InternalServerError(new IllegalStateException)
      }
  }

  val userRepoApi = addUserRepo :+: getUserRepos :+: deleteUserRepo
}
