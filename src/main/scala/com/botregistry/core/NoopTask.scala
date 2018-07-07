package com.botregistry.core

class NoopTask extends BuildTask {
  override def name: String = "noop"
  override def run(): (Boolean, String) = {
    (true, "")
  }
}

object NoopTask {
  def apply(): NoopTask = {
    new NoopTask
  }
}
