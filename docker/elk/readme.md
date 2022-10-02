### 错误处理
```
Exiting: error loading config file: config file ("filebeat.yml") can only be writable by the owner but the permissions are "-rwxrwxrwx" (to fix the permissions use: 'chmod go-w /usr/share/filebeat/filebeat.yml'

chmod 755 filebeat.yml
```

```
metrics/metrics.go:376  error getting cgroup stats: error fetching stats for controller io: error fetching IO stats: error fetching io.pressure for path /sys/fs/cgroup:: open /sys/fs/cgroup/io.pressure: no such file or directory

https://discuss.elastic.co/t/error-metrics-metrics-go-376-error-getting-cgroup-stats-error-fetching-stats-for-controller-io-error-fetching-io-stats-error-fetching-io-pressure-for-path-sys-fs-cgroup-open-sys-fs-cgroup-io-pressure-no-such-file-or-directory/307804

官方已知错误，8.x已经修复
```