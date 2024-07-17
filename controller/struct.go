package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doctor struct {
	Name      string `bson:"name,omitempty" json:"name,omitempty"`
	Specialty string `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Contact   string `bson:"contact,omitempty" json:"contact,omitempty"`
}

type MedicalRecord struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	VisitDate string             `bson:"visitdate,omitempty" json:"visitdate,omitempty"`
	Diagnosis string             `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Treatment string             `bson:"treatment,omitempty" json:"treatment,omitempty"`
}

type Biodata struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	PasienName    string             `bson:"pasienName,omitempty" json:"pasienName,omitempty"`
	Gender        string             `bson:"gender,omitempty" json:"gender,omitempty"`
	Usia          string             `bson:"usia,omitempty" json:"usia,omitempty"`
	Phone_number  string             `bson:"phonenumber,omitempty" json:"phonenumber,omitempty"`
	Alamat        string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Doctor        Doctor             `bson:"doctor,omitempty" json:"doctor,omitempty"`
	MedicalRecord MedicalRecord      `bson:"medicalRecord,omitempty" json:"medicalRecord,omitempty"`
}

type ReqPasien struct {
	PasienName    string        `bson:"pasienName,omitempty" json:"pasienName,omitempty" example:"Doni"`
	Gender        string        `bson:"gender,omitempty" json:"gender,omitempty" example:"Laki-laki"`
	Usia          string        `bson:"usia,omitempty" json:"usia,omitempty" example:"20 Tahun" `
	Phone_number  string        `bson:"phonenumber,omitempty" json:"phonenumber,omitempty" example:"08567432"`
	Alamat        string        `bson:"alamat,omitempty" json:"alamat,omitempty" example:"Sariasih 25, Bandung"`
	Doctor        Doctor        `bson:"doctor,omitempty" json:"doctor,omitempty"`
	MedicalRecord MedicalRecord `bson:"medicalRecord,omitempty" json:"medicalRecord,omitempty"`
}
