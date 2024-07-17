package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	inimodel "github.com/Nidasakinaa/BeRS/model"
	cek "github.com/Nidasakinaa/BeRS/module"
	"github.com/Nidasakinaa/ws-rumahsakit/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

// GetPasien godoc
// @Summary Get All Data Pasien.
// @Description Mengambil semua data pasien.
// @Tags Pasien
// @Accept json
// @Produce json
// @Success 200 {object} Biodata
// @Router /pasien [get]
func GetPasien(c *fiber.Ctx) error {
	ps := cek.GetAllPasien(config.Ulbimongoconn, "DataPasien")
	return c.JSON(ps)
}

// GetPasienID godoc
// @Summary Get By ID Data Pasien.
// @Description Ambil per ID data pasien.
// @Tags Pasien
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Biodata
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /pasien/{id} [get]
func GetPasienID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetPasienByID(objID, config.Ulbimongoconn, "DataPasien")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

// InsertDataPasien godoc
// @Summary Insert data pasien.
// @Description Input data pasien.
// @Tags Pasien
// @Accept json
// @Produce json
// @Param request body ReqPasien true "Payload Body [RAW]"
// @Success 200 {object} Biodata
// @Failure 400
// @Failure 500
// @Router /insert [post]
func InsertDataPasien(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var pasien inimodel.Biodata
	if err := c.BodyParser(&pasien); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := cek.InsertPasien(db, "DataPasien",
		pasien.PasienName,
		pasien.Gender,
		pasien.Usia,
		pasien.Phonenumber,
		pasien.Alamat,
		pasien.Doctor,
		pasien.MedicalRecord)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateData godoc
// @Summary Update data pasien.
// @Description Ubah data pasien.
// @Tags Pasien
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqPasien true "Payload Body [RAW]"
// @Success 200 {object} Biodata
// @Failure 400
// @Failure 500
// @Router /update/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	var pasien inimodel.Biodata
	if err := c.BodyParser(&pasien); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	err = cek.UpdatePasien(context.Background(), db, "DataPasien",
		objectID,
		pasien.PasienName,
		pasien.Gender,
		pasien.Usia,
		pasien.Phonenumber,
		pasien.Alamat,
		pasien.Doctor,
		pasien.MedicalRecord)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeletePasienByID godoc
// @Summary Delete data pasien.
// @Description Hapus data pasien.
// @Tags Pasien
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePasienByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = cek.DeletePasienByID(objID, config.Ulbimongoconn, "DataPasien")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
