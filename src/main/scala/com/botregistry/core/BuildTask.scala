package com.botregistry.core

trait BuildTask {
  def name: String
  def run(): (Boolean, String)
}
