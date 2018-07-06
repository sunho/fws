package com.botregistry.util

object TimeUtil {
  def timestamp: Int = (System.currentTimeMillis / 1000).toInt
}
