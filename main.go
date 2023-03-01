package main

import (
	"github.com/spf13/cobra"

	"fmt"
)

var rootCmd = &cobra.Command{
	Use:     "calculator",
	Version: "1.0.0",
	Short:   "calculator for add & mul",
	Long:    `calculator for addition and multiplication`,

	//In this case the root is not executable, meaning that a subcommand is required. This is accomplished by not providing a 'Run' for the 'rootCmd'.

}

var cmdAdd = &cobra.Command{
	Use:     "add",
	Version: "1.0.0",
	Short:   "Addition",
	Long:    `additon numbers`,

	// Uncomment the following line if your bare application
	Run: func(cmd *cobra.Command, args []string) {
		var sum float64
		numbers, _ := cmd.Flags().GetFloat32Slice("number")
		for _, j := range numbers {
			//fmt.Println(i)
			sum += float64(j)
		}
		fmt.Println(sum)
	},
}

var cmdMul = &cobra.Command{
	Use:     "mul",
	Version: "1.0.0",
	Short:   "Multiplication",
	Long:    `multiplication of numbers`,

	// Uncomment the following line if your bare application
	Run: func(cmd *cobra.Command, args []string) {
		//jokeTer, _ := cmd.Flags().GetString("term")
		numbers, _ := cmd.Flags().GetFloat32Slice("number")
		//fmt.Println(number)
		var prod float64
		prod = 1

		for _, j := range numbers {
			//fmt.Println(i)
			prod *= float64(j)
		}
		fmt.Println(prod)
	},
}

func Execute() {

	//var rootCmd = &cobra.Command{Use: "calculator"}
	rootCmd.AddCommand(cmdAdd)
	//rootCmd.PersistentFlags().String("term", "", "A search term for a dad joke.")
	rootCmd.PersistentFlags().Float32SliceP("number", "n", []float32{}, "enter number(required)") //P is for shorthand value
	rootCmd.MarkPersistentFlagRequired("number")
	rootCmd.AddCommand(cmdMul)
	cobra.CheckErr(rootCmd.Execute())
}

func main() {
	Execute()
}
