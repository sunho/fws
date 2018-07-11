package com.botregistry.service

import io.finch._
import io.finch.circe._
import io.finch.syntax._
import io.circe.generic.auto._
import com.botregistry.core._

trait RepoService extends StorageService with AuthService {
  protected def repoEndpoint = "repos"
  protected def repoPath: Endpoint[Repo] =
    path[Int]
      .mapOutput { i =>
        repoStore.get(i) match {
          case Some(x) => Ok(x)
          case None    => throw new IllegalArgumentException
        }
      }
      .handle {
        case e: Exception => NotFound(e)
      }

  val getRepos: Endpoint[List[Repo]] = get(authenticate :: repoEndpoint) {
    u: User =>
      if (u.isAdmin) {
        Ok(repoStore.getAll)
      } else {
        val repos: List[Repo] = u.repos.flatMap(repoStore.get).toList
        Ok(repos)
      }
  }

  val getRepo: Endpoint[Repo] = get(authenticate :: repoEndpoint :: repoPath) {
    (u: User, repo: Repo) =>
      if (u.isAdmin || u.repos.contains(repo.id)) {
        Ok(repo)
      } else {
        Forbidden(new IllegalAccessException)
      }
  }

  val putRepo: Endpoint[Unit] =
    put(admin :: "repos" :: path[Int] :: jsonBody[Repo]) {
      (id: Int, repo: Repo) =>
        repoStore.addOrUpdate(repo.copy(id = id)) match {
          case Some(_) => Ok()
          case None    => InternalServerError(new IllegalArgumentException)
        }
    }

  val deleteRepo: Endpoint[Unit] =
    delete(admin :: repoEndpoint :: repoPath) { repo: Repo =>
      repoStore.delete(repo.id) match {
        case Some(_) => Ok()
        case None    => BadRequest(new IllegalArgumentException)
      }
    }

  val repoApi = getRepos :+: getRepo :+: putRepo :+: deleteRepo
}
