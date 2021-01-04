#include<stdio.h> 
#include<signal.h> 

void handle_sigint(int sig) 
{ 
	printf("Caught signal %d\n", sig); 
} 

int main() 
{ 
	struct sigaction act;
	act.sa_handler = handle_sigint;
	sigemptyset(&act.sa_mask);
	act.sa_flags = 0;
	sigaction(SIGINT, &act, NULL); 
	while (1) ; 
	return 0; 
} 
