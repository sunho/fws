package com.botregistry.core

object CheckoutTask {
  def apply(path: String, repo: Repo): CommandTask = {
    CommandTask(s"git clone ${repo.repoURL} $path")
  }
}
