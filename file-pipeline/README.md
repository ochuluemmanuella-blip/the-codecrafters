# CodeCrafters — Operation Gopher Protocol

## Module: File Pipeline

**Author:** Ochulu Agbenu Emmanuella
**Squad:** The Interface

---

## Project Overview

This program processes raw field reports (input.txt) and transforms them into a clean, structured report (output.txt) based on a predefined pipeline of transformation rules.

It simulates SENTINEL’s archive processor, ensuring that messy, inconsistent text data is standardized, formatted, and ready for analysis.

---

## How to Run


go run . input.txt output.txt

### Requirements:

* Go installed
* Input file must exist
* Input and output files must be different

---

##  Input Format

The input file must contain at least 15 lines and include the following types:

* Lines in ALL CAPS
* Lines in all lowercase
* Lines with leading/trailing spaces
* Lines that are only dashes (`-----`) or blank


## Transformation Rules (in order)

Each line passes through the following pipeline:

1. **Trim whitespace**
   Removes leading and trailing spaces.

2. **Reverse lines containing "REVERSE"**
   Any line containing the word `REVERSE` has its word order reversed.

3. **Convert ALL CAPS → Title Case**
   Example:
   `THIS IS A LINE` → `This Is A Line`

4. **Convert lowercase → UPPERCASE**
   Example:
   `this is a line` → `THIS IS A LINE`

5. **Remove dashed or blank lines**
   Lines that are empty or contain only dashes are removed.

---

### 1. Header

```
SENTINEL FIELD REPORT — PROCESSED
```

### 2. Numbered Lines

Each processed line is numbered:

```
1. First line
2. Second line
```

### 3. Summary Block

```
===============================================
Lines read    : X
Lines written : X
Lines removed : X
Rules applied : capsToTitle, lowerToUpper, trimWhitespace, reverseIfNeeded, removeDashesOrBlanks
===============================================
```

---

## Terminal Output

After execution, the program prints:

```
> Lines read    : X
> Lines written : X
> Lines removed : X
> Rules applied : Convert ALL CAPS to Title Case, Convert lowercase to uppercase, Trim whitespace, Reverse words containing REVERSE, Remove dashes/blanks
```

---

## Edge Cases Handled

* Missing input file > error message
* Empty file > warning, no crash
* Only blank/dash lines > empty output (except header/summary)
* Same input/output file > rejected
* Output path is a directory > rejected.

---
