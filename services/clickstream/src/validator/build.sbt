lazy val root = project
  .in(file("."))
  .settings(
    name := "validator",
    version := {
      val source = scala.io.Source.fromFile("version")
      try {
        source.mkString
      } finally {
        source.close()
      }
    },
    scalaVersion := "2.12.19",
    libraryDependencies ++= Seq(
         "org.apache.spark" %% "spark-sql"  % "3.5.1",
         "org.apache.spark" %% "spark-core"  % "3.5.1",
    )
)
