#ifndef EXECUTE
#define EXECUTE
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/stat.h>
#include <fcntl.h>

static pid_t process_launch(char** argv){
  pid_t childpid;
  if ((childpid = fork()) == 0){
    if (execvp(argv[0], argv) < 0){
      fprintf(stderr, "-sh: %s not found\n", argv[0]);
      return (pid_t)-1;
    }else if (childpid < 0){
      fprintf(stderr, "error, couldn't start process" );
      return (pid_t)-1;
    }
  }
  return childpid;
}

static int waitfor(pid_t pid){
  int status;
  if (pid != -1){ //launch must have been a success
    do {
       waitpid(pid, &status, WUNTRACED);
    } while(!WIFEXITED(status) && !WIFSIGNALED(status));
    return 1;
  }else {
    return 0;
  }
}
#endif
