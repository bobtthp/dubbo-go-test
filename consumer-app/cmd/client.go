/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"flag"
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/gin-gonic/gin"
	"net/http"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

import (
	"github.com/dubbogo-test/consumer-app/api"
)

var grpcGreeterImpl = new(api.GreeterClientImpl)

var zk string

func init()  {
	flag.StringVar(&zk, "zk", "127.0.0.1:2181","-zk=127.0.0.1:2181")
}


func main() {
	flag.Parse()

	// init rootConfig with config api
	rc := config.NewRootConfigBuilder().
		SetConsumer(config.NewConsumerConfigBuilder().
			AddReference("GreeterClientImpl", config.NewReferenceConfigBuilder().
				SetProtocol("tri").
				SetInterface("com.apache.dubbo.sample.basic.IGreeter").
				Build()).
			Build()).
		AddRegistry("bob", &config.RegistryConfig{
		Protocol: "zookeeper",
		Address:  zk,
		Timeout:  "3s",
	}).
		Build()

	config.SetConsumerService(grpcGreeterImpl)
	hessian.RegisterPOJO(&api.User{})
	// load config
	if err := rc.Init(); err != nil {
		fmt.Println(err)

		panic(err)
	}

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// health check
	r.GET("/health.check", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/consumer", func(c *gin.Context) {
		req := &api.HelloRequest{
			Name: "laurence",
		}
		reply, err := grpcGreeterImpl.SayHello(context.Background(), req)
		if err != nil {
			logger.Error(err)
		}
		logger.Infof("client response result: %v\n", reply)
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8001")


}
