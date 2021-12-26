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
	//"github.com/gin-gonic/gin"
	//"net/http"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

import (
	"bob.com/dubbogo-test-app/consumer-app/cmd/server/api"
)

var grpcGreeterImpl = new(api.GreeterClientImpl)

const MeshRouteConf = "apiVersion: service.dubbo.apache.org/v1alpha1\n" +
	"kind: DestinationRule\n" +
	"metadata: { name: demo-route }\n" +
	"spec:\n" +
	"  host: demo\n" +
	"  subsets:\n" +
	"    - labels: { env-sign: xxx, tag1: hello }\n" +
	"      name: isolation\n" +
	"    - labels: { env-sign: yyy }\n" +
	"      name: testing-trunk\n" +
	"    - labels: { env-sign: zzz }\n" +
	"      name: testing\n" +
	"  trafficPolicy:\n" +
	"    loadBalancer: { simple: ROUND_ROBIN }\n\n" +
	"---\n\n" +
	"apiVersion: service.dubbo.apache.org/v1alpha1\n" +
	"kind: VirtualService\n" +
	"metadata: {name: demo-route}\n" +
	"spec:\n" +
	"  dubbo:\n" +
	"    - routedetail:\n" +
	"        - match:\n" +
	"            - sourceLabels: {trafficLabel: xxx}\n" +
	"          name: xxx-project\n" +
	"          route:\n" +
	"            - destination: {host: demo, subset: isolation}\n" +
	"        - match:\n" +
	"            - sourceLabels: {trafficLabel: testing-trunk}\n" +
	"          name: testing-trunk\n" +
	"          route:\n" +
	"            - destination: {host: demo, subset: testing-trunk}\n" +
	"        - name: testing\n" +
	"          route:\n" +
	"            - destination: {host: demo, subset: testing}\n" +
	"      services:\n" +
	"        - {exact: com.apache.dubbo.sample.basic.IGreeter}\n" +
	"  hosts: [demo]"

// export DUBBO_GO_CONFIG_PATH= PATH_TO_SAMPLES/helloworld/go-client/conf/dubbogo.yml
func main() {
	config.SetConsumerService(grpcGreeterImpl)
	if err := config.Load(); err != nil {
		panic(err)
	}
	// cofig center
	dynamicConfiguration, err := config.GetRootConfig().ConfigCenter.GetDynamicConfiguration()
	if err != nil {
		panic(err)
	}

	// publish mesh route config
	err = dynamicConfiguration.PublishConfig("dubbo.io.MESHAPPRULE", "dubbo", MeshRouteConf)
	if err != nil {
		return
	}
	//// 1.创建路由
	//r := gin.Default()
	//// 2.绑定路由规则，执行的函数
	//// gin.Context，封装了request和response
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//// health check
	//r.GET("/health.check", func(c *gin.Context) {
	//	c.String(http.StatusOK, "ok")
	//})
	//// 3.监听端口，默认在8080
	//// Run("里面不指定端口号默认为8080")
	//r.Run(":8001")

	logger.Info("start to test dubbo")
	req := &api.HelloRequest{
		Name: "laurence",
	}
	reply, err := grpcGreeterImpl.SayHello(context.Background(), req)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("client response result: %v\n", reply)
}
