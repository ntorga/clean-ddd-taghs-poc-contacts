package cliController

import (
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/dto"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/useCase"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/domain/valueObject"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra"
	"github.com/ntorga/clean-ddd-taghs-poc-contacts/src/infra/db"
	cliHelper "github.com/ntorga/clean-ddd-taghs-poc-contacts/src/presentation/cli/helper"
	"github.com/spf13/cobra"
)

type ContactController struct {
	persistentDbSvc *db.PersistentDatabaseService
}

func NewContactController(
	persistentDbSvc *db.PersistentDatabaseService,
) *ContactController {
	return &ContactController{
		persistentDbSvc: persistentDbSvc,
	}
}

func (controller *ContactController) Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "GetContacts",
		Run: func(cmd *cobra.Command, args []string) {
			contactQueryRepo := infra.NewContactQueryRepo(controller.persistentDbSvc)
			contactsList, err := useCase.GetContacts(contactQueryRepo)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, contactsList)
		},
	}

	return cmd
}

func (controller *ContactController) Create() *cobra.Command {
	var nameStr string
	var nicknameStr string
	var phoneStr string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "CreateNewContact",
		Run: func(cmd *cobra.Command, args []string) {
			name := valueObject.NewPersonNamePanic(nameStr)
			nickname := valueObject.NewNicknamePanic(nicknameStr)
			phone := valueObject.NewPhoneNumberPanic(phoneStr)

			createContactDto := dto.NewCreateContact(
				name,
				nickname,
				phone,
			)

			contactQueryRepo := infra.NewContactQueryRepo(controller.persistentDbSvc)
			contactCmdRepo := infra.NewContactCmdRepo(controller.persistentDbSvc)

			err := useCase.CreateContact(
				contactQueryRepo,
				contactCmdRepo,
				createContactDto,
			)
			if err != nil {
				cliHelper.ResponseWrapper(false, err.Error())
			}

			cliHelper.ResponseWrapper(true, "ContactCreated")
		},
	}

	cmd.Flags().StringVarP(&nameStr, "name", "n", "", "Name")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVarP(&nicknameStr, "nickname", "k", "", "Nickname")
	cmd.MarkFlagRequired("nickname")
	cmd.Flags().StringVarP(&phoneStr, "phone", "p", "", "Phone")
	cmd.MarkFlagRequired("phone")
	return cmd
}
