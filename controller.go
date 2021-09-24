package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/betNevS/build-own-web-framework/framework"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)
	ctx, cancel := context.WithTimeout(c, time.Second)
	defer cancel()
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")
		finish <- struct{}{}
	}()
	select {
	case p := <-panicChan:
		c.WriteMutex().Lock()
		defer c.WriteMutex().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-ctx.Done():
		c.WriteMutex().Lock()
		defer c.WriteMutex().Unlock()
		c.Json(500, "timeout")
		c.SetHasTimeout()
	}
	return nil
}

func BarControllerHandler(c *framework.Context) error {
	str := c.FormString("name", "bo")
	fmt.Println("post content: ", str)
	c.Json(200, str)
	return nil
}
