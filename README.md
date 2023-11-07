# gauth-go

---

[![Go Reference](https://pkg.go.dev/badge/github.com/onee-only/gauth-go.svg)](https://pkg.go.dev/github.com/onee-only/gauth-go)
[![codecov](https://codecov.io/gh/onee-only/gauth-go/graph/badge.svg?token=CG5WOHFTMX)](https://codecov.io/gh/onee-only/gauth-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/onee-only/gauth-go)](https://goreportcard.com/report/github.com/onee-only/gauth-go)

Package gauth is go version of gauth sdk.

## Description

---

광주소프트웨어 마이스터고등학교의 통합 계정 관리 서비스인 ```GAuth```를 ```Golang```으로 더 손쉽게 사용할 수 있도록 만들어진 라이브러리 입니다

## Installation

```
go get -u https://github.com/onee-only/gauth-go
```

## How To Use

### Client Initialization 
```go
package main

import(
	"github.com/onee-only/gauth-go"
)

func main() {
	client := gauth.NewDefaultClient(gauth.ClientOpts{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		RedirectURI:  "localhost:8080",
	})
}
```

### code 발급
```go
package main

import (
	"log"
	
	"github.com/onee-only/gauth-go"
)

func issueCode(client *gauth.Client) {
	code, err := client.IssueCode("email", "password")
	if err != nil {
		// You shouldn't use panic in your code.
		panic(err)
	}
	log.Println(code)
}
```

### Token 발급

```go
package main

import (
	"log"

	"github.com/onee-only/gauth-go"
)

func issueToken(client *gauth.Client) {
	accessToken, refreshToken, err := client.IssueToken("code")
	if err != nil {
		// You shouldn't use panic in your code.
		panic(err)
	}
	log.Println(accessToken)
	log.Println(refreshToken)
}
```

### Token 재발급
```go
package main

import (
	"log"
	
	"github.com/onee-only/gauth-go"
)

func reIssueToken(client *gauth.Client) {
	access, refresh, err := client.ReIssueToken("refreshToken")
	if err != nil {
		// You shouldn't use panic in your code.
		panic(err)
	}
	log.Println(access)
	log.Println(refresh)
}
```

### 유저 정보 가져오기
```go
package main

import (
	"log"
	
	"github.com/onee-only/gauth-go"
)

func getUserInfo(client *gauth.Client) {
	info, err := client.GetUserInfo("accessToken")
	if err != nil {
		// You shouldn't use panic in your code.
		panic(err)
	}
	log.Println(info)
}
```