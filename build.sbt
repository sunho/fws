name := "BotRegistry"

val circleVersion = "0.9.0"
val finchVersion = "0.21.0"

libraryDependencies ++= Seq(
  "org.scalatest" %% "scalatest" % "3.0.5" % "test",
  "com.github.finagle" %% "finch-core" % finchVersion,
  "com.github.finagle" %% "finch-circe" % finchVersion,
  "io.circe" %% "circe-generic" % circleVersion
)
