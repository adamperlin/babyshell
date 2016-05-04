#include <stdio.h>
#include <stdlib.h>
#include <readline/readline.h>
#include <readline/history.h>

static char* read_line(char* prompt){
    return readline(prompt);
}
