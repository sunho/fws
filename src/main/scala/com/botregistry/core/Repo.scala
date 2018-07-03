package com.botregistry.core

case class Repo(id: Int, repoURL: String, dockerFile: String, deployment: String)

trait RepoStorage {
  def addRepo(repo: Repo): Option[Repo]
  def getRepo(id: Int): Option[Repo]
  def updateRepo(repo: Repo): Option[Repo]
  def deleteRepo(id: int): Option[Unit]
}

