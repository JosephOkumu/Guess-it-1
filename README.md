# guess-it-1

# Description
This Go program reads a number from the standard input, calculates the average, mean, variance and standard deviation and gives a prediction of the range the next number the user inputs will fall. 
## Documentation
This section provides details on how to use the guess-it-1 program.

### Installation
To use the this program, you need to have Node.js installed on your system. You can download and install the latest version of Node.js from [here](https://nodejs.org/).

### Usage
1. Clone this repository to your local machine using the following command:
    ```bash
    git clone https://github.com/JosephOkumu/Guess-it-1
    ```
2. Navigate into the guess-it-1 directory, then navigate to the student directory:
    ```bash
    cd guess-it-1
    cd student
    ```
3. Run the program:
    ```bash
    go run .
    ```
    Enter your numbers and checkout the next prediction range.

    Example:
    ```bash
    100 --> the standard input
    189 --> the standard input
    120 200    --> the range for the next input, in this case for the number 113
    113 --> the standard input
    160 230    --> the range for the next input, in this case for the number 121
    121 --> the standard input
    110 140    --> the range for the next input, in this case for the number 114
    114 --> the standard input
    100 200    --> the range for the next input, in this case for the number 145
    145 --> the standard input
    1 99      --> the range for the next input, in this case for the number 110
    110 --> the standard input
    100 150    --> the range for the next input, in this case for the number
    ```

### Testing
1.  Download the zip file for testing [here](https://assets.01-edu.org/guess-it/guess-it-dockerized.zip)
2.  Extract the zip file.
3.  Copy the student folder from this project and paste it inside the extracted folder.
4.  Follow the instructions present in the ReadMe of the extracted folder for testing.
5. Instead of "docker composer up" command, use "npm start" to run the test.

### Features
- Displays the range for the next user input.

### Contributions
Pull requests are welcome. You can contribute to this project by adding new features, optimizing the algorithm, or fixing bugs.

For major changes, please open an issue first to discuss what you would like to change.

### Author
[JosephOkumu](https://github.com/JosephOkumu)

### License
[LICENSE](./LICENSE)


### Credits
[Zone01Kisumu](https://www.zone01kisumu.ke/)
