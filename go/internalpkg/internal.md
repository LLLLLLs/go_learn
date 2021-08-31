### internal

golang 特殊文件夹:internal

只有同个父目录下的文件夹才有访问权限

```
internalpkg
├── diffather
│   └── bar1.go
│
├── samefather
│   ├── internal
│   │   └── foo
│   │       └── foo.go
│   └── bar2.go
│
└── bar3.go
```

针对上述工程目录，bar2(与internal直属同一个包)可以顺利调用foo中的代码，而bar1（diffather与samefather为兄弟目录）
与bar3（samefather与internalpkg为父子目录）都无法直接调用foo中的代码