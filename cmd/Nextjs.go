package cmd

// Nextjs function
func Nextjs()  {
	SystemCommand("yarn create next-app . && touch tsconfig.json && yarn add --dev typescript @types/react @types/node && yarn add sass")		
}