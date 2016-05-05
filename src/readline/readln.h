#include <stdio.h>
#include <stdlib.h>
#include <readline/readline.h>
#include <readline/history.h>
#include <signal.h>

static void signal_handler(int signo){
  signal(signo, signal_handler);
}
static char* read_line(char* prompt){
    signal(SIGWINCH, signal_handler);
    signal(SIGINT, signal_handler);
    char * ret = readline(prompt);
    add_history(ret);
    return ret;
}
