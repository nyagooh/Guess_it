# Number Range Predictor

This Go program reads a series of  numbers from the standard input and predicts the range of the next input number based on the mean and standard deviation of the previous inputs. The prediction is expressed as an interval [lower, upper].

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [How It Works](#how-it-works)
- [Error Handling](#error-handling)
- [Example](#example)
- [License](#license)

## Introduction

The program takes  numbers as input and predicts the interval in which the next number might fall. It calculates the prediction interval using the mean and standard deviation of the previous numbers and provides a two-standard-deviation interval around the mean.

## Installation

- Make sure you have Go installed. If not, you can install it from [here](https://golang.org/doc/install).
- Clone this repository or copy the `main.go` file to your workspace.
- Navigate to the folder where the `main.go` file is stored.

```bash
git clone https://learn.zone01kisumu.ke/git/nymaina/guess-it-1
cd Guess_IT
cd student
```
## Usage
This program reads input from the standard input (stdin). You can manually input numbers line by line, or you can pipe data into the program.

### Running the Program

```bash
./script.sh
```
## How It Works

- The program reads numbers line by line and stores them in a slice of floats (previousnumbers). For each new input:
 - If it's the first input, it prints NaN since no range can be calculated.

- If it's the second input or beyond, the program calculates the mean and standard deviation of the previous inputs using the following formulas:

Mean: mean = sum / n

Standard Deviation (SD): SD = sqrt(variance / n)

-  Variance is calculated as the sum of squared differences from the mean.

- The prediction interval is determined by subtracting and adding 2 standard deviations from the mean:

 Lower Limit = mean - 2 * SD

 Upper Limit = mean + 2 * SD

- The interval is rounded to the nearest integer for output.

## Error Handling
 - Invalid Input: If a line does not contain a valid number, the program will print an error and skip that line. 

- Input Errors: If there's an error while reading from standard input, it will print the error message at the end.

Example Error Output:
```bash
Error: strconv.ParseFloat: parsing "abc": invalid syntax
```
### [License](/home/nymaina/Guess_it/LICENSE)
