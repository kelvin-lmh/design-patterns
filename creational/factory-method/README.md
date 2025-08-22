# Factory method pattern

## Intent

    Factory Method is a Creational Design Pattern that defines an interface for creating objects, but lets sub classes decide which specific objects to create.

## Problem

    In programming, sometimes you need to create many different types of objects (for example: Truck, Ship...).
    If you use new or switch-case directly to initialize:
        - Difficult to extend when adding new types.
        - Violates Open/Closed Principle (add new types and have to modify old code).
        - Makes code highly dependent on specific classes.

## Solution

    - Define a common interface for products (eg: Transport).
    - Each specific product type (Truck, Ship) will implement this interface.
    - Instead of letting the client create the object itself, we give this creation to the Factory.
    - The client just calls the Factory to get the object it needs → The Factory will decide which type to create.

## Structure

    - Product: interface representing a generic product.
    - ConcreteProduct: concrete implementations of Product.
    - Creator (Factory): interface that declares a factory method that returns a Product, may contain some logic that uses the Product.
    - ConcreteCreator: implementation of the Factory Method to create concrete products.
    - Client: uses Creator to get a Product without knowing the concrete type.

## Real-World Analogy

    Payment Gateway
    Similar to when you shop online:
    You can choose PayPal, Stripe, MoMo to pay.

## Applicability

    - Use Factory Method when your code needs to work with different types of objects, but you don’t want to hard-code their classes.
    - When adding a new type (e.g., new Transport, new PaymentGateway), you want to avoid modifying existing code → helps follow the Open/Closed Principle.
    - When the object creation process is complex or varies by type, centralize it in factory methods instead of spreading new and switch-case everywhere.
    - Use when you want to reduce direct depend on a concrete class.

## Pros

    - Reduces dependency between client and sub classes because client only knows and works with interface (Product).
    - Easy to extend new products just need to add sub struct/class
    - Follow the Open/Closed principle. You can introduce new product types into the program without breaking existing code.

## Cons

    - The code may become more complicated since you need to introduce a lot of new subclasses to implement the pattern.
    - Can get complicated if only a few simple types are needed — sometimes a Simple Factory or DI is enough.

## Relations with Other Patterns
