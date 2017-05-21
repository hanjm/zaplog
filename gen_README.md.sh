#!/usr/bin/env bash
echo "## zaplog是包装了[zap](https://github.com/uber-go/zap)、用于生产环境的日志输出工具"
echo "[![GoDoc](https://godoc.org/github.com/hanjm/zaplog?status.svg)](https://godoc.org/github.com/hanjm/zaplog)"
echo "[![Go Report Card](https://goreportcard.com/badge/github.com/hanjm/zaplog)](https://goreportcard.com/report/github.com/hanjm/zaplog)"
echo "[![code-coverage](http://gocover.io/_badge/github.com/hanjm/zaplog)](http://gocover.io/github.com/hanjm/zaplog)"
echo ""
echo "### Document"
echo ""
echo "\`\`\`go"
godoc . | sed '1,8d'
echo "\`\`\`"
echo ""