package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"main/ec2"
	"main/vpc"
)

func main() {
	app := fiber.New()
	app.Get("/vpcs", func(ctx *fiber.Ctx) error {
		return ctx.JSON(vpc.GetVpc())
	})
	app.Get("/vpcSubnets", func(ctx *fiber.Ctx) error {
		return ctx.JSON(vpc.GetVpcSubnets())
	})
	app.Get("/instanceTypes", func(ctx *fiber.Ctx) error {
		return ctx.JSON(ec2.GetEC2InstanceTypes())
	})
	app.Get("/ami", func(ctx *fiber.Ctx) error {
		return ctx.JSON(ec2.GetEC2AMI())
	})
	log.Fatal(app.Listen(":3000"))
}
