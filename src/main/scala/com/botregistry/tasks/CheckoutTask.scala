package com.botregistry.tasks

import com.botregistry.core.Repo

object CheckoutTask {
  def apply(path: String, repo: Repo): CommandTask = {
    CommandTask(s"git clone ${repo.repoURL} $path")
  }
}
