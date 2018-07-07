package com.botregistry.core

import com.botregistry.util.KubeUtil

case class ConfigMap(name: String, kubeName: String, key: String)

object ConfigMap {
  def getContent(namespace: String, configMap: ConfigMap): Option[String] = {
    KubeUtil.getConfigMap(namespace, configMap.kubeName, configMap.key) match {
      case Right(x) => Some(x)
      case Left(e)  => println(e); None
    }
  }

  def setContent(namespace: String,
                 configMap: ConfigMap,
                 content: String): Option[Unit] = {
    KubeUtil.setConfigMap(namespace, configMap.kubeName, configMap.key, content) match {
      case Right(x) => Some()
      case Left(e)  => println(e); None
    }
  }
}
