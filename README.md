Interactive CLI Program with Validation in Go
This Go program implements an interactive Command-Line Interface (CLI) that prompts users with questions and validates their inputs using Go's standard library functionalities. Here’s an overview of its key aspects:

Features:
Modular Structure: The program is structured with separate functions for handling different aspects of user interaction.

Input Handling: Utilizes bufio.Scanner for reading user inputs from standard input (os.Stdin).

Question Handling: Defines questions as structs with prompts and corresponding handler functions.

Validation: Implements validation functions to ensure user inputs meet specific criteria (e.g., numeric and positive integers).

Usage:
Clone the repository to your local machine.
Run the program using go run main.go.
Follow the prompts to enter your name and age, ensuring to provide valid inputs as specified.
Notes:
Validation Function: Custom validation functions ensure that user inputs are validated according to specified rules (e.g., numeric and positive integers).
Modular Design: Separate functions (handleName, handleAge) handle different aspects of user inputs, promoting code organization and reusability.
Interactive Experience: Provides an interactive CLI interface for seamless user interaction and feedback.
