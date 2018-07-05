package com.botregistry.core

import com.github.nscala_time.time.Imports._

case class RepoState(id: Int,
                     buildState: BuildState,
                     runState: RunState,
                     latestBuildTime: DateTime,
                     subsequentBuildCount: Int)
