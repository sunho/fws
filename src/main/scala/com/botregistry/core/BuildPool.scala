package com.botregistry.core

//should replace
import scala.concurrent.ExecutionContext.Implicits.global
import collection.concurrent.TrieMap
import scala.concurrent.Future

object BuildPool {
  val building = new TrieMap[Int, Build]

  // hjv;afsdilfksnam,l,dlkasdhfkjsfefb
  def run(build: Build): Future[(Boolean, Map[String, String])] = {
    building += build.repo.id -> build
    Future {
      @scala.annotation.tailrec
      def runTasks(step: Int = 0,
                   l: List[(String, String)] = List[(String, String)]())
        : List[(String, String)] = {
        building(build.repo.id) = build.copy(step = step)
        val task = build.tasks(step)
        val res = task.run()
        val tup = task.name -> res._2
        if (step + 1 == build.tasks.size || !res._1) {
          tup :: l reverse
        } else {
          runTasks(step + 1, tup :: l)
        }
      }
      val logs = runTasks()
      building -= build.repo.id
      (logs.size == build.tasks.size, logs.toMap)
    }
  }
}
