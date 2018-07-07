package com.botregistry.tasks

trait BuildTask {
  def name: String
  def run(): (Boolean, String)
}
