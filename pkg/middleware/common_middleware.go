package middleware

import (
	"cli-boilerplate-generator/constants/middleware"
	"cli-boilerplate-generator/pkg/utils/folder_utils"
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

func CommonMiddleware(handler func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		cwd, err := folder_utils.FolderNameInTerminal()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		ctx := context.WithValue(cmd.Context(), middleware.TERMINAL_PATH_GETTER, cwd)
		cmd.SetContext(ctx)

		// Actual handler
		handler(cmd, args)

	}
}
