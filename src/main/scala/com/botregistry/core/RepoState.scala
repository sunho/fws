package com.botregistry.core

case class RepoState(buildState: BuildState = BuildState.None(),
                     runState: RunState = RunState())
