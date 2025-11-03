# CLI Expense Manager
_Short Summary_
An expense manager works in command line with critical features such as adding, deleting and viewing expense data. Works with JSON file. <br>

# Commands and their usage
**to ADD an expense:** add --description _whatever_ --amount _amount_ <br>
**to DELETE an expense:** delete --id _idnumber_
**to VIEW all expenses including their ID, description, amount, and date:** list
**to UPDATE any expenses:** update --id _idnumber_ --description _whatever_ / --amount _whatever_ / both, next to each other with same respective style
**to view SUMMARY of all expenses:** summary
**to view SUMMARY of all expenses BY MONTH:** summary --month _monthnumber_ 
Note: Writing monthnumber, make sure you wrote like how you are writing in an formal and global way. e.g 05, 12, 01... etc.

