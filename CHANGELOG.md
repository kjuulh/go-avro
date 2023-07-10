# Changelog

## 1.0.0 (2023-07-10)


### Features

* change to master branch ([f84b009](https://github.com/kjuulh/go-avro/commit/f84b009ebba6b686e6753aa4965e7fb31f067d00))
* test commit ([39cac0c](https://github.com/kjuulh/go-avro/commit/39cac0c3e1f1fc8398f933db88d38c74a8e5c3af))
* update with fork name ([6814c66](https://github.com/kjuulh/go-avro/commit/6814c660ad835ef097d41afa35a12a9510dede19))
* with workflow ([9b1635d](https://github.com/kjuulh/go-avro/commit/9b1635df220136e80a6d6ccd149e29b2ea8eb597))


### Bug Fixes

* don't panic for out-of-range on corrupted union decoding. ([18b258d](https://github.com/kjuulh/go-avro/commit/18b258d6b1d69766c5cb561ae8bdd8d93d320617)), closes [#10](https://github.com/kjuulh/go-avro/issues/10)

## Changelog

#### Version 0.3 (2017-12-17)

API Changes:

 - DataFileReader gets some major updates:

   - Can now decode 'deflate' codec files.
   - Does not read the entire file into memory, for large files and streaming.
     Also added `Close()` to close the underlying file if applicable.
   - [breaking] No longer pass a DatumReader to DataFileReader, it creates it for you.
   - [breaking] `Next()` only returns a single error value, use `HasNext()` in a
     loop to proceed.

 - Removed Seek and Tell from Decoder, should have no real impact on most applications.

#### Version 0.2 (2017-12-16)

Intention: start making changes towards a 1.0 release.

API Changes:
 - The `BinaryEncoder` type is now a private type. `avro.NewBinaryEncoder()`
   now returns a value of the `Encoder` interface.
 - Decoder changes:

   - The `BinaryDecoder` type is now also a private type. `avro.NewBinaryDecoder()`
     now returns a value of the `Decoder` interface.
   - Removed `ReadFixedWithBounds`, removed the use case which dictated it.
   - Add an implementation of BinaryDecoder which can work on an io.Reader

 - Rename the `Writer` and `Reader` interfaces to `Marshaler` and `Unmarshaler` to
   be more like the JSON encoder and also use similar method names.
 - Rename error types `FooBar` to be `ErrFooBar`

Improvements:
 - Major improvement to docs and docs coverage
 - Add a singular `NewDatumWriter` which will become the replacement for the generic/specific types.


#### Version 0.1 (2017-08-23)

 - First version after forking from elodina.
 - Started a semver-considering API, using the gopkg.in interface,
   and planning for a 1.x release.

Improvements:
 - Error reporting: specify which field is missing when throwing FieldDoesNotExist
   [#5](https://github.com/go-avro/avro/pull/5)
 - Speedup encoding for strings and bools
   [#6](https://github.com/go-avro/avro/pull/6)
 - Can prepare schemas which are self-recursive and co-recursive.

Bug Fixes:
 - Can decode maps of non-primitive types [#2](https://github.com/go-avro/avro/pull/2)
 - Fix encoding of 'fixed' type [#3](https://github.com/go-avro/avro/pull/3) [elodina/#78](https://github.com/elodina/go-avro/issues/78)
 - Fix encoding of boolean when used in a type union [#4](https://github.com/go-avro/avro/pull/4)
