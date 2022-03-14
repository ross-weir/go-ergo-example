
# go-ergo-example

To generate the header and shared library follow the steps [here](https://github.com/ergoplatform/sigma-rust/tree/develop/bindings/ergo-lib-c).

Example project showing how to use the C ergo library in golang.

### Issues

```
undefined reference to `__chkstk'
more undefined references to `__CxxFrameHandler3' follow
```

- Ensure the same toolchain is used to compile in Rust and Go (GNU vs MSVC)
- Ensure additionally libraries are linked, we also link `bcrypt` in this example
