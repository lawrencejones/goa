package design

import . "goa.design/goa.v2/http/design"
import . "goa.design/goa.v2/http/dsl"

var _ = Service("storage", func() {
	Description("The storage service makes it possible to view, add or remove wine bottles.")

	HTTP(func() {
		Path("/storage")
	})

	Method("list", func() {
		Description("List all stored bottles")
		Result(CollectionOf(StoredBottle), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show bottle by ID")
		Payload(func() {
			Attribute("id", String, "ID of bottle to show")
			Required("id")
		})
		Result(StoredBottle)
		Error("not_found", NotFound, "Bottle not found")
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add", func() {
		Description("Add new bottle and return its ID.")
		Payload(Bottle)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("remove", func() {
		Description("Remove bottle from storage")
		Payload(func() {
			Attribute("id", String, "ID of bottle to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Bottle not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})