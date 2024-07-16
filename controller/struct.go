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
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	VisitDate  string             `bson:"visitdate,omitempty" json:"visitdate,omitempty"`
	DoctorName string             `bson:"doctor,omitempty" json:"doctor,omitempty"`
	Diagnosis  string             `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Treatment  string             `bson:"treatment,omitempty" json:"treatment,omitempty"`
	Notes      string             `bson:"notes,omitempty" json:"notes,omitempty"`
}

type Biodata struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	PasienName    string             `bson:"pasienName,omitempty" json:"pasienName,omitempty"`
	Gender        string             `bson:"gender,omitempty" json:"gender,omitempty"`
	TTL           string             `bson:"ttl,omitempty" json:"ttl,omitempty"`
	Status        string             `bson:"status,omitempty" json:"status,omitempty"`
	Phone_number  string             `bson:"phonenumber,omitempty" json:"phonenumber,omitempty"`
	Alamat        string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Doctor        Doctor             `bson:"doctor,omitempty" json:"doctor,omitempty"`
	MedicalRecord MedicalRecord      `bson:"medicalRecord,omitempty" json:"medicalRecord,omitempty"`
}

type ReqPasien struct {
	PasienName   string        `bson:"pasienName,omitempty" json:"pasienName,omitempty" example:"Doni"`
	Gender       string        `bson:"gender,omitempty" json:"gender,omitempty" example:"Perempuan"`
	TTL          string        `bson:"ttl,omitempty" json:"ttl,omitempty" `
	Status       string        `bson:"status,omitempty" json:"status,omitempty" example:"single"`
	Phone_number string        `bson:"phonenumber,omitempty" json:"phonenumber,omitempty" example:"08567432"`
	Alamat       string        `bson:"alamat,omitempty" json:"alamat,omitempty" example:"Sariasih 25, Bandung"`
	DoctorName   Doctor        `bson:"doctor,omitempty" json:"doctor,omitempty" example:"Ardi"`
	Diagnosis    MedicalRecord `bson:"medicalRecord,omitempty" json:"medicalRecord,omitempty" example:"Stroke"`
}
