package com.botregistry.tasks

import java.io.PrintWriter

class WriteTask(path: String, content: String) extends BuildTask {
  override def name: String = s"Write $path"
  override def run(): (Boolean, String) = {
    val writer = new PrintWriter(path)
    try writer.write(content)
    catch {
      case e: Exception => return (false, e.getMessage)
    } finally {
      writer.close()
    }
    (true, s"wrote $path")
  }
}

object WriteTask {
  def apply(path: String, content: String): WriteTask = {
    new WriteTask(path, content)
  }
}
