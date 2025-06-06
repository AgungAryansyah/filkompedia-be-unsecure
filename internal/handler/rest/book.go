package rest

import (
	"net/http"

	"github.com/AgungAryansyah/filkompedia-be-insecure/model"
	"github.com/AgungAryansyah/filkompedia-be-insecure/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r *Rest) GetBook(ctx *fiber.Ctx) error {
	bookIdString := ctx.Params("id")
	bookId, err := uuid.Parse(bookIdString)
	if err != nil {
		return err
	}

	books, err := r.service.BookService.GetBook(bookId)
	if err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", books)
	return nil
}

func (r *Rest) SearchBooks(ctx *fiber.Ctx) error {
	var bookSearch model.BookSearch
	bookSearch.Page = ctx.QueryInt("page", 1)
	bookSearch.PageSize = ctx.QueryInt("size", 9)
	bookSearch.SearchParam = ctx.Query("search", "%")

	books, err := r.service.BookService.SearchBooks(bookSearch)
	if err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", books)
	return nil
}

func (r *Rest) CreateBook(ctx *fiber.Ctx) error {
	var create model.CreateBook
	if err := ctx.BodyParser(&create); err != nil {
		return err
	}

	if err := r.service.BookService.CreateBook(&create); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", nil)
	return nil
}

func (r *Rest) DeleteBook(ctx *fiber.Ctx) error {
	bookIdString := ctx.Params("id")
	bookId, err := uuid.Parse(bookIdString)
	if err != nil {
		return err
	}

	if err := r.service.BookService.DeleteBook(bookId); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", nil)
	return nil
}

func (r *Rest) EditBook(ctx *fiber.Ctx) error {
	var edit model.EditBook
	if err := ctx.BodyParser(&edit); err != nil {
		return err
	}

	if err := r.service.BookService.EditBook(edit); err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", nil)
	return nil
}

func (r *Rest) UploadBookCover(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	url, err := r.service.BookService.UploadBookCover(file)
	if err != nil {
		return err
	}

	response.Success(ctx, http.StatusOK, "success", url)
	return nil
}
