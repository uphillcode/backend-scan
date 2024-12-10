package operations

import (
	"backend-scan/internal/models"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type CreateOperation struct {
	Service Service
}

func NewCreateOperation(service Service) *CreateOperation {
	return &CreateOperation{
		Service: service,
	}
}

func (c *CreateOperation) Execute(ctx context.Context) error {
	var table1 = "students"
	var table2 = "identities"
	var table3 = "cypher_code"
	var column = "code"

	settings, err := c.Service.GetSettings()
	if err != nil {
		return fmt.Errorf("error getting settings: %w", err)
	}

	table1Active := c.isTableActive(settings, table1)
	table2Active := c.isTableActive(settings, table2)
	table3Active := c.isTableActive(settings, table3)

	fmt.Println("table1Active: ", table1Active)
	fmt.Println("table2Active: ", table2Active)
	fmt.Println("table3Active: ", table3Active)
	if table1Active && table2Active {
		fmt.Println("inside of the  case 1: ", table3Active)
		err = c.handleTable1AndTable2(ctx, column)

	}
	if table3Active {
		fmt.Println("inside of the case table3Active: ", table3Active)
		err = c.handleTable3(ctx)
		if err != nil {
			return err
		}
	} else {
		switch {
		case table1Active && table2Active:
			fmt.Println("inside of the case 1: ", table3Active)
			err = c.handleTable1AndTable2(ctx, column)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("no valid table configuration active")
		}
	}

	if err != nil {
		return fmt.Errorf("error in table operations: %w", err)
	}

	return nil
}

func (c *CreateOperation) handleTable1AndTable2(ctx context.Context, column string) error {

	identitiesWithoutMatchingStudents, err := c.Service.FindAllIdentityWithoutMatchingStudents()
	if err != nil {
		return fmt.Errorf("error getting all identity without matching students: %w", err)
	}

	for _, identity := range identitiesWithoutMatchingStudents {
		observationAdd := models.ObservationAdd{
			Observation: models.Observation{
				Code:        fmt.Sprintf("%d", identity.Code),
				Litho:       fmt.Sprintf("%d", identity.Litho),
				Tema:        "N",
				State:       "0",
				Type:        "identities",
				CalendarsID: 2,
			},
		}

		_, err := c.Service.InsertObservation(observationAdd)
		if err != nil {
			return fmt.Errorf("error inserting observation: %w", err)
		}
	}

	err = c.processDuplicates("students", column)
	if err != nil {
		return err
	}

	err = c.processDuplicates("identities", column)
	if err != nil {
		return err
	}

	return nil
}
func (c *CreateOperation) handleTable3(ctx context.Context) error {
	fmt.Println("Handling table3")

	responses, err := c.Service.GetResponses()
	if err != nil {
		return fmt.Errorf("error getting responses: %w", err)
	}
	fmt.Printf("Responses RESPUESTAS:\n")
	for i, response := range responses {
		if i >= 5 {
			break
		}
		fmt.Printf("%+v\n", response)
	}

	cypherCodes, err := c.Service.GetClavesToCalification()
	if err != nil {
		return fmt.Errorf("error getting claves to calification: %w", err)
	}
	fmt.Printf("Cypher Codes CLAVES:\n")
	for i, cypherCode := range cypherCodes {
		if i >= 5 {
			break
		}
		fmt.Printf("%+v\n", cypherCode)
	}

	studentAndIdentities, err := c.Service.FindAllStudentAndIdentity()
	if err != nil {
		return fmt.Errorf("error getting student and identity: %w", err)
	}
	fmt.Printf("Student and Identities identity and student:\n")
	for i, studentAndIdentity := range studentAndIdentities {
		if i >= 5 {
			break
		}
		fmt.Printf("%+v\n", studentAndIdentity)
	}

	identityMap := make(map[int]models.StudentAndIdentity)
	for _, record := range studentAndIdentities {
		if record.Litho == "" {
			fmt.Printf("Skipping record with empty litho: %+v\n", record)
			continue
		}
		normalizedLitho, err := normalizeLitho(record.Litho)
		if err != nil {
			fmt.Printf("Error normalizing litho: %v\n", err)
			continue
		}
		identityMap[normalizedLitho] = record
	}
	fmt.Printf("Identity Map (first 5 entries):\n")
	count := 0
	for key, value := range identityMap {
		fmt.Printf("%d: %+v\n", key, value)
		count++
	}

	var emptyLithoCount int
	for _, response := range responses {
		if response.Litho == "" {
			fmt.Printf("Skipping response with empty litho for student ID: %d\n", response.StudentID)
			emptyLithoCount++
			fmt.Printf("Response Code: %d, Tema: N/A, Correctas: 0, Incorrectas: 0, Sin Responder: 0\n", response.StudentID)
			continue
		}

		normalizedLitho, err := normalizeLitho(response.Litho)
		if err != nil {
			fmt.Printf("Error normalizing litho for response code: %d\n", response.StudentID)
			continue
		}

		if identity, exists := identityMap[normalizedLitho]; exists {
			tema := identity.Tema

			correctas, incorrectas, sinResponder := c.calificarRespuestas(response, cypherCodes, tema, identity.Litho)

			fmt.Printf("Response Code: %d, Tema: %s, Correctas: %d, Incorrectas: %d, Sin Responder: %d\n",
				response.StudentID, tema, correctas, incorrectas, sinResponder)
		} else {
			fmt.Printf("No identity found for response code: %d\n", response.StudentID)
		}
	}

	fmt.Printf("Total responses with empty litho: %d\n", emptyLithoCount)

	return nil
}
func (c *CreateOperation) calificarRespuestas(response models.StudentResponse, cypherCodes []models.Cypher_code, tema string, litho string) (correctas int, incorrectas int, sinResponder int) {
	fmt.Printf("Calificando respuestas para Student Code: %d, Litho: %s, Tema: %s\n", response.StudentID, litho, tema)

	for _, cypherCode := range cypherCodes {
		if cypherCode.Tema == tema {
			questionsEvaluated := 0

			for _, studentResponse := range response.Responses {
				if questionsEvaluated >= int(cypherCode.Number_question) {
					break
				}

				if studentResponse.Response == "" {
					sinResponder++
				} else {
					respuestaCorrecta := false
					for _, correctResponse := range cypherCode.Response {
						if studentResponse.QuestionNumber == correctResponse.QuestionNumber && studentResponse.Response == correctResponse.Response {
							respuestaCorrecta = true
							break
						}
					}
					if respuestaCorrecta {
						correctas++
					} else {
						incorrectas++
					}
				}

				questionsEvaluated++
			}

			err := c.Service.InsertResponse(response.Code, response.Tema, correctas, incorrectas, sinResponder, litho)
			if err != nil {
				fmt.Printf("Error inserting response for Student Code: %d, Litho: %s, Tema: %s\n", response.StudentID, litho, tema)
				return 0, 0, 0
			}
		} else {
			fmt.Printf("Skipping cypherCode with Tema: %s, Litho: %s as it does not match the criteria\n", cypherCode.Tema, litho)
		}
	}

	fmt.Printf("Final Grades for Student Code: %d, Litho: %s, Tema: %s - Correctas: %d, Incorrectas: %d, Sin Responder: %d\n",
		response.StudentID, litho, tema, correctas, incorrectas, sinResponder)

	return
}

func normalizeLitho(litho string) (int, error) {
	if litho == "" {
		return 0, errors.New("empty litho")
	}
	normalizedLitho, err := strconv.Atoi(strings.TrimLeft(litho, "0"))
	if err != nil {
		return 0, err
	}
	return normalizedLitho, nil
}

func (c *CreateOperation) processDuplicates(tableName string, column string) error {
	duplicates, err := c.Service.GetGroupedColumnsCount(tableName, column)
	if err != nil {
		return fmt.Errorf("error getting duplicates from %s: %w", tableName, err)
	}

	for _, duplicate := range duplicates {
		fmt.Printf("Inserting duplicate from %s: %v with count %d\n", tableName, duplicate.ColumnValue, duplicate.Count)
		err = c.Service.InsertDuplicateInNewTable(duplicate.ColumnValue, duplicate.Count, tableName)
		if err != nil {
			return fmt.Errorf("error inserting duplicate from %s: %w", tableName, err)
		}
	}

	return nil
}

func (c *CreateOperation) Name() string {
	return "Create"
}

func (c *CreateOperation) isTableActive(settings []models.Setting, tableName string) bool {
	for _, setting := range settings {
		if setting.Table == tableName {
			return setting.State == "1"
		}
	}
	return false
}
