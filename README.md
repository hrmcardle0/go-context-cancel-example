# Go Context Cancellation Example  

In Go, when you cancel a context, the cancellation signal goes down the tree. Any contexts you create off of a parent context will become children of that parent context, and any cancellation signals
done to any context will only propegate downard words parents.

## How it works

This project is a simple example of how it works. We create a master, parent, and child context. 

- The master context creates the parent, which creates the child context. 
- We call the cancel() function for the parent context after a few seconds
- The print statements show that the parent Done() channel's value is properly received, as well as the child's Done() channel.
- The master context Done() channel never contains a value, as cancellation signals are always propegated upwards

## Why is this useful?

The concept of context cancellation signal is a bit hard to grasp, so actually implementing code which shows timestamp'd log entries helps to visualize how everything works together. Contexts
like this are useful when you need to perform some type of timeout that cancels multiple requests down the line, for example a microsrevice application that has multiple requests going through various services