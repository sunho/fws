package com.botregistry.core

abstract class RunState

object RunState {
  case class Stopped() extends RunState
  case class CrashLoopBack(times: Int) extends RunState
  case class Running() extends RunState
}
