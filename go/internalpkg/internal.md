### internal

golang 特殊文件夹:internal

只有同个父目录下的文件夹才有访问权限

```
internalpkg
├── diffather
│   └── bar1.go
│
├── samefather
│   ├── another
│   │   └── bar4.go
│   ├── internal
│   │   └── foo
│   │       └── foo.go
│   └── bar2.go
│
└── bar3.go
```

- bar1 ×
- bar2 √
- bar3 ×
- bar4 √

总结，与internal包拥有公共父包的包才可调用internal里的函数