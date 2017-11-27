package commands

import (
	"fmt"

	"github.com/pivotal-cf/jhanda/commands"
)

type StagedProducts struct {
	presenter         Presenter
	diagnosticService diagnosticService
}

func NewStagedProducts(presenter Presenter, diagnosticService diagnosticService) StagedProducts {
	return StagedProducts{
		presenter:         presenter,
		diagnosticService: diagnosticService,
	}
}

func (sp StagedProducts) Execute(args []string) error {
	diagnosticReport, err := sp.diagnosticService.Report()
	if err != nil {
		return fmt.Errorf("failed to retrieve staged products %s", err)
	}

	stagedProducts := diagnosticReport.StagedProducts

	sp.presenter.PresentStagedProducts(stagedProducts)
	return nil
}

func (sp StagedProducts) Usage() commands.Usage {
	return commands.Usage{
		Description:      "This authenticated command lists all staged products.",
		ShortDescription: "lists staged products",
	}
}
