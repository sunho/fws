package com.botregistry.tasks

import com.botregistry.util.CommandExec

class CommandTask(cmd: String) extends BuildTask {
  override def name = cmd
  override def run(): (Boolean, String) = {
    val res = CommandExec(cmd)
    (res._1 == 0, res._2)
  }
}

object CommandTask {
  def apply(cmd: String): CommandTask = {
    new CommandTask(cmd)
  }
}
