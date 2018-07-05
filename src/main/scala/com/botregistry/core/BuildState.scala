package com.botregistry.core

abstract class BuildState

object BuildState {
  case class Rejected() extends BuildState
  case class Pending() extends BuildState
  case class Building(step: Int, name: String) extends BuildState
  case class Completed(history: BuildHistory) extends BuildState
}
