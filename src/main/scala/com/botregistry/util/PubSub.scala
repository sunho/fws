package com.botregistry.util

trait Subscriber[T] {
  def handler(arg: T)
}

trait Publisher[T]{
  private var subscribers: Set[Subscriber[T]] = Set()
  def subscribe(subscriber: Subscriber[T]): Unit = subscribers += subscriber
  def unsubscribe(subscriber: Subscriber[T]): Unit = subscribers -= subscriber
  def publish(arg: T): Unit = subscribers.foreach(_.handler(arg))
}
