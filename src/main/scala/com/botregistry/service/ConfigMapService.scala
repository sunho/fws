package com.botregistry.service

import io.finch.syntax._
import io.finch._
import com.botregistry.core._
import com.botregistry.util.KubeUtil

trait ConfigMapService extends RepoService {
  val getRepoConfig: Endpoint[String] =
    get(authenticate :: repoEndpoint :: repoPath :: "configs" :: path[String]) {
      (u: User, repo: Repo, confName: String) =>
        if (u.isAdmin || u.repos.contains(repo.id)) {
          val res = repo.configMaps.find(_.name == confName) match {
            case Some(x) => x
            case None    => throw new IllegalArgumentException
          }
          ConfigMap.getContent(config.kubeNamespace, res) match {
            case Some(x) => Ok(x)
            case None    => throw new IllegalStateException("config map not found")
          }
        } else {
          throw new IllegalAccessException
        }
    }.handle {
      case e: IllegalArgumentException => NotFound(e)
      case e: IllegalAccessException   => Forbidden(e)
      case e: IllegalStateException    => InternalServerError(e)
    }

  val setRepoConfig: Endpoint[Unit] =
    put(
      authenticate :: repoEndpoint :: repoPath :: "configs" :: path[String] :: stringBody
    ) { (u: User, repo: Repo, confName: String, content: String) =>
      if (u.isAdmin || u.repos.contains(repo.id)) {
        val res = repo.configMaps.find(_.name == confName) match {
          case Some(x) => x
          case None    => throw new IllegalArgumentException
        }
        ConfigMap.setContent(config.kubeNamespace, res, content) match {
          case Some(x) => Ok()
          case None    => throw new IllegalStateException("config map not found")
        }
      } else {
        throw new IllegalAccessException
      }
    }.handle {
      case e: IllegalArgumentException => NotFound(e)
      case e: IllegalAccessException   => Forbidden(e)
      case e: IllegalStateException    => InternalServerError(e)
    }

  val getRepoLog: Endpoint[String] =
    get(authenticate :: repoEndpoint :: repoPath :: "logs") {
      (u: User, repo: Repo) =>
        if (u.isAdmin || u.repos.contains(repo.id)) {
          KubeUtil.getDeploymentLog(config.kubeNamespace, repo.kubeName) match {
            case Right(x) => Ok(x)
            case Left(x) => println(x); throw new IllegalStateException
          }
        } else {
          throw new IllegalAccessException
        }
    }.handle {
      case e: IllegalAccessException   => Forbidden(e)
      case e: IllegalStateException=> InternalServerError(e)
    }

  val configMapApi = getRepoConfig :+: setRepoConfig :+: getRepoLog
}
