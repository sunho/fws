package com.botregistry.service

import java.io.PrintWriter
import java.nio.file.{Files, Paths}

import com.botregistry.core.{Config, DockerTask}
import com.botregistry.util.CommandExec

object Bootstrap {
  def create(path: String, content: String): Unit = {
    if (!Files.exists(Paths.get(path))) {
      new PrintWriter(path) {
        write(content); close
      }
    }
  }

  def storageSetup(config: Config): Unit = {
    create("History.json", "[]")
    create("Repo.json", "[]")
    create("User.json", "[]")
    create("Token.json", "[]")
  }

  def dockerSetup(config: Config): Unit = {
    val res = CommandExec(DockerTask.loginCmd(config))
    println(res._2)
    if (res._1 != 0) {
      throw new IllegalStateException
    }
  }

  def setup(config: Config): Unit = {
    storageSetup(config)
    dockerSetup(config)
  }
}
