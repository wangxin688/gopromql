package main

import (
	"fmt"

	"github.com/wangxin688/gopromql"
)

func main() {

	pb1 := example1()

	pb2 := example2()

	pb3 := example3()

	pb4 := example4()

	fmt.Println(pb1, pb2, pb3, pb4)
}

func example1() string {

	pb, err := gopromql.NewPromqlBuilder("http_requests_total").
		WithFuncName("last_over_time").
		WithWindow("5m").
		Build()
	if err != nil {
		panic(err)
	}

	return pb
}

func example2() string {

	pb, err := gopromql.NewPromqlBuilder("http_requests_total").
		WithFuncName("sum_over_time").
		WithLabels(gopromql.Label{Name: "status_code", Matcher: "=", Value: "200"}).
		WithLabels(gopromql.Label{Name: "method", Matcher: "=~", Value: "GET|POST|PUT"}).
		WithLabels(gopromql.Label{Name: "path", Matcher: "!=", Value: "/health"}).
		WithWindow("5m").
		Build()
	if err != nil {
		panic(err)
	}
	return pb
}

func example3() string {
	pb, err := gopromql.NewPromqlBuilder("http_requests_total").
		WithFuncName("sum_over_time").
		WithLabels(gopromql.Label{Name: "status_code", Matcher: "=", Value: "200"}).
		WithLabels(gopromql.Label{Name: "method", Matcher: "=~", Value: "GET|POST|PUT"}).
		WithLabels(gopromql.Label{Name: "path", Matcher: "!=", Value: "/health"}).
		WithComp(gopromql.Compare{Op: gopromql.GreaterThan, Value: 0}).
		WithWindow("5m").
		Build()
	if err != nil {
		panic(err)
	}

	return pb
}

func example4() string {
	pb, err := gopromql.NewPromqlBuilder("http_requests_total").
		WithFuncName("sum_over_time").
		WithLabels(gopromql.Label{Name: "status_code", Matcher: "=", Value: "200"}).
		WithAgg(gopromql.Aggregation{Op: gopromql.Max, AggWay: gopromql.GroupBy, By: []string{"method", "path"}}).
		WithWindow("5m").
		Build()
	if err != nil {
		panic(err)
	}
	return pb
}
