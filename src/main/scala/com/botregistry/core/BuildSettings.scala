package com.botregistry.core

case class BuildSettings(workspacePath: String,
                         dockerRegistry: String,
                         kubeNamespace: String)
