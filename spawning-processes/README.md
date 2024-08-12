## Spawning processes

### When shouldn't use

**Performance:** If you only need to perform simple tasks without interacting with an external program, spawning a new process can be inefficient and resource-consuming.

**Complex management:** Spawning multiple processes can make resource and error management more complex, especially when these processes need to interact with each other or when unexpected errors from other processes need to be handled. that process.

### When should use

**Lightweight tasks:** When you need to perform light tasks or parallel processing but don't need a completely isolated environment, you should use goroutine instead of spawn process.

**Resource management:** Goroutine is much lighter than a new process and can be easily managed within the same Go application.
