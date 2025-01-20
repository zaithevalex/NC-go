package main

func main() {
	//ctx := context.Background()
	//
	//cfg, err := pgxpool.ParseConfig(connString)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//
	//cfg.ConnConfig.Tracer = otelpgx.NewTracer()
	//
	//conn, err := pgxpool.NewWithConfig(ctx, cfg)
	//if err != nil {
	//	log.Println(err.Error())
	//}

	//db, err := models.Connection()
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, _ = otel.InitTracer("http://localhost:14268/api/traces", "Note Service")
	////if err != nil {
	////	log.Fatal("init tracer", err)
	////}
	//
	//app := fiber.New()
	//app.Use(otelfiber.Middleware())
	//app.Post("/add", func(ctx *fiber.Ctx) error {
	//	var content models.Content
	//
	//	err := ctx.BodyParser(&content)
	//	if err != nil {
	//		return err
	//	}
	//
	//	id, err := models.Insert(content, db)
	//	if err != nil {
	//		return err
	//	}
	//
	//	return ctx.JSON(struct {
	//		Id int
	//	}{
	//		Id: id,
	//	})
	//})
	//
	//app.Get("/get", func(ctx *fiber.Ctx) error {
	//	id, err := strconv.Atoi(ctx.Query("id"))
	//	if err != nil {
	//		return err
	//	}
	//
	//	note, err := models.Get(id, db)
	//	if err != nil {
	//		return err
	//	}
	//
	//	return ctx.JSON(note)
	//})
	//
	//err = app.Listen(":8080")
	//if err != nil {
	//	panic(err)
	//}
}
