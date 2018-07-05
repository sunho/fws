package com.botregistry.service
import com.botregistry.core.{Repo, User}
import io.finch._
import io.finch.circe._
import io.finch.syntax._
import io.circe.syntax._
import io.circe.generic.auto._
import com.twitter.finagle.Http

trait RepoService extends StorageService with AuthService {
  val getRepos: Endpoint[List[Repo]] = get("repos" :: authenticate) { u: User =>
    if (u.isAdmin) {
      Ok(repoStore.getAll)
    } else {
      val repos: List[Repo] = u.repos
        .map(repoStore.get(_))
        .filter(_.isDefined)
        .map({ case Some(x) => x })
      Ok(repos)
    }
  }

  private def repoBody: Endpoint[Repo] = jsonBody[Repo]

  val addRepos: Endpoint[Int] = post("repos" :: authenticate :: repoBody) {
    (u: User, r: Repo) =>
      if (!u.isAdmin) {
        Unauthorized(new IllegalAccessException)
      } else {
        val r2 = r.copy(id = repoStore.getLastKey match {
          case Some(x) => x + 1
          case None => 0
        })
        repoStore.addOrUpdate(r2)
        Created(r2.id)
      }
  }
}
