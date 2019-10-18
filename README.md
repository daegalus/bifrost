# Bifrost

Gateway to secure data.

## Usage

### Library

Please see examples.go for examples on using Bifrost.

### CLI

Not implemented yet.

## Cloud Providers

If you wish to make new cloud providers, they need to conform to the providers.types.Provider interface. Please use the GCP implementation as a guideline.

The plugin system works on a registry that is loaded with providers when the loader function is called (TODO: Make it automatic whenever a function is used so you don't have end-user call it manually like in the example. or make it part of a constuctor?).

Once loaded, the registry acts as a gateway from the main library to the plugins. Using this method we can use any provider as long as it has matching concepts, and the provider implementation is encapsulated into the plugin.
