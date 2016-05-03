#include <stdio.h>
#include <ncurses.h>
#define UPARROW 72
#define DOWNARROW 80
#define RIGHTARROW 77
#define LEFTARROW 75

int main(){
  initscr();
  raw();
  keypad(stdscr, TRUE);
  noecho();
  int column, line;
  WINDOW *win;

  int c;
  while((c = getch()) != 'q'){
    switch(c){
      case KEY_UP:  break;
      case KEY_DOWN: break;
      case KEY_LEFT: getyx(win, line, column);
      //  printw("y: %d; x: %d ", line, column);
        move(line-1, column-1);
        refresh();
       break;
      case KEY_RIGHT:
        getyx(win, line, column);
      //  printw("y: %d; x: %d ",line, column );
        move(line-1, column+1);
        refresh();
       break;
    }
  }
  refresh();
  getch();
  endwin();

  return 0;
}
