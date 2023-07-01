<div align="center">
<h1>Project base module</h1>
<p>
</p>
<a href="https://github.com/grpc-ecosystem/grpc-gateway/blob/main/LICENSE.txt"><img src="https://img.shields.io/github/license/grpc-ecosystem/grpc-gateway?color=379c9c&style=flat-square"/></a>
<a href="https://github.com/adon988/project-base/releases"><img src="https://img.shields.io/github/v/release/adon988/project-base?color=379c9c&logoColor=ffffff&style=flat-square"/></a>

</div>
## Project base module

## Version naming

For project base core module should naming with tag like 

```
{moduleName}/{version}
```

For Example:

| Module Name | Path | Tag Name |
| -- | -- | -- | 
| base | github.com/adon988/project-base/base | base/v0.0.1 |


### Commands

| Command | Description |
| -- | -- |
| `make buf-gen` | Will execute buf-format > buf-lint > buf-go-clean > buf-go "


### Tagging sync to go pkg

Go to go pkg path(like following) and send request to sync github tag version to go pkg

## License

This module is licensed under the BSD 3-Clause License.
See [LICENSE.txt](https://github.com/grpc-ecosystem/grpc-gateway/blob/main/LICENSE.txt) for more details.
