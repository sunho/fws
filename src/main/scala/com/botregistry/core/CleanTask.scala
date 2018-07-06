package com.botregistry.core

import java.io.File
import java.nio.file.{Files, Paths}

import com.botregistry.util.FileUtil

class CleanTask(path: String) extends BuildTask {
  override def name: String = s"Clean $path"
  override def run(): (Boolean, String) = {
    if (Files.exists(Paths.get(path))) {
      try FileUtil.deleteRecursively(new File(path))
      catch {
        case e: Exception => return (false, e.getMessage)
      }
      (true, s"deleted $path")
    } else {
      (true, "")
    }
  }
}

object CleanTask {
  def apply(path: String): CleanTask = {
    new CleanTask(path)
  }
}
