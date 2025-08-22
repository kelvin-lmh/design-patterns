# Singleton pattern

## Intent

    Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.

## Problem

    The Singleton pattern solves two problems at the same time:
    - Ensure that a class has just a single instance:
        The purpose is to control access to some shared resource—for example, a database or a file.
    - Provide a global access point to that instance:
        Just like a global variable, the Singleton pattern lets you access some object from anywhere in the program. However, it also protects that instance from being overwritten by other code.

## Solution

    All implementations of the Singleton in Go have two common steps:

    - Keep the single instance in a private (unexported) package-level variable, so outside code cannot directly control it.
    - Provide a public function (like GetInstance) that creates the instance once (using sync.Once) and returns it. Every call to this function always returns the same instance.

## Real-World Analogy

    The government is an excellent example of the Singleton pattern. A country can have only one official government

## Applicability

    - Use when your program needs only one shared instance of a class (e.g., DB connection pool, logger, config).
    - The instance is stored in a private package-level variable, and accessed through a public function (GetInstance).
    - Use when need provides better control than global variables, since the instance cannot be replaced from outside the package.
    - Can be adjusted later: just change GetInstance() if you want to allow multiple instances.
    - Ensures thread-safety and prevents accidental creation of multiple objects.

## Pros

    - The singleton object is initialized only when it’s requested for the first time.
    - You can be sure that a class has only a single instance.
    - Control the number of instances, reduce memory costs & avoid unnecessary object creation.
    - Unified data state, since everyone shares the same object.
    - You gain a global access point to that instance.

## Cons

    - Violates the Single Responsibility Principle. The pattern solves two problems at the time.
    - Not safe in multithreaded environment if not implemented thread-safe. Multiple threads can create a singleton object multiple times.
    - Class/service directly depends on Singleton, making it difficult to write unit test because you always have to use real instance (cannot mock or fake dependency to control data) `Solution interface + dependency injection`

## Relations with Other Patterns
