package com.botregistry.util

import scala.sys.process._

object CommandExec {
  def apply(cmd: String): (Int, String) = {
    val std = new StringBuilder
    val logger = ProcessLogger(std append _ + "\n", std append _ + "\n")
    val status = cmd ! logger
    (status, std.toString)
  }
}
