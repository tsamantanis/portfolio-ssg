# Portfolio Generator

[![Go Report Card](https://goreportcard.com/badge/github.com/tsamantanis/portfolio-ssg)](https://goreportcard.com/report/github.com/tsamantanis/portfolio-ssg)

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Getting Started](#getting-started)
3. [Styles](#styles)
4. [Live Example](#live-example)

## Project Structure

```bash
📂 portfolio-ssg
├── README.md
├── examples
    ├── data.json
├── generate.go
├── styles.css
└── template.tmpl
```

## Getting Started

1. Visit [https://github.com/tsamantanis/portfolio-ssg](https://github.com/tsamantanis/portfolio-ssg) and create a new repository named `portfolio-ssg`.
2. Run each command line-by-line in your terminal to set up the project:

```bash
$ git clone git@github.com:tsamantanis/portfolio-ssg.git
$ cd portfolio-ssg
```

### Single File 📄
To generate a retro style portfolio from a single file use:

```bash
$ go run generate.go --file=PATH_TO_FILE 
```

replacing `PATH_TO_FILE` with the relative path to a `.json` file of your choosing. Place your `.json` file within the `data` directory for optimal results.

### Folder 📂
To generate multiple portfolio files from a directory use:

```bash
$ go run generate.go --dir=PATH_TO_DIRECTORY 
```

replacing `PATH_TO_DIRECTORY` with the relative path to a directory which contains `.json` files. The directory can contain sub-directories.

### Styles 🎨

You can edit the `styles.css` file to add your own styles. View an example of a retro styled portfolio page [here](https://tsamantanis.github.io/retro-portfolio/)

## Live Example

You can view a live example of this readme file generated using this portfolio generator [here](https://tsamantanis.github.io/portfolio-ssg/)
