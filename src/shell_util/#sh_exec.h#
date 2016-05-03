#ifndef SH_EXEC
#define SH_EXEC
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/stat.h>
#include <fcntl.h>

typedef struct {
  char** argv;
} command;

static inline pid_t process_launch(int in, int out, int err, char** argv){
  pid_t childpid;
  if ((childpid = fork()) == 0){
    if (in != 0){
      dup2(in, 0);
      close(in);
    }
    if (out != 1){
      dup2(out, 1);
      close(out);
    }
    if (err != 2){
      dup2(err, 2);
      close(err);
    }
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

static inline int waitfor(pid_t pid){
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

static inline int exec_cmd(char** argv, int pipes){
  int status;
  pid_t pid = process_launch(0, 1, 2, argv);
  status = waitfor(pid);
  if (status){
    return 1;
  }else {
    return 0;
  }
}

static inline int exec_cmd_pipe(int pipes, command* cmd){ //takes command arguments, and number of pipes as an argument
  int in, fd[2];
  in = 0;
  int i;
  for (i = 0; i < pipes - 1; i++){
    pipe(fd);
    process_launch(in, fd[1], 2, cmd[i].argv);
    close(fd[1]);
    in = fd[0];
  }
  if (in != 0) {
    dup2(in, 0);
  }
return execvp(cmd[i].argv[0], cmd[i].argv);
}

static inline int io_redirect(command* cmd, char mode, char* filepath){
  printf("mode is: %s", mode); // mode is either > or <. left associative, assumes command will be on right, place to redirect will be on left.
  int fd;
  fd = open(filepath, O_RDWR); // get a file descriptor, for rw usage (don't know which it will be used for)
  if (fd == -1) { //if error:
    fprintf(stderr, "-sh: file %s not found\n",filepath);
    return 0;
  }
  if (mode == '>'){ //redirect stdout
    return process_launch(0, fd, 2, cmd->argv);
  }else if (mode == '<'){ //redirect stdin
    return process_launch(fd, 1, 2, cmd->argv);
  }else {
    close(fd);
    return 0; //if no mode was given, nothing can be done.
  }
}

#endif
