#include <ncurses.h>

int main(){

initscr();
printw("works");
refresh();
getch();
endwin();
return 0;

}
