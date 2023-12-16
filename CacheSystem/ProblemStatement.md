# Cache System Low-Level Design

## Problem Statement

### Overview
Design a cache system with minimal concurrency support to efficiently store key-value pairs and handle scenarios where the cache becomes full, requiring the removal of entries.

### Operations
1. **Put**: Allow users to store a value associated with a specific key in the cache.
2. **Get**: Enable users to retrieve a previously stored value using a provided key.
3. **Eviction**: Implement a mechanism to remove entries when the cache reaches its capacity and a new key-value pair needs to be added.

### Minimal Concurrency Support
1. **Concurrent Access:**
    - Basic support for concurrent access is required.
    - Operations like Put, Get, and Eviction should be designed to handle simple concurrency scenarios.

2. **Atomic Operations:**
    - Ensure basic atomicity for operations to prevent common race conditions.

3. **Locking Mechanisms:**
    - Implement minimal locking mechanisms, such as simple mutexes, to provide basic thread safety.
    - Threads should acquire locks before accessing critical sections (e.g., modifying the cache).

### Expectations
* **Functionality:** The code should implement the specified operations with basic support for concurrency.
* **Modularity and Readability:** Code should be modular, readable, and maintainable, adhering to clean and professional coding standards.
* **Extensibility and Scalability:** The design should be flexible and scalable, capable of accommodating new requirements with minimal changes.
* **Eviction Policy:** The cache system should allow switching between eviction policies dynamically.
