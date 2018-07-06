name := "bot-registry"

val circeVersion = "0.9.0"
val finchVersion = "0.21.0"

libraryDependencies ++= Seq(
  "org.scalatest" %% "scalatest" % "3.0.5" % "test",
  "com.github.finagle" %% "finch-core" % finchVersion,
  "com.github.finagle" %% "finch-circe" % finchVersion,
  "io.circe" %% "circe-core" % circeVersion,
  "io.circe" %% "circe-generic" % circeVersion,
  "io.circe" %% "circe-parser" % circeVersion,
  "io.circe" %% "circe-generic-extras" % circeVersion,
  "com.twitter" %% "twitter-server" % "18.6.0",
  "com.typesafe.play" %% "play-json" % "2.6.7",
  "com.github.nscala-time" %% "nscala-time" % "2.20.0",
  "ch.qos.logback" % "logback-classic" % "1.1.3" % Runtime,
  "io.skuber" %% "skuber" % "2.0.7"
)

enablePlugins(DockerPlugin)

assemblyMergeStrategy in assembly := {
  case PathList("META-INF", xs @ _*) => MergeStrategy.discard
  case PathList("reference.conf") => MergeStrategy.concat
  case x => MergeStrategy.first
}

dockerfile in docker := {
  val artifact: File = assembly.value
  val artifactTargetPath = s"/app/${artifact.name}"
  new Dockerfile {
    from("ksunhokim/docker-kube-jre8")
    add(artifact, artifactTargetPath)
    entryPoint("java", "-jar", artifactTargetPath)
  }
}

imageNames in docker := Seq(
  ImageName(s"ksunhokim/${name.value}:latest")
)

buildOptions in docker := BuildOptions(
  cache = false
)

fork in run := true