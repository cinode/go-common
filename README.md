# Cinode - golang - common utilities

## picotestify

Minimal implementation of testify-like library for internal use in Cinode.

It has the following goals over the original testify one:

- No dependencies - it can not use external modules and the use of standard packages is reduced - no i/o, no system-specific code
- Additional type safety through generics - where possible dynamic function arguments passed through `any` are replaced by a typed version
- Only bare-minimum set of functions - only the most necessary assertions are left, no printf-like messages etc.

The interface of picotestify is a subset of the original testify one and thus switching to the original testify library should be easily achievable through go.mod's rewrite option.
